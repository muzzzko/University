package entities

type Condition struct {
	ID int64
	AdmissionCondition string
	FacultyID int64

	Faculty Faculty
}

func (c Condition) TableName() string {
	return "condition"
}