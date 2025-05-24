package oauth

import (
	"context"
	"backend/models"
	"backend/db"
	"encoding/json"
	"net/http"
	"os"
	

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)



func getGoogleOAuthConfig() *oauth2.Config {
    return &oauth2.Config{
        ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        RedirectURL:  "http://localhost:8080/auth/google/callback",
        Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
        Endpoint:     google.Endpoint,
    }
}

type GoogleUser struct {
	Email string `json:"email"`
}

func GoogleLogin(c echo.Context) error {

	url := getGoogleOAuthConfig().AuthCodeURL("randomstate") // powinno być bezpieczniej w prod
	return c.Redirect(http.StatusTemporaryRedirect, url)

}

func GoogleCallback(c echo.Context) error {
	code := c.QueryParam("code")

	token, err := getGoogleOAuthConfig().Exchange(context.Background(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Nie udało się wymienić kodu na token",
		})
	}

	

	client := getGoogleOAuthConfig().Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Nie udało się pobrać danych użytkownika",
		})
	}
	defer resp.Body.Close()

	var gUser GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&gUser); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Nie udało się zdekodować danych użytkownika",
		})
	}

	var user models.User
	result := db.DB.Where("email = ?", gUser.Email).First(&user)
	if result.Error != nil {

		user = models.User{Email: gUser.Email, Name: gUser.Email}
		db.DB.Create(&user)
	}


	sess, _ := session.Get("session", c)
	sess.Values["user_id"] = user.ID
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusSeeOther, "http://localhost:3000")
}
