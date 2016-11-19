local _Planet = {}
local _mtPlanet = {__index = _Planet} 

function NewPlanet()
    local Planet = setmetatable({}, _mtPlanet)
    Planet.Position = {}
    Planet.ScreenPosition = {}
    return Planet
end

-- 初始化星球
function _Planet.Initilize(self, position)
    self.Position = position
end
