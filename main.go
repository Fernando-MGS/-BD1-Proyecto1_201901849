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

type BestClient struct {
	ID       int     `json:"ID,omitempty"`
	Nombre   string  `json:"Nombre,omitempty"`
	Apellido string  `json:"Apellido,omitempty"`
	Pais     string  `json:"Pais,omitempty"`
	Monto    float64 `json:"Monto,omitempty"`
}

type ConsultaElemento struct {
	ID     int     `json:"ID,omitempty"`
	Nombre string  `json:"Nombre,omitempty"`
	Monto  float64 `json:"Monto,omitempty"`
}

type ConsultaMes struct {
	Mes   int     `json:"Mes,omitempty"`
	Monto float64 `json:"Monto,omitempty"`
}

type ConsultaProducto struct {
	ID        int     `json:"ID,omitempty"`
	Producto  string  `json:"Producto,omitempty"`
	Categoria string  `json:"Categoria,omitempty"`
	Unidades  int     `json:"Unidades,omitempty"`
	Monto     float64 `json:"Monto,omitempty"`
}

type ConsultaPais struct {
	Nombre string  `json:"Nombre,omitempty"`
	Monto  float64 `json:"Monto,omitempty"`
}

type ConsultaPaisCategoria struct {
	Pais   string  `json:"Pais,omitempty"`
	Nombre string  `json:"Nombre,omitempty"`
	Monto  float64 `json:"Monto,omitempty"`
}

func consulta1(c *fiber.Ctx) error {
	query := "select Cliente.Id_cliente AS ID, Cliente.Nombre AS Nombre, Cliente.Apellido AS Apellido, Pais.Nombre AS Pais, SUM(Producto.Precio * DetalleOrden.Cantidad) as Monto "
	query += " FROM Orden INNER JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente "
	query += " INNER JOIN DetalleOrden on DetalleOrden.Id_orden = Orden.Id_orden "
	query += " inner join Producto on Producto.Id_producto = DetalleOrden.Id_producto "
	query += " INNER JOIN Pais on Cliente.Id_pais = Pais.Id_pais "
	query += " group by Orden.Id_cliente order by SUM(Producto.Precio * DetalleOrden.Cantidad) DESC limit 0,1; "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := BestClient{}
	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Nombre, &result.Apellido, &result.Pais, &result.Monto); err != nil {
			return err
		}
	}

	return c.JSON(result)
}

func Consulta2(c *fiber.Ctx) error {
	query := "(SELECT Producto.Id_producto as ID, Producto.Nombre as Producto, Categoria.Nombre as Categoria, SUM(DetalleOrden.Cantidad) as Unidades, SUM(DetalleOrden.Cantidad*Producto.Precio) as Monto FROM "
	query += " Producto INNER JOIN DetalleOrden ON Producto.Id_producto = DetalleOrden.Id_producto "
	query += " INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria "
	query += " GROUP BY Producto.Id_producto order by SUM(DetalleOrden.Cantidad) DESC LIMIT 0,1) "
	query += " UNION( "
	query += " SELECT Producto.Id_producto as ID, Producto.Nombre as Producto, Categoria.Nombre as Categoria, SUM(DetalleOrden.Cantidad) as Unidades, SUM(DetalleOrden.Cantidad*Producto.Precio) as Monto FROM "
	query += " Producto INNER JOIN DetalleOrden ON Producto.Id_producto = DetalleOrden.Id_producto "
	query += " INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria "
	query += " GROUP BY Producto.Id_producto order by SUM(DetalleOrden.Cantidad) ASC LIMIT 0,1); "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var resultados []ConsultaProducto
	for rows.Next() {
		result := ConsultaProducto{}
		if err := rows.Scan(&result.ID, &result.Producto, &result.Categoria, &result.Unidades, &result.Monto); err != nil {
			return err
		}
		resultados = append(resultados, result)
	}

	return c.JSON(resultados)
}

func Consulta3(c *fiber.Ctx) error {
	query := "SELECT Vendedor.Id_vendedor as ID, Vendedor.Nombre AS Nombre, SUM(DetalleOrden.Cantidad * Producto.Precio) AS Monto "
	query += " FROM Vendedor inner join DetalleOrden ON Vendedor.Id_vendedor = DetalleOrden.Id_vendedor "
	query += " inner JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto "
	query += " group by Vendedor.Id_vendedor order by SUM(DetalleOrden.Cantidad * Producto.Precio) DESC limit 0,1; "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	result := ConsultaElemento{}
	for rows.Next() {
		if err := rows.Scan(&result.ID, &result.Nombre, &result.Monto); err != nil {
			return err
		}
	}

	return c.JSON(result)
}

