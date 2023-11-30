package schema

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Appointment struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	DoctorID        primitive.ObjectID `bson:"doctorID"`
	PatientID       primitive.ObjectID `bson:"patientID"`
	AppointmentDate time.Time          `bson:"appointmentDate"`
}

func InsertAppointment(client *mongo.Client, doctor DoctorUser, patient PatientUser, date time.Time) error {
	collection := client.Database("Debug").Collection("appointment")
	if collection == nil {
		if client.Database("Debug") == nil {
			return fmt.Errorf("could not connect to database or database does not exist")
		}
		client.Database("Debug").CreateCollection(context.Background(), "appointment")
		collection = client.Database("Debug").Collection("appointment")
	}

	app := &Appointment{
		ID:              primitive.NewObjectID(),
		DoctorID:        doctor.ID,
		PatientID:       patient.ID,
		AppointmentDate: date,
	}

	_, err := collection.InsertOne(context.Background(), app)
	if err != nil {
		return err
	}
	return nil
}

func FetchFutureAppointments(client *mongo.Client, patient *PatientUser, doctor *DoctorUser) ([]Appointment, error) {
	var field string
	var id primitive.ObjectID

	if patient != nil {
		field = "patientID"
		id = patient.ID
	} else if doctor != nil {
		field = "doctorID"
		id = doctor.ID
	} else {
		return nil, fmt.Errorf("doctor or patient is nil")
	}

	filter := bson.M{
		field: id,
		"appointmentDate": bson.M{
			"$gt": time.Now(),
		},
	}

	appointmentsCollection := client.Database("Debug").Collection("appointment")

	if appointmentsCollection == nil {
		if client.Database("Debug") == nil {
			return nil, fmt.Errorf("could not connect to database or database does not exist")
		}
		client.Database("Debug").CreateCollection(context.Background(), "appointment")
		appointmentsCollection = client.Database("Debug").Collection("appointment")
	}

	cursor, err := appointmentsCollection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	var results []Appointment

	if err := cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
