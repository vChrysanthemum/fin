local json = require("json")

_DeathStar = {}
local _mtDeathStar = {__index = _DeathStar} 

function NewDeathStar(buildingCore)
    local DeathStar = setmetatable({}, _mtDeathStar)
    DeathStar.BuildingCore = buildingCore
    DeathStar.BuildingCore.BuildingType = "DeathStar"
    DeathStar.ClientTerminal = nil
    return DeathStar
end

function BuildDeathStar(planet)
    local buildingType = "DeathStar" 
    local buildCost = 2
    local buildingCore = CreateBuildingCore(planet, buildingType, buildCost)
    if "table" ~= type(buildingCore) then
        return buildingCore
    end
    return buildingCore.Building
end

function DestroyDeathStar(planet)
    local buildingType = "DeathStar" 
    return DestroyBuildingCore(planet, buildingType)
end

function _DeathStar.SetClientTerminal(self, clientTerminal)
    self.ClientTerminal = clientTerminal
end

function _DeathStar.ExecCommand(self, command)
    local commandArr = StringSplit(command, " ")

    if "destroy" == commandArr[1] then
        if TableLength(commandArr) < 3 then
            self.ClientTerminal:ScreenErrMsg(string.format("需指定坐标"))
            return
        end

        local position = {X=tonumber(commandArr[2]), Y=tonumber(commandArr[3])}
        local planet = GRadar:GetPlanetOnScreenByPosition(position)
        if nil == planet then
            self.ClientTerminal:ScreenErrMsg(string.format("指定坐标无法探测到星球 %d %d", position.X, position.Y))
            return
        end

        if planet.Info.PlanetId == GUserSpaceship.LandingPlanetId then
            self.ClientTerminal:ScreenErrMsg(string.format("飞船在指定星球上无法摧毁"))
            return
        end

        GWorld:RemovePlanet(planet)
        self.ClientTerminal:ScreenInfoMsg(string.format("摧毁星球完成"))
        GUserSpaceship:SetNewestMsg(string.format("星球 %s 被摧毁", planet.Info.Name))

    else
        self.Terminal:ScreenErrMsg(string.format("%s %s", self.Terminal.ErrCommandNotExists, command))
    end
end