func Consulta4(c *fiber.Ctx) error {
	query := "(SELECT Pais.Nombre AS Nombre, SUM(DetalleOrden.Cantidad*Producto.Precio) AS Monto FROM "
	query += " Producto Inner JOIN DetalleOrden ON DetalleOrden.Id_producto = Producto.Id_producto "
	query += " INNER JOIN Vendedor ON DetalleOrden.Id_vendedor = Vendedor.Id_vendedor "
	query += " INNER JOIN Pais ON Pais.Id_pais = Vendedor.Id_pais "
	query += " group by Pais.Id_pais ORDER BY SUM(DetalleOrden.Cantidad*Producto.Precio) DESC limit 0,1) "
	query += " UNION( "
	query += " SELECT Pais.Nombre AS Nombre, SUM(DetalleOrden.Cantidad*Producto.Precio) AS Monto FROM "
	query += " Producto Inner JOIN DetalleOrden ON DetalleOrden.Id_producto = Producto.Id_producto "
	query += " INNER JOIN Vendedor ON DetalleOrden.Id_vendedor = Vendedor.Id_vendedor "
	query += " INNER JOIN Pais ON Pais.Id_pais = Vendedor.Id_pais "
	query += " group by Pais.Id_pais ORDER BY SUM(DetalleOrden.Cantidad*Producto.Precio) ASC limit 0,1); "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var resultados []ConsultaPais
	for rows.Next() {
		result := ConsultaPais{}
		if err := rows.Scan(&result.Nombre, &result.Monto); err != nil {
			return err
		}
		resultados = append(resultados, result)
	}

	return c.JSON(resultados)
}

func Consulta5(c *fiber.Ctx) error {
	query := "SELECT Pais.Id_pais AS ID, Pais.Nombre as Nombre, SUM(DetalleOrden.Cantidad * Producto.Precio) AS Monto FROM DetalleOrden "
	query += " inner JOIN Orden ON Orden.Id_Orden = DetalleOrden.Id_orden "
	query += " inner JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente "
	query += " inner JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto "
	query += " inner JOIN Pais ON Cliente.Id_pais = Pais.Id_pais "
	query += " group by Pais.Id_pais order by SUM(DetalleOrden.Cantidad * Producto.Precio) DESC limit 0,5; "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var resultados []ConsultaElemento
	for rows.Next() {
		result := ConsultaElemento{}
		if err := rows.Scan(&result.ID, &result.Nombre, &result.Monto); err != nil {
			return err
		}
		resultados = append(resultados, result)
	}

	return c.JSON(resultados)
}

func Consulta6(c *fiber.Ctx) error {
	query := "(SELECT Categoria.Nombre as Nombre, SUM(DetalleOrden.Cantidad) as Monto FROM DetalleOrden "
	query += " INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto "
	query += " INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria "
	query += " GROUP BY Categoria.Id_categoria ORDER BY SUM(DetalleOrden.Cantidad) DESC LIMIT 0,1) "
	query += " UNION "
	query += " (SELECT Categoria.Nombre as Nombre, SUM(DetalleOrden.Cantidad ) as Monto FROM DetalleOrden "
	query += " INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto "
	query += " INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria "
	query += " GROUP BY Categoria.Id_categoria ORDER BY SUM(DetalleOrden.Cantidad ) ASC LIMIT 0,1); "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var resultados []ConsultaPais
	for rows.Next() {
		result := ConsultaPais{}
		if err := rows.Scan(&result.Nombre, &result.Monto); err != nil {
			return err
		}
		resultados = append(resultados, result)
	}

	return c.JSON(resultados)
}

func Consulta7(c *fiber.Ctx) error {
	query := "WITH CTRY AS (SELECT Pais.Nombre AS Pais, Categoria.Nombre AS Nombre,SUM(DetalleOrden.Cantidad*Producto.Precio) As Monto FROM DetalleOrden "
	query += " INNER JOIN Orden ON DetalleOrden.Id_orden = Orden.Id_orden "
	query += " INNER JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente "
	query += " INNER JOIN Pais ON Cliente.Id_pais = Pais.Id_pais "
	query += " INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto "
	query += " INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_Categoria "
	query += " GROUP BY Pais.Id_pais, Categoria.Id_categoria  order by  Pais.Nombre ,SUM(DetalleOrden.Cantidad) DESC), "
	query += " FILTRO AS (select Pais AS Pais, MAX(Monto) as Monto FROM CTRY GROUP BY Pais) "
	query += " select FILTRO.Pais, CTRY.Nombre ,CTRY.Monto FROM FILTRO "
	query += " INNER JOIN CTRY ON FILTRO.Pais = CTRY.Pais AND FILTRO.Monto = CTRY.Monto; "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var resultados []ConsultaPaisCategoria
	for rows.Next() {
		result := ConsultaPaisCategoria{}
		if err := rows.Scan(&result.Pais, &result.Nombre, &result.Monto); err != nil {
			return err
		}
		resultados = append(resultados, result)
	}

	return c.JSON(resultados)
}

