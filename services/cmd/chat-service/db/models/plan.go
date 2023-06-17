package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Visibility string

const (
	VisibilityPublic  Visibility = "public"  // A workout plan is public and accessible by other users
	VisibilityPrivate Visibility = "private" // A workout plan is private and is NOT accessible by other users
)

// Workout plan model
type Plan struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	UserID      int          `json:"user_id"`
	Visibility  Visibility   `json:"visibility"`
	Exercises   ExerciseList `json:"exercises"`
}

// Plan Exercise model
type PlanExercise struct {
	ID               int    `json:"id"`
	PlanID           int    `json:"planID,omitempty"`
	SerialNumber     int    `json:"serialNumber,omitempty"`
	Name             string `json:"name,omitempty"`
	MuscleGroup      string `json:"muscleGroup,omitempty"`
	ExerciseEntityID int    `json:"exerciseEntityID"`
	Sets             int    `json:"sets"`
}

// Plan Exercise Collection Type
type ExerciseList []PlanExercise

// Value implements the driver Value method
func (el ExerciseList) Value() (driver.Value, error) {
	return json.Marshal(el)
}

// Scan implements the Scan method
func (el *ExerciseList) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	err := json.Unmarshal(b, &el)
	return err
}
