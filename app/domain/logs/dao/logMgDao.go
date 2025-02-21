package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"mangoTV/app/domain/logs/models"
	"time"
)

type LogsMgDao struct {
	con        *mongo.Client
	collection *mongo.Collection
}

func NewLogsMgDao(mongoClient *mongo.Client) *LogsMgDao {
	return &LogsMgDao{
		con:        mongoClient,
		collection: mongoClient.Database("dev").Collection("logs"),
	}
}

func (dao *LogsMgDao) InsertLog(logs models.LogsMg) error {
	logs.CreatedAt = time.Now()
	_, err := dao.collection.InsertOne(context.TODO(), logs)
	return err
}

func (dao *LogsMgDao) FindLogs(filter interface{}) ([]models.LogsMg, error) {
	var logs []models.LogsMg
	cursor, err := dao.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var log models.LogsMg
		err = cursor.Decode(&log)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return logs, nil
}
