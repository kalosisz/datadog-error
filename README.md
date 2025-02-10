This setup produces the error

```shell
../../go/pkg/mod/github.com/!data!dog/datadog-agent/pkg/trace@v0.62.1/traceutil/otel_util.go:413:46: not enough arguments in call to tr.ResourceToSource
        have (context.Context, pcommon.Resource, "go.opentelemetry.io/otel/attribute".Set)
        want (context.Context, pcommon.Resource, "go.opentelemetry.io/otel/attribute".Set, "github.com/DataDog/opentelemetry-mapping-go/pkg/otlp/attributes".HostFromAttributesHandler)
```

The go mod indirect dependencies were updates using `go get -u -t ./...`.
When only the `gopkg.in/DataDog/dd-trace-go.v1` was installed, 
it did not produce any error. 
It needs an update run.
