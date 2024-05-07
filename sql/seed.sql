INSERT INTO users (name, username, email, password)
VALUES
('John Doe', 'johndoe123', 'johndoe@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Jane Smith', 'janesmith456', 'janesmith@example.com',  "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Alice Johnson', 'alicej', 'alice@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG");


INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3);


INSERT INTO posts (title, content, author_id)
VALUES
("Post 1", "Post 1 Content", 1),
("Post 2", "Post 2 Content", 1),
("Post 3", "Post 3 Content", 2),
("Post 4", "Post 4 Content", 3),
("Post 5", "Post 5 Content", 3);


