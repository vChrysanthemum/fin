local _Terminal = {}
local _mtTerminal = {__index = _Terminal} 

function NewTerminal()
    local Terminal = setmetatable({}, _mtTerminal)
    Terminal.CurrentEnv = "/"

    Terminal.CommandSig = NodeTerminalMain:TerminalRegisterCommandHandle(function(nodePointer, command)
        Terminal:ExecCommand(nodePointer, command)
    end)

    Terminal.ErrCommandNotExists = "无效命令"
    Terminal.ConnentingPlanet = nil

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

function _Terminal.ExecCommand(self, nodePointer, command)
    local tmp
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
    end

    if true == self:ExecCommandMain(nodePointer, command) then
        return
    end
       

    if "/" == self.CurrentEnv then
    elseif "/spaceship" == self.CurrentEnv then
        self:ExecCommandSpaceship(nodePointer, command)
    elseif "/planet" == self.CurrentEnv then
        self:ExecCommandPlanet(nodePointer, command)
    end
end

function _Terminal.StartEnvMain(self)
    self:ScreenInfoMsg("返回 主控台 ...")
    self.CurrentEnv = "/"
    NodeTerminalMain:TerminalSetCommandPrefix("> ")
end

function _Terminal.ExecCommandMain(self, nodePointer, command)
    local tmp
    local commandArr = StringSplit(command, " ")

    if "/spaceship" == commandArr[1] then
        self:StartEnvSpaceship()
        return true

    elseif "/planet" == commandArr[1] then
        local location = {}
        if nil ~= GUserSpaceship.LoginedPlanet then
            location = GUserSpaceship.LoginedPlanet.Info.Position
        else
            if TableLength(commandArr) < 3 then
                self:ScreenErrMsg("请输入星球坐标")
                return
            end
            location = {X=tonumber(commandArr[2]), Y=tonumber(commandArr[3])}
        end

        self:StartEnvPlanet(location)
        return true
    end

    return false
end

function _Terminal.StartEnvSpaceship(self)
    self:ScreenInfoMsg("连接 飞船 ...")
    NodeTerminalMain:TerminalSetCommandPrefix(string.format("%s> ", GUserSpaceship.Info.Name))
    self.CurrentEnv = "/spaceship"
end

function _Terminal.ExecCommandSpaceship(self, nodePointer, command)
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

    elseif "jump" == commandArr[1] then
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

function _Terminal.StartEnvPlanet(self, position)
    if nil == position or nil == position.X or nil == position.Y then
        self:ScreenErrMsg(string.format("请输入有效坐标"))
        return
    end

    self:ScreenInfoMsg(string.format("连接 星球 %s ...", PointToStr(position)))
    local planet = GRadar:GetPlanetOnScreenByPosition(position)
    if nil == planet then
        self:ScreenErrMsg(string.format("无法连接星球 %s", PointToStr(position)))
        return
    end

    self.ConnentingPlanet = planet
    NodeTerminalMain:TerminalSetCommandPrefix(string.format("%s> ", planet.Info.Name))
    self.CurrentEnv = "/planet"
end

function _Terminal.ExecCommandPlanet(self, nodePointer, command)
    local tmp
    local commandArr = StringSplit(command, " ")

    if "info" == commandArr[1] then
        self:ScreenInfoMsg(string.format("名称: %s", self.ConnentingPlanet.Info.Name))
        self:ScreenInfoMsg(string.format("坐标: %s", PointToStr(self.ConnentingPlanet.Info.Position)))

    elseif "rename" == commandArr[1] then
        self.ConnentingPlanet:SetName(commandArr[2])
        NodeTerminalMain:TerminalSetCommandPrefix(string.format("%s> ", self.ConnentingPlanet.Info.Name))

    elseif "detail" == commandArr[1] then
        self:ShowPlanetDetail()

    else
        self:ScreenErrMsg(string.format("%s %s", self.ErrCommandNotExists, command))
    end
end

function _Terminal.ShowPlanetDetail(self)
    GWorld:Stop(function()
        NodeModalPlanet:ModalShow()
    end)
end
