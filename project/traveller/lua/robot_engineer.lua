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

    elseif "aboard" == commandArr[1] then
        if nil == self.RobotCore.PlanetLanding then
            self.ClientTerminal:ScreenErrMsg(string.format("机器人不在星球上"))
            return
        end

        if nil == GUserSpaceship.LoginedPlanet or 
            GUserSpaceship.LoginedPlanet.Info.PlanetId ~= self.RobotCore.PlanetLanding.Info.PlanetId then
            self.ClientTerminal:ScreenErrMsg(string.format("飞船不在星球上，机器人无法登船"))
            return
        end

        self:AboardSpaceship()
        self.ClientTerminal:ScreenInfoMsg(string.format("%s登船成功", self.RobotCore.Info.Name))
        RefreshNodeTabPlanetParPlanetInfo()

    elseif "mine" == commandArr[1] then
        if nil ~= self.RobotCore.PlanetLanding then
            self.RobotCore.Info.Action = "mine"
            self.RobotCore:FlushToDB()
            RefreshNodeTabPlanetParPlanetInfo()
            self.ClientTerminal:ScreenInfoMsg(string.format("%s开始采矿", self.RobotCore.Info.Name))
        end

    elseif "build" == commandArr[1] then
        if nil == self.RobotCore.PlanetLanding then
            self.ClientTerminal:ScreenErrMsg(string.format("机器人没有停落在任何星球，无法建造建筑"))
            return
        end

        local buildingType = commandArr[2]
        if "PowerPlant" == buildingType then
            self:BuildPowerPlant()
        end


    elseif "cleanjob" == commandArr[1] then
        self:CleanJob()
        RefreshNodeTabPlanetParPlanetInfo()
        self.ClientTerminal:ScreenInfoMsg(string.format("清空任务完成"))

    else
        self.ClientTerminal:ScreenErrMsg(string.format("%s %s", self.ClientTerminal.ErrCommandNotExists, command))
    end
end

function _RobotEngineer.CleanJob(self)
    if nil == self.RobotCore.Info.Action then
        return
    end
    self.RobotCore.Info.Action = nil
    self.RobotCore:FlushToDB()
end

function _RobotEngineer.AboardSpaceship(self)
    self:CleanJob()
    self.RobotCore.AboardSpaceship()
end

function _RobotEngineer.CollectResourceToPlanet(self)
end

function _RobotEngineer.BuildPowerPlant(self)
    local powerPlant = BuildPowerPlant(self.RobotCore.PlanetLanding)
    if "string" == type(powerPlant) then
        self.ClientTerminal:ScreenErrMsg(string.format("建造失败: %s", powerPlant))
        return
    end
    self.RobotCore.PlanetLanding:RefreshModuleDevelopedBuilding()
    RefreshNodeTabPlanetParPlanetInfo()
    self.ClientTerminal:ScreenInfoMsg(string.format("建造%s完成", powerPlant:GetBuildingTypeCh()))
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
            RefreshNodeTabPlanetParPlanetInfo()
        end
    end
end
