local json = require("json")
local rwmutex = require("rwmutex")

local _World = {}
local _mtWorld = {__index = _World} 

function NewWorld()
    local World             = setmetatable({}, _mtWorld)
    World.LoopEventSig      = nil
    World.WorkBlockWidth    = 60
    World.WorkBlockHeight   = 20
    World.WorkBlockSize     = World.WorkBlockWidth * World.WorkBlockHeight
    World.WorkBlocksCount   = 20
    World.WorkBlockColumns  = 5
    World.WorkBlockRows     = World.WorkBlocksCount / World.WorkBlockColumns
    World.CreateBlockSize   = World.WorkBlockSize * World.WorkBlocksCount
    World.CreateBlockWidth  = World.WorkBlockWidth * World.WorkBlockColumns
    World.CreateBlockHeight = World.WorkBlockHeight * World.WorkBlockRows
    World.Planets           = {}
    World.EventLocker       = rwmutex.NewRWMutex()

    World.isShouldMove = true
    World.afterStopCallback = nil
    World.LeavePlanetSpeed = 0.3

    return World
end

function _World.loopEvent(self)
    if false == self.isShouldMove then 
        if nil ~= self.afterStopCallback then
            self.afterStopCallback()
            self.afterStopCallback = nil
        end
        return
    end

    GUserSpaceship:LoopEvent()

    NodeRadar:CanvasClean()
    GRadar:DrawSpaceship()
    GRadar:DrawPlanets()
    NodeRadar:CanvasDraw()
    GUserSpaceship:RefreshNodeParGUserSpaceshipStatus()
    return
end

function _World.Stop(self, afterStopCallback)
    self.isShouldMove = false
    self.afterStopCallback = afterStopCallback
end

function _World.Resume(self)
    self.isShouldMove = true
end

function _World.LoopEvent(self)
    self.LoopEventSig = SetInterval(200, function()
        self:loopEvent()
    end)
end

-- 生成指定区域内的星球
-- 如果已存在则返回
-- createBlockIndex 为区域索引 {X, Y}
function _World.initAreaPlanets(self, createBlockIndex)
    local key = PointToStr(createBlockIndex)

    if CheckTableHasKey(self.Planets, key) then
        return
    end

    local createBlockStartPosition = {
        X = createBlockIndex.X * self.CreateBlockWidth,
        Y = createBlockIndex.Y * self.CreateBlockHeight
    }

    local k = 0
    local planet = {}
    local planets = {}
    local sql
    local queryRet = nil

    sql = string.format([[
    select created_at from b_planets_block where x=%d and y=%d limit 1
    ]], createBlockStartPosition.X, createBlockStartPosition.Y)
    local rows = DB:Query(sql)
    local row = rows:FetchOne()
    rows:Close()
    if "table" == type(row) then
        sql = string.format([[
        select planet_id, data from b_planet where planets_block_x=%d and planets_block_y=%d
        ]], createBlockStartPosition.X, createBlockStartPosition.Y)
        rows = DB:Query(sql)
        while true do
            row = rows:FetchOne()
            if "table" ~= type(row) then
                break
            end
            planet = NewPlanet()
            planet:Format(json.decode(row.data))
            planet.Info.PlanetId = tonumber(row.planet_id)
            table.insert(planets, planet)
        end
        rows:Close()
        self.Planets[key] = planets
        return
    end

    RefreshRandomSeed()
    local i, columnIndex, rowIndex = 0, 0, 0
    local rectangle = {}
    local i = 0
    while columnIndex < self.WorkBlockColumns do
        rowIndex = 0
        while rowIndex < self.WorkBlockRows do
            i = i + 1
            rectangle = {
                Min = {
                    X = createBlockStartPosition.X + columnIndex * self.WorkBlockWidth,
                    Y = createBlockStartPosition.Y + rowIndex * self.WorkBlockHeight
                },
                Max = {
                    X = createBlockStartPosition.X + (columnIndex+1) * self.WorkBlockWidth,
                    Y = createBlockStartPosition.Y + (rowIndex+1) * self.WorkBlockHeight
                }
            }
            -- planetsPosition = {rectangle.Min, rectangle.Max}
            planetsPosition = InitRandomPoints(math.random(0, 8), rectangle)

            for _, planetPosition in pairs(planetsPosition) do
                planet = NewPlanet()
                planet:Initilize({X=planetPosition.X, Y=planetPosition.Y})
                table.insert(planets, planet)
            end

            rowIndex = rowIndex + 1
        end
        columnIndex = columnIndex + 1
    end

    for k, planet in pairs(planets) do
        sql = string.format([[
        insert into b_planet (planets_block_x, planets_block_y, data) values(%d, %d, '%s')
        ]], createBlockStartPosition.X, createBlockStartPosition.Y, DB:QuoteSQL(json.encode(planet.Info)))
        queryRet = DB:Exec(sql)
        planets[k].Info.PlanetId = queryRet:LastInsertId()
    end
    sql = string.format([[
    insert into b_planets_block(x, y, created_at) values(%d, %d, %d)
    ]], createBlockStartPosition.X, createBlockStartPosition.Y, TimeNow())
    DB:Exec(sql)

    self.Planets[key] = planets
end

-- 获取指定区域内的星球
-- rectangle 宇宙位置
function _World.GetPlanetsByRectangle(self, rectangle)
    local blockIndexs = {
        Min = {
            X = math.floor(rectangle.Min.X / self.CreateBlockWidth),
            Y = math.floor(rectangle.Min.Y / self.CreateBlockHeight),
        },
        Max = {
            X = math.floor(rectangle.Max.X / self.CreateBlockWidth),
            Y = math.floor(rectangle.Max.Y / self.CreateBlockHeight),
        }
    }

    if blockIndexs.Min.X == blockIndexs.Max.X and blockIndexs.Min.Y == blockIndexs.Max.Y then
        blockIndexs = {blockIndexs.Min}
    else
        local newblockIndexs = {}
        local columnsMax = blockIndexs.Max.X
        local rowsMax = blockIndexs.Max.Y
        local columnIndex = blockIndexs.Min.X
        local rowIndex = blockIndexs.Min.Y
        while columnIndex <= columnsMax do
            rowIndex = blockIndexs.Min.Y
            while rowIndex <= rowsMax do
                table.insert(newblockIndexs, {X=columnIndex, Y=rowIndex})
                rowIndex = rowIndex + 1
            end
            columnIndex = columnIndex + 1
        end
        blockIndexs = newblockIndexs
    end

    local key = nil
    local planets = {}
    for _, block in pairs(blockIndexs) do
        self:initAreaPlanets(block)
    end
    for _, block in pairs(blockIndexs) do
        key = PointToStr(block)
        for _, planet in pairs(self.Planets[key]) do

            if rectangle.Min.X <= planet.Info.Position.X and planet.Info.Position.X <= rectangle.Max.X and 
                rectangle.Min.Y <= planet.Info.Position.Y and planet.Info.Position.Y <= rectangle.Max.Y then
                table.insert(planets, planet)
            end

        end
    end

    return planets
end
