package mongo

import (
	"context"
	"fmt"
	"os"
	"spo_task_3/model"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

func ConnectToDB() (*mongo.Client, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	MongoUrl := os.Getenv("MONGO_DB_URL")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(MongoUrl).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client, nil

}

func OpenCollection(client *mongo.Client, collectionName string, databaseName string) *mongo.Collection {
	collection := client.Database(databaseName).Collection(collectionName)
	return collection
}

func InsertUserIntoDB(user model.User) error {
	client, _ := ConnectToDB()
	collection := OpenCollection(client, "user", "SPO_TASK")
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func IsUserExist(email string) (model.User, bool) {
	var user model.User
	client, _ := ConnectToDB()
	collection := OpenCollection(client, "user", "SPO_TASK")
	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No documents found")

		} else {
			panic(err)

		}
		return user, false
	}
	return user, true
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
