local _TerminalPlanet = {}
local _mtTerminalPlanet = {__index = _TerminalPlanet} 

function NewTerminalPlanet(terminal)
    local TerminalPlanet = setmetatable({}, _mtTerminalPlanet)
    TerminalPlanet.Env = "/planet"
    TerminalPlanet.ConnetingPlanet = nil
    TerminalPlanet.Terminal = terminal

    return TerminalPlanet
end

function _TerminalPlanet.StartEnv(self, command)
    local commandArr = StringSplit(command, " ")

    local position = {}
    local planet = {}

    planet = GWorld:GetPlanetByPlanetId(GUserSpaceship.LandingPlanetId)
    if TableLength(commandArr) >= 3 then
        position = {X=tonumber(commandArr[2]), Y=tonumber(commandArr[3])}
    elseif nil ~= planet then
        position = planet.Info.Position
    else 
        self.Terminal:ScreenErrMsg("请输入星球坐标")
        return
    end

    if nil == position or nil == position.X or nil == position.Y then
        self.Terminal:ScreenErrMsg(string.format("请输入有效坐标"))
        return
    end

    self.Terminal:ScreenInfoMsg(string.format("连接 星球 %s ...", PointToStr(position)))
    planet = GRadar:GetPlanetOnScreenByPosition(position)
    if nil == planet then
        self.Terminal:ScreenErrMsg(string.format("无法连接星球 %s", PointToStr(position)))
        return
    end

    self.ConnetingPlanet = planet
    self.Terminal.Port:TerminalSetCommandPrefix(string.format("%s> ", planet.Info.Name))
end

function _TerminalPlanet.ExecCommand(self, nodePointer, command)
    local commandArr = StringSplit(command, " ")

    if "info" == commandArr[1] then
        self.Terminal:ScreenInfoMsg(string.format("名称: %s", self.ConnetingPlanet.Info.Name))
        self.Terminal:ScreenInfoMsg(string.format("坐标: %s", PointToStr(self.ConnetingPlanet.Info.Position)))
        self.Terminal:ScreenInfoMsg(string.format("资源: %d", self.ConnetingPlanet.Info.Resource))

    elseif "rename" == commandArr[1] then
        self.ConnetingPlanet:SetName(commandArr[2])
        self.Terminal.Port:TerminalSetCommandPrefix(string.format("%s> ", self.ConnetingPlanet.Info.Name))

    elseif "detail" == commandArr[1] then
        self:ShowPlanetDetail()

    else
        self.Terminal:ScreenErrMsg(string.format("%s %s", self.Terminal.ErrCommandNotExists, command))
    end
end

function _TerminalPlanet.ShowPlanetDetail(self)
    local planet = self.Terminal.CmdExcuter["/planet"].ConnetingPlanet
    if nil == planet then
        return
    end
    StartTabPlanet(planet.Info.PlanetId)
end
