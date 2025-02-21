package config

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var MongoClient *mongo.Client

func InitMongoDB() {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", Cfg.MongoDB.User, Cfg.MongoDB.Password, Cfg.MongoDB.Host, Cfg.MongoDB.Ports)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		logrus.Errorf("InitMongoDB err:%+v", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logrus.Errorf("err: %+v", err)
		return
	}

	MongoClient = client
	logrus.Debug("MongoDB connected successfully")
}
