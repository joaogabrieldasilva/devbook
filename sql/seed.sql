INSERT INTO users (name, username, email, password)
VALUES
('John Doe', 'johndoe123', 'johndoe@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Jane Smith', 'janesmith456', 'janesmith@example.com',  "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Alice Johnson', 'alicej', 'alice@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Michael Brown', 'mikebrown789', 'mike@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Emily Davis', 'emilyd', 'emily@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Christopher Wilson', 'chriswilson', 'chris@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Jessica Martinez', 'jessmartinez', 'jessica@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('David Anderson', 'davidanderson', 'david@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Sarah Thompson', 'sarah.t', 'sarah@example.com', "$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG"),
('Matthew Lee', 'mattlee', 'matt@example.com', '$2a$10$eLCfxl6J5hPr5ZMTwMUIw.TxHH8tzQtXEu7yuUDI3SQGcZnMGuFCG');


INSERT INTO followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3),
(2, 4),
(4, 5),
(6, 1);