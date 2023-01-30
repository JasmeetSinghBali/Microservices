package repository

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Jasmeet-1998/Microservices/intermediate-go-microservices/graphql_mongodb_go/graph/model"
	"github.com/joho/godotenv"
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
	var donuthub model.DonutHub
	return &donuthub
}

// method - GetHubs to the Repository struct, to get all donut hubs
func (db *Repository) GetHubs() []*model.DonutHub {
	var donuthubs []*model.DonutHub
	return donuthubs
}

// method - CreateDonutHub to the Repository struct, to create donut hub
func (db *Repository) CreateDonutHub(hubInfo model.CreateDonutHubInput) *model.DonutHub {
	var newDonutHub model.DonutHub
	return &newDonutHub
}

// method - UpdateDonut to the Repository struct, to update donut hub
func (db *Repository) UpdateDonutHub(hubId string, hubInfo model.UpdateDonutHubInput) *model.DonutHub {
	var donutHub model.DonutHub
	return &donutHub
}

func (db *Repository) DeleteDonutHub(hubId string) *model.DeleteDonutHubResponse {
	return &model.DeleteDonutHubResponse{DeleteHubID: hubId}
}
