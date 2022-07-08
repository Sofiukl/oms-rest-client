# oms-rest-client #
Order Management System (OMS) REST Util

[![oms-rest-client release (latest SemVer)](https://img.shields.io/github/v/release/sofiukl/oms-rest-client?sort=semver)](https://github.com/sofiukl/oms-rest-client/releases)
[![GoDoc](https://img.shields.io/static/v1?label=godoc&message=reference&color=blue)](https://pkg.go.dev/github.com/sofiukl/oms-rest-client/)

oms-rest-client is a Go client library for rest call.

Currently, **oms-rest-client requires Go version 1.18 or greater**.  oms-rest-client tracks
[Go's version support policy][support-policy].  We do our best not to break
older versions of Go if we don't have to, but due to tooling constraints, we
don't always test older versions.

[support-policy]: https://golang.org/doc/devel/release.html#policy

## Installation ##

oms-rest-client is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/sofiukl/oms-rest-client/
```

will resolve and add the package to the current development module, along with its dependencies.

Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/sofiukl/oms-rest-client/rest-util"
```

and run `go get` without parameters.

Finally, to use the top-of-trunk version of this repo, use the following command:

```bash
go get github.com/sofiukl/oms-rest-client@master
```

## Usage ##

```go
import "github.com/sofiukl/oms-rest-client/restutil"
```

Call Function like Get, Post etc

```go
resp, err := restutil.Get(url)
```

### Integration Tests ###
PENDING

## Contributing ##
I would like to cover the all type of API calls and contributions are of course always welcome. The
calling pattern is pretty well established, so adding new methods is relatively
straightforward. See [`CONTRIBUTING.md`](CONTRIBUTING.md) for details.

## Versioning ##

In general, oms-rest-client follows [semver](https://semver.org/) as closely as we
can for tagging releases of the package. For self-contained libraries, the
application of semantic versioning is relatively straightforward and generally
understood.

* We increment the **major version** with any incompatible change.
* We increment the **minor version** with any backwards-compatible changes to
	functionality.
* We increment the **patch version** with any backwards-compatible bug fixes.

## License ##

This library is distributed under the BSD-style license found in the [LICENSE](./LICENSE)
file.
