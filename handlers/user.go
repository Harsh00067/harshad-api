package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Harsh00067/harshad-api/models"
	"github.com/Harsh00067/harshad-api/services"
	"github.com/Harsh00067/harshad-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateUser(c *gin.Context) {

	fmt.Println("Handler: CreateUser is called")

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("User in handler:", user)

	err := services.CreateUser(user)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created",
	})

}
func GetUsers(c *gin.Context) {
	fmt.Println("Handler: getUsers is called")

	pageStr := c.Query("page")   //string
	limitStr := c.Query("limit") //string
	name := c.Query("name")

	if pageStr == "" || limitStr == "" {
		users, err := services.GetUser()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, users)
		return
	}

	page, err := strconv.Atoi(pageStr) //convert to int
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid page number"})
		return
	}
	limit, err := strconv.Atoi(limitStr) //convert to int
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid limit number"})
		return
	}

	users, err := services.SearchUsers(name, page, limit)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}

func GetUserByID(c *gin.Context) {
	fmt.Println("Handler: getUserByID is called")

	idstr := c.Param("id") // string

	fmt.Println("ID received:", idstr)
	id, err := strconv.Atoi(idstr) //convert to int
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err1 := services.GetUserByID(id)

	if err1 != nil {
		c.JSON(404, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	fmt.Println("Handler: DeleteUser is called")

	idstr := c.Param("id") // string

	fmt.Println("ID received:", idstr)
	id, err := strconv.Atoi(idstr) //convert to int
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}

	err1 := services.DeleteUser(id)

	if err1 != nil {
		c.JSON(404, gin.H{"error": err1.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user Deleted",
	})
}
func UpdateUser(c *gin.Context) {
	fmt.Println("Handler: UpdateUser is called")

	idstr := c.Param("id") // string

	id, err := strconv.Atoi(idstr) //convert to int
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user id"})
		return
	}
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.UpdateUser(id, user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})
}

func Login(c *gin.Context) {

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	// Trim spaces
	req.Username = strings.TrimSpace(req.Username)
	req.Password = strings.TrimSpace(req.Password)

	if req.Username == "" || req.Password == "" {
		c.JSON(400, gin.H{"error": "username and password required"})
		return
	}

	// Dummy check (replace with DB later)
	if req.Username != "admin" || req.Password != "admin" {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	//token, err := utils.GenerateToken(req.Username)

	accessToken, err := utils.GenerateToken(req.Username, 1*time.Minute)
	refreshToken, err := utils.GenerateToken(req.Username, 7*24*time.Hour)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(200, gin.H{
		"message":      "login successful",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {

	var req struct {
		RefreshToken string `json:"refreshtoken"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	token, err := utils.ValidateToken(req.RefreshToken)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refreshtoken"})
		c.Abort()
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	newToken, _ := utils.GenerateToken(username, 15*time.Minute)

	c.JSON(200, gin.H{"token": newToken})
}
