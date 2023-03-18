
CREATE TABLE users (
	user_id serial PRIMARY KEY,
	username VARCHAR ( 50 ) UNIQUE NOT NULL,
    email VARCHAR ( 255 ) UNIQUE NOT NULL,
	password VARCHAR ( 50 ) NOT NULL
	
);

INSERT INTO users(user_id,username,password,email)  VALUES
 (1,'Carlos Silva ', 'testando123','carlos.silva@gmail.com'),
 (2,'Monica Araujo', 'testando456','monica.araujo@gmail.com'),
 (3,'André Bonafonte', 'testando789','andré.bonafonte@outlook.com');


CREATE TABLE course (
	id serial PRIMARY KEY,
    description VARCHAR ( 50 ) NOT NULL,
	course VARCHAR ( 150 ) NOT NULL,
    department  VARCHAR ( 50 ) NOT NULL,
    duration VARCHAR ( 15 ) NOT NULL
	
);
INSERT INTO course(id, description, course, department, duration)  VALUES
 (1,'Curso de redes e protocolos', 'redes de computadores','TI','48h'),
 (2,'Curso de manutenção e suporte', 'Curso de manutenção e suporte','TI','128h'),
 (3,'Curso de Programação', 'Lógica da programação','TI','300h'),
 (4,'Curso de secretariado ', 'Curso de secretariado','administração','150h'),
 (5,'Curso de excel', 'Excel','TI','48h'),
 (6,'Curso de costura', 'Curso de costura','Industrial','48h');
