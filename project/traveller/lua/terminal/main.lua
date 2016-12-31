function _Terminal.StartEnvMain(self)
    self:ScreenInfoMsg("返回 主控台 ...")
    self.CurrentEnv = "/"
    NodeTerminalMain:TerminalSetCommandPrefix("> ")
end

function _Terminal.ExecCommand(self, nodePointer, command)
    local commandArr = StringSplit(command, " ")
    if nil == commandArr or 0 == TableLength(commandArr) then
        return
    end

    if "/" == commandArr[1] then
        self:StartEnvMain()
        return

    elseif "clear" == commandArr[1] then
        NodeTerminalMain:TerminalClearLines()
        return

    elseif "clearhistory" == commandArr[1] then
        NodeTerminalMain:TerminalClearCommandHistory()
        return

    elseif "quit" == commandArr[1] then
        Quit()
        return
    end

    if true == CheckTableHasKey(self.CmdExcuter, commandArr[1]) then
        self.CurrentEnv = commandArr[1]
        self.CmdExcuter[commandArr[1]]:StartEnv(command)
        return
    end

    if "/" ~= self.CurrentEnv then
        self.CmdExcuter[self.CurrentEnv]:ExecCommand(nodePointer, command)
        return
    end

    GTerminal:ScreenErrMsg(string.format("%s %s", GTerminal.ErrCommandNotExists, command))

    return true
end
