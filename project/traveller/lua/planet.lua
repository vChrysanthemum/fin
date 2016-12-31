local json = require("json")

local _Planet = {}
local _mtPlanet = {__index = _Planet} 

function NewPlanet()
  local Planet = setmetatable({}, _mtPlanet)
  Planet:Format({
    Name      = "未命名星球",
    Position  = {},
    Resource  = 0,
    Character = "*",
    ColorFg   = "blue"
  }, 0)
  Planet.ScreenPosition = {}
  Planet.ColorBg        = ""
  return Planet
end

-- 初始化星球
function _Planet.Initilize(self, position)
  self.Info.Position = position
  RefreshRandomSeed()
  local a, b, multiplyNumber

  multiplyNumber = 1
  if position.X < 0 then
    multiplyNumber = -1
  end
  a = "a" .. string.lpad(tostring(position.X*multiplyNumber), 3, '0')

  multiplyNumber = 1
  if position.Y < 0 then
    multiplyNumber = -1
  end
  b = "b" .. string.lpad(tostring(position.Y*multiplyNumber), 3, '0')

  self.Info.Name = a .. b
  self.Info.Resource = math.random(0,10000)
end

function _Planet.SetName(self, name)
  self.Info.Name = name
  self:FlushToDB()
end

function _Planet.Format(self, planetInfo, planetId)
  self.Info = {
    PlanetId  = tonumber(planetId),
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

-- 被机器人挖矿，资源变动
function _Planet.MineByRobot(self)
    self.Info.Resource = self.Info.Resource - 1
end
