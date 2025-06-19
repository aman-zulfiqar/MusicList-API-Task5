package functions

import (
	"log"
	"net/http"

	"github.com/aman-zulfiqar/MusicList-API/config"
	"github.com/aman-zulfiqar/MusicList-API/models"
	"github.com/aman-zulfiqar/MusicList-API/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func RegisterUser(c echo.Context) error {
	var input RegisterInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	var exists bool
	err := config.DB.QueryRow(`SELECT EXISTS (SELECT 1 FROM users WHERE email = $1)`, input.Email).Scan(&exists)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error checking user"})
	}
	if exists {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Email already registered"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	}

	var userID string
	err = config.DB.QueryRow(
		`INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id`,
		input.Email, string(hashedPassword),
	).Scan(&userID)

	if err != nil {
		log.Println("CREATE USER ERROR:", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Error creating user"})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "User registered successfully",
		"user": map[string]interface{}{
			"id":    userID,
			"email": input.Email,
		},
	})
}

func Login(c echo.Context) error {
	var input LoginInput

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	var user models.User
	err := config.DB.QueryRow(
		`SELECT id, password FROM users WHERE email = $1`,
		input.Email,
	).Scan(&user.ID, &user.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Wrong password"})
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Could not create token"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Login successful",
		"token":   token,
		"user": map[string]interface{}{
			"id":    user.ID,
			"email": input.Email,
		},
	})
}
