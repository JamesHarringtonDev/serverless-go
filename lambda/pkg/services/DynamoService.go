package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)
import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DynamoService struct {
}

type DynamoPutItemAPI interface {
	PutItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

func PutItem(ctx context.Context, api DynamoPutItemAPI, table string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	item, err := api.PutItem(ctx, &dynamodb.GetItemInput{
		TableName: &table,
		Key:       key,
	})
	if err != nil {
		return nil, err
	}
	return item.Item, nil
}

type DynamoGetItemAPI interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

func GetItem(ctx context.Context, api DynamoGetItemAPI, table string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	item, err := api.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &table,
		Key:       key,
	})
	if err != nil {
		return nil, err
	}
	return item.Item, nil
}

type DynamoDeleteItemApi interface {
	DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
}

func DeleteItem(ctx context.Context, api DynamoDeleteItemApi, table string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	output, err := api.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &table,
		Key:       key,
	})
	if err != nil {
		return nil, err
	}
	return output.Attributes, nil
}

type DynamoQueryItemsApi interface {
	QueryItems(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

func QueryItems(ctx context.Context, api DynamoQueryItemsApi, table string, key map[string]types.AttributeValue) ([]map[string]types.AttributeValue, error) {
	output, err := api.QueryItems(ctx, &dynamodb.QueryInput{
		TableName: &table,
	})
	if err != nil {
		return nil, err
	}
	return output.Items, nil
}
