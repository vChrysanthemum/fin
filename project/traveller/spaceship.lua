local _SpaceShip = {}
local _mtSpaceShip = {__index = _SpaceShip}

function NewSpaceShip()
    local SpaceShip = setmetatable({}, _mtSpaceShip)
    SpaceShip.Location = {X=0.0, Y=0.0}
    SpaceShip.Speed = {X=0.0, Y=0.0}
    SpaceShip.CenterRectangle = {}
    local Warehouse = {}
    SpaceShip.Warehouse = Warehouse 

    return SpaceShip
end

-- 刷新飞船为中心的指定大小区域所在的宇宙位置
function _SpaceShip.refreshCenterRectangle(self, rectangleWidth, rectangleHeight)
    local x = GetIntPart(self.Location.X)
    local y = GetIntPart(self.Location.Y)
    local rectangle = {}
    rectangle.Min = {
        X = GetIntPart(x - rectangleWidth / 2),
        Y = GetIntPart(y - rectangleHeight / 2)
    }
    rectangle.Max = {
        X = rectangle.Min.X + rectangleWidth,
        Y = rectangle.Min.Y + rectangleHeight,
    }

    self.CenterRectangle = rectangle

    return rectangle
end

-- 更新飞船位置
function _SpaceShip.SetPosition(self, positionX, positionY)
    self.Location.X = positionX
    self.Location.Y = positionY
    self:refreshCenterRectangle(NodeRadar:Width(), NodeRadar:Height())
end
