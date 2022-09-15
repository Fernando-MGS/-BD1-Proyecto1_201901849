use Samaritano;
#CONSULTA 1
select Cliente.Id_cliente, Cliente.Nombre, Cliente.Apellido, Pais.Nombre ,SUM(Producto.Precio * DetalleOrden.Cantidad) as Monto
FROM Orden INNER JOIN Cliente ON Cliente.Id_cliente = Orden.Id_cliente 
INNER JOIN DetalleOrden on DetalleOrden.Id_orden = Orden.Id_orden
inner join Producto on Producto.Id_producto = DetalleOrden.Id_producto
INNER JOIN Pais on Cliente.Id_pais = Pais.Id_pais
group by Orden.Id_cliente order by SUM(Producto.Precio * DetalleOrden.Cantidad) DESC limit 0,1;
#CONSULTA 2

#CONSULTA 3
SELECT V
#CONSULTA 4

#CONSULTA 5

#CONSULTA 6

#CONSULTA 7

#CONSULTA 8

#CONSULTA 9
SELECT SUM(DetalleOrden.Cantidad) AS Contador, DetalleOrden.Id_producto  From DetalleOrden INNER JOIN Producto ON Producto.Id_categoria=15 GROUP BY DetalleOrden.Id_producto;
SELECT DetalleOrden.Id_producto, DetalleOrden.Cantidad From DetalleOrden INNER JOIN Producto ON Producto.Id_categoria=15 GROUP BY DetalleOrden.Id_producto;
SELECT DetalleOrden.Id_producto, DetalleOrden.Cantidad, DetalleOrden.Id_detalle From DetalleOrden INNER JOIN Producto ON Producto.Id_categoria=15 order by DetalleOrden.Id_producto;
SELECT DISTINCT COUNT(DetalleOrden.Id_detalle) From DetalleOrden 
INNER JOIN Producto ON Producto.Id_categoria=15;
SELECT COUNT(DetalleOrden.Id_detalle) from DetalleOrden;

SELECT Producto.Id_producto, Producto.Nombre from Categoria INNER JOIN Producto ON Producto.Id_categoria = Categoria.Id_categoria AND Categoria.Nombre='Deportes';

#CONSULTA 10
SELECT Producto.Nombre, Producto.Id_producto, SUM(Producto.Precio*DetalleOrden.Cantidad) as Monto
FROM DetalleOrden inner JOIN Producto ON DetalleOrden.Id_producto = Producto.Id_producto
inner JOIN Categoria ON Categoria.Id_categoria = Producto.Id_categoria && Categoria.Nombre='Deportes'
group by (Producto.Id_producto) order by SUM(Producto.Precio*DetalleOrden.Cantidad) DESC;
