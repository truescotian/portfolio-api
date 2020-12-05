package schema

import (
	"errors"
	"fmt"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func MinMaxString(maxLen int) func(s string) error {
	return func(s string) error {
		if len(s) == 0 {
			return errors.New("Value cannot be empty")
		} else if len(s) > 255 {
			return errors.New("Value cannot be > " + fmt.Sprintf("%d", maxLen) + " characters")
		}
		return nil
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Validate(MinMaxString(255)),
		field.String("firstName").
			Validate(MinMaxString(255)),
		field.String("lastName").
			Validate(MinMaxString(255)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
	}
}
