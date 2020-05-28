// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewRobot struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Robot struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}
type RobotDb struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	Name string `json:"name"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserDb struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Text string `db:"text"`
}
