use Samaritano;
SET GLOBAL local_infile=1;
#Cargar datos de tabla Pais
load data local infile '/home/fernando/Universidad/BD1/pais.csv' into table Pais fields terminated by ',' enclosed by '"' lines terminated by '\n' IGNORE 1 ROWS (Id_pais, Nombre);
#Cargar datos de tabla Categoria
load data local infile '/home/fernando/Universidad/BD1/categoria.csv' into table Categoria fields terminated by ',' enclosed by '"' lines terminated by '\n' IGNORE 1 ROWS (Id_categoria, Nombre);
#Cargar datos de tabla Producto
load data local infile '/home/fernando/Universidad/BD1/producto.csv' into table Producto fields terminated by ',' enclosed by '"' lines terminated by '\n' IGNORE 1 ROWS (Id_producto, Nombre, Precio, Id_categoria);
#Cargar datos de tabla Vendedor
load data local infile '/home/fernando/Universidad/BD1/vendedores.csv' into table Vendedor fields terminated by ',' enclosed by '"' lines terminated by '\n' IGNORE 1 ROWS (Id_vendedor, Nombre, Id_pais);
#Cargar datos de tabla Cliente
load data local infile '/home/fernando/Universidad/BD1/cliente.csv' into table Cliente fields terminated by ',' enclosed by '"' lines terminated by '\n' IGNORE 1 ROWS (Id_cliente, Nombre, Apellido, Direccion, Telefono, Tarjeta, Edad, Salario, Genero, Id_pais);
#Cargar datos de tabla Orden_Aux
load data local infile '/home/fernando/Universidad/BD1/orden.csv' into table Orden_Aux fields terminated by ',' enclosed by '"' lines terminated by '\n' IGNORE 1 ROWS (Id_orden, Linea_orden,@date_time_variable, Id_cliente, Id_vendedor, Id_producto, Cantidad) SET fecha_orden = STR_TO_DATE(@date_time_variable, '%d/%m/%Y');
#Cargar datos de tabla Orden
INSERT INTO Orden (Id_orden, fecha, Id_cliente) SELECT DISTINCT Id_orden, fecha_orden, Id_cliente  FROM Orden_Aux;
#Cargar datos de tabla DetalleOrden
INSERT INTO DetalleOrden (Id_orden, Id_producto, Cantidad, Id_vendedor) SELECT Id_orden, Id_producto, Cantidad, Id_vendedor  FROM Orden_Aux;
