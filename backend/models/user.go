package models

type User struct {
	ID       string   `bson:"_id" json:"id"`
	Username string   `bson:"username" json:"username"`
	Password string   `bson:"password" json:"-"`
	Role     string   `bson:"role" json:"role"` // visitor, evaluator, reviewer, admin
	Email    string   `bson:"email" json:"email"`
}