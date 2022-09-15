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
	dbname   = "Samaritano"
)

type Logs struct {
	Types string `json:"type,omitempty"`
	Time  string `json:"time,omitempty"`
}

type Categoria struct {
	IdCategoria int    `json:"Id_categoria,omitempty"`
	Nombre      string `json:"Nombre,omitempty"`
}

type Producto struct {
	IdProducto int    `json:"Id_producto,omitempty"`
	Nombre     string `json:"Nombre,omitempty"`
	Precio     int    `json:"Precio,omitempty"`
	Categoria  int    `json:"Id_categoria,omitempty"`
}

type Pais struct {
	IdPais int    `json:"Id_pais,omitempty"`
	Nombre string `json:"Nombre,omitempty"`
}

type Vendedor struct {
	IdVendedor int    `json:"Id_vendedor,omitempty"`
	Nombre     string `json:"Nombre,omitempty"`
	IdPais     int    `json:"Id_pais,omitempty"`
}

type Cliente struct {
	IdClient  int    `json:"Id_cliente,omitempty"`
	Nombre    string `json:"Nombre,omitempty"`
	Apellido  string `json:"Apellido,omitempty"`
	Direccion string `json:"Direccion,omitempty"`
	Telefono  int    `json:"Telefono,omitempty"`
	Tarjeta   int    `json:"Tarjeta,omitempty"`
	Edad      int    `json:"Edad,omitempty"`
	Salario   int    `json:"Salario,omitempty"`
	Genero    string `json:"Genero,omitempty"`
	IdPais    int    `json:"Id_pais,omitempty"`
}

type Orden struct {
	IdOrden    int    `json:"Id_orden,omitempty"`
	Fecha      string `json:"fecha,omitempty"`
	Id_cliente int    `json:"Id_cliente,omitempty"`
}

type Orden_Aux struct {
	IdOrden     int    `json:"Id_orden,omitempty"`
	Linea_Orden int    `json:"Linea_orden,omitempty"`
	Fecha       string `json:"fecha_orden,omitempty"`
	Id_cliente  int    `json:"Id_cliente,omitempty"`
	Id_vendedor int    `json:"Id_vendedor,omitempty"`
	Id_producto int    `json:"Id_producto,omitempty"`
	Cantidad    int    `json:"Cantidad,omitempty"`
}

type DetalleOrden struct {
	Id_detalle  int `json:"Id_detalle,omitempty"`
	Id_orden    int `json:"Id_orden,omitempty"`
	Id_producto int `json:"Id_producto,omitempty"`
	Cantidad    int `json:"Cantidad,omitempty"`
	Id_vendedor int `json:"Id_vendedor,omitempty"`
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
	rows, err := db.Query("SELECT * from Pais")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var result []Pais

	for rows.Next() {
		user := Pais{}
		if err := rows.Scan(&user.IdPais, &user.Nombre); err != nil {
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
