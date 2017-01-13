local json = require("json")

local _BuildingCenter = {}
local _mtBuildingCenter = {__index = _BuildingCenter} 

function NewBuildingCenter()
    local BuildingCenter = setmetatable({}, _mtBuildingCenter)
    -- BuildingCenter.buildingIdToBuilding = {}
    BuildingCenter.Buildings = {}
    return BuildingCenter
end

function _BuildingCenter.GetBuildingByPlanetIdAndBuildingType(self, planetId, buildingType)
    local buildings = self:GetBuildingByPlanetId(planetId)
    for _, building in pairs(buildings) do
        if buildingType == building.BuildingCore.Info.BuildingType then
            return building
        end
    end
    return nil
end

function _BuildingCenter.GetBuildingByPlanetId(self, planetId)
    local sql = string.format([[
    select building_id, data from b_building where planet_id=%d
    ]], planetId)
    local rows = DB:Query(sql)
    local row = nil
    local building
    local buildingCore
    local buildings = {}
    while true do
        row = rows:FetchOne()
        if "table" ~= type(row) then
            break
        end
        if nil ~= self.Buildings[row.building_id] then
            building = self.Buildings[row.building_id]
        else
            buildingCore = NewBuildingCore()
            buildingCore:Format(json.decode(row.data), row.building_id)
            self.Buildings[row.building_id] = buildingCore.Building
            building = buildingCore.Building
        end
        table.insert(buildings, building)
    end
    rows:Close()
    return buildings
end
