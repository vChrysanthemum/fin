local json = require("json")

_PowerPlant = {}
local _mtPowerPlant = {__index = _PowerPlant} 

function NewPowerPlant(buildingCore)
    local PowerPlant = setmetatable({}, _mtPowerPlant)
    PowerPlant.BuildingCore = buildingCore
    return PowerPlant
end

function BuildPowerPlant(planet)
    local buildingType = "PowerPlant" 
    local buildCost = 1
    local buildingCore = CreateBuildingCore(planet, buildingType, buildCost)
    if "table" ~= type(buildingCore) then
        return buildingCore
    end
    return buildingCore.Building
end

function DestroyPowerPlant(planet)
    local buildingType = "PowerPlant" 
    return DestroyBuildingCore(planet, buildingType)
end
