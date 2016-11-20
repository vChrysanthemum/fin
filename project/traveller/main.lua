WorldLoop = function()
end 

NodeRadar = Node("CanvasRadar")
NodeTerminalMain = Node("TerminalMain")
NodeParInfo = Node("ParInfo")
NodeInputTextNamePlanet = Node("InputTextNamePlanet")
NodeParGUserSpaceshipStatus = Node("ParGUserSpaceshipStatus")

NodeRadar:SetAttribute("height", tostring(WindowHeight()-NodeTerminalMain:Height()))

GUserSpaceship = GetSpaceshipFromDB(1)
GUserSpaceship:RefreshNodeParGUserSpaceshipStatus()
GRadar = NewRadar()
GTerminal = NewTerminal()
GWorld = NewWorld()

NodeParGUserSpaceshipStatus:SetAttribute("borderlabel", " " .. GUserSpaceship.Info.Name .. " ")

NodeInputTextNamePlanet:RegisterKeyPressEnterHandler(function(nodePointer)
  if nil ~= GRadar.FocusTarget then
    GRadar.FocusTarget:SetName(Node(nodePointer):GetValue())
  end
end)

NodeTerminalMain:SetActive()
GWorld:LoopEvent()

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
