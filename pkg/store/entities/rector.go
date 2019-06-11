package entities

type Rector struct {
	ID int64
	FirstName string
	FamilyName string
	Patronymic string
}

func (r *Rector) TableName() string {
	return "rector"
}
