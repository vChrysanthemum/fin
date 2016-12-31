_Terminal = {}
local _mtTerminal = {__index = _Terminal} 

function NewTerminal()
    local Terminal = setmetatable({}, _mtTerminal)
    Terminal.CurrentEnv = "/"

    Terminal.ErrCommandNotExists = "无效命令"
    Terminal.ConnentingPlanet = nil

    local _terminalPlanet = NewTerminalPlanet()
    local _terminalSpaceship = NewTerminalSpaceship()
    Terminal.CmdExcuter = {}
    Terminal.CmdExcuter[_terminalPlanet.Env] = _terminalPlanet
    Terminal.CmdExcuter[_terminalSpaceship.Env] = _terminalSpaceship

    Terminal.KeyPressSig = NodeTerminalMain:RegisterKeyPressHandler(function(nodePointer, keystr)
        if "<tab>" == keystr then
        end
    end)

    Terminal.CommandSig = NodeTerminalMain:TerminalRegisterCommandHandle(function(nodePointer, command)
        Terminal:ExecCommand(nodePointer, command)
    end)

    return Terminal
end

function _Terminal.ScreenErrMsg(self, msg)
    NodeTerminalMain:TerminalWriteNewLine(string.format("[%s](fg-red)", msg))
end

function _Terminal.ScreenSuccessMsg(self, msg)
    NodeTerminalMain:TerminalWriteNewLine(string.format("[%s](fg-green)", msg))
end

function _Terminal.ScreenInfoMsg(self, msg)
    NodeTerminalMain:TerminalWriteNewLine(string.format("[%s](fg-green)", msg))
end
