package models

type (
	Todo struct {
		Name        string `bson:"name" json:"name"`
		Description string `bson:"description" json:"description"`
		Status      bool   `bson:"status" json:"status"`
	}
)
