package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type DoctorUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
}

func NewDoctorUser(email, firstName, lastName, username, password string) *DoctorUser {
	return &DoctorUser{
		ID:        primitive.NewObjectID(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Password:  password,
	}
}
