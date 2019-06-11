package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"univers/pkg/converter"
	"univers/pkg/httperrors"
	"univers/pkg/models"
	university_pkg "univers/pkg/restapi/operations/university"
	"univers/pkg/store"
)

type UniversityHandler struct {
	st *store.Store
}

var universityHandler *UniversityHandler

func GetUniversityHandler() *UniversityHandler {
	if nil == universityHandler {
		universityHandler = &UniversityHandler{}
		universityHandler.init()
	}

	return universityHandler
}

func (h *UniversityHandler) init() {
	h.st = store.GetStore()
}

func (h *UniversityHandler) GetUniversity(params university_pkg.GetUniversityParams) middleware.Responder {
	universityEntity, ok := h.st.GetUniversity(params.UniversityID)
	if !ok {
		return university_pkg.NewGetUniversitiesUnprocessableEntity().WithPayload(httperrors.ObjectNotFound)
	}

	return university_pkg.NewGetUniversityOK().WithPayload(converter.ConvertUniversityFromEntityToModel(universityEntity))
}

func (h *UniversityHandler) GetUniversities(params university_pkg.GetUniversitiesParams) middleware.Responder {
	universities, _ := h.st.GetUniversities(params.Offset, params.Limit, params.RegionID, params.Status)

	return university_pkg.NewGetUniversitiesOK().WithPayload(&models.GetUniversitiesOKBody{
		Universities: converter.ConvertUniversitiesFromEntityToModel(universities),
		CountUniversity: h.st.GetCountUniversity(params.RegionID, params.Status),
	})
}

func (h *UniversityHandler) PostUniversity(params university_pkg.PostUniversityParams, principal interface{}) middleware.Responder {
	university := converter.ConvertUniversityFromModelToEntity(params.Body)

	h.st.CreateUniversity(university)
	return university_pkg.NewPostUniversityOK().WithPayload(&models.PostUniversityOKBody{ID: &university.ID})
}

func (h *UniversityHandler) PatchUniversity(params university_pkg.PatchUniversityParams, principal interface{}) middleware.Responder {
	university := converter.ConvertUniversityUpdateDataFromModelToEntity(params.Body, params.UniversityID)

	h.st.UpdateUniversity(university)
	return university_pkg.NewPatchUniversityOK().WithPayload(converter.ConvertUniversityFromEntityToModel(university))
}