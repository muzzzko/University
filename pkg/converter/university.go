package converter

import (
	"univers/pkg/models"
	"univers/pkg/store/entities"
)

func ConvertUniversityFromEntityToModel(university *entities.University) *models.University {
	faculties := make([]*models.UniversityFacultiesItems, len(university.Faculties))
	for i, faculty := range university.Faculties {
		conditions := make([]*models.UniversityFacultiesItemsConditionsItems, len(faculty.Conditions))
		for i, condition := range faculty.Conditions {
			conditions[i] = &models.UniversityFacultiesItemsConditionsItems{
				ID: condition.ID,
				AdmissionCondition: condition.AdmissionCondition,
			}
		}
		faculties[i] = &models.UniversityFacultiesItems{
			ID: faculty.ID,
			Name: faculty.FacultyName,
			Conditions: conditions,
		}
	}

	return &models.University{
		ID: university.ID,
		Name: university.UniversityName,
		Faculties: faculties,
		Shape: university.Shape,
		Status: university.Status,
		StudentNumber: university.StudentNumber,
		Address: university.Address,
		Rector: &models.UniversityRector{
			ID: university.Rector.ID,
			FirstName: university.Rector.FirstName,
			FamilyName: university.Rector.FamilyName,
			Patronymic: university.Rector.Patronymic,
		},
		Region: &models.UniversityRegion{
			ID: university.Region.ID,
			Name: university.Region.RegionName,
		},
	}
}

func ConvertUniversitiesFromEntityToModel(universities []*entities.University) []*models.GetUniversitiesOKBodyUniversitiesItems {
	body := make([]*models.GetUniversitiesOKBodyUniversitiesItems, len(universities))
	for i, university := range universities {
		body[i] = &models.GetUniversitiesOKBodyUniversitiesItems{
			ID: university.ID,
			Name: university.UniversityName,
		}
	}

	return body
}

func ConvertUniversityFromModelToEntity(universityData *models.UniversityData) *entities.University {
	faculties := make([]entities.Faculty, len(universityData.Faculties))
	for i, faculty := range universityData.Faculties {
		conditions := make([]entities.Condition, len(faculty.Conditions))
		for j, condition := range faculty.Conditions {
			conditions[j] = entities.Condition{
				AdmissionCondition: condition,
			}
		}
		faculties[i] = entities.Faculty{
			FacultyName: faculty.Name,
			Conditions: conditions,
		}
	}

	return &entities.University{
		UniversityName: universityData.Name,
		Faculties: faculties,
		Shape: universityData.Shape,
		Status: universityData.Status,
		StudentNumber: universityData.StudentNumber,
		Address: universityData.Address,
		RegionID: universityData.RegionID,
		Rector: entities.Rector{
			FamilyName: universityData.Rector.FamilyName,
			FirstName: universityData.Rector.FirstName,
			Patronymic: universityData.Rector.Patronymic,
		},
	}
}


func ConvertUniversityUpdateDataFromModelToEntity(universityData *models.UniversityUpdateData, ID int64) *entities.University {
	university := &entities.University{
		ID: ID,
		UniversityName: universityData.Name,
		Shape: universityData.Shape,
		Status: universityData.Status,
		StudentNumber: universityData.StudentNumber,
		Address: universityData.Address,
		RegionID: universityData.RegionID,
	}

	if nil != universityData.Faculties {
		faculties := make([]entities.Faculty, len(universityData.Faculties))
		for i, faculty := range universityData.Faculties {
			faculties[i] = entities.Faculty{
				ID:          faculty.ID,
				FacultyName: faculty.Name,
			}

			if nil != faculty.Conditions {
				conditions := make([]entities.Condition, len(faculty.Conditions))
				for j, condition := range faculty.Conditions {
					conditions[j] = entities.Condition{
						ID:                 condition.ID,
						AdmissionCondition: condition.AdmissionCondition,
					}
				}
				faculties[i].Conditions = conditions
			}
		}
		university.Faculties = faculties
	}

	if nil != universityData.Rector {
		university.Rector = entities.Rector{
			ID:         universityData.Rector.ID,
			FamilyName: universityData.Rector.FamilyName,
			FirstName:  universityData.Rector.FirstName,
			Patronymic: universityData.Rector.Patronymic,
		}
	}

	return university
}