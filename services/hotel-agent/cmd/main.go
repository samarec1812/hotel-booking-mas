package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samarec1812/hotel-booking-mas/services/hotel-agent/models"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
)

var ctx context.Context
var err error
var client *mongo.Client
var db *mongo.Database
var (
	RoomCollection    = "rooms"
	BookingCollection = "bookings"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func (cfg Config) SetConfig() string {
	MONGO_URI := fmt.Sprintf("mongodb://%s:%s@%s:%s/",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	return MONGO_URI
}

func init() {
	ctx = context.Background()
	if err := initConfig(); err != nil {
		log.Fatal(err)
	}
	mongoURI := Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
	}.SetConfig()
	fmt.Println(mongoURI)
	client, err = mongo.Connect(ctx,
		options.Client().ApplyURI(mongoURI))

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	db = client.Database(viper.GetString("db.dbname"))
	log.Println("Connected to MongoDB")
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func NewRoomHandler(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	room.ID = primitive.NewObjectID()
	_, err = db.Collection(RoomCollection).InsertOne(ctx, room)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while inserting a new room",
		})
		return
	}
	c.JSON(http.StatusOK, room)
}

func ListRoomHandler(c *gin.Context) {
	cur, err := db.Collection(RoomCollection).Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	defer cur.Close(ctx)

	rooms := make([]models.Room, 0)
	for cur.Next(ctx) {
		var room models.Room
		cur.Decode(&room)
		rooms = append(rooms, room)
	}

	c.JSON(http.StatusOK, rooms)
}

func NewBookingHandler(c *gin.Context) {
	var booking models.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	booking.ID = primitive.NewObjectID()
	_, err = db.Collection(BookingCollection).InsertOne(ctx, booking)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while inserting a new room",
		})
		return
	}
	c.JSON(http.StatusOK, booking)
}

func ListBookingHandler(c *gin.Context) {
	cur, err := db.Collection(BookingCollection).Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	defer cur.Close(ctx)

	bookings := make([]models.Booking, 0)
	for cur.Next(ctx) {
		var booking models.Booking
		cur.Decode(&booking)
		bookings = append(bookings, booking)
	}

	c.JSON(http.StatusOK, bookings)
}

func main() {
	router := gin.Default()
	router.GET("/rooms", ListRoomHandler)
	router.POST("/rooms/create", NewRoomHandler)
	router.GET("/booking", ListBookingHandler)
	router.POST("/booking/create", NewBookingHandler)

	router.Run()
}

func initConfig() error {
	viper.AddConfigPath("../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
