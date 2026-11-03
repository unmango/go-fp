<!-- markdownlint-disable-file MD010 -->

# fp

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/unmango/go-fp/ci.yml)
![Codecov](https://img.shields.io/codecov/c/github/unmango/go-fp)
![GitHub Release](https://img.shields.io/github/v/release/unmango/go-fp)

Small functional programming types for Go, built on go1.27 generic methods.

```shell
go get go.unmango.dev/fp
```

Every type converts to and from the ordinary Go shape it replaces,
such as `(T, bool)`, `(T, error)`, `*T` and `iter.Seq`, so it can be used at the edges of existing code.
The library depends only on the standard library.

## maybe

`Maybe[T]` is a value that may be absent. The zero value is `None`.

```go
port := maybe.FromOk(os.LookupEnv("PORT")).
	Map(strings.TrimSpace).
	OrElse("8080")
```

## result

`Result[T]` is a value or an error.

```go
n, err := result.Of(strconv.Atoi(s)).
	Map(func(x int) int { return x * 2 }).
	Get()
```

`result.Collect` turns a sequence of results into a result of a slice, stopping at the first error.

## either

`Either[L, R]` holds one of two typed values.

```go
e := either.Left[string](69) // Either[int, string]

label := e.Fold(strconv.Itoa, strings.ToUpper)
```

## lazy

`Lazy[T]` is a value computed when it is first needed.
`lazy.Once` memoizes it and is safe to force from several goroutines.

```go
config := lazy.Once(loadConfig)

config.Get()
```

## pipe

`pipe.Pipe` through `pipe.Pipe6` compose functions left to right, with every type inferred.

```go
parse := pipe.Pipe(strings.TrimSpace, strings.ToLower)
```

## semigroup and monoid

`Semigroup[T]` combines two values. `Monoid[T]` adds an identity, and `monoid.Fold` folds an `iter.Seq` with it.

```go
total := monoid.Fold(monoid.Sum[int](), maps.Values(stock))
```

Instances are provided for `Sum`, `Product`, `String`, `Slice`, `All` and `Any`.
A type with its own `Combine(T) T` method can be used directly with `semigroup.Method` and `monoid.Method`.

## Why no Functor or Monad

Earlier versions in [unmango/go](https://github.com/unmango/go) emulated higher-kinded types with function-type constraints.
Nothing used them, and Go interface methods still cannot have type parameters, so they are not part of this module.
Whether go1.27 makes any of them worthwhile is tracked in [unmango/go#198](https://github.com/unmango/go/issues/198).

## Development

```shell
nix develop   # or: direnv allow
make test     # Ginkgo suites
make check    # nix flake check: flake evaluation and formatting
make build    # nix build: runs the Ginkgo suites inside the Nix build
make tidy     # go mod tidy and gomod2nix
```
