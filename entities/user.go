package entities

type User struct {
	Username   string      `bson:"username" json:"username"`
	Identity   string      `bson:"identity" json:"identity"`
	FullName   string      `bson:"fullName" json:"full_name"`
	Position   string      `bson:"position,omitempty" json:"position,omitempty"`
	Department *Department `bson:"department,omitempty" json:"department,omitempty"`
}
