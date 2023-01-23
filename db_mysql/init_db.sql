CREATE TABLE `user`
(
    `id`       int          NOT NULL AUTO_INCREMENT,
    `name`     varchar(24)  NOT NULL DEFAULT '',
    `uid`      int          NOT NULL DEFAULT '0',
    `password` varchar(256) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB