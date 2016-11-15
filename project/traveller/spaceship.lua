local _SpaceShip = {}
local _mtSpaceShip = {__index = _SpaceShip}

function InitSpaceShip()
  local SpaceShip = setmetatable({}, _mtSpaceShip)
  SpaceShip.LocationX = 0.0
  SpaceShip.LocationY = 0.0
  SpaceShip.SpeedX = 0.0
  SpaceShip.SpeedY = 0.0
  local Warehouse = {}
  SpaceShip.Warehouse = Warehouse 

  return SpaceShip
end

function _SpaceShip.Run(self)
end
