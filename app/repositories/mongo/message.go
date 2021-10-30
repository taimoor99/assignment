package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/taimoor99/assignment/app/entities"
)

func GetSession(ctx context.Context, db, url string) (*mongo.Database, error) {
	// Connect to our mongo
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI("mongodb://"+url+":27017/"))

	if err != nil {
		return nil, err
	}

	return client.Database(db), err
}

type message struct {
	mgo *mongo.Database
}

type MessageRepository interface {
	FindMessageById(ctx context.Context, id string) (entities.Messages, error)
	AddMessage(ctx context.Context, body entities.Messages) (entities.Messages, error)
	GetAllMessages(ctx context.Context, limit, offset int64) ([]entities.Messages, error)
	DeleteMessageById(ctx context.Context, id string) (int64, error)
}

func NewMessageRepository(db *mongo.Database) MessageRepository {
	return &message{
		mgo: db,
	}
}

func (u *message) FindMessageById(ctx context.Context, id string) (entities.Messages, error) {
	var msg entities.Messages
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Messages{}, err
	}
	if err := u.mgo.Collection(msg.MessagesCollection()).
		FindOne(ctx, bson.M{"_id": objID}).Decode(&msg); err != nil {
		return entities.Messages{}, err
	}
	return msg, nil
}

func (u *message) AddMessage(ctx context.Context, message entities.Messages) (entities.Messages, error) {
	message.ID = primitive.NewObjectID()
	message.CreatedAt = time.Now()
	message.UpdatedAt = time.Now()
	_, err := u.mgo.Collection(message.MessagesCollection()).InsertOne(ctx, message)
	if err != nil {
		return entities.Messages{}, err
	}
	return message, nil
}

func (u *message) GetAllMessages(ctx context.Context, limit, offset int64) ([]entities.Messages, error) {
	var msgs []entities.Messages
	option := options.FindOptions{
		Limit: &limit,
		Skip:  &offset,
	}
	cur, err := u.mgo.Collection(entities.Messages{}.MessagesCollection()).Find(ctx, bson.D{{}}, &option)
	if err != nil {
		return []entities.Messages{}, err
	}

	for cur.Next(context.TODO()) {
		var elem entities.Messages
		err := cur.Decode(&elem)
		if err != nil {
			return []entities.Messages{}, err
		}
		msgs = append(msgs, elem)
	}

	if err := cur.Err(); err != nil {
		return []entities.Messages{}, err
	}
	cur.Close(context.TODO())

	return msgs, nil
}

func (u *message) DeleteMessageById(ctx context.Context, id string) (int64, error) {
	idPrimitive, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	res, err := u.mgo.Collection(entities.Messages{}.MessagesCollection()).DeleteOne(ctx, bson.M{"_id": idPrimitive})
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
