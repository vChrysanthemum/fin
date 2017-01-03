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
    local building = GBuildingCenter:GetBuildingByPlanetIdAndBuildingType(planet.Info.PlanetId, buildingType)
    if nil ~= building then
        return nil
    end

    local PowerPlantCost = 10
    if planet.Info.ModuleDeveloped.Resource < PowerPlantCost then
        return string.format("资源小于%d，无法建造能源站", PowerPlantCost)
    end

    planet.Info.ModuleDeveloped.Resource = planet.Info.ModuleDeveloped.Resource - PowerPlantCost
    planet:FlushToDB()

    local buildingCore = NewBuildingCore()
    buildingCore.Info.BuildingType = buildingType
    buildingCore.Info.PlanetId = planet.Info.PlanetId
    local sql = string.format([[
    insert into b_building (planet_id, data) values (%d, '%s')
    ]], planet.Info.PlanetId, DB:QuoteSQL(json.encode(buildingCore.Info)))
    local queryRet = DB:Exec(sql)
    if "string" == queryRet then
        return queryRet
    end
    buildingCore.Info.BuildingId = queryRet:LastInsertId()
    buildingCore:Format(buildingCore.Info, buildingCore.Info.BuildingId)
    return buildingCore.Building
end

function _PowerPlant.GetBuildingTypeCh(self)
    return "能源站"
end
