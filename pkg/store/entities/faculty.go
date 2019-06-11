package entities

type Faculty struct {
	ID int64
	FacultyName string
	UniversityID int64
	Conditions []Condition

	University University
}

func (f *Faculty) TableName() string {
	return "faculty"
}