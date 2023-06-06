package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// Workout model
type Workout struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	UserID    int                 `json:"userID"`
	CreatedAt time.Time           `json:"createdAt"`
	Exercises WorkoutExerciseList `json:"exercises"`
}

// Workout Exercise model
type WorkoutExercise struct {
	ID               int                    `json:"ID"`
	ExerciseEntityID int                    `json:"exerciseEntityID"`
	Name             string                 `json:"name,omitempty"`
	Number           int                    `json:"number,omitempty"`
	Sets             WorkoutExerciseSetList `json:"workoutSets"`
}

// Workout Set type
type SetType string

// Definitions of different set types
const (
	RegularSetType SetType = "regular"
	DropSetType    SetType = "dropset"
	SuperSetType   SetType = "superset"
	WarmupSet      SetType = "warmup"
)

// Workout Set model
type WorkoutSet struct {
	ID          int     `json:"id"`
	Number      int     `json:"number,omitempty"`
	Type        SetType `json:"type"`
	Weight      float32 `json:"weight"`
	Repetitions uint    `json:"reps"`
}

// Workout Exercise Collection Type
type WorkoutExerciseSetList []WorkoutSet

// Value implements the driver Value method
func (wesl WorkoutExerciseSetList) Value() (driver.Value, error) {
	return json.Marshal(wesl)
}

// Scan implements the Scan method
func (wesl *WorkoutExerciseSetList) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	err := json.Unmarshal(b, &wesl)
	return err
}

// Workout Exercise Collection Type
type WorkoutExerciseList []WorkoutExercise

// Value implements the driver Value method
func (wel WorkoutExerciseList) Value() (driver.Value, error) {
	return json.Marshal(wel)
}

// Scan implements the Scan method
func (wel *WorkoutExerciseList) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	err := json.Unmarshal(b, &wel)
	return err
}
