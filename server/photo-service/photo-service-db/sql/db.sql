DROP DATABASE IF EXISTS Photos;
CREATE DATABASE Photos CHARACTER SET utf8 COLLATE UTF8_GENERAL_CI;

USE Photos;

CREATE TABLE Profile(
Id INT  unsigned NOT NULL auto_increment,
 UserEmail VARCHAR(64) not null, 
 TimeStmp DATETIME  DEFAULT NOW(),
 Url VARCHAR(128) not null,
 Selected BOOLEAN DEFAUlT TRUE,
  PRIMARY KEY(Id)

)Engine=InnoDB;

CREATE  INDEX email_index 
ON Profile(UserEmail);

CREATE TABLE City(
Id INT  unsigned NOT NULL auto_increment,
 CityId INT not null, 
 TimeStmp DATETIME  DEFAULT NOW(),
 Url VARCHAR(128) not null,
 Selected BOOLEAN DEFAUlT TRUE,
  PRIMARY KEY(Id)

)Engine=InnoDB;

CREATE INDEX CityId_index 
ON City(CityId);

CREATE TABLE Place(
Id INT  unsigned NOT NULL auto_increment,
 PlaceId INT not null, 
 TimeStmp DATETIME  DEFAULT NOW(),
 Url VARCHAR(128) not null,
 Selected BOOLEAN DEFAUlT TRUE,
  PRIMARY KEY(Id)

)Engine=InnoDB;

CREATE INDEX PlaceId_index 
ON Place(PlaceId);

CREATE TABLE Post(
Id INT  unsigned NOT NULL auto_increment,
 PostId VARCHAR(64) not null, 
 TimeStmp DATETIME  DEFAULT NOW(),
 Url VARCHAR(128) not null,
 Selected BOOLEAN DEFAUlT TRUE,
  PRIMARY KEY(Id)

)Engine=InnoDB;


CREATE  INDEX PostId_index 
ON Post(PostId);

insert into Profile(UserEmail,Url) values("my@gmail.com","http:/one-the.two-tr/12344543-1232331-12312-123123-123213-123213.png")