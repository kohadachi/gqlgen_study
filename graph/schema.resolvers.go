package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/kouheiadachi/gqlgen_study/graph/generated"
	"github.com/kouheiadachi/gqlgen_study/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateRobot(ctx context.Context, input model.NewRobot) (*model.Robot, error) {
	robot := &model.Robot{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.robots = append(r.robots, robot)
	return robot, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	// db, err := sqlx.Open("mysql", "root@/go")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	tx := r.DB.MustBegin()
	tx.MustExec(tx.Rebind("INSERT INTO user (name, text) VALUES (?, ?)"), input.Name, input.Text)
	tx.Commit()
	user := &model.User{
		Name: input.Name,
		ID:   fmt.Sprintf("T%d", rand.Int()),
	}
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	tx := r.DB.MustBegin()
	tx.MustExec(tx.Rebind("UPDATE user SET name=?,text=? WHERE id=?"), input.Name, input.Text, input.ID)
	tx.Commit()
	user := &model.User{
		Name: input.Name,
		ID:   input.ID,
	}
	return user, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Robots(ctx context.Context) ([]*model.Robot, error) {
	return r.robots, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.UserDb, error) {
	users := []model.UserDb{}
	r.DB.Select(&users, "SELECT * FROM user where id = ?", id)

	user := &model.UserDb{
		Name: users[0].Name,
		ID:   users[0].ID,
	}
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.UserDb, error) {
	users := []model.UserDb{}
	r.DB.Select(&users, "SELECT * FROM user")
	results := []*model.UserDb{}
	for _, user := range users {
		results = append(results, &model.UserDb{
			Name: user.Name,
			ID:   user.ID,
			Text: user.Text,
		})
	}
	return results, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
