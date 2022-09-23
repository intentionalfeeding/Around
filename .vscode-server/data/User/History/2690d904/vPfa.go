package backend

import (
	"context"
    "fmt"

    "around/constants"

	"github.com/olivere/elastic/v7"
)

var (
	ESBackend *ElasticsearchBackend
)

type ElasticsearchBackend struct {
	client *elastic.Client
}

func InitElasticsearchBackend() {
	client, err := elastic.NewClient(
		elastic.SetURL(constants.ES_URL),
		elastic.SetBasicAuth(constants.ES_USERNAME,constants.ES_PASSWORD))
		if err != nil {
			panic(err)
		}

		if !exists{
			mapping := `{
				"mappings":{
					"id":  { "type": "keyword" },
					"user": { "type": "keyword" },
					"message": { "type": "text" },
					"
				}
			}`
		}
}
