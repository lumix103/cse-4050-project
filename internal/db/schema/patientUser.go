package schema

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PatientUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	FirstName string             `bson:"firstName"`
	LastName  string             `bson:"lastName"`
	DOB       time.Time          `bson:"dob"`
	Gender    string             `bson:"gender,omitempty"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
}

func NewPatientUser(email, firstName, lastName string, dob time.Time, gender, username, password string) *PatientUser {
	return &PatientUser{
		ID:        primitive.NewObjectID(),
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		DOB:       dob,
		Gender:    gender,
		Username:  username,
		Password:  password,
	}
}

func CheckIfPatientExists(client *mongo.Client, email string) (bool, error) {
	collection := client.Database("Debug").Collection("patientUser")

	if collection == nil {
		return false, fmt.Errorf("failed to get the collection Debug:patientUser")
	}

	filter := bson.M{"email": email}

	count, err := collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func InsertPatientUser(client *mongo.Client, patient *PatientUser) error {
	collection := client.Database("Debug").Collection("patientUser")

	if collection == nil {
		return fmt.Errorf("failed to get the collection Debug:patientUser")
	}

	_, err := collection.InsertOne(context.Background(), patient)
	if err != nil {
		return err
	}

	return nil
}
