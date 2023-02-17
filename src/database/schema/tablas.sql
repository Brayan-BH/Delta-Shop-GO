CREATE TABLE Productos(
	idProducto varchar(8) PRIMARY KEY NOT NULL,
	nombre varchar(100) NOT NULL,
	nivel int NOT NULL,
	descripcion varchar (200) NOT NULL,
	modelo varchar(100) NOT NULL,
	color varchar(100) NOT NULL,
	nombreImagen varchar(40) NOT NULL,
	urls varchar(100) NOT NULL,
	idCategoria varchar(3) NOT NULL
);

CREATE TABLE Categorias(
    idCategoria varchar(3) PRIMARY KEY NOT NULL,
    nombreCategoria varchar(50) NOT NULL
);

CREATE TABLE Pedidos(
	idPedido varchar(8) PRIMARY KEY NOT NULL,
	idDomicilio varchar(8) NOT NULL,
	total float NOT NULL,
	horaPedido datetime NOT NULL,
	fechaPedido datetime NOT NULL,
	fechaEntrega datetime NOT NULL,
	estado int NOT NULL,
    idCliente varchar(10) NOT NULL,
    FOREIGN KEY (idCliente) REFERENCES Clientes(idCliente),
    FOREIGN KEY (idDomicilio) REFERENCES Domicilios(idDomicilio),
);

CREATE TABLE DetallePagos(
	idDetallePago varchar PRIMARY KEY NOT NULL,
    metodoPago varchar(50) NOT NULL,
    detallePago varchar(100) NOT NULL,
    oservaciones varchar(50) NOT NULL,
    idPedido varchar (8) NOT NULL,
	FOREIGN KEY (idPedido) REFERENCES Pedidos(idPedido)
);

CREATE TABLE Domicilios(
    idDomicilio varchar(36) PRIMARY KEY NOT NULL,
    idPedido varchar(8) NOT NULL,
    direccion varchar(50) NOT NULL,
    distrito varchar(50) NOT NULL,
    provincia varchar(50) NOT NULL,
    departamento varchar(50) NOT NULL,
    ubigeo char(6) NOT NULL,
    referencia varchar(300) NOT NULL,
);

CREATE TABLE Clientes(
    idCliente varchar(36) PRIMARY KEY NOT NULL,
    nombres varchar(100) NOT NULL,
    apellidos varchar(100) NOT NULL,
    email varchar(40) NOT NULL,
    genero int NOT NULL,
    passwords varchar(200) NOT NULL,
    telefono varchar(20) NOT NULL
);

