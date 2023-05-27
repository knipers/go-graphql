package graph

import "github.com/knipers/go-graphql/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	AuthorDB *database.Author
}
