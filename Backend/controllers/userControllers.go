package controllers

import (
	"net/http"
	"spo_task_3/model"
	"spo_task_3/mongo"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginHelper(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	UserData, isUser := mongo.IsUserExist(user.Email)
	if isUser {
		if mongo.CheckPasswordHash(user.Password, UserData.Password) {
			ctx.JSON(200, gin.H{"message": "Logged In"})

		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Credentials!"})
		}

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Email Not Found!"})
	}

}

func RegisterHelper(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	_, isUser := mongo.IsUserExist(user.Email)
	if !isUser {

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashedPassword)
		if err := mongo.InsertUserIntoDB(user); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		} else {
			ctx.JSON(201, gin.H{"message": "User Created"})
		}

	} else {
		ctx.JSON(409, gin.H{"message": "User Already Exist!"})
	}

}
