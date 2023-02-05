package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/graphql_mongodb_go/graph/model"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Repository struct {
	client *mongo.Client
}

func Connect() *Repository {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MAIN_DB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	// a 30 sec timeout context, to allow connection for mongodb to get establish
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &Repository{
		client: client,
	}

}

// method- GetHub to the Repository struct , to get a donut hub by id
func (db *Repository) GetHub(id string) *model.DonutHub {

	hubCollection := db.client.Database("mockgqlgo").Collection("hubs")
	// context to handle request-scoped data ,cancel signals among goroutines
	// ensures that db communication are not just blocked has max context session of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// cancel signals gets called just before this function returns
	defer cancel()

	// grab and store id passedin request as ObjectId
	_id, _ := primitive.ObjectIDFromHex(id)
	// create filter option as the bson mongoDB document by _id
	filter := bson.M{"_id": _id}
	var donuthub model.DonutHub
	// perform findOne on hubColeection wrapped under 30 seconds ctx with filter and decode the recieved document into golang supported data type(struct) and store in donuthub
	err := hubCollection.FindOne(ctx, filter).Decode(&donuthub)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(donuthub)
	return &donuthub
}

// method - GetHubs to the Repository struct, to get all donut hubs
func (db *Repository) GetHubs() []*model.DonutHub {
	hubCollection := db.client.Database("mockgqlgo").Collection("hubs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var donuthubs []*model.DonutHub
	cursor, err := hubCollection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cursor)
	// All iterates the cursor and decodes each document into results parameter i.e donuthubs , note- the results parameter must be pointer to a slice
	if err = cursor.All(context.TODO(), &donuthubs); err != nil {
		panic(err)
	}
	return donuthubs
}

// method - CreateDonutHub to the Repository struct, to create donut hub & returns the newly created donut hub
func (db *Repository) CreateDonutHub(hubInfo model.CreateDonutHubInput) *model.DonutHub {
	hubCollection := db.client.Database("mockgqlgo").Collection("hubs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	inserted, err := hubCollection.InsertOne(ctx, bson.M{
		"title":             hubInfo.Title,
		"description":       hubInfo.Description,
		"shop_location":     hubInfo.ShopLocation,
		"bestselling":       hubInfo.Bestselling,
		"price_bestselling": hubInfo.PriceBestselling,
	})
	if err != nil {
		log.Fatal(err)
	}

	insertedID := inserted.InsertedID.(primitive.ObjectID).Hex()
	log.Println(insertedID)
	newDonutHub := model.DonutHub{
		ID:               insertedID,
		Title:            hubInfo.Title,
		Description:      hubInfo.Description,
		ShopLocation:     hubInfo.ShopLocation,
		Bestselling:      hubInfo.Bestselling,
		PriceBestselling: hubInfo.PriceBestselling,
	}
	return &newDonutHub
}

// method - UpdateDonut to the Repository struct, to update donut hub
func (db *Repository) UpdateDonutHub(hubId string, hubInfo model.UpdateDonutHubInput) *model.DonutHub {
	hubCollection := db.client.Database("mockgqlgo").Collection("hubs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	updateDonutHub := bson.M{}

	// check for nil existance, as input createHubInput allows nullish values reff: schema.graphqls
	if hubInfo.Bestselling != nil {
		updateDonutHub["bestselling"] = hubInfo.Bestselling
	}
	if hubInfo.PriceBestselling != nil {
		updateDonutHub["price_bestselling"] = hubInfo.PriceBestselling
	}

	_id, _ := primitive.ObjectIDFromHex(hubId)
	filter := bson.M{"_id": _id}
	update := bson.M{"$set": updateDonutHub}

	results := hubCollection.FindOneAndUpdate(ctx, filter, update, options.FindOneAndUpdate().SetReturnDocument(1))

	var donutHub model.DonutHub

	if err := results.Decode(&donutHub); err != nil {
		log.Fatal(err)
	}
	return &donutHub
}

func (db *Repository) DeleteDonutHub(hubId string) *model.DeleteDonutHubResponse {

	hubCollection := db.client.Database("mockgqlgo").Collection("hubs")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(hubId)
	filter := bson.M{"_id": _id}

	_, err := hubCollection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	return &model.DeleteDonutHubResponse{DeleteHubID: hubId}
}
