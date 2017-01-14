_Terminal = {}
local _mtTerminal = {__index = _Terminal} 

function NewTerminal(port)
    local Terminal = setmetatable({}, _mtTerminal)
    Terminal.CurrentEnv = "/"

    Terminal.ErrCommandNotExists = "无效命令"
    Terminal.ConnetingPlanet = nil

    local _terminalPlanet = NewTerminalPlanet(Terminal)
    local _terminalSpaceship = NewTerminalSpaceship(Terminal)
    local _terminalRobot = NewTerminalRobot(Terminal)
    local _terminalBuilding = NewTerminalBuilding(Terminal)
    Terminal.CmdExcuter = {}
    Terminal.CmdExcuter[_terminalPlanet.Env] = _terminalPlanet
    Terminal.CmdExcuter[_terminalSpaceship.Env] = _terminalSpaceship
    Terminal.CmdExcuter[_terminalRobot.Env] = _terminalRobot
    Terminal.CmdExcuter[_terminalBuilding.Env] = _terminalBuilding

    Terminal.Port = port

    return Terminal
end

function _Terminal.ScreenErrMsg(self, msg)
    self.Port:TerminalWriteNewLine(string.format("[%s](fg-red)", msg))
end

function _Terminal.ScreenSuccessMsg(self, msg)
    self.Port:TerminalWriteNewLine(string.format("[%s](fg-green)", msg))
end

function _Terminal.ScreenInfoMsg(self, msg)
    self.Port:TerminalWriteNewLine(string.format("[%s](fg-green)", msg))
end
