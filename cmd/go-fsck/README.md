# go-fsck

The `go fmt` for your package layout.

Currently, `go-fsck extract --pretty-json` will render the schema for a
package into a local `go-fsck.json` file. No work on restore yet.

## Motivation

The go/ast package is essentially very simple. There are only a few
declaration types in the language, `var`, `const`, `type` and `func`, and
that's about it for possible global symbols an application developer
cares about. A special case is the package level documentation, a
comment. There are a few other edge cases where the declaration may not
make sense, but for the most part, this encompases the go type system.

### Goals:

- group all `var` declarations into `vars.go`,
  - optional: group `var Err...` into `errors.go`.
  - any good convention to follow to know ErrSomething belongs to Something{} struct?
- group all `const` declarations into `const.go`,
- group all functions without receivers into `funcs.go`,
  - classify if there's a pattern we can follow to see if some of it belogs to struct internals.
- group all types into `<name>.go`,
- group all `Test<Name>*` functions into `<name>_test.go`,
- group remaining functions into `funcs_test.go`,
- group all the interfaces into `interfaces.go`,
- store package doc in `doc.go`.

### Non-goals:

- build tags?
- unnamed `_` vars?
- supporting `./...` to reformat the world (do we need it?)

Things that are enabled by this:

Restructuring the package to above conventions would let us surface
bounded contexts for individual declarations. Surfacing bounded context
for declarations uses `go test` to reduce scope only to particular files.
Code may not be coupled to anything in the package (strict) and if we can
test for that, we can move it out. Moving things out lets us test better.

For each resulting declaration, we can surface bounded contexts like so:

- strict: `go test <name>*.go const.go`
- with vars: `go test <name>*.go const.go vars.go`
- with funcs: `go test <name>*.go const.go funcs.go`
- with funcs and vars: `go test <name>*.go const.go vars.go funcs.go`
- additional cases for all with `interfaces.go`.

Now, code, with small adjustments, may be possible to become strictly
bounded. For example, it may implement an internal function that landed
in `funcs`, but is not used otherwise. Running the strict check will
surface these explicit couplings and let us know which declarations
depend on others, and what the coupling level inside the package is.

Anything that's not a public declaration inside `vars.go` is a code
smell, hinting at global singletons. It takes additional conventions to
make singletons safe (e.g. interfaces, mutexes, pointer swaps, etc.).
Having those grouped in a nice little `vars.go` file is nice. Globals
need to be understood and protected and testing with t.Parallel is going
to be a pain if the data is shared. Even reusing global loggers is a
code smell, because you can never move that file out without changes.

# Summary

This tool will let us pick code apart more safely. We can see what's
already implemented in ways that let us extract it from large package
scope. The benefit of smaller package is focus when addressing defects,
and this is the main goal of the tool, to enable that analysis and act on
the data. We often don't know how large problems are due to large package
scopes and couplings, this gives us data.

# Fidelity

As it may produce unwanted results, the way to use the tool is to
generate it from a package, and output to a new package. Using it
is expected to have bugs (I am my own QA), but - here's a few caveats:

- the premise is simple: the package would compile if we had all the
  symbols in one file, or if we had them scattered in a thousand,
- when we essentially restructure the package, this is a significant
  automated change. The change will be attributed to the commiter,
- if you'd use it, i'd suggest a git hook to check it on pre_commit,
  or even better, run it by hand in `task fmt` or something,
- it may not work in various use cases, things like go version may be
  problematic, generally we build it on a recent one and see,
- just consider it an academics tool, rather than a CI one. I don't
  expect this to be stable, so control the invocation.
- i mean, it's in the experimental repo...
