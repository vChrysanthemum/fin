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
  if "punch" == command then
    NodeTerminalMain:TerminalWriteNewLine("  punch success.")
  elseif "clear" == command then
    NodeTerminalMain:TerminalClearLines()
  elseif "destory" == command then
    NodeTerminalMain:TerminalRemoveCommandHandle(self.CommandSig)
  end
end
