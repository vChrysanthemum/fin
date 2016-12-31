function StartTabPlanet(planet)
    GTabPlanet = planet
    NodeTabPlanetTerminalMain:SetAttribute("borderlabel", string.format(" 大脑 - %s ", GTabPlanet.Info.Name))
    NodeTabpaneMain:TabpaneSetActiveTab("planet")
    RefreshNodeTabPlanetParPlanetInfo()
    NodeTabPlanetTerminalMain:SetActive()
end

function RefreshNodeTabPlanetParPlanetInfo()
    local value = string.format([[
星球: %s
坐标: %d, %d
资源: %d

已着陆机器人:
]], GTabPlanet.Info.Name, GTabPlanet.Info.Position.X, GTabPlanet.Info.Position.Y, GTabPlanet.Info.Resource)
    for k, robot in pairs(GRobotCenter.Robots) do
        if nil ~= robot.RobotCore.PlanetLanding and
            robot.RobotCore.PlanetLanding.Info.PlanetId == GTabPlanet.Info.PlanetId then
            value = value .. string.format("%s ", robot.RobotCore.Info.Name)
            if nil ~= robot.RobotCore.Info.Action then
                value = value .. string.format("%s ", robot:GetActionCh())
            end
            value = value .. string.format("\n")
        end
    end
    NodeTabPlanetParPlanetInfo:SetValue(value)
end

NodeTabPlanetTerminalMain:SetAttribute("top", tostring(WindowHeight()-NodeTabPlanetTerminalMain:Height()))
