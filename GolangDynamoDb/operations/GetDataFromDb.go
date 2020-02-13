package operations

// snippet-start:[dynamodb.go.read_item.imports]
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws/credentials"
    "fmt"
)
// snippet-end:[dynamodb.go.read_item.imports]

// snippet-start:[dynamodb.go.read_item.struct]
// Create struct to hold info about new item

// snippet-end:[dynamodb.go.read_item.struct]

func GetDataFromDb() {
    // snippet-start:[dynamodb.go.read_item.session]
    // Initialize a session that the SDK will use to load
    // credentials from the shared credentials file ~/.aws/credentials
    // and region from the shared configuration file ~/.aws/config.
    sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials("AKIARNVQA3QVTS6ORTOC", "L1dCDoD8jkApCYiOeVnUhU7I82VXpuP4YZSAlsMR", ""),
	})
	if(err!=nil){
		fmt.Println("there is an error getting")
	}

    // Create DynamoDB client
    svc := dynamodb.New(sess)
    // snippet-end:[dynamodb.go.read_item.session]

    // snippet-start:[dynamodb.go.read_item.call]
    tableName := "Movies"
    movieName := "The Big New Movie"
    movieYear := "2015"

    result, err := svc.GetItem(&dynamodb.GetItemInput{
        TableName: aws.String(tableName),
        Key: map[string]*dynamodb.AttributeValue{
            "Year": {
                N: aws.String(movieYear),
            },
            "Title": {
                S: aws.String(movieName),
            },
        },
    })
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    // snippet-end:[dynamodb.go.read_item.call]

    // snippet-start:[dynamodb.go.read_item.unmarshall]
    item := Item{}

    err = dynamodbattribute.UnmarshalMap(result.Item, &item)
    if err != nil {
        panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
    }

    if item.Title == "" {
        fmt.Println("Could not find '" + movieName + "' (" + movieYear + ")")
        return
    }

    fmt.Println("Found item:")
    fmt.Println("Year:  ", item.Year)
    fmt.Println("Title: ", item.Title)
    fmt.Println("Plot:  ", item.Plot)
    fmt.Println("Rating:", item.Rating)
    // snippet-end:[dynamodb.go.read_item.unmarshall]
}