package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"univers/pkg/converter"
	"univers/pkg/restapi/operations/region"
	"univers/pkg/store"
)

var regionHandler *RegionHandler

type RegionHandler struct {
	st *store.Store
}

func GetRegionHandler() *RegionHandler {
	if nil == regionHandler {
		regionHandler = &RegionHandler{}
		regionHandler.init()
	}

	return regionHandler
}

func (h *RegionHandler) init() {
	h.st = store.GetStore()
}

func (h *RegionHandler) GetRegions(params region.GetRegionsParams) middleware.Responder {
	regions := h.st.GetRegions()

	return region.NewGetRegionsOK().WithPayload(converter.ConvertRegionFromEntityToModel(regions))
}