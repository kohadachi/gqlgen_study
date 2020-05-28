package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/kouheiadachi/gqlgen_study/graph/model"
)

type Resolver struct {
	todos  []*model.Todo
	robots []*model.Robot
	users  []*model.User
}
