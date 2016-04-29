CREATE DATABASE `dz`;
USE dz;
CREATE TABLE `users`(
	id VARCHAR(20) NOT NULL, --mobile number
	mobile VARCHAR(20) NOT NULL,
	name VARCHAR(36) NOT NULL,
	gender ENUME('male','female') NOT NULL,
	email VARCHAR(36) NOT NULL,
	birthday VARCHAR(36) NOT NULL,
	password VARCHAR(100) NOT NULL,
	status TINYINT(1) NOT NULL DEFAULT 0,
	PRIMARY KEY(`id`),
	KEY `users_name` (`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `permissions`(
	name VARCHAR(36) NOT NULL, --英文名
	value  VARCHAR(36) NOT NULL, --中文名
	description  VARCHAR(200) NOT NULL DEFAULT "",  --解释
	PRIMARY KEY(`name`),
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `roles`(
	name VARCHAR(36) NOT NULL,
	value  VARCHAR(36) NOT NULL,
	description  VARCHAR(200) NOT NULL DEFAULT "",  --解释
	PRIMARY KEY(`name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `role_permissions`(
	id INT(36) UNSIGNED NOT NULL AUTO_INCREMENT,
	role_name VARCHAR(36) NOT NULL,
	permission_name  VARCHAR(36) NOT NULL,
	PRIMARY KEY (`id`),
	KEY `role_permission_role_name` (`role_name`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_roles`(
	id INT(36) UNSIGNED NOT NULL AUTO_INCREMENT,
	user_id VARCHAR(20) NOT NULL,
	role_name VARCHAR(36) NOT NULL,
	PRIMARY KEY (`id`),
	KEY `user_role_user_id` (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `signups`(
	id INT(36) UNSIGNED NOT NULL AUTO_INCREMENT,
	mobile VARCHAR(20) NOT NULL,
	name VARCHAR(36) NOT NULL,
	gender ENUME('male','female') NOT NULL,
	email VARCHAR(36) NOT NULL,
	birthday VARCHAR(36) NOT NULL,
	address VARCHAR(100) NOT NULL,
	province VARCHAR(100) NOT NULL,
	city VARCHAR(100) NOT NULL,
	when VARCHAR(30) NOT NULL,
	PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `classes`(
	id VARCHAR(100) NOT NULL, --nanjing20161201
	address VARCHAR(100) NOT NULL, --详细地址
	province VARCHAR(100) NOT NULL,--省份
	city VARCHAR(100) NOT NULL,  --城市
	start datetime  NOT NULL,  --开始时间
	end datetime NOT NULL,  --结束时间
	contact_mobile VARCHAR(20) NOT NULL,--联系人电话
	contact_name VARCHAR(20) NOT NULL,--联系人姓名
	notice TEXT NOT NULL,
	PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `posts`(
	id INT(36) UNSIGNED NOT NULL,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
