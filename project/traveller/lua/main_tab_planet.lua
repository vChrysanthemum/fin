function StartTabPlanet(planetId)
    local planet = GWorld:GetPlanetByPlanetId(planetId)
    if nil == planet then
        return
    end
    GTabPlanetId = planet.Info.PlanetId
    NodeTabPlanetTerminalMain:SetAttribute("borderlabel", string.format(" 大脑 - %s ", planet.Info.Name))
    NodeTabpaneMain:TabpaneSetActiveTab("planet")
    planet:RefreshModuleDevelopedBuilding()
    RefreshNodeTabPlanetParPlanetInfo()
    NodeTabPlanetTerminalMain:SetActive()
end

function RefreshNodeTabPlanetParPlanetInfo()
    local planet = GWorld:GetPlanetByPlanetId(GTabPlanetId)
    if nil == planet then
        return
    end
    local robotsStr = ""
    local RobotCorePlanetLanding = {}
    for _, robot in ipairs(GRobotCenter.Robots) do
        RobotCorePlanetLanding = GWorld:GetPlanetByPlanetId(robot.RobotCore.Info.LandingPlanetId)
        if nil ~= RobotCorePlanetLanding and
            RobotCorePlanetLanding.Info.PlanetId == planet.Info.PlanetId then
            robotsStr = robotsStr .. string.format("%s ", robot.RobotCore.Info.Name)
            if nil ~= robot.RobotCore.Info.Action then
                robotsStr = robotsStr .. string.format("%s ", robot:GetActionCh())
            end
            robotsStr = robotsStr .. string.format("\n")
        end
    end

    local moduleDeveloped = string.format("资源: %f\n", planet.Info.ModuleDeveloped.Resource)
    local building = nil
    if nil ~= planet.ModuleDevelopedBuilding then
        for _, building in ipairs(planet.ModuleDevelopedBuilding) do
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
    planet.Info.Name, planet.Info.Position.X, planet.Info.Position.Y, planet.Info.Resource,
    moduleDeveloped, robotsStr)

    NodeTabPlanetParPlanetInfo:SetValue(value)
end

NodeTabPlanetTerminalMain:SetAttribute("top", tostring(WindowHeight()-NodeTabPlanetTerminalMain:Height()))