func Consulta8(c *fiber.Ctx) error {
	query := "SELECT DATE_FORMAT(Orden.fecha, \"%m\") as Mes, SUM(Producto.Precio * DetalleOrden.Cantidad) AS Monto FROM DetalleOrden "
	query += " INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto "
	query += " INNER JOIN Orden ON Orden.Id_orden = DetalleOrden.Id_orden "
	query += " INNER JOIN Vendedor ON Vendedor.Id_vendedor = DetalleOrden.Id_vendedor "
	query += " INNER JOIN Pais ON Pais.Id_pais = Vendedor.Id_pais AND Pais.Nombre = 'Inglaterra' "
	query += " group by DATE_FORMAT(Orden.fecha, \"%m\"); "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var resultados []ConsultaMes
	for rows.Next() {
		result := ConsultaMes{}
		if err := rows.Scan(&result.Mes, &result.Monto); err != nil {
			return err
		}
		resultados = append(resultados, result)
	}

	return c.JSON(resultados)
}

func Consulta9(c *fiber.Ctx) error {
	query := "(SELECT DATE_FORMAT(Orden.fecha, \"%m\") AS Mes, SUM(DetalleOrden.Cantidad * Producto.Precio) as Monto FROM DetalleOrden "
	query += " INNER JOIN Producto ON DetalleOrden.Id_producto = Producto.Id_producto "
	query += " INNER JOIN Orden ON DetalleOrden.Id_Orden = Orden.Id_orden "
	query += " GROUP BY DATE_FORMAT(Orden.fecha, \"%m\") ORDER BY  SUM(DetalleOrden.Cantidad * Producto.Precio) DESC LIMIT 0,1) "
	query += " UNION( "
	query += "SELECT DATE_FORMAT(Orden.fecha, \"%m\") AS Mes, SUM(DetalleOrden.Cantidad * Producto.Precio) as Monto FROM DetalleOrden "
	query += " INNER JOIN Producto ON DetalleOrden.Id_producto = Producto.Id_producto "
	query += " INNER JOIN Orden ON DetalleOrden.Id_Orden = Orden.Id_orden "
	query += " GROUP BY DATE_FORMAT(Orden.fecha, \"%m\") ORDER BY  SUM(DetalleOrden.Cantidad * Producto.Precio) ASC LIMIT 0,1); "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var resultados []ConsultaMes
	for rows.Next() {
		result := ConsultaMes{}
		if err := rows.Scan(&result.Mes, &result.Monto); err != nil {
			return err
		}
		resultados = append(resultados, result)
	}

	return c.JSON(resultados)
}

func Consulta10(c *fiber.Ctx) error {
	query := "SELECT Producto.Id_producto AS ID, Producto.Nombre AS Nombre, SUM(Producto.Precio*DetalleOrden.Cantidad) as Monto "
	query += " FROM DetalleOrden inner JOIN Producto ON DetalleOrden.Id_producto = Producto.Id_producto "
	query += " inner JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria && Categoria.Nombre='Deportes' "
	query += " group by (Producto.Id_producto) order by SUM(Producto.Precio*DetalleOrden.Cantidad) DESC; "
	rows, err := db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()
	var resultados []ConsultaElemento
	for rows.Next() {
		result := ConsultaElemento{}
		if err := rows.Scan(&result.ID, &result.Nombre, &result.Monto); err != nil {
			return err
		}
		resultados = append(resultados, result)
	}

	return c.JSON(resultados)
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
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/consulta1", consulta1)
	app.Get("/consulta2", Consulta2)
	app.Get("/consulta3", Consulta3)
	app.Get("/consulta4", Consulta4)
	app.Get("/consulta5", Consulta5)
	app.Get("/consulta6", Consulta6)
	app.Get("/consulta7", Consulta7)
	app.Get("/consulta8", Consulta8)
	app.Get("/consulta9", Consulta9)
	app.Get("/consulta10", Consulta10)
	log.Fatal(app.Listen(":4000"))
}
