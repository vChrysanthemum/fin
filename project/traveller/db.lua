local database = require("database")
local DB = database.OpenDB("main.db")

local sql = [[
create table if not exists `b_planet` (
planet_id integer primary key not null,
position_x integer,
position_y integer
);

create table if not exists `b_inited_area` (
min_x integer,
min_y integer,
max_x integer,
max_y integer,
primary key(min_x, min_y, max_x, max_y)
);

]]
local ret = DB:Exec(sql)
