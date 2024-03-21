package database

import (
	"context"
	"log"

	"github.com/era-n/nic-test/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct {
	Client *mongo.Client
	ctx    context.Context
}

type Car struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Make  string             `bson:"make" json:"make"`
	Model string             `bson:"model" json:"model"`
	Year  string             `bson:"year" json:"year"`
	Image string             `bson:"image" json:"image"`
}

func NewDB() *Db {
	return &Db{}
}

func (d *Db) Disconnect() error {
	err := d.Client.Disconnect(d.ctx)
	return err
}

func (d *Db) InitMongo() error {
	cfg, err := config.LoadConfig()

	client, err := mongo.Connect(d.ctx, options.Client().ApplyURI(cfg.MongoUrl))

	d.Client = client

	if err != nil {
		panic(err)
	}

	err = d.Client.Ping(d.ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (d *Db) GetCars() []Car {
	collection := d.Client.Database("nic-test").Collection("cars")
	cursor, err := collection.Find(d.ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(d.ctx)

	cars := make([]Car, 0)
	for cursor.Next(d.ctx) {
		car := Car{}
		err := cursor.Decode(&car)
		cars = append(cars, car)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return cars
}
