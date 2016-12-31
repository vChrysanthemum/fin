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

create table if not exists `b_robot` (
robot_id integer primary key not null,
data text
);

]]
local ret = DB:Exec(sql)

local spaceship = NewSpaceship()
sql = string.format([[
insert into b_spaceship (spaceship_id, data) values (1, '%s');
]], json.encode(spaceship.Info))
ret = DB:Exec(sql)

local robot = NewRobotCore()
robot.Info.RobotOS = "engineer"

robot.Info.RobotId = 1
robot.Info.Name = "黄鹂"
robot.Info.ServiceAddress = "a1"
sql = string.format([[
insert into b_robot (robot_id, data) values (%d, '%s');
]], robot.Info.RobotId, json.encode(robot.Info))
ret = DB:Exec(sql)

robot.Info.RobotId = 2
robot.Info.Name = "大象"
robot.Info.ServiceAddress = "a2"
sql = string.format([[
insert into b_robot (robot_id, data) values (%d, '%s');
]], robot.Info.RobotId, json.encode(robot.Info))
ret = DB:Exec(sql)
