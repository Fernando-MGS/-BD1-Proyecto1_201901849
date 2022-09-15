use Samaritano;

DELETE FROM Orden_Aux WHERE Id_orden>0;

DROP table Orden_Aux;
DROP table Orden;
DROP table DetalleOrden;

SELECT * FROM Orden;

SELECT DISTINCT Id_orden, fecha_orden, Id_cliente  FROM Orden_Aux;

SELECT COUNTDISTINCT Id_orden, fecha_orden, Id_cliente  FROM Orden_Aux;

SELECT Id_orden, Id_producto, Cantidad, Id_vendedor FROM Orden_Aux;

INSERT INTO Orden (Id_orden, fecha, Id_cliente) SELECT DISTINCT Id_orden, fecha_orden, Id_cliente  FROM Orden_Aux;

INSERT INTO Orden (Id_orden, fecha, Id_cliente) SELECT DISTINCT Id_orden, fecha_orden, Id_cliente  FROM Orden_Aux;

