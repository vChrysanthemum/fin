local json = require("json")

local _RobotCore = {}
local _mtRobotCore = {__index = _RobotCore} 

function NewRobotCore()
    local RobotCore = setmetatable({}, _mtRobotCore)
    RobotCore.Info = {
        RobotId         = 0,
        Name            = "",
        ServiceAddress  = "",
        RobotOS       = "",
        LandingPlanetId = nil,
    }
    RobotCore.Robot = nil
    return RobotCore
end

function _RobotCore.FlushToDB(self)
    if "number" ~= type(self.Info.RobotId) then
        return nil
    end

    sql = string.format([[
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
        RobotOS       = robotInfo.RobotOS,
        LandingPlanetId = robotInfo.LandingPlanetId,
    }

    if "engineer" == self.Info.RobotOS then
        self.Robot = NewRobotEngineer(self)
    end

    if "number" == type(self.Info.LandingPlanetId) and self.Info.LandingPlanetId > 0 then
        self.PlanetLanding = GWorld:GetPlanetByPlanetId(self.Info.LandingPlanetId)
    end
end

function _RobotCore.LandingPlanet(self, planet)
    self.PlanetLanding = planet
    self.Info.LandingPlanetId = planet.Info.PlanetId
    self:FlushToDB()
end
