package main

import (
	"github.com/graphql-go/graphql"
)

var s = graphql.SchemaConfig{
	Query: graphql.NewObject(graphql.ObjectConfig{
		Name: "Root",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "test", nil
				},
			},
			"hello2": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "test", nil
				},
			},
		},
	}),
}

var scm *graphql.Schema

func schema() (*graphql.Schema, error) {
	if scm != nil {
		return scm, nil
	}
	ns, err := graphql.NewSchema(s)
	scm := &ns
	if err != nil {
		panic(err)
	}
	return scm, nil
}
