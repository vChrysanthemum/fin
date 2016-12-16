local json = require("json")
local database = require("database")
DB = database.OpenDB("main/main.db")

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

local spaceship = NewSpaceshipInfo()
sql = string.format([[
insert into b_spaceship (spaceship_id, data) values (1, '%s');
]], json.encode(spaceship))
ret = DB:Exec(sql)
