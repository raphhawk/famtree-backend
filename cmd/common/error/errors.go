package error

import "net/http"

type ArtificialErrors struct {
	Type       string `json:"type"`
	Desc       string `json:"desc"`
	StatusCode int    `json:"status"`
	Error      error  `json:"error"`
	Others     any    `json:"log"`
}

var (
	InsertSuccess   = ArtificialErrors{Type: "[INFO]", Desc: "Data saved successfully", StatusCode: http.StatusCreated}
	AlterSuccess    = ArtificialErrors{Type: "[INFO]", Desc: "Data updated successfully", StatusCode: http.StatusOK}
	InternalDbError = ArtificialErrors{Type: "[ERROR]", Desc: "Unable to access DB", StatusCode: http.StatusInternalServerError}
	InvalidDBId     = ArtificialErrors{Type: "[ERROR]", Desc: "The requested data is invalid", StatusCode: http.StatusNotFound}
	InvalidDate     = ArtificialErrors{Type: "[ERROR]", Desc: "The requested time is invalid", StatusCode: http.StatusInternalServerError}
)
