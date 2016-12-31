_RobotEngineer = {}
local _mtRobotEngineer = {__index = _RobotEngineer} 

function NewRobotEngineer(robotCore)
    local RobotEngineer = setmetatable({}, _mtRobotEngineer)
    RobotEngineer.RobotCore = robotCore
    RobotEngineer.ClientTerminal = nil
    return RobotEngineer
end

function _RobotEngineer.SetClientTerminal(self, clientTerminal)
    self.ClientTerminal = clientTerminal
end

function _RobotEngineer.ExecCommand(self, command)
    local commandArr = StringSplit(command, " ")

    if "info" == commandArr[1] then
        self.ClientTerminal:ScreenInfoMsg(string.format("OS: %s",self.RobotCore.Info.RobotOS))

    elseif "landing" == commandArr[1] then
        if nil == GUserSpaceship.LoginedPlanet then
            self.ClientTerminal:ScreenErrMsg(string.format("飞船没有停落在任何星球，机器人无法着陆"))
            return
        end

        GUserSpaceship:SetNewestMsg(string.format("%s着陆星球%s", 
        self.RobotCore.Info.Name, GUserSpaceship.LoginedPlanet.Info.Name))

        self.RobotCore:LandingPlanet(GUserSpaceship.LoginedPlanet)
        self.ClientTerminal:ScreenErrMsg(string.format("%s着陆%s成功",
        self.RobotCore.Info.Name, self.RobotCore.PlanetLanding.Info.Name))

    elseif "location" then
        if "number" == type(self.RobotCore.Info.LandingPlanetId) then
            if nil ~= self.RobotCore.LandingPlanet then
                self.ClientTerminal:ScreenInfoMsg(string.format("%s位于%s",
                self.RobotCore.Info.Name, self.RobotCore.LandingPlanet.Info.Name))
            end
        end

    else
        self.ClientTerminal:ScreenErrMsg(string.format("%s %s", self.ClientTerminal.ErrCommandNotExists, command))
    end
end
