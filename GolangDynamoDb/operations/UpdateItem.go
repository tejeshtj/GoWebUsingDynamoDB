package operations

// snippet-start:[dynamodb.go.update_item.imports]
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/credentials"
    "fmt"
)
// snippet-end:[dynamodb.go.update_item.imports]

func UpdateItem() {
    // snippet-start:[dynamodb.go.update_item.session]
    // Initialize a session that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials
    // and region from the shared configuration file ~/.aws/config.
    sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("access key", "secretkey", ""),	})
	if(err!=nil){
		fmt.Println("there is an error getting")
	}

    // Create DynamoDB client
    svc := dynamodb.New(sess)
    // snippet-end:[dynamodb.go.update_item.session]

    // snippet-start:[dynamodb.go.update_item.call]
    // Create item in table Movies
    tableName := "Movies"
    movieName := "The Big New Movie"
    movieYear := "2015"
    movieRating := "1.3"

    input := &dynamodb.UpdateItemInput{
        ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
            ":r": {
                N: aws.String(movieRating),
            },
        },
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Year": {
                N: aws.String(movieYear),
            },
            "Title": {
                S: aws.String(movieName),
            },
        },
        ReturnValues:     aws.String("UPDATED_NEW"),
        UpdateExpression: aws.String("set Rating = :r"),
    }

    _, error := svc.UpdateItem(input)
    if error != nil {
        fmt.Println(error.Error())
        return
    }

    fmt.Println("Successfully updated '" + movieName + "' (" + movieYear + ") rating to " + movieRating)
    // snippet-end:[dynamodb.go.update_item.call]
}
