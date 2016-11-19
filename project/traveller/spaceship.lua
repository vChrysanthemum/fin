local _SpaceShip = {}
local _mtSpaceShip = {__index = _SpaceShip}

function NewSpaceShip()
  local SpaceShip = setmetatable({}, _mtSpaceShip)
  SpaceShip.Location = {X=0.0, Y=0.0}
  SpaceShip.Speed = {X=0.0, Y=0.0}
  local Warehouse = {}
  SpaceShip.Warehouse = Warehouse 

  return SpaceShip
end

function _SpaceShip.Run(self)
end
