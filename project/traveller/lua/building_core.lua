local json = require("json")

local _BuildingCore = {}
local _mtBuildingCore = {__index = _BuildingCore} 

function NewBuildingCore()
    local BuildingCore = setmetatable({}, _mtBuildingCore)
    BuildingCore.Info = {
        BuildingId       = nil,
        BuildingType     = "",
        PlanetId         = nil,
    }
    BuildingCore.Building = nil
    return BuildingCore
end

function _BuildingCore.Format(self, buildingInfo, building_id)
    self.Info = {
        BuildingId   = tonumber(building_id),
        BuildingType = buildingInfo.BuildingType,
        PlanetId     = buildingInfo.PlanetId,
    }

    if "PowerPlant" == self.Info.BuildingType then
        self.Building = NewPowerPlant(self)
    elseif "DeathStar" == self.Info.BuildingType then
        self.Building = NewDeathStar(self)
    end
end

function _BuildingCore.FlushToDB(self)
    if "number" ~= type(self.Info.BuildingId) then
        return nil
    end

    local sql = string.format([[
    update b_building set planet_id=%d, data = '%s' where building_id=%d
    ]], self.Info.PlanetId, DB:QuoteSQL(json.encode(self.Info)), self.Info.BuildingId)
    local queryRet = DB:Exec(sql)
    if "string" == type(queryRet) then
        Log(queryRet)
    end
end

function CreateBuildingCore(planet, buildingType, buildCost)
    local building = GBuildingCenter:GetBuildingByPlanetIdAndBuildingType(planet.Info.PlanetId, buildingType)
    if nil ~= building then
        return string.format("%s已建造", GDictE2C[buildingType])
    end

    if planet.Info.ModuleDeveloped.Resource < buildCost then
        return string.format("资源小于%d，无法建造%s", buildCost, GDictE2C[buildingType])
    end

    planet.Info.ModuleDeveloped.Resource = planet.Info.ModuleDeveloped.Resource - buildCost
    planet:FlushToDB()

    local buildingCore = NewBuildingCore()
    buildingCore.Info.BuildingType = buildingType
    buildingCore.Info.PlanetId = planet.Info.PlanetId

    local sql = string.format(
    "insert into b_building (planet_id, data) values (%d, '%s')",
    planet.Info.PlanetId, DB:QuoteSQL(json.encode(buildingCore.Info)))

    local queryRet = DB:Exec(sql)
    if "string" == type(queryRet) then
        return queryRet
    end
    buildingCore.Info.BuildingId = queryRet:LastInsertId()
    buildingCore:Format(buildingCore.Info, buildingCore.Info.BuildingId)
    return buildingCore
end

function DestroyBuildingCore(planet, buildingType)
    local building = GBuildingCenter:GetBuildingByPlanetIdAndBuildingType(planet.Info.PlanetId, buildingType)
    if nil == building then
        return string.format("%s未建造", GDictE2C[buildingType])
    end

    local sql = string.format(
    "delete from b_building where building_id=%d",
    building.BuildingCore.Info.BuildingId)
    local queryRet = DB:Exec(sql)
    if "string" == type(queryRet) then
        return queryRet
    end
end
