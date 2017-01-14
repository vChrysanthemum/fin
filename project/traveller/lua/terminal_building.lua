local _TerminalBuilding = {}
local _mtTerminalBuilding = {__index = _TerminalBuilding} 

function NewTerminalBuilding(terminal)
    local TerminalBuilding = setmetatable({}, _mtTerminalBuilding)
    TerminalBuilding.Env = "/building"
    TerminalBuilding.ConnentingBuilding = nil
    TerminalBuilding.Terminal = terminal

    return TerminalBuilding
end

function _TerminalBuilding.StartEnv(self, command)
    local commandArr = StringSplit(command, " ")

    local position = {}
    if nil == GUserSpaceship.LandingPlanetId then
        self.Terminal:ScreenErrMsg("飞船未登录星球，无法连接建筑")
        return false
    end 
    buildingType = commandArr[2]

    self.Terminal:ScreenInfoMsg(string.format("连接 %s ...", buildingType))
    local building = GBuildingCenter:GetBuildingByPlanetIdAndBuildingType(GUserSpaceship.LandingPlanetId, buildingType)
    if nil == building then
        self.Terminal:ScreenErrMsg(string.format("%s不存在", buildingType))
        return false
    end

    self.ConnentingBuilding = building
    self.ConnentingBuilding:SetClientTerminal(self.Terminal)
    self.Terminal.Port:TerminalSetCommandPrefix(string.format("%s> ", GDictE2C[building.BuildingCore.BuildingType]))
end

function _TerminalBuilding.ExecCommand(self, nodePointer, command)
    self.ConnentingBuilding:ExecCommand(command)
end
