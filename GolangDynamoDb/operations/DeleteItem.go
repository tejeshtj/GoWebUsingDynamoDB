package operations

// snippet-start:[dynamodb.go.delete_item.imports]
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
    "fmt"
)
// snippet-end:[dynamodb.go.delete_item.imports]

func DeleteItem() {
    // snippet-start:[dynamodb.go.delete_item.session]
    // Initialize a session that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials
    // and region from the shared configuration file ~/.aws/config.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("access key", "secretkey", ""),
	})
	if(err!=nil){
		fmt.Println("there is an error getting")
	}
    // Create DynamoDB client
    svc := dynamodb.New(sess)
    // snippet-end:[dynamodb.go.delete_item.session]

    // snippet-start:[dynamodb.go.delete_item.call]
    tableName := "Movies"
    movieName := "The Big New Movie"
    movieYear := "2015"

    input := &dynamodb.DeleteItemInput{
        Key: map[string]*dynamodb.AttributeValue{
            "Year": {
                N: aws.String(movieYear),
            },
            "Title": {
                S: aws.String(movieName),
            },
        },
        TableName: aws.String(tableName),
    }

    _, error:= svc.DeleteItem(input)
    if error != nil {
        fmt.Println("Got error calling DeleteItem")
        fmt.Println(error.Error())
        return
    }

    fmt.Println("Deleted '" + movieName + "' (" + movieYear + ") from table " + tableName)
    // snippet-end:[dynamodb.go.delete_item.call]
}
