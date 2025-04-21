package datasource 

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"fmt"
	"log"
	"time"
)

type (
	User struct {
		ID           string    `json:"id" db:"id"`                     // User ID (UUID, Partition Key)
		Username     string    `json:"username" db:"username"`         // Username
		Email        string    `json:"email" db:"email"`           // Email address
		PasswordHash string    `json:"passwordHash" db:"password_hash"` // Hashed password
		FirstName    string    `json:"firstName" db:"first_name"`       // First name
		LastName     string    `json:"lastName" db:"last_name"`         // Last name
		CreatedAt    time.Time `json:"createdAt" db:"created_at"`         // ISO 8601 timestamp
		UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`         // ISO 8601 timestamp
		TenantID     string    `json:"tenantId" db:"tenant_id"`           // Tenant ID
	}

	Db interface {
		CreateUser(u User) (User, error)
		GetUser(userID string) (User, error)
		UpdateUser(u User) (User, error)
		DeleteUser(u User) error
	}

	dbService struct {
		dynamo *dynamodb.Session
		usersTable string
	}
)

func (db *dbService) CreateUser(u User) (User, error) {
	u.CreatedAt = time.Now()

	item, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		log.Fatalf("Got error marshalling new item: %s", err)
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(db.usersTable),
	}

	_, err = db.dbService.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return nil, err
	}


	fmt.Printf("Successfully added on table: %s", db.usersTable)
	return u, nil 
}

func (db *dbService) GetUser(userID string) (User, error) {
	u.CreatedAt = time.Now()

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(userID),
			},
		},
		TableName: aws.String(db.usersTable),
	}

	_, err = db.dbService.GetItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
		return nil, err
	}


	fmt.Printf("Successfully get user: %s", u.Username)
	return u, nil 
}


func (db *dbService) UpdateUser(u User) (User, error) {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"UserName": {
				N: aws.String(u.Username),
			},
			"Email": {
				N: aws.String(u.Email),
			},
		},
		TableName: aws.String(db.usersTable),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(u.ID),
			}
		},
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
		return err
	}

	fmt.Println("User successfully updated")
	return u
}

func (db *dbService) DeleteUser(u User) error {
	item, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		log.Fatalf("Got error marshalling new item: %s", err)
		return err
	}

	input := &dynamodb.DeleteItemInput{
		Item:      item,
		TableName: aws.String(db.usersTable),
	}

	_, err = db.dbService.DeleteItem(input)
	if err != nil {
		log.Fatalf("Got error calling DeleteItem: %s", err)
		return err
	}

	return nil 
}


func NewDatabase(tableName string) *dbService {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	dbs := dbService{
		dynamo: svc,
		usersTable: tableName,
	}

	return dbs
}
