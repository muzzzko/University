package entities

type Region struct {
	ID int64
	RegionName string
}

func (r *Region) TableName() string {
	return "region"
}