package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"server/config"
)

var es *elastic.Client

func Init() {
	client, err := elastic.NewClient(elastic.SetURL(config.Configs.EsHost))
	if err != nil {
		panic(err.Error())
	}
	_, _, err = client.Ping(config.Configs.EsHost).Do(context.Background())
	if err != nil {
		panic(err.Error())
	}
	es = client
}

func GetConn() *elastic.Client {
	return es
}
