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
    if TableLength(command) < 2 then
        return nil
    end

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
    end
        --GUserSpaceship.Info.Speed = tonumber(command[1])
    --[[
    if "punch" == command then
    NodeTerminalMain:TerminalWriteNewLine("  punch success.")
    elseif "clear" == command then
    NodeTerminalMain:TerminalClearLines()
    elseif "destory" == command then
    NodeTerminalMain:TerminalRemoveCommandHandle(self.CommandSig)
    end
    ]]
end
