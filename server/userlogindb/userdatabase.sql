DROM DATABASE IF EXISTS user_login;
CREATE DATABASE user_login CHARACTER SET utf8 COLLATE UTF8_GENERAL_CI;

USE user_login;

CREATE TABLE users(
  id INT  unsigned NOT NULL auto_increment,
  email VARCHAR(64) UNIQUE,
  username VARCHAR(64),
  hashedPassword BINARY(32),
  salt BINARY(32),
  PRIMARY KEY(id)
)Engine=InnoDB;

CREATE UNIQUE INDEX email_index 
ON users(email);

