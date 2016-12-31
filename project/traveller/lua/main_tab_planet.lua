NodeTabPlanetSelectEnterMain = Node("TabPlanetSelectEnterMain")
NodeTabPlanetParPlanetInfo = Node("TabPlanetParPlanetInfo")

NodeTabPlanetSelectEnterMainKeyPressEnterKeySig = 
NodeTabPlanetSelectEnterMain:RegisterKeyPressEnterHandler(function(nodePointer)
  NodeTabpaneMain:TabpaneSetActiveTab("main")
end)

function StartTabPlanet(planet)
    NodeTabpaneMain:TabpaneSetActiveTab("planet")
    NodeTabPlanetParPlanetInfo:SetValue(string.format([[
星球: %s
坐标: %d, %d
资源: %d]], planet.Info.Name, planet.Info.Position.X, planet.Info.Position.Y, planet.Info.Resource))
end
