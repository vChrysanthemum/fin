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
    spaceship.Fuel = 100
    return spaceship
end

function NewSpaceship()
    local Spaceship           = setmetatable({}, _mtSpaceship)
    Spaceship.ScreenPosition  = {X=GetIntPart(NodeRadar:Width()/2), Y=GetIntPart(NodeRadar:Height()/2)}
    Spaceship.CenterRectangle = {}
    Spaceship:Format({
        SpaceshipId = nil,
        Name        = "鹦鹉螺号",
        Position    = {X = 0.0, Y = 0.0},
        Speed       = {X = 0.0, Y = 0.0},
        Character   = "x",
        ColorFg     = "blue",
        StartAt     = 0
    })
    Spaceship.ColorBg   = ""
    local Warehouse     = {}
    Spaceship.Warehouse = Warehouse

    return Spaceship
end

function _Spaceship.SetSpeedX(self, speedx)
    self:UpdateFuel(-1 * GetIntPart(math.abs(self.Info.Speed.X - speedx)))
    self.Info.Speed.X = speedx
    self:FlushToDB()
    self:RefreshNodeParGUserSpaceshipStatus()
end

function _Spaceship.SetSpeedY(self, speedy)
    self:UpdateFuel(-1 * GetIntPart(math.abs(self.Info.Speed.Y - speedy)))
    self.Info.Speed.Y = speedy
    self:FlushToDB()
    self:RefreshNodeParGUserSpaceshipStatus()
end

function _Spaceship.SetName(self, name)
    self.Info.Name = name
    self:FlushToDB()
    NodeParGUserSpaceshipStatus:SetAttribute("borderlabel", " " .. self.Info.Name .. " ")
    self:RefreshNodeParGUserSpaceshipStatus()
    NodeRadar:SetActive()
end

function _Spaceship.RefreshNodeParGUserSpaceshipStatus(self)
    NodeParGUserSpaceshipStatus:SetText(string.format([[
X: %d
Y: %d
速度X: %f/s
速度Y: %f/s
船员: 7
飞行历时: 09:53

仓库:
炮弹: 612
时空跳跃者:3
    ]], self.Info.Position.X, self.Info.Position.Y, self.Info.Speed.X, self.Info.Speed.Y))
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

function _Spaceship.Format(self, spaceshipInfo)
    self.Info       = {
        SpaceshipId = spaceshipInfo.SpaceshipId,
        Name        = spaceshipInfo.Name,
        Position    = spaceshipInfo.Position,
        Speed       = spaceshipInfo.Speed,
        Character   = spaceshipInfo.Character,
        ColorFg     = spaceshipInfo.ColorFg,
        StartAt     = spaceshipInfo.StartAt
    }
    self:refreshCenterRectangle(NodeRadar:Width(), NodeRadar:Height())
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

-- 飞船飞行，改变 position
function _Spaceship.RunOneStep(self)
    self.Info.Position.X = self.Info.Position.X + self.Info.Speed.X
    self.Info.Position.Y = self.Info.Position.Y + self.Info.Speed.Y
    self:refreshCenterRectangle(NodeRadar:Width(), NodeRadar:Height())
end

function _Spaceship.UpdateFuel(self, number)
    self.Fuel = self.Fuel + number
    NodeGaugeFuel:SetAttribute("percent", tostring(self.Fuel))
    if self.Fuel < 20 then
        NodeGaugeFuel:SetAttribute("percentcolor_highlighted", "blue")
        NodeGaugeFuel:SetAttribute("barcolor", "red")
    elseif self.Fuel < 50 then
        NodeGaugeFuel:SetAttribute("percentcolor_highlighted", "blue")
        NodeGaugeFuel:SetAttribute("barcolor", "yellow")
    elseif self.Fuel < 80 then
        NodeGaugeFuel:SetAttribute("percentcolor_highlighted", "blue")
        NodeGaugeFuel:SetAttribute("barcolor", "green")
    else
        NodeGaugeFuel:SetAttribute("percentcolor_highlighted", "white")
        NodeGaugeFuel:SetAttribute("barcolor", "blue")
    end
end
