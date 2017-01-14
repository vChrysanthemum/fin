local json = require("json")

local _RobotCore = {}
local _mtRobotCore = {__index = _RobotCore} 

function NewRobotCore()
    local RobotCore = setmetatable({}, _mtRobotCore)
    RobotCore.Info = {
        RobotId         = 0,
        Name            = "",
        ServiceAddress  = "",
        RobotOS         = "",
        Location        = "",
        LandingPlanetId = nil,
        Action          = nil,
    }
    RobotCore.Robot = nil
    return RobotCore
end

function _RobotCore.FlushToDB(self)
    if "number" ~= type(self.Info.RobotId) then
        return nil
    end

    local sql = string.format([[
    update b_robot set data = '%s' where robot_id=%d
    ]], DB:QuoteSQL(json.encode(self.Info)), self.Info.RobotId)
    local queryRet = DB:Exec(sql)
    if "string" == type(queryRet) then
        Log(queryRet)
    end
end

function _RobotCore.Format(self, robotInfo, robot_id)
    self.Info = {
        RobotId         = tonumber(robot_id),
        Name            = robotInfo.Name,
        ServiceAddress  = robotInfo.ServiceAddress,
        RobotOS         = robotInfo.RobotOS,
        Location        = robotInfo.Location,
        LandingPlanetId = robotInfo.LandingPlanetId,
        Action          = robotInfo.Action,
    }

    if "Engineer" == self.Info.RobotOS then
        self.Robot = NewRobotEngineer(self)
    end
end

function _RobotCore.LandingPlanet(self, planet)
    self.Info.Location = "planet"
    self.Info.LandingPlanetId = planet.Info.PlanetId
    self:FlushToDB()
end

function _RobotCore.AboardSpaceship(self, planet)
    self.Info.Location = nil
    self.Info.LandingPlanetId = nil
    self:FlushToDB()
end
