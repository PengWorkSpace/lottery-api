CREATE DATABASE IF NOT EXISTS lotteryrecord;
USE lotteryrecord;

CREATE TABLE IF NOT EXISTS `lottery_prizes` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
 `prize` tinyint(4) NOT NULL DEFAULT '0' COMMENT '奖品类型 0-贴纸 1-电话卡 2-手机',
 `total` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '奖品总数',
 `stock` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '奖品剩余数量',
 `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
 `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
 PRIMARY KEY (`id`),
 UNIQUE KEY `uk_prize` (`prize`),
 KEY `update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='奖品列表';

-- 初始化奖品数据
INSERT IGNORE INTO `lottery_prizes` (`prize`, `total`, `stock`)
VALUES (0, 0, 0), (1, 100, 100), (2, 5, 5);

CREATE TABLE IF NOT EXISTS `lottery_records` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
 `phone` varchar(11) NOT NULL DEFAULT '0' COMMENT '手机号',
 `prize_id` int(11) NOT NULL DEFAULT '0' COMMENT 'prize',
 `draw_date` date NOT NULL DEFAULT '1970-01-01' COMMENT '抽奖日期',
 `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
 `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
 PRIMARY KEY (`id`),
 KEY `phone_prize_id_draw_date` (`phone`, `prize_id`, `draw_date`),
 KEY `update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='抽奖记录';


CREATE TABLE IF NOT EXISTS `lottery_users` (
 `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增长主键',
 `phone` bigint(11) unsigned NOT NULL DEFAULT '0' COMMENT '手机号',
 `draw_right` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否可以抽奖 0否 1是',
 `article` varchar(500) NOT NULL DEFAULT '0' COMMENT '文章',
 `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
 `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
 `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除',
 PRIMARY KEY (`id`),
 UNIQUE KEY `uk_phone` (`phone`),
 KEY `update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户报名信息';
 