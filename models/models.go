package models

type (
	Todo struct {
		ID          uint   `bson:"id"`
		Name        string `bson:"name"`
		Description string `bson:"description"`
		// true -> done, false -> undone
		Status bool `bson:"status"`
	}
)
