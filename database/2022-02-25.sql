ALTER TABLE `chujian_db`.`user`
    ADD COLUMN `sex` tinyint(2) NOT NULL DEFAULT 0 AFTER `nickname`,
ADD COLUMN `country` varchar(128) NOT NULL DEFAULT '' AFTER `state`;