package httperrors

import "univers/pkg/models"

var ObjectNotFound = &models.ErrorResult{
	Code: "002-001",
	Message: "Object not found",
}
var WrongApiKeyError = &models.ErrorResult{
	Code: "001-002",
	Message: "Wrong api key",
}
var ServiceUnavailable = &models.ErrorResult{
	Code: "000-000",
	Message: "Service unavailable",
}