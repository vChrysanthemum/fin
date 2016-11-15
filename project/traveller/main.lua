WorldLoop = function()
end 

NodeRadar = Node("CanvasRadar")
NodeTerminalMain = Node("TerminalMain")

NodeRadar:SetAttribute("height", tostring(WindowHeight()-NodeTerminalMain:Height()))

GUserSpaceShip = InitSpaceShip()
GRadar = InitRadar()
GTerminal = InitTerminal()
GWorld = InitWorld()

--[[
local NodeRadar = Node("CanvasRadar")
NodeRadar:CanvasSet(12, 3, "#", "green", "")
NodeRadar:CanvasSet(48, 10, "#", "green", "red")
NodeRadar:CanvasSet(16, 15, "#", "green", "")
NodeRadar:CanvasSet(75, 5, "#", "green", "")
NodeRadar:CanvasSet(56, 13, "#", "green", "")
NodeRadar:CanvasSet(44, 12, "*", "blue", "")
NodeRadar:CanvasDraw(12,3)
]]
