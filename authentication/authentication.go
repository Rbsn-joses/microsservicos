package authentication

import (
	"fmt"
	"log"
	"strings"

	"github.com/Rbsn-joses/microsservicos/database"
	"github.com/Rbsn-joses/microsservicos/jwt"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
)

type User struct {
	Username string `json: "username"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func init() {

}
func MicroserviceAuthentication() {
	app := fiber.New()
	var err error

	app.Post("/api/servidorweb/login/authentication", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(200)
		}
		var token string
		userNotFound, logado := login(user)
		if logado {
			token, err = jwt.GenerateJWT(user.Username)
			if err != nil {
				fmt.Println("error on jwt.GenerateJWT", err)
			}
			c.Cookie(&fiber.Cookie{
				Name:  "jwt-token",
				Value: token,
				Path:  "/",
			})
		} else {
			return c.SendString(fmt.Sprintf(`{"messsage": %s,"logado": false}`, userNotFound))
		}

		return c.SendString("autenticado " + token)

	})

	app.Post("/api/servidorweb/login/authentication-new", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			fmt.Println("error = ", err)
			return c.SendStatus(200)
		}
		var token string
		_, exist := checkLogin(user)
		if !exist {
			token, err = jwt.GenerateJWT(user.Username)
			if err != nil {
				fmt.Println("error on jwt.GenerateJWT", err)
			}
		}
		c.Cookie(&fiber.Cookie{
			Name:  "jwt-token",
			Value: token,
			Path:  "/",
		})

		return c.SendString(`{"usuário": "criado"}`)

	})
	fmt.Println("Rodando")
	app.Listen(":3000")
}
func createUser(user *User) (string, bool) {
	var email string
	var username string
	var userNotFound string
	var exist bool
	db := database.Connection()
	defer db.Close()
	var err error
	err = db.QueryRow("SELECT username,email FROM users WHERE email=$1 AND username=$2", user.Email, user.Username).Scan(&username, &email)
	switch {
	case err != nil:
		userNotFound = "Usuário ou email não encontrado no banco de dados"
		log.Printf(userNotFound)
	default:
		exist = true
		log.Printf("username is %s, account created on %s\n", username, email)
	}

	return userNotFound, exist
}
func checkLogin(user *User) (string, bool) {
	var email string
	var username string
	var userNotFound string
	var exist bool
	db := database.Connection()
	defer db.Close()
	var err error
	err = db.QueryRow("SELECT username,email FROM users WHERE email=$1 AND username=$2", user.Email, user.Username).Scan(&username, &email)
	switch {
	case err != nil:
		userNotFound = "Usuário ou email não encontrado no banco de dados"
		log.Printf(userNotFound)
	default:
		exist = true
		log.Printf("username is %s, account created on %s\n", username, email)
	}

	return userNotFound, exist
}
func login(user *User) (string, bool) {
	var logado bool

	userNotFound, exist := checkLogin(user)
	if exist {
		db := database.Connection()
		defer db.Close()
		rows, err := db.Query(`SELECT "username", "password","email" FROM "users"`)
		if err != nil {
			fmt.Println("error em login.db.Query ", err)
		}
		defer rows.Close()
		for rows.Next() {
			var username string
			var email string
			var password string

			err = rows.Scan(&username, &password, &email)

			fmt.Println(email, password, username)
			if strings.Contains(user.Username, username) || strings.Contains(user.Email, email) && user.Password == password {
				logado = true
			}
		}
	}

	return userNotFound, logado
}
