package utills

import (
	"encoding/json"
	"net/http"

	"github.com/taimoor99/assignment/app/entities"
)

const MessageCreated = "message created"
const MessageDeleted = "message deleted"
const MessageIdNotFoundInParam = "message id not not found param"
const LimitNotFoundInParam = "limit not not found param"
const OffsetNotFoundInParam = "offset not not found param"

func WriteJsonRes(w http.ResponseWriter, statusCode int, body interface{}, message string) {
	res := entities.JsonResponse{Message: message, Body: body}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(uj)
	return
}
