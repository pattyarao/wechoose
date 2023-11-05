package controllers

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/pattyarao/wechoose/database"
	"github.com/pattyarao/wechoose/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {

	if err := database.Connect(); err != nil {
		return c.Status(500).SendString("Server error!")
	}

	collection := database.MG.Db.Collection("users")

	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return c.SendString("There was an error in handling the password")
	}

	user.Password = string(hashedPassword)

	insersionResult, err := collection.InsertOne(c.Context(), user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insersionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdUser := &models.User{}
	createdRecord.Decode(createdUser)

	database.Disconnect()
	createdUser.Password = ""
	return c.Status(201).JSON(createdUser)

}

func SignIn(c *fiber.Ctx) error {
	godotenv.Load()
	if err := database.Connect(); err != nil {
		return c.Status(500).SendString("Server error!")
	}

	type LoginInput struct {
		Username string `json:"user_name" bson:"user_name"`
		Password string `json:"password" bson:"password"`
	}

	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	collection := database.MG.Db.Collection("users")

	username := input.Username
	password := input.Password

	filter := bson.M{"user_name": username}
	record := collection.FindOne(c.Context(), filter)

	if record.Err() != nil {
		return c.Status(404).SendString("User not found")
	}

	user := &models.User{}
	if err := record.Decode(user); err != nil {
		return c.Status(500).SendString("Error retrieving the user: " + err.Error())
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	finalToken, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	database.Disconnect()
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Login Successful!", "token": finalToken})
}
