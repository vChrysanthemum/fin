local json = require("json")

local _Spaceship = {}
local _mtSpaceship = {__index = _Spaceship}

function GetSpaceshipFromDB(spaceshipId)
    local sql = string.format([[
    select data from b_spaceship where spaceship_id=%d limit 1
    ]], spaceshipId)
    local rows = DB:Query(sql)
    local row = rows:FetchOne()
    rows:Close()
    if "table" ~= type(row) then
        return nil
    end

    local spaceship = NewSpaceship()
    spaceship:Format(json.decode(row.data), spaceshipId)

    spaceship.LandingPlanetId = nil
    spaceship.NewestMsg = ""

    return spaceship
end

function NewSpaceship()
    local Spaceship           = setmetatable({}, _mtSpaceship)
    Spaceship.CenterRectangle = {}
    Spaceship.ScreenPosition  = {}
    Spaceship:Format({
        Name        = "鹦鹉螺号",
        Position    = {X = 0.0, Y = 0.0},
        Speed       = {X = 0.0, Y = 0.0},
        Character   = "x",
        ColorFg     = "blue",
        StartAt     = TimeNow(),
        Life        = 82,
        Fuel        = 100,
        Cabin       = {
            Jumpers     = 6,
            Resource    = 20,
        },
    }, nil)
    Spaceship.ColorBg   = ""

    Spaceship.Cabin = {}

    Spaceship.lastFlushToDBForRunOneStepAt = 0

    return Spaceship
end

-- 设置 x 方向速度
function _Spaceship.SetSpeedX(self, speedx)
    self:UpdateFuel(-1 * math.floor(math.abs(self.Info.Speed.X - speedx)))
    self.Info.Speed.X = speedx
    self:FlushToDB()
    self:RefreshNodeParGUserSpaceshipStatus()
end

-- 设置 y 方向速度
function _Spaceship.SetSpeedY(self, speedy)
    self:UpdateFuel(-1 * math.floor(math.abs(self.Info.Speed.Y - speedy)))
    self.Info.Speed.Y = speedy
    self:FlushToDB()
    self:RefreshNodeParGUserSpaceshipStatus()
end

-- 设置飞船名称
function _Spaceship.SetName(self, name)
    self.Info.Name = name
    self:FlushToDB()
    NodeParGUserSpaceshipStatus:SetAttribute("borderlabel", string.format(" %s状态 ", self.Info.Name))
    self:RefreshNodeParGUserSpaceshipStatus()
    NodeRadar:SetActive()
end

-- 刷新飞船为中心的指定大小区域所在的宇宙位置
function _Spaceship.refreshCenterRectangle(self)
    local rectangleWidth = GRadar.ScreenRectangleSize.Width
    local rectangleHeight = GRadar.ScreenRectangleSize.Height
    self.CenterRectangle = {}
    self.CenterRectangle.Min = {
        X = math.floor(self.Info.Position.X - rectangleWidth / 2),
        Y = math.floor(self.Info.Position.Y - rectangleHeight / 2)
    }

    -- 矫正屏幕位置
    self.ScreenPosition.X = math.floor(self.Info.Position.X) - self.CenterRectangle.Min.X
    self.ScreenPosition.Y = math.floor(self.Info.Position.Y) - self.CenterRectangle.Min.Y
    -- 如果飞船不在屏幕中央，则将其放回中央，并给 CenterRectangle 补差
    self.CenterRectangle.Min.X = self.CenterRectangle.Min.X + self.ScreenPosition.X - GRadar.ScreenCenterPosition.X
    self.CenterRectangle.Min.Y = self.CenterRectangle.Min.Y + self.ScreenPosition.Y - GRadar.ScreenCenterPosition.Y
    self.ScreenPosition.X = GRadar.ScreenCenterPosition.X
    self.ScreenPosition.Y = GRadar.ScreenCenterPosition.Y

    self.CenterRectangle.Max = {
        X = self.CenterRectangle.Min.X + rectangleWidth,
        Y = self.CenterRectangle.Min.Y + rectangleHeight,
    }

    return
end

