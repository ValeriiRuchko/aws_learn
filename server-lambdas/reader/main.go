package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"reader-writer/reader/db_ops"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// TODO:
// - move functions to appropriate dirs and organize folders
// - rename 1-letter variables
// - rename upper-dir and also directory of the project here
// - choose theme of the app and what the other lambda will look like
// - think about logging and add some pretifier to look for colors

type User struct {
	ID   int
	Name string
}

func queryUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	fmt.Println("I AM A READER")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to execute query: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		users = append(users, user)
		fmt.Println(user.ID, user.Name)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during rows iteration:", err)
	}
	return users, nil
}

var ginLambda *ginadapter.GinLambdaV2

func init() {
	// db.Exec("CREATE TABLE test (id SERIAL, user TEXT NOT NULL, name TEXT NOT NULL);")
	// db.Exec("INSERT INTO test (user, name) VALUES ('Test user 1', 'Name is not important');")
	//
	// fmt.Println("Db started successfully")

	fmt.Println(os.Getenv("TursoDatabaseURL"))
	fmt.Println(os.Getenv("TursoAuthToken"))
	log.Printf("Gin cold start")

	app := gin.Default()

	// TODO: uncomment line below for release
	// gin.SetMode(gin.ReleaseMode);

	app.GET("reader/hello", func(ctx *gin.Context) {
		db := db_ops.Initialize_Connection()
		defer fmt.Println("DB connection was closed")
		defer db.Close()

		allUsers, err := queryUsers(db)
		if err != nil {
			fmt.Println("FUCK")
		}

		ctx.JSON(200, gin.H{
			"message": "Yo lets try",
			"users":   allUsers,
		})
	})
	app.GET("reader/hello/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Trying x2",
		})
	})

	ginLambda = ginadapter.NewV2(app)
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
