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
