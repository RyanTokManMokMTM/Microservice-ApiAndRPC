CREATE TABLE `pay` (
    `id` bigint unsigned  NOT NULL AUTO_INCREMENT,
    `uid` bigint unsigned  NOT NULL DEFAULT '0' COMMENT  'User ID',
    `oid` bigint unsigned  NOT NULL DEFAULT '0' COMMENT  'Order ID',
    `amount` int(10) unsigned  NOT NULL DEFAULT '0' COMMENT  'Product Amount',
    `source` tinyint(3) unsigned  NOT NULL DEFAULT '0' COMMENT  'Pay method',
    `status` tinyint(3) unsigned  NOT NULL DEFAULT '0' COMMENT  'Pay status',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    PRIMARY KEY(`id`),
    KEY `idx_uid`(`uid`),
    KEY `idx_oid`(`oid`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;
