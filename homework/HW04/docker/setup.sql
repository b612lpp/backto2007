use blog;
CREATE TABLE posts(ID int, Title TEXT, Tetxt TEXT);
INSERT INTO
  posts
values(1, "Title 1", "Text 1");
ALTER USER 'root' @'%' IDENTIFIED WITH mysql_native_password BY 'password';