USE tiktok_db;

CREATE TABLE vedio(
    id INT NOT NULL AUTO_INCREMENT,
    author INT NOT NULL,    -- id的话还需要去user里查询，太麻烦了，直接user_name?
    play_url VARCHAR(255) NOT NULL,
    cover_url VARCHAR(255) NOT NULL,
    favorite_count INT DEFAULT 0,
    comment_count INT DEFAULT 0,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(50) NOT NULL,
    PRIMARY KEY (id),
    KEY (author)
);

DROP TABLE IF EXISTS `comments`;
CREATE TABLE comments(
    id INT NOT NULL AUTO_INCREMENT,
    vedio_id INT NOT NULL,
    user_id INT NOT NULL,
    content VARCHAR(255) NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (id),
    KEY (vedio_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ;

DROP TABLE IF EXISTS `user_favourites`;
CREATE TABLE user_favourites(
    id INT NOT NULL AUTO_INCREMENT,
    vedio_id INT NOT NULL,
    user_id INT NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY favorite_relation(vedio_id, user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ;