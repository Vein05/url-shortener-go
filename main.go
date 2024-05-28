package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/big"
	"net/http"
	"os"
)

type RequestBody struct {
	URL      string `json:"url"`
	SHORTURL string `json:"shorturl"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	MONGOURL := os.Getenv("MONGOURL")
	opts := options.Client().ApplyURI(MONGOURL).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	database := client.Database("AbodeDB")
	if err := client.Database("AbodeDB").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged AbodeDB. You successfully connected to MongoDB!")
	g := gin.Default()

	g.POST("/add-url", func(c *gin.Context) {
		var body RequestBody
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newURL := getNewURL(database, body.URL)
		c.JSON(http.StatusOK, gin.H{
			"newURL": newURL,
		})
	})
	g.Static("/", "./static/")
	err = g.Run()
	if err != nil {
		return
	}

}

func getNewURL(database *mongo.Database, url string) string {
	collection := database.Collection("url-list")
	status := checkUrl(collection, url)
	if status == false {
		newUrl := getRandom()
		doc := RequestBody{URL: url, SHORTURL: newUrl}
		res, err := collection.InsertOne(context.TODO(), doc)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("The new SHORTURL is: " + newUrl + " And is updated with ID: ")
		log.Println(res.InsertedID)
		return newUrl

	} else {
		var result RequestBody
		filter := bson.M{"url": url}
		err := collection.FindOne(context.TODO(), filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		} else {
			return result.SHORTURL
		}

	}
	return url
}

func checkUrl(collection *mongo.Collection, url string) bool {
	var result RequestBody
	filter := bson.M{"url": url}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func getRandom() string {
	chars := "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	url := make([]byte, 4)
	for i := range url {
		randomChar, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			panic(err)
		}
		url[i] = chars[randomChar.Int64()]
	}

	return string(url)
}
