package models

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/taimoor99/assignment/app/entities"
)

const msg = "test"

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) DeleteMessageById(ctx context.Context, id string) (int64, error) {
	args := m.Called(id)
	return args.Get(0).(int64), args.Error(1)
}

func (m *RepositoryMock) GetAllMessages(ctx context.Context, limit, offset int64) ([]entities.Messages, error) {
	args := m.Called(limit, offset)
	return args.Get(0).([]entities.Messages), args.Error(1)
}

func (m *RepositoryMock) FindMessageById(ctx context.Context, id string) (entities.Messages, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Messages), args.Error(1)
}

func (m *RepositoryMock) AddMessage(ctx context.Context, messages entities.Messages) (entities.Messages, error) {
	args := m.Called(messages)
	return args.Get(0).(entities.Messages), args.Error(1)
}

var ums MessageModel
var rm = new(RepositoryMock)

func init() {
	ums = NewMessageModel(rm)
}

func TestMessageModel_AddMessage(t *testing.T) {
	message := entities.Messages{
		Message: msg,
	}
	messageReq := entities.MessagesCreateInput{
		Message: msg,
	}
	rm.On("AddMessage", message).Return(entities.Messages{}, nil)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := ums.AddMessage(ctx, messageReq)
	assert.Nil(t, err)
}

func TestMessageModel_FindMessageByID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rm.On("FindMessageById", "test").Return(entities.Messages{Message: msg}, nil)
	messages, err := ums.FindMessageByID(ctx, "test")
	assert.NoError(t, err)
	assert.Equal(t, messages.Message, msg)
}

func TestMessageModel_GetAllMessages(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rm.On("GetAllMessages", int64(10), int64(0)).Return([]entities.Messages{}, nil)
	_, err := ums.GetAllMessages(ctx, 10, 0)
	assert.NoError(t, err)
}

func TestMessageModel_DeleteMessageByID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	rm.On("DeleteMessageById", "test").Return(int64(1), nil)
	count, err := ums.DeleteMessage(ctx, "test")
	assert.NoError(t, err)
	assert.Equal(t, int64(1), count)
}