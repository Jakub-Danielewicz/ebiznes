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
    "golang.org/x/oauth2/github"
)

func getGithubOAuthConfig() *oauth2.Config {
    return &oauth2.Config{
        ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
        ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
        RedirectURL:  "http://localhost:8080/auth/github/callback",
        Scopes:       []string{"user:email"},
        Endpoint:     github.Endpoint,
    }
}

type GithubUser struct {
    Login string `json:"login"`
    Email string `json:"email"`
}

func GithubLogin(c echo.Context) error {
    url := getGithubOAuthConfig().AuthCodeURL("randomstate")
    return c.Redirect(http.StatusTemporaryRedirect, url)
}

func GithubCallback(c echo.Context) error {
    code := c.QueryParam("code")

    token, err := getGithubOAuthConfig().Exchange(context.Background(), code)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Nie udało się wymienić kodu na token",
        })
    }

    client := getGithubOAuthConfig().Client(context.Background(), token)
    resp, err := client.Get("https://api.github.com/user")
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Nie udało się pobrać danych użytkownika",
        })
    }
    defer resp.Body.Close()

    var ghUser GithubUser
    if err := json.NewDecoder(resp.Body).Decode(&ghUser); err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Nie udało się zdekodować danych użytkownika",
        })
    }


    if ghUser.Email == "" {
        emailResp, err := client.Get("https://api.github.com/user/emails")
        if err == nil {
            defer emailResp.Body.Close()
            var emails []struct {
                Email   string `json:"email"`
                Primary bool   `json:"primary"`
            }
            if err := json.NewDecoder(emailResp.Body).Decode(&emails); err == nil {
                for _, e := range emails {
                    if e.Primary {
                        ghUser.Email = e.Email
                        break
                    }
                }
                if ghUser.Email == "" && len(emails) > 0 {
                    ghUser.Email = emails[0].Email
                }
            }
        }
    }

    if ghUser.Email == "" {
        return c.JSON(http.StatusInternalServerError, echo.Map{
            "error": "Nie udało się pobrać adresu email z GitHuba",
        })
    }

    var user models.User
    result := db.DB.Where("email = ?", ghUser.Email).First(&user)
    if result.Error != nil {

        user = models.User{Email: ghUser.Email, Name: ghUser.Login}
        db.DB.Create(&user)
    }

    sess, _ := session.Get("session", c)
    sess.Values["user_id"] = user.ID
    sess.Save(c.Request(), c.Response())

    return c.Redirect(http.StatusSeeOther, "http://localhost:3000")
}