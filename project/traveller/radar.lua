local _Radar = {}
local _mtRadar = {__index = _Radar} 

function InitRadar()
  local Radar = setmetatable({}, _mtRadar)
  Radar.KeyPressEnterChans = {}
  Radar.KeyPressSig = NodeRadar:RegisterKeyPressHandler(function(nodePointer, keyStr)
    Radar:KeyPressHandle(nodePointer, keyStr)
  end)
  return Radar
end

function _Radar.KeyPressHandle(self, nodePointer, keyStr)
  if "<enter>" == keyStr then
    for k, ch in pairs(GRadar.KeyPressEnterChans) do
      GRadar.KeyPressEnterChans[k]:send()
    end
    return
  end

  if "<left>" == keyStr then
    GUserSpaceShip.LocationX = GUserSpaceShip.LocationX - 1
  elseif "<right>" == keyStr then
    GUserSpaceShip.LocationX = GUserSpaceShip.LocationX + 1
  elseif "<up>" == keyStr then
    GUserSpaceShip.LocationY = GUserSpaceShip.LocationY - 1
  elseif "<down>" == keyStr then
    GUserSpaceShip.LocationY = GUserSpaceShip.LocationY + 1
  end

  NodeRadar:SetCursor(GUserSpaceShip.LocationX, GUserSpaceShip.LocationY)
end
