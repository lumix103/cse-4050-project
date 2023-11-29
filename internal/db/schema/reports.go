package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Report struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	DoctorID        primitive.ObjectID `bson:"doctorID"`
	PatientID       primitive.ObjectID `bson:"patientID"`
	AppointmentDate time.Time          `bson:"appointmentDate"`
	Notes           string             `bson:"notes"`
}
