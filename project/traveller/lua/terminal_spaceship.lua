local _TerminalSpaceship = {}
local _mtTerminalSpaceship = {__index = _TerminalSpaceship} 

function NewTerminalSpaceship(terminal)
    local TerminalSpaceship = setmetatable({}, _mtTerminalSpaceship)
    TerminalSpaceship.Env = "/spaceship"
    TerminalSpaceship.Spaceship = nil
    TerminalSpaceship.Terminal = terminal

    return TerminalSpaceship
end

function _TerminalSpaceship.StartEnv(self, command)
    self.Terminal:ScreenInfoMsg("连接 飞船 ...")
    self.Spaceship = GUserSpaceship
    self.Terminal.Port:TerminalSetCommandPrefix(string.format("%s> ", self.Spaceship.Info.Name))
end

function _TerminalSpaceship.ExecCommand(self, nodePointer, command)
    local tmp
    local commandArr = StringSplit(command, " ")

    if "info" == commandArr[1] then
        self.Terminal:ScreenInfoMsg(string.format("名称: %s", self.Spaceship.Info.Name))

    elseif "speedx" == commandArr[1] then
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
            self.Terminal.Port:TerminalWriteNewLine(string.format("[%s](fg-red)", ret))
        else
            self.Terminal:ScreenSuccessMsg("跳跃成功")
        end

    elseif "landing" == commandArr[1] then
        if nil ~= self.Spaceship.LoginedPlanet then
            self.Terminal:StartEnv("/planet")
        end

    else
        self.Terminal:ScreenErrMsg(string.format("%s %s", self.Terminal.ErrCommandNotExists, command))
    end
end
