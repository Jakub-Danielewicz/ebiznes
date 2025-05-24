package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"


	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"github.com/labstack/echo-contrib/session"
)

type RegisterInput struct {
	Username string `json:"username`
	Email		 string	`json:"email"`
	Password string	`json:password`
}



func RegisterUser(c echo.Context) error {
	var input RegisterInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Nieprawidłowe dane wejściowe"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password),
												 bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
									"error": "Nie udało się zaszyfrować hasła"})
	} 

	user := models.User{
		Name: input.Username,
		Email:		input.Email,
		PassHash:	string(hashedPassword),
	}

	if result := db.DB.Create(&user); result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error":
									"Nie udało się utworzyć użytkownika"})
	}

	return c.JSON(http.StatusCreated, echo.Map {
		"message": "Użytkownik zarejestrowany pomyślnie",
	})
}

type LoginInput struct {
	Email			string  `json:"email"`
	Password	string	`json:"password"`
}

func Login(c echo.Context) error {
	var input LoginInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error":
						"Nieprawidłowe dane"})
	}
	
	var user models.User
	result := db.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error":
						"Błędny email lub hasło"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error":
						"Błędny email lub hasło"})
	}
	sess,_ := session.Get("session", c)
	sess.Values["user_id"] = user.ID
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}
	

	return c.JSON(http.StatusOK, echo.Map{"Message": "Zalogowano pomyślnie"})
	


}


