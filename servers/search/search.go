package search

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/sundogrd/search-grpc/grpc_gen/search"
	"strconv"
	"strings"
)

func (server SearchServiceServer) Search(ctx context.Context, req *search.SearchRequest) (*search.SearchResponse, error) {
	searchRes, err := server.ElasticsearchClient.Search(
		server.ElasticsearchClient.Search.WithIndex(req.AppId+"_"+req.IndexName),
		server.ElasticsearchClient.Search.WithBody(strings.NewReader(req.QueryJson)),
		server.ElasticsearchClient.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		logrus.Errorf("[search-grpc/servers/search] Search err: %s", err.Error())
		return nil, err
	}
	defer searchRes.Body.Close()

	if searchRes.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(searchRes.Body).Decode(&e); err != nil {
			return nil, err
		} else {
			// Print the response status and error information.
			return nil, errors.New(fmt.Sprintf(
				"type: %s, reason: %s",
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			))
		}
	}

	type SearchResBody struct {
		Hits map[string]interface{} `json:"hits"`
	}
	var r SearchResBody
	if err := json.NewDecoder(searchRes.Body).Decode(&r); err != nil {
		return nil, err
	}
	logrus.Infof("%s", r)

	var total int64
	var maxScore float64
	var hits []*search.SearchResult
	if r.Hits["total"] != nil {
		total = int64(r.Hits["total"].(map[string]interface{})["value"].(float64))
	}
	if r.Hits["max_score"] != nil {
		maxScore = r.Hits["max_score"].(float64)
	}
	if r.Hits["hits"] != nil {
		for _, hit := range r.Hits["hits"].([]interface{}) {
			idNum, _ := strconv.ParseInt(hit.(map[string]interface{})["_id"].(string), 10, 64)
			sourceBytes, _ := json.Marshal(hit.(map[string]interface{})["_source"].(map[string]interface{}))
			hits = append(hits, &search.SearchResult{
				Index: hit.(map[string]interface{})["_index"].(string),
				Type: hit.(map[string]interface{})["_type"].(string),
				Id: idNum,
				Score: hit.(map[string]interface{})["_score"].(float64),
				SourceJson: string(sourceBytes),
			})
		}
	}
	return &search.SearchResponse{
		Total: total,
		MaxScore: maxScore,
		Hits: hits,
	}, nil
}
