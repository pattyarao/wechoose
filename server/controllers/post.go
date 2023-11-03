package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pattyarao/wechoose/database"
	"github.com/pattyarao/wechoose/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(c *fiber.Ctx) error {
	if err := database.Connect(); err != nil {
		return c.Status(500).SendString("Server error!")
	}

	collection := database.MG.Db.Collection("posts")

	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	insertionResult, err := collection.InsertOne(c.Context(), post)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdPost := &models.Post{}
	createdRecord.Decode(createdPost)

	database.Disconnect()
	return c.Status(200).JSON(createdPost)
}

func FindPost(c *fiber.Ctx) error {
	if err := database.Connect(); err != nil {
		return c.Status(500).SendString("Server error!")
	}

	paramsId := c.Params("id")
	postId, err := primitive.ObjectIDFromHex(paramsId)
	if err != nil {
		return c.Status(400).SendString("Invalid Id")
	}
	collection := database.MG.Db.Collection("posts")

	filter := bson.M{"_id": postId}
	record := collection.FindOne(c.Context(), filter)

	if record.Err() != nil {
		return c.Status(404).SendString("Post not found")
	}

	post := &models.Post{}
	if err := record.Decode(post); err != nil {
		return c.Status(500).SendString("Error retrieving the post: " + err.Error())
	}

	database.Disconnect()
	return c.Status(200).JSON(post)
}

func GetPosts(c *fiber.Ctx) error {
	if err := database.Connect(); err != nil {
		return c.Status(500).SendString("Server error!")
	}

	query := bson.D{{}}

	collection := database.MG.Db.Collection("posts")

	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString("Server Error!")
	}

	var posts []models.Post = make([]models.Post, 0)

	if err := cursor.All(c.Context(), &posts); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.Disconnect()
	return c.JSON(posts)
}
