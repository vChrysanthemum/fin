function StartTabPlanet(planet)
  NodeTabPlanetTerminalMain:SetAttribute("borderlabel", string.format(" 大脑 - %s ", planet.Info.Name))
  NodeTabpaneMain:TabpaneSetActiveTab("planet")
  NodeTabPlanetParPlanetInfo:SetValue(string.format([[
星球: %s
坐标: %d, %d
资源: %d]], planet.Info.Name, planet.Info.Position.X, planet.Info.Position.Y, planet.Info.Resource))
  NodeTabPlanetTerminalMain:SetActive()
end

NodeTabPlanetTerminalMain:SetAttribute("top", tostring(WindowHeight()-NodeTabPlanetTerminalMain:Height()))
