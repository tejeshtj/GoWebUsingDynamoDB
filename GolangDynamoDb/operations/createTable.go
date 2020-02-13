package operations


import (
    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "fmt"
    "os"
)
func CreateTable(){
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("AKIARNVQA3QVTS6ORTOC", "L1dCDoD8jkApCYiOeVnUhU7I82VXpuP4YZSAlsMR", ""),
	})
	if(err!=nil){
		fmt.Println("there is an error getting")
	}
	svc := dynamodb.New(sess)

	// Create table Movies
tableName := "Movies"

input := &dynamodb.CreateTableInput{
    AttributeDefinitions: []*dynamodb.AttributeDefinition{
        {
            AttributeName: aws.String("Year"),
            AttributeType: aws.String("N"),
        },
        {
            AttributeName: aws.String("Title"),
            AttributeType: aws.String("S"),
        },
    },
    KeySchema: []*dynamodb.KeySchemaElement{
        {
            AttributeName: aws.String("Year"),
            KeyType:       aws.String("HASH"),
        },
        {
            AttributeName: aws.String("Title"),
            KeyType:       aws.String("RANGE"),
        },
    },
    ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
        ReadCapacityUnits:  aws.Int64(10),
        WriteCapacityUnits: aws.Int64(10),
    },
    TableName: aws.String(tableName),
}

_, error := svc.CreateTable(input)
if error != nil {
    fmt.Println("Got error calling CreateTable:")
    fmt.Println(error.Error())
    os.Exit(1)
}

fmt.Println("Created the table", tableName)

}