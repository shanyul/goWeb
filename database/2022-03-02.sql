ALTER TABLE `chujian_db`.`works_tag`
    ADD COLUMN `is_delete` tinyint(3) NOT NULL DEFAULT 0 AFTER `tag_name`;