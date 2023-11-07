package schema

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

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

func CheckIfDoctorExists(client *mongo.Client, email, username string) (bool, error) {
	collection := client.Database("Debug").Collection("doctorUser")

	if collection == nil {
		return false, fmt.Errorf("failed to get the collection Debug:doctorUser")
	}

	filter := bson.M{
		"$or": []bson.M{
			{"email": email},
			{"username": username},
		},
	}

	count, err := collection.CountDocuments(context.Background(), filter)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func FetchDoctorBy(client *mongo.Client, field, value string) (*DoctorUser, error) {
	collection := client.Database("Debug").Collection("doctorUser")

	if collection == nil {
		return nil, fmt.Errorf("failed to get the collection Debug:doctorUser")
	}

	filter := bson.M{field: value}

	var patient DoctorUser
	if err := collection.FindOne(context.Background(), filter).Decode(&patient); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &patient, nil
}

func InsertDoctorUser(client *mongo.Client, patient *DoctorUser) error {
	collection := client.Database("Debug").Collection("doctorUser")

	if collection == nil {
		return fmt.Errorf("failed to get the collection Debug:doctorUser")
	}

	_, err := collection.InsertOne(context.Background(), patient)
	if err != nil {
		return err
	}

	return nil
}
