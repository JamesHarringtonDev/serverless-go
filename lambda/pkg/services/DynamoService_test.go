package services

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
)

// Dummy DynamoDB API implementation for testing.
type DummyDynamoDBAPI struct{}
type ErroneousDummyDynamoDBAPI struct{}

func (d DummyDynamoDBAPI) PutItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return &dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"key1": &types.AttributeValueMemberS{Value: "value1"},
			"key2": &types.AttributeValueMemberN{Value: "42"},
		},
	}, nil
}

func (ed ErroneousDummyDynamoDBAPI) PutItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return nil, errors.New("there's been an error, something went wrong")
}

func (d DummyDynamoDBAPI) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return &dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"key1": &types.AttributeValueMemberS{Value: "value1"},
			"key2": &types.AttributeValueMemberN{Value: "42"},
		},
	}, nil
}

func (ed ErroneousDummyDynamoDBAPI) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return nil, errors.New("there's been an error, something went wrong")
}

func (d DummyDynamoDBAPI) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	return &dynamodb.DeleteItemOutput{
		Attributes: map[string]types.AttributeValue{
			"key1": &types.AttributeValueMemberS{Value: "value1"},
			"key2": &types.AttributeValueMemberN{Value: "42"},
		},
	}, nil
}

func (ed ErroneousDummyDynamoDBAPI) DeleteItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return nil, errors.New("there's been an error, something went wrong")
}

func (d DummyDynamoDBAPI) QueryItems(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	return &dynamodb.QueryOutput{
		Items: []map[string]types.AttributeValue{
			{
				"key1": &types.AttributeValueMemberS{Value: "value1"},
				"key2": &types.AttributeValueMemberN{Value: "42"},
			},
		},
	}, nil
}

func (ed ErroneousDummyDynamoDBAPI) QueryItems(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	return nil, errors.New("there's been an error, something went wrong")
}

func TestPutItem(t *testing.T) {
	api := DummyDynamoDBAPI{}
	ctx := context.Background()
	table := "test_table"
	key := map[string]types.AttributeValue{
		"key1": &types.AttributeValueMemberS{Value: "value1"},
		"key2": &types.AttributeValueMemberN{Value: "42"},
	}

	item, err := PutItem(ctx, api, table, key)

	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, key, item)
}

func TestPutItemReturnsError(t *testing.T) {
	api := ErroneousDummyDynamoDBAPI{}
	ctx := context.Background()
	table := "test_table"
	key := map[string]types.AttributeValue{
		"key1": &types.AttributeValueMemberS{Value: "value1"},
		"key2": &types.AttributeValueMemberN{Value: "42"},
	}

	item, err := PutItem(ctx, api, table, key)

	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestGetItem(t *testing.T) {
	api := DummyDynamoDBAPI{}
	ctx := context.Background()
	table := "test_table"
	key := map[string]types.AttributeValue{
		"key1": &types.AttributeValueMemberS{Value: "value1"},
		"key2": &types.AttributeValueMemberN{Value: "42"},
	}

	item, err := GetItem(ctx, api, table, key)

	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, key, item)
}

func TestDeleteItem(t *testing.T) {
	api := DummyDynamoDBAPI{}
	ctx := context.Background()
	table := "test_table"
	key := map[string]types.AttributeValue{
		"key1": &types.AttributeValueMemberS{Value: "value1"},
		"key2": &types.AttributeValueMemberN{Value: "42"},
	}

	attributes, err := DeleteItem(ctx, api, table, key)

	assert.NoError(t, err)
	assert.NotNil(t, attributes)
	assert.Equal(t, key, attributes)
}

func TestQueryItems(t *testing.T) {
	api := DummyDynamoDBAPI{}
	ctx := context.Background()
	table := "test_table"
	key := map[string]types.AttributeValue{
		"key1": &types.AttributeValueMemberS{Value: "value1"},
		"key2": &types.AttributeValueMemberN{Value: "42"},
	}

	items, err := QueryItems(ctx, api, table, key)

	assert.NoError(t, err)
	assert.NotNil(t, items)
	assert.Len(t, items, 1)
	assert.Equal(t, key, items[0])
}
