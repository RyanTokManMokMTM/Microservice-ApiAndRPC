CREATE TABLE `user`(
   `id` bigint unsigned NOT NULL AUTO_INCREMENT,
   `name` varchar(255) NOT NULL DEFAULT '' COMMENT 'user name',
   `gender` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT 'user gender',
   `mobile` varchar(255) NOT NULL DEFAULT '' COMMENT 'user mobile number',
   `password` varchar(255) NOT NULL DEFAULT '' COMMENT 'user password',
   `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
   `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
   UNIQUE KEY `idx_mobile_unique` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;