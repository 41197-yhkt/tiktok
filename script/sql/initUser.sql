USE tiktok_db;

-- user_name作为登录的唯一账号
CREATE TABLE user(
    id INT NOT NULL AUTO_INCREMENT,
    user_name VARCHAR(20) NOT NULL,
    user_pwd_hash  VARCHAR(32) NOT NULL,
    follow_count INT,
    follower_count INT,
    PRIMARY KEY (id),
    UNIQUE(user_name)
);

-- 考虑到用户改名，所以follow用id定位用户
-- 增加单向唯一约束，即'用户A只能follow用户B一次'
-- 传入时应用层不需排序
-- 未设置外键
CREATE TABLE user_follow_relation(
    id INT NOT NULL AUTO_INCREMENT,
    follow_from INT NOT NULL,
    follow_to INT NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY follow_relation (follow_from, follow_to)
);

-- friend是双向的，但是为了减少索引，应用层将user_a与user_b按大小排序之后再传入
CREATE TABLE user_friend_relation(
    id INT NOT NULL AUTO_INCREMENT,
    user_a INT NOT NULL,
    user_b INT NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY friend_relation (user_a, user_b)
);

