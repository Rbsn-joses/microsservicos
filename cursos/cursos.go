package cursos

import (
	"encoding/json"
	"fmt"

	"github.com/Rbsn-joses/microsservicos/database"
	"github.com/Rbsn-joses/microsservicos/jwt"
	"github.com/gofiber/fiber/v2"
)

type Course struct {
	Description string
	Course      string
	Department  string
	Duration    string
}
type Courses []Course

func MicroserviceCursos() {
	app := fiber.New()
	//var err error

	app.Get("/api/servidorweb/cursos", func(c *fiber.Ctx) error {
		token := c.Cookies("jwt-token")
		var courseArrayJson []byte
		if token == "" {
			fmt.Println("Requisição sem token")
			return c.SendString("Requisição sem token")

		} else {
			err := jwt.ValidateToken(token)
			if err != nil {
				fmt.Println("erro em jwt.Validate", err)
				return c.SendString(err.Error())

			} else {
				courseArray := getCursos()
				courseArrayJson, err = json.Marshal(courseArray)
				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println("token")
		}

		return c.SendString(string(courseArrayJson))

	})
	fmt.Println("Rodando")
	app.Listen(":3000")
}
func getCookies() {

}
func getCursos() Courses {
	var course = Course{}
	var courses = Courses{}
	db := database.Connection()
	defer db.Close()
	rows, err := db.Query(`SELECT "description", "course","department","duration" FROM "course"`)
	if err != nil {
		fmt.Println("error em login.db.Query ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var description string
		var courseInDB string
		var department string
		var duration string

		err = rows.Scan(&description, &courseInDB, &department, &duration)
		course.Description = description
		course.Course = courseInDB
		course.Department = department
		course.Duration = duration
		courses = append(courses, course)

	}
	return courses
}
