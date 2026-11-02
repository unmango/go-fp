{
  description = "Functional programming for Go";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    systems.url = "github:UnstoppableMango/nix-systems";
    flake-parts.url = "github:hercules-ci/flake-parts";

    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;
      imports = with inputs; [
        systems.flakeModule or { }
        treefmt-nix.flakeModule
      ];

      perSystem =
        {
          inputs',
          pkgs,
          lib,
          system,
          ...
        }:
        let
          inherit (inputs'.gomod2nix.legacyPackages) buildGoApplication;
        in
        {
          _module.args.pkgs = import inputs.nixpkgs {
            inherit system;
            overlays = [
              inputs.gomod2nix.overlays.default
              (final: prev: { gomod2nix = prev.gomod2nix.override { go = final.go_1_27; }; })
            ];
          };

          packages.default = pkgs.callPackage ./nix {
            inherit buildGoApplication;
            go = pkgs.go_1_27;
            pname = "fp";
            version = (lib.importJSON ./.release-please-manifest.json).".";
          };

          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              git
              gnumake
              go_1_27
              gomod2nix
              gopls
              ginkgo
              nixfmt
            ];

            GINKGO = "${pkgs.ginkgo}/bin/ginkgo";
            GO = "${pkgs.go_1_27}/bin/go";
            GOMOD2NIX = "${pkgs.gomod2nix}/bin/gomod2nix";
          };

          treefmt = {
            programs.actionlint.enable = true;
            programs.nixfmt.enable = true;
            programs.gofmt = {
              enable = true;
              package = pkgs.go_1_27;
            };
          };
        };
    };
}
