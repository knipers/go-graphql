// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewAuthor struct {
	Name string `json:"name"`
}

type NewBook struct {
	Title  string `json:"title"`
	Type   string `json:"type"`
	Author string `json:"author"`
}
