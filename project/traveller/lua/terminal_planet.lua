local _TerminalPlanet = {}
local _mtTerminalPlanet = {__index = _TerminalPlanet} 

function NewTerminalPlanet()
  local TerminalPlanet = setmetatable({}, _mtTerminalPlanet)
  TerminalPlanet.Env = "/planet"
  TerminalPlanet.ConnentingPlanet = nil

  return TerminalPlanet
end

function _TerminalPlanet.StartEnv(self, command)
  local commandArr = StringSplit(command, " ")

  local position = {}
  if nil ~= GUserSpaceship.LoginedPlanet then
    position = GUserSpaceship.LoginedPlanet.Info.Position
  else
    if TableLength(commandArr) < 3 then
      GTerminal:ScreenErrMsg("请输入星球坐标")
      return
    end
    position = {X=tonumber(commandArr[2]), Y=tonumber(commandArr[3])}
  end

  if nil == position or nil == position.X or nil == position.Y then
    GTerminal:ScreenErrMsg(string.format("请输入有效坐标"))
    return
  end

  GTerminal:ScreenInfoMsg(string.format("连接 星球 %s ...", PointToStr(position)))
  local planet = GRadar:GetPlanetOnScreenByPosition(position)
  if nil == planet then
    GTerminal:ScreenErrMsg(string.format("无法连接星球 %s", PointToStr(position)))
    return
  end

  self.ConnentingPlanet = planet
  NodeTerminalMain:TerminalSetCommandPrefix(string.format("%s> ", planet.Info.Name))
end

function _TerminalPlanet.ExecCommand(self, nodePointer, command)
  local commandArr = StringSplit(command, " ")

  if "info" == commandArr[1] then
    GTerminal:ScreenInfoMsg(string.format("名称: %s", self.ConnentingPlanet.Info.Name))
    GTerminal:ScreenInfoMsg(string.format("坐标: %s", PointToStr(self.ConnentingPlanet.Info.Position)))

  elseif "rename" == commandArr[1] then
    self.ConnentingPlanet:SetName(commandArr[2])
    NodeTerminalMain:TerminalSetCommandPrefix(string.format("%s> ", self.ConnentingPlanet.Info.Name))

  elseif "detail" == commandArr[1] then
    self:ShowPlanetDetail()

  else
    GTerminal:ScreenErrMsg(string.format("%s %s", GTerminal.ErrCommandNotExists, command))
  end
end

function _TerminalPlanet.ShowPlanetDetail(self)
  GWorld:Stop(function()
    local planet = GTerminal.CmdExcuter["/planet"].ConnentingPlanet
    if nil == planet then
      return
    end
    StartTabPlanet(planet)
  end)
end
