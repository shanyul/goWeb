ALTER TABLE `chujian_db`.`user` ADD COLUMN `union_id` varchar(128) NOT NULL DEFAULT '' AFTER `session_key`;