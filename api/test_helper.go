package api

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/zacksfF/Hotel-Reservation-Backend/storeer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	client *mongo.Client
	*storeer.Store
}

func (tdb *testdb) teardown(t *testing.T) {
	dbname := os.Getenv(storeer.MongoDBNameEnvName)
	if err := tdb.client.Database(dbname).Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	if err := godotenv.Load("../.env"); err != nil {
		t.Error(err)
	}
	dburi := os.Getenv("MONGO_DB_URL_TEST")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := storeer.NewMongoHotelStore(client)
	return &testdb{
		client: client,
		Store: &storeer.Store{
			Hotel:   hotelStore,
			User:    storeer.NewMongoUserStore(client),
			Room:    storeer.NewMongoRoomStore(client, hotelStore),
			Booking: storeer.NewMongoBookingStore(client),
		},
	}
}
