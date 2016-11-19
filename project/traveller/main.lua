WorldLoop = function()
end 

NodeRadar = Node("CanvasRadar")
NodeTerminalMain = Node("TerminalMain")

NodeRadar:SetAttribute("height", tostring(WindowHeight()-NodeTerminalMain:Height()))

GUserSpaceShip = NewSpaceShip()
GRadar = NewRadar()
GTerminal = NewTerminal()
GWorld = NewWorld()

GUserSpaceShip.Location.X = -100
GUserSpaceShip.Location.Y = -100
function DisplayPlanet()
  local width = NodeRadar:Width()
  local height = NodeRadar:Height()
  local x = GetIntPart(GUserSpaceShip.Location.X)
  local y = GetIntPart(GUserSpaceShip.Location.Y)
  local rectangle = {}

  rectangle.Min = {
    X = GetIntPart(x - width / 2),
    Y = GetIntPart(x - height / 2)
  }
  rectangle.Max = {
    X = rectangle.Min.X + width,
    Y = rectangle.Min.Y + height,
  }

  GWorld:DrawPlanets(rectangle)
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
