create table user_equipment
(
    `id`                int not null auto_increment primary key,
    `user_id`           int not null default 0, -- 用户 id
    `equipment_id`      int not null default 0, -- 装备 id
    `enhancement_level` int not null default 0, -- 强化等级
    `is_permanent`      int not null default 0, -- 是否永久
    `end_time`          int not null default 0  -- 结束时间
)