local json = require("json")

local _World = {}
local _mtWorld = {__index = _World} 

function NewWorld()
    local World = setmetatable({}, _mtWorld)
    World.LoopEventSig = nil
    World.WorkBlockWidth = 60
    World.WorkBlockHeight = 20
    World.WorkBlockSize = World.WorkBlockWidth * World.WorkBlockHeight
    World.WorkBlocksCount = 20
    World.WorkBlockColumns = 5
    World.WorkBlockRows = World.WorkBlocksCount / World.WorkBlockColumns
    World.CreateBlockSize = World.WorkBlockSize * World.WorkBlocksCount
    World.CreateBlockWidth = World.WorkBlockWidth * World.WorkBlockColumns
    World.CreateBlockHeight = World.WorkBlockHeight * World.WorkBlockRows
    World.Planets = {}
    return World
end

function _World.loopEvent(self)
    return
end

function _World.LoopEvent(self)
    self.LoopEventSig = SetInterval(200, function()
        World:loopEvent()
    end)
    SetTimeout(3000, function()
        SendCancelSig(World.LoopEventSig)
    end)
end

function _World.getPlanetsKey(self, block)
    return tostring(block.X) .. ":" .. tostring(block.Y)
end

-- 生成指定区域内的星球
-- 如果已存在则返回
-- createBlockIndex 为区域索引 {X, Y}
function _World.initAreaPlanets(self, createBlockIndex)
    local key = self:getPlanetsKey(createBlockIndex)

    if CheckTableHasKey(self.Planets, key) then
        return
    end

    local createBlockStartPosition = {
        X = createBlockIndex.X * self.CreateBlockWidth,
        Y = createBlockIndex.Y * self.CreateBlockHeight
    }

    RefreshRandomSeed()
    local i, columnIndex, rowIndex = 0, 0, 0
    local rectangle = {}
    local planet = {}
    local planets = {}
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

    self.Planets[key] = planets
end

-- 获取指定区域内的星球
-- rectangle 宇宙位置
function _World.GetPlanetsByRectangle(self, rectangle)
    local blockIndexs = {
        Min = {
            X = GetIntPart(rectangle.Min.X / self.CreateBlockWidth),
            Y = GetIntPart(rectangle.Min.Y / self.CreateBlockHeight),
        },
        Max = {
            X = GetIntPart(rectangle.Max.X / self.CreateBlockWidth),
            Y = GetIntPart(rectangle.Max.Y / self.CreateBlockHeight),
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
        key = self:getPlanetsKey(block)
        for _, planet in pairs(self.Planets[key]) do

            if rectangle.Min.X <= planet.Position.X and planet.Position.X <= rectangle.Max.X and 
                rectangle.Min.Y <= planet.Position.Y and planet.Position.Y <= rectangle.Max.Y then
                table.insert(planets, planet)
            end

        end
    end

    return planets
end
