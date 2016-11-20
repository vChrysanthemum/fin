local _SpaceShip = {}
local _mtSpaceShip = {__index = _SpaceShip}

function NewSpaceShip()
    local SpaceShip           = setmetatable({}, _mtSpaceShip)
    SpaceShip.ScreenPosition  = {X=GetIntPart(NodeRadar:Width()/2), Y=GetIntPart(NodeRadar:Height()/2)}
    SpaceShip.CenterRectangle = {}
    SpaceShip:Format({
        Name      = "鹦鹉螺号",
        Position  = {X = 0.0, Y = 0.0},
        Speed     = {X = 0.0, Y = 0.0},
        Character = "x",
        ColorFg   = "blue"
    })
    SpaceShip.ColorBg   = ""
    local Warehouse     = {}
    SpaceShip.Warehouse = Warehouse

    return SpaceShip
end

function _SpaceShip.SetName(self, name)
    self.Info.Name = name
    NodeParGUserSpaceShipStatus:SetAttribute("borderlabel", self.Info.Name)
    NodeRadar:SetActive()
end

-- 刷新飞船为中心的指定大小区域所在的宇宙位置
function _SpaceShip.refreshCenterRectangle(self, rectangleWidth, rectangleHeight)
    local x = GetIntPart(self.Info.Position.X)
    local y = GetIntPart(self.Info.Position.Y)
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
    self.Info.Position.X = positionX
    self.Info.Position.Y = positionY
    self:refreshCenterRectangle(NodeRadar:Width(), NodeRadar:Height())
end

function _SpaceShip.Format(self, spaceshipInfo)
    self.Info     = {
        Name      = spaceshipInfo.Name,
        Position  = spaceshipInfo.Position,
        Speed     = spaceshipInfo.Speed,
        Character = spaceshipInfo.Character,
        ColorFg   = spaceshipInfo.ColorFg
    }
end
