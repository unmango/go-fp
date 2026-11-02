{
  pname,
  buildGoApplication,
  ginkgo,
  go,
  lib,
  version,
}:
buildGoApplication {
  inherit go pname version;

  src = lib.cleanSource ../.;
  modules = ./gomod2nix.toml;
  nativeBuildInputs = [ ginkgo ];

  checkPhase = ''
    ginkgo -r --race
  '';
}
