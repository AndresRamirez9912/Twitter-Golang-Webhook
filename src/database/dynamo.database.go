package database

import (
	"context"
	"log"
	"twitter-webhook/src/constants"
	"twitter-webhook/src/models"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

func CreateTableDynamodb() error {
	dynamoClient, err := createDynamoCLient()
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Know if a table exist
	exist, err := tableExist(dynamoClient)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if exist {
		log.Println("Table Already Exist")
		return nil
	}
	_, err = dynamoClient.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(constants.TABLE_NAME),
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Println("Table Created Succesffuly")
	return nil
}

func GetItemById(id string) (*dynamodb.GetItemOutput, error) {
	dynamoClient, err := createDynamoCLient()
	if err != nil {

	}
	items, err := dynamoClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(constants.TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return items, nil
}

func CreateItem(dataField models.TwitterField) error {
	dynamoClient, err := createDynamoCLient()
	if err != nil {
		log.Fatal(err)
		return err
	}

	dataValue, err := attributevalue.MarshalMap(dataField)
	if err != nil {
		log.Fatal(err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      dataValue,
		TableName: aws.String(constants.TABLE_NAME),
	}

	_, err = dynamoClient.PutItem(context.TODO(), input)

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func ChangeStatus(id string, status bool) error {
	dynamoClient, err := createDynamoCLient()
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = dynamoClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(constants.TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("set Active = :active"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":active": &types.AttributeValueMemberBOOL{Value: status},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func UpdateLastMessageId(id string, messageId string) error {
	dynamoClient, err := createDynamoCLient()
	if err != nil {
		log.Fatal(err)
		return err
	}

	_, err = dynamoClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(constants.TABLE_NAME),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("set LastMessageId = :lastMessage"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":lastMessage": &types.AttributeValueMemberS{Value: messageId},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func createDynamoCLient() (*dynamodb.Client, error) {
	// Set AWS
	config, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
		opts.Region = constants.AWS_REGION
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// Create the dynamodb Client
	dynamoClient := dynamodb.NewFromConfig(config)

	return dynamoClient, nil
}

func tableExist(dynamoClient *dynamodb.Client) (bool, error) {
	tables, err := dynamoClient.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatal("ListTables failed", err)
		return false, err
	}
	for _, name := range tables.TableNames {
		if name == constants.TABLE_NAME {
			return true, nil
		}
	}

	return false, nil
}
