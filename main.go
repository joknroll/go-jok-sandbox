package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
	
    //"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/gin-gonic/gin"
)

// You will be using this Trainer type later in the program
type Trainer struct {
    Name string
    Age  int
    City string
}

func main() {
// Set client options
clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

// Connect to MongoDB
client, err := mongo.Connect(context.TODO(), clientOptions)

if err != nil {
    log.Fatal(err)
}

// Check the connection
err = client.Ping(context.TODO(), nil)

if err != nil {
    log.Fatal(err)
}

fmt.Println("Connected to MongoDB!")

r := gin.Default()
r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
})

r.LoadHTMLGlob("templates/*")
//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
r.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "Main website",
    })
})

r.Run(":5000") // listen and serve on 0.0.0.0:8080

}