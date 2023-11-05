package controllers

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/pattyarao/wechoose/database"
	"github.com/pattyarao/wechoose/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func findUser(user *models.User, c *fiber.Ctx, post models.Post, wg *sync.WaitGroup) error {
	userCollection := database.MG.Db.Collection("users")
	userFilter := bson.M{"user_name": post.Name}
	record := userCollection.FindOne(c.Context(), userFilter)

	record.Decode(user)

	if user.Username != post.Name {
		return c.Status(404).SendString("User not found")
	}
	wg.Done()
	return nil
}

func insertPost(createdPost *models.Post, post models.Post, c *fiber.Ctx, wg *sync.WaitGroup) error {
	postCollection := database.MG.Db.Collection("posts")

	insertionResult, err := postCollection.InsertOne(c.Context(), post)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := postCollection.FindOne(c.Context(), filter)
	createdRecord.Decode(createdPost)
	wg.Done()
	return nil
}

func CreatePost(c *fiber.Ctx) error {
	if err := database.Connect(); err != nil {
		return c.Status(500).SendString("Server error!")
	}

	post := new(models.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	wg := &sync.WaitGroup{}

	wg.Add(2)

	userFilter := bson.M{"user_name": post.Name}
	userFound := &models.User{}
	createdPost := &models.Post{}

	go findUser(userFound, c, *post, wg)
	go insertPost(createdPost, *post, c, wg)
	wg.Wait()

	userFound.Polls = append(userFound.Polls, createdPost.Id)

	userUpdate := bson.D{{Key: "$set", Value: bson.D{
		{Key: "polls", Value: userFound.Polls},
	}}}

	err := database.MG.Db.Collection("users").FindOneAndUpdate(c.Context(), userFilter, userUpdate).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).SendString("No document found!")
		}
		return c.Status(500).SendString(err.Error())
	}

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

func FindUserPosts(c *fiber.Ctx) error {
	if err := database.Connect(); err != nil {
		return c.Status(500).SendString("Server error!")
	}
	userCollection := database.MG.Db.Collection("users")
	postsCollection := database.MG.Db.Collection("posts")
	paramsId := c.Params("userId")

	userId, err := primitive.ObjectIDFromHex(paramsId)
	if err != nil {
		return c.Status(400).SendString("Invalid Id")
	}
	userFilter := bson.M{"_id": userId}
	record := userCollection.FindOne(c.Context(), userFilter)

	if record.Err() != nil {
		return c.Status(404).SendString("User not found")
	}
	user := &models.User{}
	record.Decode(user)

	newUserPolls := []primitive.ObjectID{}
	for _, id := range user.Polls {
		objectID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(500).SendString("Server Error")
		}
		newUserPolls = append(newUserPolls, objectID)
	}

	filter := bson.M{"_id": bson.M{"$in": newUserPolls}}
	posts := make([]models.Post, 0)
	cursor, err := postsCollection.Find(c.Context(), filter)
	if err != nil {
		return c.Status(500).SendString("Server Error!")
	}
	if err := cursor.All(c.Context(), &posts); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.Disconnect()
	return c.Status(200).JSON(posts)
}
