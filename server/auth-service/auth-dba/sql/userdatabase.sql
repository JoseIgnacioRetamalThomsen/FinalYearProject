DROP DATABASE IF EXISTS UserAuth;
CREATE DATABASE UserAuth CHARACTER SET utf8 COLLATE UTF8_GENERAL_CI;

USE UserAuth;

CREATE TABLE Users(
  Id INT  unsigned NOT NULL auto_increment,
  Email VARCHAR(64) UNIQUE,
  PasswordHash BINARY(32),
  PasswordSalt BINARY(32),
  IsEmail BOOLEAN DEFAULT FALSE,
  PRIMARY KEY(id)
)Engine=InnoDB;

CREATE UNIQUE INDEX email_index 
ON Users(Email);

CREATE TABLE UserSessions(
 SessionKey VARCHAR(64),
 Email VARCHAR(64) not null, 
 LoginTime DATETIME not null,
 LastSeenTime DATETIME not null,
PRIMARY KEY(SessionKey)
)Engine=InnoDB;
