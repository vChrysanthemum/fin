local _Terminal = {}
local _mtTerminal = {__index = _Terminal} 

function NewTerminal()
    local Terminal = setmetatable({}, _mtTerminal)
    Terminal.CurrentCommand = ""

    Terminal.CommandSig = NodeTerminalMain:TerminalRegisterCommandHandle(function(nodePointer, command)
        Terminal:ExecCommand(nodePointer, command)
    end)
    return Terminal
end

function _Terminal.ExecCommand(self, nodePointer, command)
    local tmp
    command = StringSplit(command, " ")

    if "speedx" == command[1] then
        tmp = tonumber(command[2])
        if "number" == type(tmp) then
            GUserSpaceship:SetSpeedX(tmp)
        end

    elseif "speedy" == command[1] then
        tmp = tonumber(command[2])
        if "number" == type(tmp) then
            GUserSpaceship:SetSpeedY(tmp)
        end

    elseif "speed" == command[1] then
        tmp = tonumber(command[2])
        if "number" == type(tmp) then
            GUserSpaceship:SetSpeedX(tmp)
            GUserSpaceship:SetSpeedY(tmp)
        end

    elseif "landing" == command[1] then
        GUserSpaceship:Landing()

    elseif "clear" == command[1] then
        NodeTerminalMain:TerminalClearLines()
    end
end
