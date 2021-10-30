package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/taimoor99/assignment/app/entities"
	"github.com/taimoor99/assignment/app/models"
	"github.com/taimoor99/assignment/utills"
)

type message struct {
	Models models.MessageModel
}

type MessageControllers interface {
	GetMessageDetailsByIdHandler(w http.ResponseWriter, r *http.Request)
	PostMessageHandler(w http.ResponseWriter, r *http.Request)
	GetAllMessagesHandler(w http.ResponseWriter, r *http.Request)
	DeleteMessageHandler(w http.ResponseWriter, r *http.Request)
}

func NewMessage(messageModels models.MessageModel) MessageControllers {
	return &message{
		Models: messageModels,
	}
}

func (m *message) GetMessageDetailsByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) < 1 {
		utills.WriteJsonRes(w, http.StatusBadRequest, nil, utills.MessageIdNotFoundInParam)
		return
	}

	msg, err := m.Models.FindMessageByID(context.Background(), id)
	if err != nil {
		utills.WriteJsonRes(w, http.StatusNotFound, nil, err.Error())
		return
	}

	utills.WriteJsonRes(w, http.StatusOK, msg, "")
	return
}

func (m *message) PostMessageHandler(w http.ResponseWriter, r *http.Request) {
	var messageReq entities.MessagesCreateInput
	if err := entities.DecodeAndValidate(r, &messageReq); err != nil {
		utills.WriteJsonRes(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	messageId, err := m.Models.AddMessage(context.Background(), messageReq)
	if err != nil {
		utills.WriteJsonRes(w, http.StatusNotFound, nil, err.Error())
		return
	}

	utills.WriteJsonRes(w, http.StatusOK, messageId, utills.MessageCreated)
	return
}

func (m *message) GetAllMessagesHandler(w http.ResponseWriter, r *http.Request) {
	limit, err := strconv.Atoi(chi.URLParam(r, "limit"))
	if err != nil {
		utills.WriteJsonRes(w, http.StatusBadRequest, nil, utills.LimitNotFoundInParam)
		return
	}
	offset, err := strconv.Atoi(chi.URLParam(r, "offset"))
	if err != nil {
		utills.WriteJsonRes(w, http.StatusBadRequest, nil, utills.OffsetNotFoundInParam)
		return
	}

	messages, err := m.Models.GetAllMessages(context.Background(), int64(limit), int64(offset))
	if err != nil {
		utills.WriteJsonRes(w, http.StatusNotFound, nil, err.Error())
		return
	}

	utills.WriteJsonRes(w, http.StatusOK, messages, "")
	return
}

func (m *message) DeleteMessageHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if len(id) < 1 {
		utills.WriteJsonRes(w, http.StatusBadRequest, nil, utills.MessageIdNotFoundInParam)
		return
	}

	rowCount, err := m.Models.DeleteMessage(context.Background(), id)
	if err != nil {
		utills.WriteJsonRes(w, http.StatusNotFound, nil, err.Error())
		return
	}

	utills.WriteJsonRes(w, http.StatusOK, rowCount, utills.MessageDeleted)
	return}