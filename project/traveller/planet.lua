local _Planet = {}
local _mtPlanet = {__index = _Planet} 

function NewPlanet()
    local Planet = setmetatable({}, _mtPlanet)
    Planet.Name = "未命名星球"
    Planet.Position = {}
    Planet.ScreenPosition = {}
    Planet.Resource = 0
    return Planet
end

-- 初始化星球
function _Planet.Initilize(self, position)
    self.Position = position
    RefreshRandomSeed()
    self.Resource = math.random(0,10000)
end

function _Planet.SetName(self, name)
    self.Name = name
    NodeRadar:SetActive()
end
