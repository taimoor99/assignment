package mongo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/taimoor99/assignment/app/entities"
)

const msg = "aba"

type MessageRepositoryTestSuite struct {
	suite.Suite
	MessageRepository MessageRepository
}

func (suite *MessageRepositoryTestSuite) SetupTest() {
	m, err := GetSession(context.Background(), "test_db", "127.0.0.1")
	if err != nil {
		panic(err)
	}
	m.Collection(entities.Messages{}.MessagesCollection()).Drop(context.Background())
	suite.MessageRepository = NewMessageRepository(m)
}

func TestMessageRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(MessageRepositoryTestSuite))
}

func (suite *MessageRepositoryTestSuite) TestAddMessage() {
	req := entities.Messages{}
	res, err := suite.MessageRepository.AddMessage(context.Background(), req)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), res.ID)
}

func (suite *MessageRepositoryTestSuite) TestFindMessageById() {
	messageBody := entities.Messages{
		Message: msg,
	}
	msg, err := suite.MessageRepository.AddMessage(context.Background(), messageBody)
	assert.Nil(suite.T(), err)
	messageFind, err := suite.MessageRepository.FindMessageById(context.Background(), msg.ID.Hex())
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), messageFind.ID)
}

func (suite *MessageRepositoryTestSuite) TestGetAllMessages() {
	body := entities.Messages{
		Message: msg,
	}
	_, err := suite.MessageRepository.AddMessage(context.Background(), body)
	assert.Nil(suite.T(), err)
	res, err := suite.MessageRepository.GetAllMessages(context.Background(), int64(10), int64(0))
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), len(res) == 1)
}

func (suite *MessageRepositoryTestSuite) TestDeleteMessage() {
	req := entities.Messages{}
	res, err := suite.MessageRepository.AddMessage(context.Background(), req)
	assert.Nil(suite.T(), err)
	count, err := suite.MessageRepository.DeleteMessageById(context.Background(), res.ID.Hex())
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), int64(1), count)
}