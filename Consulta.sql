use Samaritano;
#CONSULTA 1
select Cliente.Id_cliente AS ID, Cliente.Nombre AS Nombre, Cliente.Apellido AS Apellido, Pais.Nombre AS Pais, SUM(Producto.Precio * DetalleOrden.Cantidad) as Monto
FROM Orden INNER JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente 
INNER JOIN DetalleOrden on DetalleOrden.Id_orden = Orden.Id_orden
inner join Producto on Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Pais on Cliente.Id_pais = Pais.Id_pais
group by Orden.Id_cliente order by SUM(Producto.Precio * DetalleOrden.Cantidad) DESC limit 0,1;

#CONSULTA 2
(SELECT Producto.Id_producto as ID, Producto.Nombre as Producto, Categoria.Nombre as Categoria, SUM(DetalleOrden.Cantidad) as Unidades, SUM(DetalleOrden.Cantidad*Producto.Precio) as Monto FROM
Producto INNER JOIN DetalleOrden ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria
GROUP BY Producto.Id_producto order by SUM(DetalleOrden.Cantidad*Producto.Precio) DESC LIMIT 0,1)
UNION(
SELECT Producto.Id_producto as ID, Producto.Nombre as Producto, Categoria.Nombre as Categoria, SUM(DetalleOrden.Cantidad) as Unidades, SUM(DetalleOrden.Cantidad*Producto.Precio) as Monto FROM
Producto INNER JOIN DetalleOrden ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria
GROUP BY Producto.Id_producto order by SUM(DetalleOrden.Cantidad*Producto.Precio) ASC LIMIT 0,1)
;
###
(SELECT Producto.Id_producto as ID, Producto.Nombre as Producto, Categoria.Nombre as Categoria, SUM(DetalleOrden.Cantidad) as Unidades, SUM(DetalleOrden.Cantidad*Producto.Precio) as Monto FROM
Producto INNER JOIN DetalleOrden ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria
GROUP BY Producto.Id_producto order by SUM(DetalleOrden.Cantidad) DESC LIMIT 0,1)
UNION(
SELECT Producto.Id_producto as ID, Producto.Nombre as Producto, Categoria.Nombre as Categoria, SUM(DetalleOrden.Cantidad) as Unidades, SUM(DetalleOrden.Cantidad*Producto.Precio) as Monto FROM
Producto INNER JOIN DetalleOrden ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria
GROUP BY Producto.Id_producto order by SUM(DetalleOrden.Cantidad) ASC LIMIT 0,1);
#CONSULTA 3
SELECT Vendedor.Id_vendedor as ID, Vendedor.Nombre AS Nombre, SUM(DetalleOrden.Cantidad * Producto.Precio) AS Monto
FROM Vendedor inner join DetalleOrden ON Vendedor.Id_vendedor = DetalleOrden.Id_vendedor
inner JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto 
group by Vendedor.Id_vendedor order by SUM(DetalleOrden.Cantidad * Producto.Precio) DESC limit 0,1;

#CONSULTA 4
(SELECT Pais.Nombre AS Nombre, SUM(DetalleOrden.Cantidad*Producto.Precio) AS Monto FROM
Producto Inner JOIN DetalleOrden ON DetalleOrden.Id_producto = Producto.Id_producto
INNER JOIN Vendedor ON DetalleOrden.Id_vendedor = Vendedor.Id_vendedor
INNER JOIN Pais ON Pais.Id_pais = Vendedor.Id_pais
group by Pais.Id_pais ORDER BY SUM(DetalleOrden.Cantidad*Producto.Precio) DESC limit 0,1)
UNION(
SELECT Pais.Nombre AS Nombre, SUM(DetalleOrden.Cantidad*Producto.Precio) AS Monto FROM
Producto Inner JOIN DetalleOrden ON DetalleOrden.Id_producto = Producto.Id_producto
INNER JOIN Vendedor ON DetalleOrden.Id_vendedor = Vendedor.Id_vendedor
INNER JOIN Pais ON Pais.Id_pais = Vendedor.Id_pais
group by Pais.Id_pais ORDER BY SUM(DetalleOrden.Cantidad*Producto.Precio) ASC limit 0,1
);

