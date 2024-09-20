insert into Usuarios (nome, nick, email, senha) values 
("Usuario 1", "User1", "user1@gmail.com", "$2a$10$nsRKkdmAt1X/vxgbGKcQ0uAferru7z6qGXi/FxoN82qAUsxLidghy"),
("Usuario 2", "User2", "user2@gmail.com", "$2a$10$nsRKkdmAt1X/vxgbGKcQ0uAferru7z6qGXi/FxoN82qAUsxLidghy"),
("Usuario 3", "User3", "user3@gmail.com", "$2a$10$nsRKkdmAt1X/vxgbGKcQ0uAferru7z6qGXi/FxoN82qAUsxLidghy")


insert into Seguidores (usuario_id, seguidor_id) values 
(1,2),
(3,1),
(1,3)

insert into Publicacoes (titulo, conteudo, autor_id) values
("Publicação do Usuário 1", "Essa é a publicação do usuário 1!", 1),
("Publicação do Usuário 2", "Essa é a publicação do usuário 2!", 2),
("Publicação do Usuário 3", "Essa é a publicação do usuário 3!", 3);
