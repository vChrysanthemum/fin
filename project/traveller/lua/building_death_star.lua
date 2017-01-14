local json = require("json")

_DeathStar = {}
local _mtDeathStar = {__index = _DeathStar} 

function NewDeathStar(buildingCore)
    local DeathStar = setmetatable({}, _mtDeathStar)
    DeathStar.BuildingCore = buildingCore
    return DeathStar
end

function BuildDeathStar(planet)
    local buildingType = "DeathStar" 
    local buildCost = 2
    local buildingCore = CreateBuildingCore(planet, buildingType, buildCost)
    if "table" ~= type(buildingCore) then
        return buildingCore
    end
    return buildingCore.Building
end

function DestroyDeathStar(planet)
    local buildingType = "DeathStar" 
    return DestroyBuildingCore(planet, buildingType)
end
