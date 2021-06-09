package entities

import "time"

type Audit struct {
	Id          string    `bson:"_id,omitempty" json:"id"`
	CreatedBy   *User     `bson:"createdBy,omitempty"`
	CreatedDate time.Time `bson:"createdDate"`
	UpdatedBy   *User     `bson:"updatedBy,omitempty"`
	UpdatedDate time.Time `bson:"updatedDate"`
	Version     uint64    `bson:"version"`
	Delete      bool      `bson:"delete"`
}
