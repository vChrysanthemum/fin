WorldLoop = function()
end 

NodeRadar = Node("CanvasRadar")
NodeTerminalMain = Node("TerminalMain")
NodeParInfo = Node("ParInfo")
NodeInputTextNamePlanet = Node("InputTextNamePlanet")
NodeParGUserSpaceShipStatus = Node("ParGUserSpaceShipStatus")

NodeRadar:SetAttribute("height", tostring(WindowHeight()-NodeTerminalMain:Height()))

GUserSpaceShip = NewSpaceShip()
GRadar = NewRadar()
GTerminal = NewTerminal()
GWorld = NewWorld()

NodeParGUserSpaceShipStatus:SetAttribute("borderlabel", GUserSpaceShip.Info.Name)

NodeInputTextNamePlanet:RegisterKeyPressEnterHandler(function(nodePointer)
  if nil ~= GRadar.FocusTarget then
    GRadar.FocusTarget:SetName(Node(nodePointer):GetValue())
  end
end)

local planets = {}
function Display()
    GUserSpaceShip:SetPosition(-100, 100)
    planets = GWorld:GetPlanetsByRectangle(GUserSpaceShip.CenterRectangle)

    GRadar:RefreshScreenPlanets(planets, GUserSpaceShip.CenterRectangle)
    GRadar:Draw()
end
Display()
NodeTerminalMain:SetActive()

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
