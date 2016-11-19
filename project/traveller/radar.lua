local _Radar = {}
local _mtRadar = {__index = _Radar} 

function NewRadar()
    local Radar = setmetatable({}, _mtRadar)
    Radar.KeyPressEnterChans = {}
    Radar.ScreenPlanets = {}
    Radar.CursorScreenPosition = {X=GetIntPart(NodeRadar:Width()/2), Y=GetIntPart(NodeRadar:Height()/2)}

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

    local isMove = false
    if "<left>" == keyStr then
        isMove = true
        self.CursorScreenPosition.X = self.CursorScreenPosition.X - 1
    elseif "<right>" == keyStr then
        isMove = true
        self.CursorScreenPosition.X = self.CursorScreenPosition.X + 1
    elseif "<up>" == keyStr then
        isMove = true
        self.CursorScreenPosition.Y = self.CursorScreenPosition.Y - 1
    elseif "<down>" == keyStr then
        isMove = true
        self.CursorScreenPosition.Y = self.CursorScreenPosition.Y + 1
    end

    if true == isMove then
        self:renewCursor()
    end
end

-- 计算星球所在屏幕的位置
-- rectangle 为指定宇宙位置
function _Radar.renewCursor(self)
    self.CursorScreenPosition.X, self.CursorScreenPosition.Y =
    NodeRadar:SetCursor(self.CursorScreenPosition.X, self.CursorScreenPosition.Y)
end

-- 计算星球所在屏幕的位置
-- rectangle 为指定宇宙位置
function _Radar.calculatePlanetScreenPosition(self, rectangle)
    local startPosition = {
        X = rectangle.Min.X,
        Y = rectangle.Min.Y
    }
    for k, _ in pairs(self.ScreenPlanets) do
        self.ScreenPlanets[k].ScreenPosition.X = self.ScreenPlanets[k].Position.X - startPosition.X
        self.ScreenPlanets[k].ScreenPosition.Y = self.ScreenPlanets[k].Position.Y - startPosition.Y
    end
end

-- 更新 Radar 的 ScreenPlanets
-- ScreenPlanets 新的屏幕上需要显示的 planets
-- rectangle 屏幕上显示宇宙位置区域
function _Radar.refreshScreenPlanets(self, planets, rectangle)
    self.ScreenPlanets = planets
    self:calculatePlanetScreenPosition(rectangle)
end

-- 画指定区域内的的星球
function _Radar.DrawPlanets(self, planets, rectangle)
    self:refreshScreenPlanets(planets, rectangle)

    for _, planet in pairs(self.ScreenPlanets) do
        NodeRadar:CanvasSet(
        planet.ScreenPosition.X,
        planet.ScreenPosition.Y,
        "*", "blue", "")
    end

    NodeRadar:CanvasDraw()
end
