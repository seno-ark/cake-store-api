CREATE TABLE `cakes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `title` varchar(256) CHARACTER SET utf8 NOT NULL DEFAULT '',
  `description` mediumtext,
  `rating` float NOT NULL DEFAULT 0,
  `image` varchar(256) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;