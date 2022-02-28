ALTER TABLE `chujian_db`.`user`
    ADD COLUMN `wechat_code` varchar(128) NOT NULL DEFAULT '' AFTER `bg_image`;