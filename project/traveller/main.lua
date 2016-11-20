WorldLoop = function()
end 

NodeRadar = Node("CanvasRadar")
NodeTerminalMain = Node("TerminalMain")
NodeParInfo = Node("ParInfo")
NodeInputTextNamePlanet = Node("InputTextNamePlanet")

NodeRadar:SetAttribute("height", tostring(WindowHeight()-NodeTerminalMain:Height()))

GUserSpaceShip = NewSpaceShip()
GRadar = NewRadar()
GTerminal = NewTerminal()
GWorld = NewWorld()

NodeInputTextNamePlanet:RegisterKeyPressEnterHandler(function(nodePointer)
  if nil ~= GRadar.FocusPlanet then
    GRadar.FocusPlanet:SetName(Node(nodePointer):GetValue())
  end
end)

local planets = {}
function DisplayPlanet()
    GUserSpaceShip:SetPosition(-100, 100)
    planets = GWorld:GetPlanetsByRectangle(GUserSpaceShip.CenterRectangle)

    GRadar:RefreshScreenPlanets(planets, GUserSpaceShip.CenterRectangle)
    GRadar:DrawPlanets()
    NodeTerminalMain:SetActive()
end
DisplayPlanet()

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
