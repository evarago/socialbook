CREATE DATABASE IF NOT EXISTS socialbook;

USE socialbook;

DROP TABLE IF EXISTS Usuarios;

CREATE TABLE Usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criacao timestamp default current_timestamp()
) ENGINE=INNODB;


USE socialbook;

DROP TABLE IF EXISTS Seguidores;

CREATE TABLE Seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id)
    REFERENCES Usuarios(id)
    ON DELETE CASCADE,
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES Usuarios(id)
    ON DELETE CASCADE,
    criacao timestamp default current_timestamp(),
    PRIMARY KEY (usuario_id, seguidor_id) 
) ENGINE=INNODB;


USE socialbook;

DROP TABLE IF EXISTS Publicacoes;

CREATE TABLE Publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES Usuarios(id)
    ON DELETE CASCADE,
    curtidas int default 0,
    criacao timestamp default current_timestamp()
) ENGINE=INNODB;