#CONSULTA 5
SELECT Pais.Id_pais AS ID, Pais.Nombre as Nombre, SUM(DetalleOrden.Cantidad * Producto.Precio) AS Monto FROM DetalleOrden
inner JOIN Orden ON Orden.Id_Orden = DetalleOrden.Id_orden
inner JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente
inner JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
inner JOIN Pais ON Cliente.Id_pais = Pais.Id_pais
group by Pais.Id_pais order by SUM(DetalleOrden.Cantidad * Producto.Precio) DESC limit 0,5;
#####
SELECT Pais.Id_pais AS ID, Pais.Nombre as Nombre, SUM(DetalleOrden.Cantidad * Producto.Precio) AS Monto FROM
Vendedor inner JOIN DetalleOrden ON Vendedor.Id_vendedor = DetalleOrden.Id_vendedor
inner JOIN Orden ON Orden.Id_Orden = DetalleOrden.Id_orden
inner JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
inner JOIN Pais ON Vendedor.Id_pais = Pais.Id_pais
group by Pais.Id_pais order by SUM(DetalleOrden.Cantidad * Producto.Precio) DESC limit 0,5;

#CONSULTA 6
(SELECT Categoria.Nombre as Nombre, SUM(DetalleOrden.Cantidad * Producto.Precio) as Monto FROM DetalleOrden
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria
GROUP BY Categoria.Id_categoria ORDER BY SUM(DetalleOrden.Cantidad * Producto.Precio) DESC LIMIT 0,1)
UNION
(SELECT Categoria.Nombre as Nombre, SUM(DetalleOrden.Cantidad * Producto.Precio) as Monto FROM DetalleOrden
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria
GROUP BY Categoria.Id_categoria ORDER BY SUM(DetalleOrden.Cantidad * Producto.Precio) ASC LIMIT 0,1);
####

(SELECT Categoria.Nombre as Nombre, SUM(DetalleOrden.Cantidad) as Monto FROM DetalleOrden
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria
GROUP BY Categoria.Id_categoria ORDER BY SUM(DetalleOrden.Cantidad) DESC LIMIT 0,1)
UNION
(SELECT Categoria.Nombre as Nombre, SUM(DetalleOrden.Cantidad ) as Monto FROM DetalleOrden
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria
GROUP BY Categoria.Id_categoria ORDER BY SUM(DetalleOrden.Cantidad ) ASC LIMIT 0,1);

#CONSULTA 7
SELECT Pais.Nombre AS Pais, MIN(Categoria.Nombre) ,SUM(DetalleOrden.Cantidad) As Unidades FROM DetalleOrden
INNER JOIN Orden ON DetalleOrden.Id_orden = Orden.Id_orden
INNER JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente
INNER JOIN Pais ON Cliente.Id_pais = Pais.Id_pais
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_Categoria
GROUP BY Pais.Id_pais order by  SUM(DetalleOrden.Cantidad) DESC;

SELECT Pais.Nombre AS Pais, Categoria.Nombre AS Nombre,SUM(DetalleOrden.Cantidad) As Monto FROM DetalleOrden
INNER JOIN Orden ON DetalleOrden.Id_orden = Orden.Id_orden
INNER JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente
INNER JOIN Pais ON Cliente.Id_pais = Pais.Id_pais
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_Categoria
GROUP BY Pais.Id_pais, Categoria.Id_categoria  order by  Pais.Nombre ,SUM(DetalleOrden.Cantidad) DESC;
###
SELECT Pais.Nombre AS Pais, Categoria.Nombre AS Nombre,SUM(DetalleOrden.Cantidad) As Monto FROM DetalleOrden
INNER JOIN Orden ON DetalleOrden.Id_orden = Orden.Id_orden
INNER JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente
INNER JOIN Pais ON Cliente.Id_pais = Pais.Id_pais
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_Categoria
GROUP BY Pais.Id_pais, Categoria.Id_categoria  order by  Pais.Nombre ,SUM(DetalleOrden.Cantidad) DESC;

