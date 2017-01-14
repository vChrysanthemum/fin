function _Terminal.StartEnvMain(self)
    self:ScreenInfoMsg("返回 主控台 ...")
    self.CurrentEnv = "/"
    self.Port:TerminalSetCommandPrefix("> ")
end

function _Terminal.StartEnv(self, command)
    local commandArr = StringSplit(command, " ")
    if nil == commandArr or 0 == TableLength(commandArr) then
        return false
    end
    local ret = self.CmdExcuter[commandArr[1]]:StartEnv(command)
    if false == ret then
        return false
    end
    self.CurrentEnv = commandArr[1]
    return true
end

function _Terminal.ExecCommand(self, nodePointer, command)
    local commandArr = StringSplit(command, " ")
    if nil == commandArr or 0 == TableLength(commandArr) then
        return
    end

    if "/" == commandArr[1] then
        self:StartEnvMain()
        return

    elseif "tab" == commandArr[1] then
        self:StartTab(command)
        return

    elseif "clear" == commandArr[1] then
        self.Port:TerminalClearLines()
        return

    elseif "clearhistory" == commandArr[1] then
        self.Port:TerminalClearCommandHistory()
        return

    elseif "quit" == commandArr[1] then
        Quit()
        return
    end

    if true == CheckTableHasKey(self.CmdExcuter, commandArr[1]) then
        self:StartEnv(command)
        return
    end

    if "/" ~= self.CurrentEnv then
        self.CmdExcuter[self.CurrentEnv]:ExecCommand(nodePointer, command)
        return
    end

    self:ScreenErrMsg(string.format("%s %s", self.ErrCommandNotExists, command))

    return true
end

function _Terminal.StartTab(self, command)
    local commandArr = StringSplit(command, " ")

    if "main" == commandArr[2] then
        StartTabMain()
    elseif "planet" == commandArr[2] then
      if nil ~= GTabPlanetId then
        StartTabPlanet(GTabPlanetId)
      end
    end
end
