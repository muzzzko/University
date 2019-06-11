package entities

type University struct {
	ID int64
	UniversityName string
	Address string
	RegionID int64
	RectorID int64
	Status string
	Shape string
	StudentNumber int64

	Faculties []Faculty
	Region Region
	Rector Rector

}

func (u *University) TableName() string {
	return "university"
}