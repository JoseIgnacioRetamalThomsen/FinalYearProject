insert into users (email, username, hashedpassword, salt) values("email9","name","hashed","salt");
GRANT ALL ON user_login TO 'test'@'localhost';
select hashedPassword, salt from users where email =  "email1"