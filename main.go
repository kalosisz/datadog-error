package main

import "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"

func main() {
	tracer.Start(
		tracer.WithRuntimeMetrics(),
	)
	defer tracer.Stop()
}
