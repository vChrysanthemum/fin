local _mtPlanet = {__index = _Planet} 

function InitPlanet()
  local Planet = setmetatable({}, _mtPlanet)
  return Planet
end
