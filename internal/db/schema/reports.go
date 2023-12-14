package schema

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Report struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	DoctorID        primitive.ObjectID `bson:"doctorID"`
	PatientID       primitive.ObjectID `bson:"patientID"`
	AppointmentDate time.Time          `bson:"appointmentDate"`
	Notes           string             `bson:"notes"`
}

func InsertReport(client *mongo.Client, report *Report) error {
	db := client.Database("Debug")
	collection := db.Collection("report")

	insertResult, err := collection.InsertOne(context.TODO(), report)
	if err != nil {
		return err
	}

	fmt.Println("Inserted new report with ID:", insertResult.InsertedID)
	return nil
}

func DeleteReportByID(client *mongo.Client, reportID primitive.ObjectID) error {
	db := client.Database("Debug")
	collection := db.Collection("report")

	deleteResult, err := collection.DeleteOne(context.TODO(), bson.M{"_id": reportID})
	if err != nil {
		return err
	}

	fmt.Printf("Deleted %v document(s)\n", deleteResult.DeletedCount)
	return nil
}

func GetReportsByUserID(client *mongo.Client, userID primitive.ObjectID) ([]Report, error) {
	db := client.Database("Debug")
	collection := db.Collection("report")

	filter := bson.M{"$or": []bson.M{
		{"doctorID": userID},
		{"patientID": userID},
	}}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var reports []Report
	for cursor.Next(context.TODO()) {
		var report Report
		if err := cursor.Decode(&report); err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}
