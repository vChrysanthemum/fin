local json = require("json")

local _Planet = {}
local _mtPlanet = {__index = _Planet} 

function NewPlanet()
    local Planet = setmetatable({}, _mtPlanet)
    Planet:Format({
        PlanetId  = nil,
        Name      = "未命名星球",
        Position  = {},
        Resource  = 0,
        Character = "*",
        ColorFg   = "blue"
    })
    Planet.ScreenPosition = {}
    Planet.ColorBg        = ""
    return Planet
end

-- 初始化星球
function _Planet.Initilize(self, position)
    self.Info.Position = position
    RefreshRandomSeed()
    local a, b
    if position.X < 0 then
        a = string.format("a%d", position.X*-1)
    else
        a = string.format("b%d", position.X)
    end
    if position.Y < 0 then
        b = string.format("a%d", position.Y*-1)
    else
        b = string.format("b%d", position.Y)
    end
    self.Info.Name = a .. b
    self.Info.Resource = math.random(0,10000)
end

function _Planet.SetName(self, name)
    self.Info.Name = name
    self:FlushToDB()
    NodeRadar:SetActive()
end

function _Planet.Format(self, planetInfo)
    self.Info = {
        PlanetId  = planetInfo.PlanetId,
        Name      = planetInfo.Name,
        Position  = planetInfo.Position,
        Resource  = planetInfo.Resource,
        Character = planetInfo.Character,
        ColorFg   = planetInfo.ColorFg,
    }
end

function _Planet.FlushToDB(self)
    if "number" ~= type(self.Info.PlanetId) then
        return nil
    end

    sql = string.format([[
    update b_planet set data = '%s' where planet_id=%d
    ]], DB:QuoteSQL(json.encode(self.Info)), self.Info.PlanetId)
    local queryRet = DB:Exec(sql)
    if "string" == type(queryRet) then
        Log(queryRet)
    end
end
