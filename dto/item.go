package dto

import "time"

type URLToShortURLItem struct {
	URL string `dynamodbav:"pk"`
	URLSK string `dynamodbav:"sk"`
	ShortURL string `dynamodbav:"gs1pk"`
	ShortURLSk string `dynamodbav:"gs1sk"`
	TTL time.Duration `dynamodbav:"ttl"`
}