package controllers

import (
    "backend-go/config"
    "backend-go/models"
    "net/http"
    "time"
    "os"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

var JwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func Register(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
        return
    }

    // ตรวจสอบว่ามี email, username และ password
    if input.Email == "" || input.Username == "" || input.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email, username, and password are required"})
        return
    }

    // สร้าง user จาก input
    user := models.User{
        Email:    input.Email,
        Username: input.Username,
        Password: input.Password,
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }
    user.Password = string(hashedPassword)

    if err := config.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register: " + err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    var dbUser models.User
    if err := config.DB.Where("email = ?", input.Email).First(&dbUser).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    expirationTime := time.Now().Add(time.Hour * 24)
    claims := &Claims{
        Username: dbUser.Username, // ยังคงเก็บ username ใน token
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(JwtKey)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func GetMe(c *gin.Context) {
    username, exists := c.Get("username")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
        return
    }

    var user models.User
    if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}