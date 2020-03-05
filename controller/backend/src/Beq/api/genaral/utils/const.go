package utils

import (
	"Beq/api/genaral/model"
	"net/http"
)

// State handling
var (
	errorEmailValidation = model.InnerResponse{
		Status:  http.StatusForbidden,
		Message: "Email not Valid",
	}
	errorDataBase = model.InnerResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal Error",
	}
	errorDuplicateEmail = model.InnerResponse{
		Status:  http.StatusForbidden,
		Message: "Email is already used.",
	}
	stateCreated = model.InnerResponse{
		Status:  http.StatusCreated,
		Message: "User has been registered successfully.",
	}
)
