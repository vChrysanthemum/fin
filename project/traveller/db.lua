local json = require("json")
local database = require("database")
DB = database.OpenDB("main.db")

local sql = [[
create table if not exists `b_planets_block` (
x integer,
y integer,
created_at int,
primary key(x, y)
);

create table if not exists `b_planet` (
planet_id integer primary key not null,
planets_block_x integer,
planets_block_y integer,
data text
);

create table if not exists `b_spaceship` (
spaceship_id integer primary key not null,
data text
);

]]
local ret = DB:Exec(sql)

local spaceship = {
    Name      = "鹦鹉螺号",
    Position  = {X = 0.0, Y = 0.0},
    Speed     = {X = 0.02, Y = 0.03},
    Character = "x",
    ColorFg   = "blue"
}
sql = string.format([[
insert into b_spaceship (spaceship_id, data) values (1, '%s');
]], json.encode(spaceship))
DB:Exec(sql)
