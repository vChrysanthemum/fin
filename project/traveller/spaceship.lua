local json = require("json")

local _Spaceship = {}
local _mtSpaceship = {__index = _Spaceship}

function GetSpaceshipFromDB(spaceshipId)
    sql = string.format([[
    select data from b_spaceship where spaceship_id=%d limit 1
    ]], spaceshipId)
    local rows = DB:Query(sql)
    local row = rows:FetchOne()
    rows:Close()
    if "table" ~= type(row) then
        return nil
    end
    local spaceship = NewSpaceship()
    spaceship:Format(json.decode(row.data))
    spaceship.Info.SpaceshipId = spaceshipId
    return spaceship
end

function NewSpaceship()
    local Spaceship           = setmetatable({}, _mtSpaceship)
    Spaceship.ScreenPosition  = {X=GetIntPart(NodeRadar:Width()/2), Y=GetIntPart(NodeRadar:Height()/2)}
    Spaceship.CenterRectangle = {}
    Spaceship:Format({
        SpaceshipId = nil,
        Name        = "鹦鹉螺号",
        Position    = {X          = 0.0, Y = 0.0},
        Speed       = {X          = 0.0, Y = 0.0},
        Character   = "x",
        ColorFg     = "blue"
    })
    Spaceship.ColorBg   = ""
    local Warehouse     = {}
    Spaceship.Warehouse = Warehouse

    return Spaceship
end

function _Spaceship.SetName(self, name)
    self.Info.Name = name
    self:FlushToDB()
    NodeParGUserSpaceshipStatus:SetAttribute("borderlabel", " " .. self.Info.Name .. " ")
    NodeRadar:SetActive()
end

-- 刷新飞船为中心的指定大小区域所在的宇宙位置
function _Spaceship.refreshCenterRectangle(self, rectangleWidth, rectangleHeight)
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
function _Spaceship.SetPosition(self, positionX, positionY)
    self.Info.Position.X = positionX
    self.Info.Position.Y = positionY
    self:refreshCenterRectangle(NodeRadar:Width(), NodeRadar:Height())
end

function _Spaceship.Format(self, spaceshipInfo)
    self.Info       = {
        SpaceshipId = spaceshipInfo.SpaceshipId,
        Name        = spaceshipInfo.Name,
        Position    = spaceshipInfo.Position,
        Speed       = spaceshipInfo.Speed,
        Character   = spaceshipInfo.Character,
        ColorFg     = spaceshipInfo.ColorFg
    }
end

function _Spaceship.FlushToDB(self)
    if "number" ~= type(self.Info.SpaceshipId) then
        return nil
    end

    sql = string.format([[
    update b_spaceship set data = '%s' where spaceship_id=%d
    ]], DB:QuoteSQL(json.encode(self.Info)), self.Info.SpaceshipId)
    local queryRet = DB:Exec(sql)
    if "string" == type(queryRet) then
        Log(queryRet)
    end
end
