package entities

type User struct {
	Username   string      `bson:"username,omitempty" json:"username,omitempty"`
	Identity   string      `bson:"identity,omitempty" json:"identity,omitempty"`
	FullName   string      `bson:"fullName,omitempty" json:"full_name,omitempty"`
	Position   *string     `bson:"position,omitempty" json:"position,omitempty"`
	Department *Department `bson:"department,omitempty" json:"department,omitempty"`
}