function _Spaceship.Format(self, spaceshipInfo, spaceship_id)
    self.Info       = {
        SpaceshipId = tonumber(spaceship_id),
        Name        = spaceshipInfo.Name,
        Position    = spaceshipInfo.Position,
        Speed       = spaceshipInfo.Speed,
        Character   = spaceshipInfo.Character,
        ColorFg     = spaceshipInfo.ColorFg,
        StartAt     = spaceshipInfo.StartAt,
        Fuel        = spaceshipInfo.Fuel,
        Life        = spaceshipInfo.Life,
        Cabin = {
            Jumpers     = spaceshipInfo.Cabin.Jumpers,
            Resource    = spaceshipInfo.Cabin.Resource,
        },
    }
end

function _Spaceship.FlushToDB(self)
    if "number" ~= type(self.Info.SpaceshipId) then
        return nil
    end

    local sql = string.format([[
    update b_spaceship set data = '%s' where spaceship_id=%d
    ]], DB:QuoteSQL(json.encode(self.Info)), self.Info.SpaceshipId)
    local queryRet = DB:Exec(sql)
    if "string" == type(queryRet) then
        Log(queryRet)
    end
end

-- 更改燃料
function _Spaceship.UpdateFuel(self, number)
    local newFuelValue = self.Info.Fuel + number
    if newFuelValue < 0 then
        newFuelValue = 0
    end

    if math.abs(math.ceil(self.Info.Fuel) - math.ceil(newFuelValue)) < 1 then
        self.Info.Fuel = newFuelValue
        return
    end

    if newFuelValue > 100 then
        newFuelValue = 100
    end
    self.Info.Fuel = newFuelValue
    self:RefreshGaugeFuel()
    self:FlushToDB()
end

--- 更改生命值
function _Spaceship.UpdateLife(self, number)
    self.Info.Life = self.Info.Life + number
    if self.Info.Life < 0 then
        self.Info.Life = 0
    end
    self:RefreshGaugeLife()
    self:FlushToDB()
end

function _Spaceship.RefreshGaugeLife(self)
    NodeGaugeLife:SetAttribute("percent", tostring(self.Info.Life))
    if self.Info.Life < 20 then
        NodeGaugeLife:SetAttribute("percentcolor_highlighted", "black")
        NodeGaugeLife:SetAttribute("barcolor", "red")
    elseif self.Info.Life < 80 then
        NodeGaugeLife:SetAttribute("percentcolor_highlighted", "blue")
        NodeGaugeLife:SetAttribute("barcolor", "yellow")
    else
        NodeGaugeLife:SetAttribute("percentcolor_highlighted", "black")
        NodeGaugeLife:SetAttribute("barcolor", "green")
    end
end

function _Spaceship.RefreshGaugeFuel(self)
    NodeGaugeFuel:SetAttribute("percent", tostring(math.ceil(self.Info.Fuel)))
    if self.Info.Fuel < 20 then
        NodeGaugeFuel:SetAttribute("percentcolor_highlighted", "black")
        NodeGaugeFuel:SetAttribute("barcolor", "red")
    elseif self.Info.Fuel < 80 then
        NodeGaugeFuel:SetAttribute("percentcolor_highlighted", "blue")
        NodeGaugeFuel:SetAttribute("barcolor", "yellow")
    else
        NodeGaugeFuel:SetAttribute("percentcolor_highlighted", "white")
        NodeGaugeFuel:SetAttribute("barcolor", "blue")
    end
end

-- 刷新 NodeParGUserSpaceshipStatus 显示
function _Spaceship.RefreshNodeParGUserSpaceshipStatus(self)
    NodeParGUserSpaceshipStatus:SetValue(string.format([[
X: %f
Y: %f
速度X: %f/s
速度Y: %f/s
飞行历时: %d]], self.Info.Position.X, self.Info.Position.Y, self.Info.Speed.X, self.Info.Speed.Y, TimeNow() - self.Info.StartAt))
    
    local robotStr = {}
    for _, robot in ipairs(GRobotCenter.Robots) do
        if nil == RobotCorePlanetLanding then
            table.insert(robotStr, string.format("%s", robot.RobotCore.Info.Name))
        end
    end
    if TableLength(robotStr) > 0 then
        robotStr = table.concat(robotStr, ', ')
    else 
        robotStr = ''
    end
    NodeParGUserSpaceshipCabin:SetValue(string.format([[
时空跳跃者: %d
资源: %f
机器人: %s]], self.Info.Cabin.Jumpers, self.Info.Cabin.Resource, robotStr))
end

