USE tiktok_db;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         `deleted_at` datetime(3) DEFAULT NULL,
                         `user_name` longtext,
                         `user_pwd_hash` longtext,
                         `follow_count` bigint DEFAULT NULL,
                         `follower_count` bigint DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;


-- 考虑到用户改名，所以follow用id定位用户
-- 传入时应用层不需排序
-- 未设置外键
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_relations
-- ----------------------------
DROP TABLE IF EXISTS `user_relations`;
CREATE TABLE `user_relations` (
                                  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
                                  `created_at` datetime(3) DEFAULT NULL,
                                  `updated_at` datetime(3) DEFAULT NULL,
                                  `deleted_at` datetime(3) DEFAULT NULL,
                                  `follow_from` bigint unsigned DEFAULT NULL,
                                  `follow_to` bigint unsigned DEFAULT NULL,
                                  PRIMARY KEY (`id`),
                                  KEY `idx_user_relations_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;

