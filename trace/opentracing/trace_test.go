package opentracing_test

import (
	"testing"

	"go.ddollar.dev/graphql-go"
	"go.ddollar.dev/graphql-go/example/starwars"
	"go.ddollar.dev/graphql-go/trace/opentracing"
	"go.ddollar.dev/graphql-go/trace/tracer"
)

func TestInterfaceImplementation(t *testing.T) {
	var _ tracer.ValidationTracer = &opentracing.Tracer{}
	var _ tracer.Tracer = &opentracing.Tracer{}
}

func TestTracerOption(t *testing.T) {
	_, err := graphql.ParseSchema(starwars.Schema, nil, graphql.Tracer(opentracing.Tracer{}))
	if err != nil {
		t.Fatal(err)
	}
}
