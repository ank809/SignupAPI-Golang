package controllers

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/ank809/SignupAPI-Golang/database"
	"github.com/ank809/SignupAPI-Golang/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func SignupUser(c *gin.Context) {
	var user models.User
	id := primitive.NewObjectID()
	pass := user.Password
	hashed_password, errs := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if errs != nil {
		fmt.Println(errs)
		return
	}
	if err := c.BindJSON(&user); err != nil {
		fmt.Println(err)
		return
	}
	user.ID = id
	user.Password = string(hashed_password)
	collection_name := "users"
	collection := database.OpenCollection(database.Client, collection_name)
	// Email checker
	if !ValidateEmail(user.Email) {
		fmt.Println("Email is not valid")
		c.JSON(http.StatusBadRequest, "Email is not valid")
		return
	}
	//Username
	filter := bson.M{"username": user.Username}
	usercount, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	if usercount > 0 {
		c.JSON(http.StatusBadRequest, "Username already exists please choose a new one")
		return
	}

	// Password
	if len(user.Password) < 6 {
		a := "Password length should be greater than 6 "
		fmt.Println(a)
		c.JSON(http.StatusBadRequest, a)
		return
	}

	//Adding to database
	_, error := collection.InsertOne(context.Background(), user)
	if error != nil {
		fmt.Println(error)
	}

	// fmt.Println(bcrypt.CompareHashAndPassword(hashed_password, []byte(pass))) // nil its a match
	msg := "User added Successfully"
	fmt.Println(msg)
	c.JSON(200, msg)
	fmt.Println(user.Username, user.Password, user.Email, user.Name, user.ID)

}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
