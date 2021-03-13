package graphql

import "go.rayyildiz.dev/todo/pkg/port"

type resolver struct {
	service port.Service
}

func NewGraphqlResolver(svc port.Service) *resolver {
	return &resolver{svc}
}

func (r resolver) Mutation() MutationResolver {
	return newMutationResolver(r.service)
}

func (r resolver) Query() QueryResolver {
	return newQueryResolver(r.service)
}