###
WITH CTRY AS (SELECT Pais.Nombre AS Pais, Categoria.Nombre AS Nombre,SUM(DetalleOrden.Cantidad) As Monto FROM DetalleOrden
INNER JOIN Orden ON DetalleOrden.Id_orden = Orden.Id_orden
INNER JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente
INNER JOIN Pais ON Cliente.Id_pais = Pais.Id_pais
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_Categoria
GROUP BY Pais.Id_pais, Categoria.Id_categoria  order by  Pais.Nombre ,SUM(DetalleOrden.Cantidad) DESC),
FILTRO AS (select Pais AS Pais, MAX(Monto) as Monto FROM CTRY GROUP BY Pais)
select FILTRO.Pais, CTRY.Nombre ,CTRY.Monto FROM FILTRO 
INNER JOIN CTRY ON FILTRO.Pais = CTRY.Pais AND FILTRO.Monto = CTRY.Monto;
###
SELECT Pais.Nombre AS Pais, Categoria.Nombre ,SUM(DetalleOrden.Cantidad) As Unidades FROM DetalleOrden
INNER JOIN Orden ON DetalleOrden.Id_orden = Orden.Id_orden
INNER JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente
INNER JOIN Pais ON Cliente.Id_pais = Pais.Id_pais
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Categoria ON Categoria.Id_categoria = Producto.Id_Categoria
GROUP BY Pais.Id_pais, Categoria.Id_categoria  
order by  Pais.Nombre ,SUM(DetalleOrden.Cantidad) DESC;
#CONSULTA 8
SELECT DATE_FORMAT(Orden.fecha, "%m") as Mes, SUM(Producto.Precio * DetalleOrden.Cantidad) AS Monto FROM DetalleOrden
INNER JOIN Producto ON Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Orden ON Orden.Id_orden = DetalleOrden.Id_orden
INNER JOIN Vendedor ON Vendedor.Id_vendedor = DetalleOrden.Id_vendedor
INNER JOIN Pais ON Pais.Id_pais = Vendedor.Id_pais AND Pais.Nombre = 'Inglaterra'
group by DATE_FORMAT(Orden.fecha, "%m");

#CONSULTA 9
(SELECT DATE_FORMAT(Orden.fecha, "%m") AS Mes, SUM(DetalleOrden.Cantidad * Producto.Precio) as Monto FROM DetalleOrden
INNER JOIN Producto ON DetalleOrden.Id_producto = Producto.Id_producto
INNER JOIN Orden ON DetalleOrden.Id_Orden = Orden.Id_orden
GROUP BY DATE_FORMAT(Orden.fecha, "%m") ORDER BY  SUM(DetalleOrden.Cantidad * Producto.Precio) DESC LIMIT 0,1)
UNION(
SELECT DATE_FORMAT(Orden.fecha, "%m") AS Mes, SUM(DetalleOrden.Cantidad * Producto.Precio) as Monto FROM DetalleOrden
INNER JOIN Producto ON DetalleOrden.Id_producto = Producto.Id_producto
INNER JOIN Orden ON DetalleOrden.Id_Orden = Orden.Id_orden
GROUP BY DATE_FORMAT(Orden.fecha, "%m") ORDER BY  SUM(DetalleOrden.Cantidad * Producto.Precio) ASC LIMIT 0,1
);

#CONSULTA 10
SELECT Producto.Id_producto AS ID, Producto.Nombre AS Nombre, SUM(Producto.Precio*DetalleOrden.Cantidad) as Monto
FROM DetalleOrden inner JOIN Producto ON DetalleOrden.Id_producto = Producto.Id_producto
inner JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria AND Categoria.Nombre='Deportes'
group by (Producto.Id_producto) order by SUM(Producto.Precio*DetalleOrden.Cantidad) DESC;
