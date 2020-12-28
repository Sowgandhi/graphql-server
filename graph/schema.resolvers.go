package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphql-server/graph/generated"
	"graphql-server/graph/model"
	"math/rand"
)

// This is where your application code lives. generated.go will call into this to get the data the user has requested.

func (r *mutationResolver) AddEmployee(ctx context.Context, input model.NewEmployee) (*model.Employee, error) {
	employee := &model.Employee{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Name:   input.Name,
		Age:    input.Age,
		Salary: input.Salary,
		Address: &model.Address{
			Street:  input.Street,
			City:    input.City,
			State:   input.State,
			PinCode: input.PinCode,
		},
	}
	r.employees = append(r.employees, employee)
	return employee, nil
}

func (r *queryResolver) Employees(ctx context.Context) ([]*model.Employee, error) {
	return r.employees, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
