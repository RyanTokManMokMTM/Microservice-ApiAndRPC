CREATE TABLE `product`(
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL DEFAULT  '' COMMENT 'product name',
    `desc` varchar(255) NOT NULL DEFAULT  '' COMMENT 'product description',
    `stock` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'product stock',
    `amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'amount of product',
    `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT 'product status',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
)  ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;