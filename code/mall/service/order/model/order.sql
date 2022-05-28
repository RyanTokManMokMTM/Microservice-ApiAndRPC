CREATE TABLE `order`(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'User ID',
    `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'Product ID',
    `amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'Order amount',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT 'order status',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    PRIMARY KEY(`id`),
    KEY `idx_uid`(`uid`),
    KEY `idx_pid`(`pid`)
);