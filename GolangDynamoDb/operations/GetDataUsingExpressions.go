package operations

// snippet-start:[dynamodb.go.scan_items.imports]
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/aws/aws-sdk-go/aws/credentials"
    "fmt"
    "os"
)
// snippet-end:[dynamodb.go.scan_items.imports]

// snippet-start:[dynamodb.go.scan_items.struct]
// Create struct to hold info about new item

// snippet-end:[dynamodb.go.scan_items.struct]

// Get the movies with a minimum rating of 8.0 in 2011
func GetDataUsingExp() {
    // snippet-start:[dynamodb.go.scan_items.session]
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
    // snippet-end:[dynamodb.go.scan_items.session]

    // snippet-start:[dynamodb.go.scan_items.vars]
    tableName := "Movies"
    minRating := 4.0
    year := 2013
    // snippet-end:[dynamodb.go.scan_items.vars]

    // snippet-start:[dynamodb.go.scan_items.expr]
    // Create the Expression to fill the input struct with.
    // Get all movies in that year; we'll pull out those with a higher rating later
    filt := expression.Name("Year").Equal(expression.Value(year))

    // Or we could get by ratings and pull out those with the right year later
    //    filt := expression.Name("info.rating").GreaterThan(expression.Value(min_rating))

    // Get back the title, year, and rating
    proj := expression.NamesList(expression.Name("Title"), expression.Name("Year"), expression.Name("Rating"))

    expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
    if err != nil {
        fmt.Println("Got error building expression:")
        fmt.Println(err.Error())
        os.Exit(1)
    }
    // snippet-end:[dynamodb.go.scan_items.expr]

    // snippet-start:[dynamodb.go.scan_items.call]
    // Build the query input parameters
    params := &dynamodb.ScanInput{
        ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
        FilterExpression:          expr.Filter(),
        ProjectionExpression:      expr.Projection(),
        TableName:                 aws.String(tableName),
    }

    // Make the DynamoDB Query API call
    result, err := svc.Scan(params)
    if err != nil {
        fmt.Println("Query API call failed:")
        fmt.Println((err.Error()))
        os.Exit(1)
    }
    // snippet-end:[dynamodb.go.scan_items.call]

    // snippet-start:[dynamodb.go.scan_items.process]
    numItems := 0

    for _, i := range result.Items {
        item := Item{}

        err = dynamodbattribute.UnmarshalMap(i, &item)

        if err != nil {
            fmt.Println("Got error unmarshalling:")
            fmt.Println(err.Error())
            os.Exit(1)
        }

        // Which ones had a higher rating than minimum?
        if item.Rating > minRating {
            // Or it we had filtered by rating previously:
            //   if item.Year == year {
            numItems++

            fmt.Println("Title: ", item.Title)
            fmt.Println("Rating:", item.Rating)
            fmt.Println()
        }
    }

    fmt.Println("Found", numItems, "movie(s) with a rating above", minRating, "in", year)
    // snippet-end:[dynamodb.go.scan_items.process]
}
