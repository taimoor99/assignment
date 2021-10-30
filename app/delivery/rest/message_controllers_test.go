package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/taimoor99/assignment/app/entities"
	"github.com/taimoor99/assignment/utills"
)

type ModelsMock struct {
	mock.Mock
}

const msg = "test"

func (m *ModelsMock) GetAllMessages(ctx context.Context, limit, offset int64) ([]entities.Messages, error) {
	res := []entities.Messages{entities.Messages{Message: msg}}
	return res, nil
}

func (m *ModelsMock) FindMessageByID(ctx context.Context, id string) (entities.MessageDetails, error) {
	var res entities.MessageDetails
	res.ID = primitive.NewObjectID()
	return res, nil
}

func (m *ModelsMock) AddMessage(ctx context.Context, data entities.MessagesCreateInput) (entities.Messages, error) {
	res := entities.Messages{
		Message: data.Message,
	}
	return res, nil
}

func (m *ModelsMock) DeleteMessage(ctx context.Context, id string) (int64, error) {
	return 1, nil
}

var r *chi.Mux

func init() {
	r = chi.NewRouter()
	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // AllowOriginFunc:  func(r *rest.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	ctrl := NewMessage(new(ModelsMock))
	NewRouter(r, ctrl)
}

func TestMessage_PostMessageHandler(t *testing.T) {
	body := entities.MessagesCreateInput{
		Message: msg,
	}

	payload, _ := json.Marshal(body)
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("POST", ts.URL+"/message", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response entities.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, response.Message, utills.MessageCreated)
}

func TestMessage_GetMessageDetailsByIdHandler(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/message/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response entities.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.NotEmpty(t, response.Body)
}

func TestMessage_GetAllUsersHandler(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/messages/10/0", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	var response entities.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.NotEmpty(t, response.Body)
}

func TestMessage_DeleteMessageHandler(t *testing.T) {
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, err := http.NewRequest("DELETE", ts.URL+"/message/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response entities.JsonResponse
	if err := json.Unmarshal(respBody, &response); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	assert.Equal(t, response.Message, utills.MessageDeleted)
}
