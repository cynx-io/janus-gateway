package elastic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cynxees/janus-gateway/internal/dependencies/logger"
	"github.com/elastic/go-elasticsearch"
	"io"
)

type Client struct {
	Client *elasticsearch.Client
}

func NewClient(cfg elasticsearch.Config) (*Client, error) {
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Client{Client: client}, nil
}

func (c *Client) LogToElasticsearch(entry LogEntry) error {
	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	logger.Debug("Logging to Elasticsearch: ", string(data))
	res, err := c.Client.Index(
		"log-janus-gateway",
		bytes.NewReader(data),
		c.Client.Index.WithDocumentType("_doc"),
	)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error("Error closing response body: ", err)
		}
	}(res.Body)

	if res.IsError() {
		return fmt.Errorf("elasticsearch index error: %s", res.String())
	}
	return nil
}
