package operations

// snippet-start:[dynamodb.go.create_item.imports]
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws/credentials"
    "fmt"
    "os"
    "strconv"
)
// snippet-end:[dynamodb.go.create_item.imports]

// snippet-start:[dynamodb.go.create_item.struct]
// Create struct to hold info about new item
type Item struct {
    Year   int
    Title  string
    Plot   string
    Rating float64
}
// snippet-end:[dynamodb.go.create_item.struct]

func PutData() {
    // snippet-start:[dynamodb.go.create_item.session]
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
    // snippet-end:[dynamodb.go.create_item.session]

    // snippet-start:[dynamodb.go.create_item.assign_struct]
    item := Item{
        Year:   2015,
        Title:  "The Big New Movie",
        Plot:   "Nothing happens at all.",
        Rating: 0.0,
    }

    av, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
        fmt.Println("Got error marshalling new movie item:")
        fmt.Println(err.Error())
        os.Exit(1)
    }
    // snippet-end:[dynamodb.go.create_item.assign_struct]

    // snippet-start:[dynamodb.go.create_item.call]
    // Create item in table Movies
    movieName := "The Big New Movie"
    movieYear := 2015
    tableName := "Movies"

    input := &dynamodb.PutItemInput{
        Item:      av,
        TableName: aws.String(tableName),
    }

    _, err = svc.PutItem(input)
    if err != nil {
        fmt.Println("Got error calling PutItem:")
        fmt.Println(err.Error())
        os.Exit(1)
    }

    year := strconv.Itoa(movieYear)

    fmt.Println("Successfully added '" + movieName + "' (" + year + ") to table " + tableName)
    // snippet-end:[dynamodb.go.create_item.call]
}
