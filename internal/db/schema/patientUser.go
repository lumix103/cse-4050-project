package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatientUser struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email        string             `json:"email"`
	FirstName    string             `json:"firstName"`
	LastName     string             `json:"lastName"`
	DOB          time.Time          `json:"dob"`
	Gender       string             `json:"gender,omitempty"`
	Username     string             `json:"username"`
	Password     string             `json:"password"`
	Salt         string             `json:"salt"`
	AccessToken  []string           `json:"accessToken,omitempty"`
	RefreshToken []string           `json:"refreshToken,omitempty"`
}
