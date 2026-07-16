# fetchd

`fetchd` is the continuous example used by the Go runtime source-notes series.
It accepts one to eight `url` query parameters, fetches them
concurrently under the incoming request deadline, and returns a compact JSON
summary.

Run the service:

```bash
go run .
```

Then issue a request from another terminal:

```bash
curl --get 'http://127.0.0.1:8080/fetch' \
  --data-urlencode 'url=https://go.dev/' \
  --data-urlencode 'url=https://pkg.go.dev/net/http'
```

Run the tests with `go test ./...`. The focused labs are grouped by article:

- `value_semantics_test.go` explores copying and aliasing.
- `scheduler_lab_test.go` creates runnable bursts and reads scheduler metrics.
- `concurrency_lab_test.go` makes channel backpressure, cancellation, and
  mutex-protected state observable.
- `interface_lab_test.go` compares interface values, typed nil, generic type
  preservation, and reflection's addressability boundary.
- `context_lab_test.go` traces cancellation causes, worker joining, timeout and
  detached lifetimes, `AfterFunc`, and graceful HTTP server shutdown.
- `memory_lab_test.go` compares stack-shaped values, escaping pointers, slice
  growth and preallocation, then verifies the runtime metrics used for GC evidence.
- `transport_lab_test.go` makes HTTP/1 response-body reuse, per-host connection
  limits, cancellation, `httptrace`, and HTTP/2 stream multiplexing observable.
