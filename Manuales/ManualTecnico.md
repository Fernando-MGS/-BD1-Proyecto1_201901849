# SBD1_P1_201901849
# Manual Técnico
## Fernando Mauricio Gómez Santos
## 201901849

Primer proyecto del curso de Sistemas de Bases de Datos 1. Consiste en un API para realizar consultas hacia una base de datos de MySQL.

## Tecnologías usadas

* Go 1.19
* MySQL

## Esquema Conceptual

![Esquema Conceptual](/Manuales/IMG/e_conceptual.png)

## Esquema Lógico

![Esquema lógico](/Manuales/IMG/e_logico.jpeg)

## Descripción de las tablas

* Categoria: Almacena la información de las categorías de cada producto. Sus columnas son Id_categoria y Nombre. Id_categoria es un integer y la llave primaria y Nombre es un varchar de 50 caracteres.

* Producto: Almacena la información de cada producto disponible. Sus columnas son Id_producto, Nombre, Precio, Id_categoria. Id_producto es la llave primaria y Id_categoria es una llave foránea de la tabla Categoria.

* Pais: Almacena la información de los países de los clientes y vendedores. Sus columnas son Id_pais y Nombre. Id_pais es la llave primaria.

* Vendedor: Almacena la información de los vendedores. Sus columnas son Id_vendedor, Nombre y Id_pais. Id_vendedor es la llave primaria y Id_pais es una llave foránea de la tabla Pais.

* Cliente: Almacena la información de los clientes. Sus columnas son Id_cliente, Nombre, Apellido, Direccion, Telefono, Tarjeta, Edad, Salario, Genero y Id_pais. Id_cliente es la llave primaria y Id_pais es la llave foránea de la tabla Pais.

* Orden: Almacena la información de las ordenes de compra. Sus columnas son Id_orden, fecha, Id_cliente. Id_orden es la llave primaria y Id_cliente es una llave foránea de la tabla Cliente.

* DetalleOrden: Almacena la información de cada producto y vendedor relacionado a una orden. Sus columnas son Id_detalle, Id_orden, Id_producto, Cantidad, Id_vendedor. Id_detalle es la llave primaria de la tabla, Id_orden es una llave foránea de la tabla Orden, Id_producto es una llave foránea de la tabla producto, Id_vendedor es la llave foránea de la tabla Vendedor.
  
## Descripción de la API

La API fue creada utilizando el lenguaje de progamación Golang. Para crear el servidor se hizo uso del framework llamado Fiber y de los drivers de mysql para completar la conexión a la base de datos.


## Descripción de los endpoint utilizados

* '/consulta1': Muestra el cliente que mas ha comprado. Incluye su id, nombre, apellido, país y el monto total de sus compras.
* '/consulta2': Muestra el producto más y menos comprado. Incluye el id, nombre, categoría, unidades y monto de los productos.
* '/consulta3': Muestra al vendedor con mayores ventas. Incluye el id y nombre del vendedor junto con el monto total de ventas.
* '/consulta4': Muestra los países con los mayores y menores ventas. Incluye el nombre del país y el monto de sus ventas.
* '/consulta5': Busca el top 5 de países que mas han comprado. Los datos de cada país son el id, nombre y el monto de ventas.
* '/consulta6': Realiza una búsqueda de las categorías mas y y menos compradas. Los datos que incluyen son el nombre y la cantidad de unidades de cada uno.
* '/consulta7': Muestra  la categoría mas comprada por cada país. Cada resultado incluye el nombre del país, nombre de la categoría y la cantidad de unidades compradas de cada país.
* '/consulta8': Muestra las ventas por mes de Inglaterra. Incluye el número del mes y el monto.
* '/consulta9': Muestra el mes con mas y menos ventas. Cada resultado incluye el número de mes y el monto vendido.
* '/consulta10': Muestra las ventas de cada producto de la categoría deportes. Cada fila incluye el id del producto, nombre y monto.
