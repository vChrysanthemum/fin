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
