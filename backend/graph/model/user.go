// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserDb struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Text string `db:"text"`
}