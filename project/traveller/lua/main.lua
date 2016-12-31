WorldLoop = function()
end 

-- node tab main
NodeRadar                   = Node("CanvasRadar")
NodeTerminalMain            = Node("TerminalMain")
NodeParInfo                 = Node("ParInfo")
NodeParNewestMsg            = Node("ParNewestMsg")
NodeParGUserSpaceshipStatus = Node("ParGUserSpaceshipStatus")
NodeParGUserSpaceshipCabin  = Node("ParGUserSpaceshipCabin")
NodeGaugeFuel               = Node("GaugeFuel")
NodeGaugeLife               = Node("GaugeLife")
NodeTabpaneMain             = Node("TabpaneMain")

-- node tab planet
NodeTabPlanetParPlanetInfo   = Node("TabPlanetParPlanetInfo")
NodeTabPlanetTerminalMain    = Node("TabPlanetTerminalMain")

-- 初始化所有全局变量
NodeRadar:SetAttribute("height", tostring(WindowHeight()-NodeTerminalMain:Height()))
GRadar         = NewRadar()
GTerminal      = NewTerminal(NodeTerminalMain)
GWorld         = NewWorld()
GUserSpaceship = GetSpaceshipFromDB(1)
GTabTerminal   = NewTerminal(NodeTabPlanetTerminalMain)
GUserSpaceship:RefreshNodeParGUserSpaceshipStatus()
GUserSpaceship:RefreshGaugeLife()
GUserSpaceship:RefreshGaugeFuel()

-- 调整页面组件大小
NodeTabPlanetParPlanetInfo:SetAttribute("height", tostring(WindowHeight()-NodeTabPlanetTerminalMain:Height()))

-- 初始化所有信号
GTerminal.CommandSig = NodeTerminalMain:TerminalRegisterCommandHandle(function(nodePointer, command)
  GTerminal:ExecCommand(nodePointer, command)
end)

GTabTerminal.TabPlanetCommandSig = NodeTabPlanetTerminalMain:TerminalRegisterCommandHandle(function(nodePointer, command)
    GTabTerminal:ExecCommand(nodePointer, command)
end)

function StartTabMain()
  NodeTabpaneMain:TabpaneSetActiveTab("main")
end

-- 世界开始运行
NodeParGUserSpaceshipStatus:SetAttribute("borderlabel", string.format(" %s状态 ", GUserSpaceship.Info.Name))
NodeTerminalMain:SetActive()
GWorld:LoopEvent()

--[[
SetTimeout(200, function()
    GTerminal.CmdExcuter["/planet"]:ShowPlanetDetail()
end)
]]
