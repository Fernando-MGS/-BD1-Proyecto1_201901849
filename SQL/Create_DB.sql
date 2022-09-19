CREATE DATABASE Samaritano;

USE Samaritano;

CREATE TABLE Categoria(
	Id_categoria INT PRIMARY KEY NOT NULL,
    Nombre varchar(50)
);

CREATE TABLE Producto(
	Id_producto INT PRIMARY KEY,
    Nombre varchar(50),
    Precio float,
    Id_categoria INT NOT NULL,
    foreign key (Id_categoria) references Categoria(Id_categoria)
);

CREATE TABLE Pais(
	Id_pais INT primary KEY, 
    Nombre varchar(50)
);

CREATE TABLE Vendedor(
	Id_vendedor INT primary key,
    Nombre varchar(50),
	Id_pais INT,
    foreign key (Id_pais) references Pais(Id_pais)
);

CREATE TABLE Cliente(
	Id_cliente INT primary key,
    Nombre varchar(50),
    Apellido varchar(50),
    Direccion varchar(100),
    Telefono INT,
    Tarjeta INT,
    Edad INT,    
    Salario INT,
    Genero varchar(1),
	Id_pais INT,
    foreign key (Id_pais) references Pais(Id_pais)
);

CREATE TABLE Orden(
	Id_orden INT primary key,
    fecha date,
    Id_cliente INT,
    foreign key (Id_cliente) references Cliente(Id_cliente)
);

CREATE TABLE Orden_Aux(
	Id_orden INT,
    Linea_orden INT,
    fecha_orden date,
    Id_cliente INT,
    Id_vendedor INT,
    Id_producto INT,
    Cantidad INT
);

CREATE TABLE DetalleOrden(
	Id_detalle INT primary key auto_increment,
    Id_orden INT,
    Id_producto INT,
    Cantidad INT,
    Id_vendedor INT,
	foreign key (Id_orden) references Orden(Id_orden),
    foreign key (Id_producto) references Producto(Id_producto),
    foreign key (Id_vendedor) references Vendedor(Id_vendedor)
);