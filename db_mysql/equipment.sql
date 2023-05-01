create table equipment
(
    `id` int not null auto_increment,
    `equipment_id` int not null default 0, -- 装备 id
    `atk` int not null default 0, -- 攻击力
    `def` int not null default 0, -- 防御力
    `hp` int not null default 0, -- 血量
    `level` int not null default 0, -- 等级
    `crit` int not null default 0, -- 暴击率
    `pierce` int not null default 0, -- 穿甲
    `agile` int not null default 0, -- 敏捷
    -- 品质
    `quality` int not null default 0,
    -- 类型
    `type` int not null default 0,
)