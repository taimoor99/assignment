package models

import (
	"context"
	"github.com/taimoor99/assignment/app/entities"
	"github.com/taimoor99/assignment/app/repositories/mongo"
	"github.com/taimoor99/assignment/utills"
)

type messageRepo struct {
	Repo mongo.MessageRepository
}

type MessageModel interface {
	FindMessageByID(ctx context.Context, id string) (entities.MessageDetails, error)
	AddMessage(ctx context.Context, body entities.MessagesCreateInput) (entities.Messages, error)
	GetAllMessages(ctx context.Context, limit, offset int64) ([]entities.Messages, error)
	DeleteMessage(ctx context.Context, id string) (int64, error)
}

func NewMessageModel(messageRepository mongo.MessageRepository) MessageModel {
	return &messageRepo{
		Repo: messageRepository,
	}
}

func (m *messageRepo) FindMessageByID(ctx context.Context, userId string) (entities.MessageDetails, error) {
	msg, err := m.Repo.FindMessageById(ctx, userId)
	if err != nil {
		return entities.MessageDetails{}, err
	}
	var res entities.MessageDetails
	res.IsPalindrome = utills.IsPalindrome(msg.Message)
	res.Messages = msg
	return res, nil
}

func (m *messageRepo) AddMessage(ctx context.Context, req entities.MessagesCreateInput) (entities.Messages, error) {
	data := entities.Messages{
		Message: req.Message,
	}
	res, err := m.Repo.AddMessage(ctx, data)
	if err != nil {
		return entities.Messages{}, err
	}
	return res, nil
}

func (m *messageRepo) GetAllMessages(ctx context.Context, limit, offset int64) ([]entities.Messages, error) {
	res, err := m.Repo.GetAllMessages(ctx, limit, offset)
	if err != nil {
		return []entities.Messages{}, err
	}
	return res, nil
}

func (m *messageRepo) DeleteMessage(ctx context.Context, id string) (int64, error) {
	count, err := m.Repo.DeleteMessageById(ctx, id)
	if err != nil {
		return 0, err
	}
	return count, nil
}
