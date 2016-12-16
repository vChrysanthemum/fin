local _Terminal = {}
local _mtTerminal = {__index = _Terminal} 

function NewTerminal()
    local Terminal = setmetatable({}, _mtTerminal)
    Terminal.CurrentEnv = "/"
    Terminal.CurrentCommand = ""

    Terminal.CommandSig = NodeTerminalMain:TerminalRegisterCommandHandle(function(nodePointer, command)
        Terminal:ExecCommand(nodePointer, command)
    end)

    Terminal.ErrCommandNotExists = "无效命令"

    return Terminal
end

function _Terminal.ExecCommand(self, nodePointer, command)
    local tmp
    local commandArr = StringSplit(command, " ")
    if 0 == TableLength(commandArr) then
        return
    end

    if "/" == commandArr[1] then
        self:StartEnvMain()
        return

    elseif "quit" == commandArr[1] then
        Quit()
    end

    if "/" == self.CurrentEnv then
        self:ExecCommandMain(nodePointer, command)
    elseif "/jumper" == self.CurrentEnv then
        self:ExecCommandJump(nodePointer, command)
    end
end

function _Terminal.StartEnvMain(self)
    self:ScreenSuccessMsg("连接 主控台 ...")
    self.CurrentEnv = "/"
    NodeTerminalMain:TerminalSetCommandPrefix("> ")
end

function _Terminal.ExecCommandMain(self, nodePointer, command)
    local tmp
    local commandArr = StringSplit(command, " ")

    if "speedx" == commandArr[1] then
        tmp = tonumber(commandArr[2])
        if "number" == type(tmp) then
            GUserSpaceship:SetSpeedX(tmp)
        end

    elseif "speedy" == commandArr[1] then
        tmp = tonumber(commandArr[2])
        if "number" == type(tmp) then
            GUserSpaceship:SetSpeedY(tmp)
        end

    elseif "speed" == commandArr[1] then
        tmp = tonumber(commandArr[2])
        if "number" == type(tmp) then
            GUserSpaceship:SetSpeedX(tmp)
            GUserSpaceship:SetSpeedY(tmp)
        end

    elseif "landing" == commandArr[1] then
        GUserSpaceship:Landing()

    elseif "clear" == commandArr[1] then
        NodeTerminalMain:TerminalClearLines()

    elseif "/jumper" == commandArr[1] then
        self:StartEnvJumper()

    else
        self:ScreenErrMsg(string.format("%s %s", self.ErrCommandNotExists, command))
    end
end

function _Terminal.StartEnvJumper(self)
    self:ScreenSuccessMsg("连接 jumper ...")
    NodeTerminalMain:TerminalSetCommandPrefix("jumper> ")
    self.CurrentEnv = "/jumper"
end

function _Terminal.ExecCommandJump(self, nodePointer, command)
    local tmp
    local commandArr = StringSplit(command, " ")

    if "jump" == commandArr[1] then
        local ret = GUserSpaceship:JumperRun({X=tonumber(commandArr[2]), Y=tonumber(commandArr[3])})
        if "string" == type(ret) then
            NodeTerminalMain:TerminalWriteNewLine(string.format("[%s](fg-red)", ret))
        else
            self:ScreenSuccessMsg("跳跃成功")
        end

    else
        self:ScreenErrMsg(string.format("%s %s", self.ErrCommandNotExists, command))
    end
end

function _Terminal.ScreenErrMsg(self, msg)
    NodeTerminalMain:TerminalWriteNewLine(string.format("[%s](fg-red)", msg))
end

function _Terminal.ScreenSuccessMsg(self, msg)
    NodeTerminalMain:TerminalWriteNewLine(string.format("[%s](fg-green)", msg))
end
