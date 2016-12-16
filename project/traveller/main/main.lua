WorldLoop = function()
end 

NodeRadar = Node("CanvasRadar")
NodeTerminalMain = Node("TerminalMain")
NodeParInfo = Node("ParInfo")
NodeParNewestMsg = Node("ParNewestMsg")
NodeParGUserSpaceshipStatus = Node("ParGUserSpaceshipStatus")
NodeGaugeFuel = Node("GaugeFuel")
NodeGaugeLife = Node("GaugeLife")
NodeModalPlanet = Node("ModalPlanet")

NodeRadar:SetAttribute("height", tostring(WindowHeight()-NodeTerminalMain:Height()))

GUserSpaceship = GetSpaceshipFromDB(1)
GUserSpaceship:RefreshNodeParGUserSpaceshipStatus()
GRadar = NewRadar()
GTerminal = NewTerminal()
GWorld = NewWorld()

NodeParGUserSpaceshipStatus:SetAttribute("borderlabel", " " .. GUserSpaceship.Info.Name .. " ")

NodeTerminalMain:SetActive()
GUserSpaceship:UpdateFuel(0)
GUserSpaceship:UpdateLife(0)
GWorld:LoopEvent()
