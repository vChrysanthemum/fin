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

function _RobotEngineer.LoopEvent(self)
    if nil == self.RobotCore.Info.LandingPlanetId then
        return
    end

    if "mine" == self.RobotCore.Info.Action then
        self:ActionMine()
    end
end

function _RobotEngineer.ExecCommand(self, command)
    local commandArr = StringSplit(command, " ")

    if "info" == commandArr[1] then
        self.ClientTerminal:ScreenInfoMsg(string.format("OS: %s",self.RobotCore.Info.RobotOS))
        if "planet" == self.RobotCore.Info.Location and nil ~= self.RobotCore.Info.LandingPlanetId then
            local planet = GWorld:GetPlanetByPlanetId(self.RobotCore.Info.LandingPlanetId)
            self.ClientTerminal:ScreenInfoMsg(string.format("位于%s",
            planet.Info.Name))
        end
        if nil ~= self.RobotCore.Info.Action then
            self.ClientTerminal:ScreenInfoMsg(string.format("正在%s中", self:GetActionCh()))
        end

    elseif "landing" == commandArr[1] then
        local GUserSpaceshipPlanetLanding = GWorld:GetPlanetByPlanetId(GUserSpaceship.LandingPlanetId)
        if nil == GUserSpaceshipPlanetLanding then
            self.ClientTerminal:ScreenErrMsg(string.format("飞船没有停落在任何星球，机器人无法着陆"))
            return
        end

        GUserSpaceship:SetNewestMsg(string.format("%s着陆星球%s", 
        self.RobotCore.Info.Name, GUserSpaceshipPlanetLanding.Info.Name))

        self.RobotCore:LandingPlanet(GUserSpaceshipPlanetLanding)
        self.ClientTerminal:ScreenInfoMsg(string.format("%s着陆%s成功",
        self.RobotCore.Info.Name, GUserSpaceshipPlanetLanding.Info.Name))
        RefreshNodeTabPlanetParPlanetInfo()

    elseif "aboard" == commandArr[1] then
        if nil == self.RobotCore.Info.LandingPlanetId then
            self.ClientTerminal:ScreenErrMsg(string.format("机器人不在星球上"))
            return
        end

        if nil == GUserSpaceship.LandingPlanetId or 
            GUserSpaceship.LandingPlanetId ~= self.RobotCore.Info.LandingPlanetId then
            self.ClientTerminal:ScreenErrMsg(string.format("飞船不在星球上，机器人无法登船"))
            return
        end

        self:AboardSpaceship()
        self.ClientTerminal:ScreenInfoMsg(string.format("%s登船成功", self.RobotCore.Info.Name))
        RefreshNodeTabPlanetParPlanetInfo()

    elseif "mine" == commandArr[1] then
        if nil ~= self.RobotCore.Info.LandingPlanetId then
            self.RobotCore.Info.Action = "mine"
            self.RobotCore:FlushToDB()
            RefreshNodeTabPlanetParPlanetInfo()
            self.ClientTerminal:ScreenInfoMsg(string.format("%s开始采矿", self.RobotCore.Info.Name))
        end

    elseif "collect" == commandArr[1] then
        local GUserSpaceshipPlanetLanding = GWorld:GetPlanetByPlanetId(GUserSpaceship.LandingPlanetId)
        local RobotCorePlanetLanding = GWorld:GetPlanetByPlanetId(self.RobotCore.Info.LandingPlanetId)

        if nil == self.RobotCore.Info.LandingPlanetId or 
            nil == GUserSpaceshipPlanetLanding or 
            RobotCorePlanetLanding.Info.PlanetId ~= GUserSpaceshipPlanetLanding.Info.PlanetId then
            self.ClientTerminal:ScreenErrMsg(string.format("无法收集"))
            return
        end

        if "resource" == commandArr[2] then
            self:CollectResourceToPlanet(commandArr[3])
            RefreshNodeTabPlanetParPlanetInfo()
        else 
            self.ClientTerminal:ScreenErrMsg(string.format("需指定收集对象"))
        end


    elseif "build" == commandArr[1] then
        if nil == self.RobotCore.Info.LandingPlanetId then
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
    self.RobotCore:AboardSpaceship()
end

function _RobotEngineer.CollectResourceToPlanet(self, resourceNum)
    resourceNum = tonumber(resourceNum)
    if nil == resourceNum or resourceNum <= 0 then
        self.ClientTerminal:ScreenErrMsg("需输入正确资源数量")
        return
    end

    local RobotCorePlanetLanding = GWorld:GetPlanetByPlanetId(self.RobotCore.Info.LandingPlanetId)
    if nil == RobotCorePlanetLanding then
        self.ClientTerminal:ScreenErrMsg("星球不存在")
        return
    end

    local ret = RobotCorePlanetLanding:ChangeDevelopedResource(-1*resourceNum)
    if true ~= ret then
        self.ClientTerminal:ScreenErrMsg(string.format(ret))
        return
    end

    ret = GUserSpaceship:ChangeCabinResource(resourceNum)
    if true ~= ret then
        self.ClientTerminal:ScreenErrMsg(string.format(ret))
        return
    end

    self.ClientTerminal:ScreenInfoMsg("收集资源完成")
end

function _RobotEngineer.BuildPowerPlant(self)
    local RobotCorePlanetLanding = GWorld:GetPlanetByPlanetId(self.RobotCore.Info.LandingPlanetId)
    if nil == RobotCorePlanetLanding then
        self.ClientTerminal:ScreenErrMsg("星球不存在")
        return
    end

    local powerPlant = BuildPowerPlant(RobotCorePlanetLanding)

    if "string" == type(powerPlant) then
        self.ClientTerminal:ScreenErrMsg(string.format("建造失败: %s", powerPlant))
        return
    end
    RobotCorePlanetLanding:RefreshModuleDevelopedBuilding()
    RefreshNodeTabPlanetParPlanetInfo()
    self.ClientTerminal:ScreenInfoMsg(string.format("建造%s完成", powerPlant:GetBuildingTypeCh()))
end

-- 挖矿
function _RobotEngineer.ActionMine(self)
    local RobotCorePlanetLanding = GWorld:GetPlanetByPlanetId(self.RobotCore.Info.LandingPlanetId)
    if nil == RobotCorePlanetLanding then
        self.ClientTerminal:ScreenErrMsg("星球不存在")
        return
    end

    if RobotCorePlanetLanding.Info.Resource <= 0 then
        self.RobotCore.Info.Action = nil
        self.RobotCore:FlushToDB()
    else
        RobotCorePlanetLanding:MineByRobot()
        RefreshNodeTabPlanetParPlanetInfo()
    end
end
