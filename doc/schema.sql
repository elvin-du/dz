CREATE DATABASE `dz`;

USE dz;

CREATE TABLE `users`(
	id VARCHAR(36) NOT NULL,
	name VARCHAR(36) NOT NULL,
	PRIMARY KEY(`id`),
	KEY `users_name` (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


INSERT `users` SET id='1',name='elvin';