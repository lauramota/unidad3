create database if not exists libros1;
use libros1;
create table if not exists libros1(
    id bigint unsigned not null auto_increment,
    nombre varchar(100) not null,
    descripcion varchar(450) not null,
    autor varchar (200) not null ,
    editorial varchar(200) not null,
    primary key(id)
    )