-- spaceship tools

-- 运行跳跃者
function _Spaceship.JumperRun(self, position)
    if self.Info.Cabin.Jumpers <= 0 then
        return "没有可用跳跃者"
    end

    self.Info.Position.X = position.X
    self.Info.Position.Y = position.Y
    self:refreshCenterRectangle()
    self.Info.Cabin.Jumpers = self.Info.Cabin.Jumpers - 1
    self:FlushToDB()

    return nil
end

function _Spaceship.SetNewestMsg(self, msg)
    self.NewestMsg = msg
    NodeParNewestMsg:SetValue(string.format("%s", self.NewestMsg))
end

-- 被星球捕获事件
function _Spaceship.EventCachedByPlanet(self, planet)
    self.Info.Speed.X = 0
    self.Info.Speed.Y = 0
    self.LandingPlanetId = planet.Info.PlanetId
    self:SetNewestMsg(string.format("飞船被 %s 引力捕获", planet.Info.Name))

    self:FlushToDB()
end

-- 离开星球捕获事件
function _Spaceship.EventLeavePlanet(self)
    local planet = GWorld:GetPlanetByPlanetId(self.LandingPlanetId)
    self:SetNewestMsg(string.format("飞船离开 %s", planet.Info.Name))
    self.LandingPlanetId = nil

    self:FlushToDB()
end

-- 飞船飞行一次，改变 position
function _Spaceship.RunOneStep(self)
    self.Info.Position.X = self.Info.Position.X + self.Info.Speed.X
    self.Info.Position.Y = self.Info.Position.Y + self.Info.Speed.Y
    self:refreshCenterRectangle()

    if nil == self.LandingPlanetId then
        self:UpdateFuel(-0.01)
    end

    if TimeNow() - self.lastFlushToDBForRunOneStepAt > 3 then
        self.lastFlushToDBForRunOneStepAt = TimeNow()
        self:FlushToDB()
    end
end

function _Spaceship.LoopEvent(self)
    local planet = {}

    -- 检查飞船是否被星球引力捕获
    if nil == self.LandingPlanetId then
        planet = GRadar:GetPlanetOnScreenByScreenPosition(self.ScreenPosition)
        if nil ~= planet and
            math.abs(self.Info.Speed.X) < GWorld.LeavePlanetSpeed and
            math.abs(self.Info.Speed.Y) < GWorld.LeavePlanetSpeed then
            GUserSpaceship:EventCachedByPlanet(planet)
            return
        end
    end

    -- 检查飞船是否离开星球
    if nil ~= self.LandingPlanetId then
        planet = GWorld:GetPlanetByPlanetId(self.LandingPlanetId)
        if nil ~= planet then
            -- 检查飞船是否会被星球引力捕获
            if (math.abs(self.Info.Speed.X) > 0 or
                math.abs(self.Info.Speed.Y) > 0) and
                math.abs(self.Info.Speed.X) < GWorld.LeavePlanetSpeed and
                math.abs(self.Info.Speed.Y) < GWorld.LeavePlanetSpeed then
                GUserSpaceship:EventCachedByPlanet(planet)
                return
            end

            -- 检查飞船是否足够动力离开星球
            if (math.abs(self.Info.Speed.X) >= GWorld.LeavePlanetSpeed or
                math.abs(self.Info.Speed.Y) >= GWorld.LeavePlanetSpeed) then
                GUserSpaceship:EventLeavePlanet()
            end
        end
    end

    GUserSpaceship:RunOneStep()
    GRadar:RefreshScreenPlanets(GWorld:GetPlanetsByRectangle(self.CenterRectangle),
    self.CenterRectangle)
end

-- 更改仓库资源数量
function _Spaceship.ChangeCabinResource(self, delta)
    if self.Info.Cabin.Resource + delta < 0 then
        return "资源不足"
    end
    self.Info.Cabin.Resource = self.Info.Cabin.Resource + delta
    self:FlushToDB()
    return true
end
