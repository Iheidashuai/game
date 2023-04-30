create table user
(
    `id`       int         NOT NULL AUTO_INCREMENT,
    `name`     varchar(24) NOT NULL DEFAULT '',
    `password` varchar(24) NOT NULL DEFAULT '',
    `uid`      int         NOT NULL DEFAULT '',
)