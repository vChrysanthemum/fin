local _World = {}
local _mtWorld = {__index = _World} 

function InitWorld()
    local World = setmetatable({}, _mtWorld)
    World.LoopEventSig = nil
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

function _World.RenderArea(self, area)
    -- (area.Max.X-area.Min.X) * (area.Max.Y-area.Min.Y)
end
