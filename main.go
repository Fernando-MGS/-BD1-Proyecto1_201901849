package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 3306 // Default port
	user     = "root"
	password = "323125"
	dbname   = "practicas"
)

type Logs struct {
	Types string `json:"type,omitempty"`
	Time  string `json:"time,omitempty"`
}

type usuario struct {
	Carnet     int    `json:"Carnet,omitempty"`
	Nombres    string `json:"Nombres,omitempty"`
	Apellidos  string `json:"Apellidos,omitempty"`
	Contrasena string `json:"Contrasena,omitempty"`
	Correo     string `json:"Correo,omitempty"`
	Creditos   int    `json:"Creditos,omitempty"`
}

func getLog(c *fiber.Ctx) error {
	return c.JSON(Logs{Types: "hola", Time: "mundo!"})
}

func postLog(c *fiber.Ctx) error {
	u := new(Logs)
	if err := c.BodyParser(u); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	fmt.Println(u)
	return c.JSON(u)
}

func getUsers(c *fiber.Ctx) error {
	rows, err := db.Query("SELECT * from usuario")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []usuario

	for rows.Next() {
		user := usuario{}
		if err := rows.Scan(&user.Carnet, &user.Nombres, &user.Apellidos, &user.Contrasena, &user.Correo, &user.Creditos); err != nil {
			return err
		}

		result = append(result, user)
	}

	return c.JSON(result)
}

func Connect() error {
	var err error
	// Use DSN string to open
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func main() {
	fmt.Println("Hola mundo")
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", getLog)
	app.Get("/users", getUsers)
	app.Post("/employee", postLog)
	log.Fatal(app.Listen(":4000"))
}
