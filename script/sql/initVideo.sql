USE tiktok_db;

CREATE TABLE video(
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

CREATE TABLE comment(
    id INT NOT NULL AUTO_INCREMENT,
    vedio_id INT NOT NULL,
    user_id INT NOT NULL,
    content VARCHAR(255) NOT NULL,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    KEY (vedio_id)
);

CREATE TABLE user_favourite_video(
    id INT NOT NULL AUTO_INCREMENT,
    vedio_id INT NOT NULL,
    user_id INT NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY favorite_relation(vedio_id, user_id)
);