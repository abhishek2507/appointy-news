package main

//Importing libraries

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Declaring client globally
var client *mongo.Client

//DB Model for every article
type Post struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"title,omitempty"`
	SubTitle string             `json:"subtitle,omitempty" bson:"subtitle,omitempty"`
	Content  string             `json:"content,omitempty" bson:"content,omitempty"`
}

//Function to create an Article
func CreateArticleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var post Post
	_ = json.NewDecoder(request.Body).Decode(&post)
	collection := client.Database("appointy-news").Collection("appointy-news")

	//5 seconds context to timeout
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	//Inserting the details in the MongoDB
	result, _ := collection.InsertOne(ctx, post)

	//Returning the last entered Article as the response
	json.NewEncoder(response).Encode(result)
}

//Function to Fetch all the articles
func GetArticlesEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var posts []Post
	collection := client.Database("appointy-news").Collection("appointy-news")

	//30 Seconds context to tiemout
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	//Cursor is similar to the one we use in JAVA And Oracle DB
	cursor, err := collection.Find(ctx, bson.M{})

	//Check for errors
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	//Closing the cursor
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}

	//Checking for errors again
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	//Encodes and sends all the Articles as a JSON Response
	json.NewEncoder(response).Encode(posts)
}

//Function to fetch a particular article through ID
func GetArticleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var post Post
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Post{ID: id}).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(post)
}

//Main function
func main() {

	//Printing the initial startup in the terminal
	fmt.Println("Starting the application...")

	//Connecting to MongoDB atlas through URI and checking for errors
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://abhiscluster:Abhi2507@abhiscluster.tyxhs.mongodb.net/appointy-news?retryWrites=true&w=majority"))

	//If any error in DB connection log it into the terminal
	if err != nil {
		log.Fatal(err)
	}

	//10 seconds context to tiemout
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	//Initialising the router from gorilla mux
	router := mux.NewRouter()

	//Post request for creating an article using MUX
	router.HandleFunc("/articles", CreateArticleEndpoint).Methods("POST")
	//GET Request to get all the articles using MUX
	router.HandleFunc("/articles", GetArticlesEndpoint).Methods("GET")
	//GET Request to get a queried article using its ID using MUX
	router.HandleFunc("/articles/{id}", GetArticleEndpoint).Methods("GET")

	//Lisetning to port 12345 (Local Host)
	http.ListenAndServe(":12345", router)
}
