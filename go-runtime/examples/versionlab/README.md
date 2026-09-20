# Go loop-variable language-version experiment

Run the same source with the same installed Go toolchain and two different
module language versions. Go 1.22 or newer is required; the comparison was also
verified with Go 1.27.1.

From this directory:

```sh
GOTOOLCHAIN=local GOWORK=off GOEXPERIMENT= GOFLAGS= go run .
```

The runner copies [loop.go.txt](./loop.go.txt) into a temporary module, changes
only its `go.mod` directive, and checks the complete output. It uses the Go
executable from the toolchain that built the runner. It clears experiment,
workspace, and compiler-flag overrides for the child runs, prevents automatic
toolchain downloads, and removes the temporary module on exit.

Expected output:

```text
go.mod: go 1.21
closures: [3 3 3]
pointers: [3 3 3]
pointers equal: true
assignment closures: [3 3 3]
shared backing array: [99 99 99]

go.mod: go 1.22
closures: [1 2 3]
pointers: [1 2 3]
pointers equal: false
assignment closures: [3 3 3]
shared backing array: [99 99 99]
```

The installed toolchain is the compiler and standard library that build and
run the program. The module's `go` directive supplies the language version for
its packages. A new compiler therefore does not automatically give every old
module the new loop-variable semantics. This fixture has no per-file build
constraint that changes its language version.

In Go 1.22 language mode, a `range` clause that declares variables with `:=`
creates separate variables for each iteration. Delayed closures and retained
pointers consequently observe `1`, `2`, and `3`. A clause using `=` assigns to
an existing variable and continues to share it. Copying a slice into separate
variables still shares its backing array; the language change is not a deep
copy. All closures run sequentially after their loops, so these results do not
depend on goroutine scheduling and introduce no data race.

These assertions isolate loop-variable semantics. They do not emulate all
behavior of historical Go runtimes, and they do not compare timer-channel
compatibility modes: a module directive can select language rules, but it
cannot restore an implementation removed from a newer runtime.

References: [Go 1.22 language changes](https://go.dev/doc/go1.22#language),
[loop-variable transition](https://go.dev/blog/loopvar-preview), and the
[module `go` directive](https://go.dev/ref/mod#go-mod-file-go).
