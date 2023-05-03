create table user
(
    `id`                    int         NOT NULL AUTO_INCREMENT primary key,
    `name`                  varchar(24) NOT NULL DEFAULT '',
    `password`              varchar(24) NOT NULL DEFAULT '',
    `uid`                   int         NOT NULL DEFAULT 0,

    -- 攻击力
    `atk`                   int         NOT NULL DEFAULT 0,
    -- 防御力
    `def`                   int         NOT NULL DEFAULT 0,
    -- 血量
    `hp`                    int         NOT NULL DEFAULT 0,
    -- 金币
    `gold`                  int         NOT NULL DEFAULT 0,
    -- 钻石
    `diamond`               int         NOT NULL DEFAULT 0,
    -- 经验
    `exp`                   int         NOT NULL DEFAULT 0,
    -- 等级
    `level`                 int         NOT NULL DEFAULT 0,
    -- 暴击率
    `crit`                  int         NOT NULL DEFAULT 0,
    -- 穿甲
    `pierce`                int         NOT NULL DEFAULT 0,
    -- 敏捷
    `agile`                 int         NOT NULL DEFAULT 0,
    `clothes_equipment_id`  int         not null default 0, -- 衣服 id
    `weapon_equipment_id`   int         not null default 0, -- 武器 id
    `shoes_equipment_id`    int         not null default 0, -- 鞋子 id
    `helmet_equipment_id`   int         not null default 0, -- 头盔 id
    `trousers_equipment_id` int         not null default 0, -- 裤子 id
    `bracers_equipment_id`  int         not null default 0  -- 护腕 id
)