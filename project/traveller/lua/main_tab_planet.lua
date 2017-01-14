function StartTabPlanet(planet)
    GTabPlanet = planet
    NodeTabPlanetTerminalMain:SetAttribute("borderlabel", string.format(" 大脑 - %s ", GTabPlanet.Info.Name))
    NodeTabpaneMain:TabpaneSetActiveTab("planet")
    GTabPlanet:RefreshModuleDevelopedBuilding()
    RefreshNodeTabPlanetParPlanetInfo()
    NodeTabPlanetTerminalMain:SetActive()
end

function RefreshNodeTabPlanetParPlanetInfo()
    if nil == GTabPlanet then
        return
    end
    local robotsStr = ""
    for _, robot in ipairs(GRobotCenter.Robots) do
        if nil ~= robot.RobotCore.PlanetLanding and
            robot.RobotCore.PlanetLanding.Info.PlanetId == GTabPlanet.Info.PlanetId then
            robotsStr = robotsStr .. string.format("%s ", robot.RobotCore.Info.Name)
            if nil ~= robot.RobotCore.Info.Action then
                robotsStr = robotsStr .. string.format("%s ", robot:GetActionCh())
            end
            robotsStr = robotsStr .. string.format("\n")
        end
    end

    local moduleDeveloped = string.format("资源: %f\n", GTabPlanet.Info.ModuleDeveloped.Resource)
    local building = nil
    if nil ~= GTabPlanet.ModuleDevelopedBuilding then
        for _, building in ipairs(GTabPlanet.ModuleDevelopedBuilding) do
            moduleDeveloped = moduleDeveloped .. string.format("%s\n", building:GetBuildingTypeCh())
        end
    end

    local value = string.format([[
星球: %s
坐标: %d, %d
资源: %f

已开发模块:
%s
已着陆机器人:
%s]], 
    GTabPlanet.Info.Name, GTabPlanet.Info.Position.X, GTabPlanet.Info.Position.Y, GTabPlanet.Info.Resource,
    moduleDeveloped, robotsStr)

    NodeTabPlanetParPlanetInfo:SetValue(value)
end

NodeTabPlanetTerminalMain:SetAttribute("top", tostring(WindowHeight()-NodeTabPlanetTerminalMain:Height()))
