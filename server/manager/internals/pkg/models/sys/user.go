package sys

type User struct {
	UserName string `bson:"user_name" json:"user_name"`
	Password string `bson:"password" json:"password"`
}
