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

function _RobotEngineer.GetActionCh(self)
    if "mine" == self.RobotCore.Info.Action then
        return "挖矿"
    end
    return self.RobotCore.Info.Action
end

function _RobotEngineer.ExecCommand(self, command)
    local commandArr = StringSplit(command, " ")

    if "info" == commandArr[1] then
        self.ClientTerminal:ScreenInfoMsg(string.format("OS: %s",self.RobotCore.Info.RobotOS))
        if "planet" == self.RobotCore.Info.Location and nil ~= self.RobotCore.PlanetLanding then
            self.ClientTerminal:ScreenInfoMsg(string.format("位于%s",
            self.RobotCore.PlanetLanding.Info.Name))
        end
        if nil ~= self.RobotCore.Info.Action then
            self.ClientTerminal:ScreenInfoMsg(string.format("正在%s中", self:GetActionCh()))
        end

    elseif "landing" == commandArr[1] then
        if nil == GUserSpaceship.LoginedPlanet then
            self.ClientTerminal:ScreenErrMsg(string.format("飞船没有停落在任何星球，机器人无法着陆"))
            return
        end

        GUserSpaceship:SetNewestMsg(string.format("%s着陆星球%s", 
        self.RobotCore.Info.Name, GUserSpaceship.LoginedPlanet.Info.Name))

        self.RobotCore:LandingPlanet(GUserSpaceship.LoginedPlanet)
        self.ClientTerminal:ScreenInfoMsg(string.format("%s着陆%s成功",
        self.RobotCore.Info.Name, self.RobotCore.PlanetLanding.Info.Name))
        RefreshNodeTabPlanetParPlanetInfo()

    elseif "mine" == commandArr[1] then
        if nil ~= self.RobotCore.PlanetLanding then
            self.ClientTerminal:ScreenInfoMsg(string.format("%s开始采矿",
            self.RobotCore.Info.Name))
            self.RobotCore.Info.Action = "mine"
            self.RobotCore:FlushToDB()
            RefreshNodeTabPlanetParPlanetInfo()
        end

    else
        self.ClientTerminal:ScreenErrMsg(string.format("%s %s", self.ClientTerminal.ErrCommandNotExists, command))
    end
end

function _RobotEngineer.LoopEvent(self)
    if nil == self.RobotCore.PlanetLanding then
        return
    end

    if "mine" == self.RobotCore.Info.Action then
        if self.RobotCore.PlanetLanding.Info.Resource <= 0 then
            self.RobotCore.Info.Action = nil
            self.RobotCore:FlushToDB()
        else
            self.RobotCore.PlanetLanding:MineByRobot()
        end
    end
end
