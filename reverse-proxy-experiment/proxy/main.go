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

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type User struct {
	ID   int
	Name string
}

func queryUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
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
	url := os.Getenv("TursoDatabaseURL") + "?authToken=" + os.Getenv("TursoAuthToken")

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	}
	// defer db.Close()

	// db.Exec("CREATE TABLE test (id SERIAL, user TEXT NOT NULL, name TEXT NOT NULL);")
	// db.Exec("INSERT INTO test (user, name) VALUES ('Test user 1', 'Name is not important');")
	//
	// fmt.Println("Db started successfully")

	log.Printf("Gin cold start")

	r := gin.Default()

	// TODO: uncomment line below for release
	// gin.SetMode(gin.ReleaseMode);

	r.GET("/hello", func(c *gin.Context) {
		all_users, err := queryUsers(db)
		if err != nil {
			fmt.Println("FUCK")
		}

		c.JSON(200, gin.H{
			"message": "Yo lets try",
			"users":   all_users,
		})
	})
	r.GET("/hello/me", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Trying x2",
		})
	})

	ginLambda = ginadapter.NewV2(r)
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
