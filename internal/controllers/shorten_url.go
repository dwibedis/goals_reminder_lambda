package controllers

import (
	"context"
	"dwibedis/goal_reminder_lambda/dto"
	"dwibedis/goal_reminder_lambda/internal/repository"
	"dwibedis/goal_reminder_lambda/internal/services"
	"dwibedis/goal_reminder_lambda/pkg/dynamodb"
	"fmt"
	"time"
)

const tableName = "url-shortner"
const retryMax = 3

type ShortenURLRequest struct {
	URL string
	TTL time.Duration
}

type ShortenURLResponse struct {
	URL string
	ShortenedURL string
	TTL time.Duration
}

func HandleURLShortenRequest(ctx context.Context, req ShortenURLRequest) (*ShortenURLResponse, error) {
	client := dynamodb.NewDynamoDBClient(ctx)
	shortUrl := services.GenerateShortUrl(req.URL)
	retryCnt := 0
	var err error
	var status bool
	for retryCnt = 0; retryCnt < retryMax; retryCnt++ {
		item := dto.URLToShortURLItem{
			URL:        req.URL,
			URLSK:      req.URL,
			ShortURL:   shortUrl,
			ShortURLSk: shortUrl,
			TTL:        req.TTL,
		}
		status, err = repository.PutItem(ctx, client, item, tableName)
		if err != nil || !status {
			fmt.Printf("retry count: %v, error occurred, %v \n", retryCnt, err)
			continue
		}
		break
	}

	if err != nil || !status {
		fmt.Printf("retry count: %v, error occurred, %v \n", retryCnt, err)
		return nil, err
	}
	return &ShortenURLResponse{
		URL:          req.URL,
		ShortenedURL: shortUrl,
		TTL:          req.TTL,
	}, nil
}

func validateRequest(req ShortenURLRequest) bool {
	return false
}