package operations

// snippet-start:[dynamodb.go.load_items.imports]
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/aws/credentials"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)
// snippet-end:[dynamodb.go.load_items.imports]

// snippet-start:[dynamodb.go.load_items.struct]
// Create struct to hold info about new item

// snippet-end:[dynamodb.go.load_items.struct]

// snippet-start:[dynamodb.go.load_items.func]
// Get table items from JSON file
func getItems() []Item {
    raw, err := ioutil.ReadFile("./operations/movie_data.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var items []Item
    json.Unmarshal(raw, &items)
    return items
}
// snippet-end:[dynamodb.go.load_items.func]

func PutItemsFromJsonFile() {
    // snippet-start:[dynamodb.go.load_items.session]
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
    // snippet-end:[dynamodb.go.load_items.session]

    // snippet-start:[dynamodb.go.load_items.call]
    // Get table items from movie_data.json
    items := getItems()

    // Add each item to Movies table:
    tableName := "Movies"

    for _, item := range items {
        av, err := dynamodbattribute.MarshalMap(item)
        if err != nil {
            fmt.Println("Got error marshalling map:")
            fmt.Println(err.Error())
            os.Exit(1)
        }

        // Create item in table Movies
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

        year := strconv.Itoa(item.Year)

        fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)
        // snippet-end:[dynamodb.go.load_items.call]
    }
}