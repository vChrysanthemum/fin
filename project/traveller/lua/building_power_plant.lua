local json = require("json")

_PowerPlant = {}
local _mtPowerPlant = {__index = _PowerPlant} 

function NewPowerPlant(buildingCore)
    local PowerPlant = setmetatable({}, _mtPowerPlant)
    PowerPlant.BuildingCore = buildingCore
    PowerPlant.BuildingCore.BuildingType = "PowerPlant"
    PowerPlant.ClientTerminal = nil
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

function _PowerPlant.SetClientTerminal(self, clientTerminal)
    self.ClientTerminal = clientTerminal
end

function _PowerPlant.ExecCommand(self, command)
    self.ClientTerminal:ScreenErrMsg("PowerPlant")
end
