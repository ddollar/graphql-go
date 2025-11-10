package noop_test

import (
	"testing"

	"go.ddollar.dev/graphql-go"
	"go.ddollar.dev/graphql-go/example/starwars"
	"go.ddollar.dev/graphql-go/trace/noop"
	"go.ddollar.dev/graphql-go/trace/tracer"
)

func TestInterfaceImplementation(t *testing.T) {
	var _ tracer.ValidationTracer = &noop.Tracer{}
	var _ tracer.Tracer = &noop.Tracer{}
}

func TestTracerOption(t *testing.T) {
	_, err := graphql.ParseSchema(starwars.Schema, nil, graphql.Tracer(noop.Tracer{}))
	if err != nil {
		t.Fatal(err)
	}
}
