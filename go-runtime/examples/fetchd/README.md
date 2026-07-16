# fetchd

`fetchd` is the continuous example used by the first Go runtime source-notes
chapter. It accepts one to eight `url` query parameters, fetches them
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

Run the tests with `go test ./...`.
