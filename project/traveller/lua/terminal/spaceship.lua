local _TerminalSpaceship = {}
local _mtTerminalSpaceship = {__index = _TerminalSpaceship} 

function NewTerminalSpaceship()
    local TerminalSpaceship = setmetatable({}, _mtTerminalSpaceship)
    TerminalSpaceship.Env = "/spaceship"
    TerminalSpaceship.Spaceship = nil

    return TerminalSpaceship
end

function _TerminalSpaceship.StartEnv(self, command)
    GTerminal:ScreenInfoMsg("连接 飞船 ...")
    self.Spaceship = GUserSpaceship
    NodeTerminalMain:TerminalSetCommandPrefix(string.format("%s> ", self.Spaceship.Info.Name))
end

function _TerminalSpaceship.ExecCommand(self, nodePointer, command)
    local tmp
    local commandArr = StringSplit(command, " ")

    if "speedx" == commandArr[1] then
        tmp = tonumber(commandArr[2])
        if "number" == type(tmp) then
            self.Spaceship:SetSpeedX(tmp)
        end

    elseif "speedy" == commandArr[1] then
        tmp = tonumber(commandArr[2])
        if "number" == type(tmp) then
            self.Spaceship:SetSpeedY(tmp)
        end

    elseif "speed" == commandArr[1] then
        tmp = tonumber(commandArr[2])
        if "number" == type(tmp) then
            self.Spaceship:SetSpeedX(tmp)
            self.Spaceship:SetSpeedY(tmp)
        end

    elseif "jump" == commandArr[1] then
        local ret = self.Spaceship:JumperRun({X=tonumber(commandArr[2]), Y=tonumber(commandArr[3])})
        if "string" == type(ret) then
            NodeTerminalMain:TerminalWriteNewLine(string.format("[%s](fg-red)", ret))
        else
            GTerminal:ScreenSuccessMsg("跳跃成功")
        end

    elseif "landing" == commandArr[1] then
        if nil ~= self.Spaceship.LoginedPlanet then
            GWorld:Stop(function()
                NodeModalPlanet:ModalShow()
            end)
        end

    else
        GTerminal:ScreenErrMsg(string.format("%s %s", GTerminal.ErrCommandNotExists, command))
    end
end
