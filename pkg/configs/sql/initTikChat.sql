USE tiktok_db;

DROP TABLE IF EXISTS `messages`;
CREATE TABLE messages(
    id INT NOT NULL AUTO_INCREMENT,
    from_user_id INT NOT NULL,
    to_user_id INT NOT NULL,
    msg_content VARCHAR(255) NOT NULL,
    is_send TINYINT(1),
    PRIMARY KEY (id),
    KEY (to_user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ;
