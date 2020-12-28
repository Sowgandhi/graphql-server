package graph

import (
	"graphql-server/graph/model"
)

//go: generate go run github.com/99designs/gqlgen

// Resolver ...
type Resolver struct {
	employees []*model.Employee
}
