DROP DATABASE IF EXISTS Sessions;
CREATE DATABASE user_login CHARACTER SET utf8 COLLATE UTF8_GENERAL_CI;

USE user_login;

CREATE TABLE user_sessions(
 SessionKey text primary key,
 UserEmail VARCHAR(64) not null, 
 LoginTime DATETIME not null,
 LastSeenTime DATETIME not null
)Engine=InnoDB;

