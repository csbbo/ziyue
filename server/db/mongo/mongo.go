package mongo

import (
	"context"
	"github.com/juju/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"server/config"
	"server/loger"
	"time"
)

type Mongo struct {
	conn *mongo.Client
}

func Default() (*Mongo, error) {
	mg := &Mongo{}
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(config.Configs.MongoConnTimeout)*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Configs.MongoHost))
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		return nil, err
	}
	ctx, _ = context.WithTimeout(context.Background(), time.Duration(config.Configs.MongoConnTimeout)*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		loger.Loger.Error(errors.ErrorStack(errors.Trace(err)))
		return nil, err
	}
	mg.conn = client
	return mg, nil
}

func (this *Mongo) GetConn() *mongo.Client {
	return this.conn
}
