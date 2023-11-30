package schema

import (
	"context"
	"fmt"
	"time"

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

func FetchFutureAppointments(client *mongo.Client, patient PatientUser, doctor DoctorUser) ([]Appointment, error) {

	return nil, nil
}
