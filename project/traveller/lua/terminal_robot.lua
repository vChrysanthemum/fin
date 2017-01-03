local _TerminalRobot = {}
local _mtTerminalRobot = {__index = _TerminalRobot} 

function NewTerminalRobot(terminal)
    local TerminalRobot = setmetatable({}, _mtTerminalRobot)
    TerminalRobot.Env = "/robot"
    TerminalRobot.ConnentingRobot = nil
    TerminalRobot.Terminal = terminal

    return TerminalRobot
end

function _TerminalRobot.StartEnv(self, command)
    local commandArr = StringSplit(command, " ")

    local position = {}
    if TableLength(commandArr) < 2 then
        self.Terminal:ScreenErrMsg("请输入机器人监听地址")
        return
    end 
    robotServiceAddress = commandArr[2]

    self.Terminal:ScreenInfoMsg(string.format("连接 机器人 %s ...", robotServiceAddress))
    local robot = GRobotCenter:GetRobotByServiceAddress(robotServiceAddress)
    if nil == robot then
        self.Terminal:ScreenErrMsg(string.format("无法连接机器人 %s", robotServiceAddress))
        return
    end

    self.ConnentingRobot = robot
    self.ConnentingRobot:SetClientTerminal(self.Terminal)
    self.Terminal.Port:TerminalSetCommandPrefix(string.format("%s> ", robot.RobotCore.Info.Name))
end

function _TerminalRobot.ExecCommand(self, nodePointer, command)
    self.ConnentingRobot:ExecCommand(command)
end
