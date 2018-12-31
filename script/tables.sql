create table if not exists `thing` (
    `_id` INT UNSIGNED AUTO_INCREMENT,
    `title` VARCHAR(128),
    `create_time` int,
    `breif` VARCHAR(256),
    `images` VARCHAR(2048),
    `publish_time` int,
    `source` VARCHAR(64),
    `type` VARCHAR(64),
    `origin_url` VARCHAR(128),
    `author` VARCHAR(128),
    `avatar` VARCHAR(128) DEFAULT "",
    `commmets` int(11) DEFAULT 0,
    `stars` int(10) DEFAULT 0,
    PRIMARY KEY (`_id`, `title`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

create table if not exists `index_data` (
    `_id` INT UNSIGNED AUTO_INCREMENT,
    `banner_url_list` VARCHAR(1024) DEFAULT "",
    `banner_url_target` VARCHAR(1024) DEFAULT "",
    `channel_icon_list` VARCHAR(1024) DEFAULT "",
    `channel_icon_target` VARCHAR(1024) DEFAULT "",
    `thing_ids` VARCHAR(32),
    PRIMARY KEY (`_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
