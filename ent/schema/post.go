package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Validate(MinMaxString(255)),
		field.String("description").
			Validate(MinMaxString(255)),
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("posts").
			Unique().
			// Edge is required on entity creation.
			// Post cannot be created without its owner.
			Required(),
	}
}
