CREATE TABLE `tags` (
    `tag_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `tag_name` varchar(32) NOT NULL DEFAULT '',
    `user_id` int(11) unsigned NOT NULL DEFAULT '0',
    `username` varchar(128) NOT NULL DEFAULT '',
    `count` int(11) NOT NULL DEFAULT '0',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `delete_timestamp` int(11) NOT NULL DEFAULT '0',
    PRIMARY KEY (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `works_tag` (
    `tag_id` int(11) unsigned NOT NULL DEFAULT '0',
    `works_id` int(11) unsigned NOT NULL DEFAULT '0',
    `works_name` varchar(255) NOT NULL DEFAULT '',
    `tag_name` varchar(32) NOT NULL DEFAULT '',
    UNIQUE KEY `idx_works_tag` (`tag_id`,`works_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

ALTER TABLE `chujian_db`.`user` ADD COLUMN `profession` VARCHAR(512) DEFAULT '' NOT NULL AFTER `remark`, ADD COLUMN `charge` VARCHAR(255) DEFAULT '' NOT NULL AFTER `profession`, ADD COLUMN `introduction` VARCHAR(512) NULL AFTER `charge`, CHANGE `create_time` `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP NOT NULL AFTER `introduction`, CHANGE `update_time` `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL AFTER `create_time`, CHANGE `delete_timestamp` `delete_timestamp` INT(11) DEFAULT 0 NOT NULL AFTER `update_time`;