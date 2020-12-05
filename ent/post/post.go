// Code generated by entc, DO NOT EDIT.

package post

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"

	// Table holds the table name of the post in the database.
	Table = "posts"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "posts"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_posts"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Post type.
var ForeignKeys = []string{
	"user_posts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
)