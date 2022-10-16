package repository

import (
	"context"
	"dwibedis/goal_reminder_lambda/dto"
	l_dynamo "dwibedis/goal_reminder_lambda/pkg/dynamodb"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func PutItem(
	ctx context.Context,
	client *l_dynamo.Client,
	item dto.URLToShortURLItem,
	tableName string,
	) (bool, error) {
	marshalledItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return false, err
	}

	exprStr := aws.String(fmt.Sprintf("SET pk = if_not_exists(pk, :%v)", item.URL))


	strTableName := aws.String(tableName)
	putItemInput := &dynamodb.PutItemInput{
		ConditionExpression:         exprStr,
		Item:                        marshalledItem,
		TableName:                   strTableName,
	}
	_, err = client.PutItemWithContext(ctx, putItemInput)
	if err != nil {
		return false, err
	}
	return true, nil
}
