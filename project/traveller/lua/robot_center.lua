local json = require("json")

local _RobotCenter = {}
local _mtRobotCenter = {__index = _RobotCenter} 

function NewRobotCenter()
  local RobotCenter = setmetatable({}, _mtRobotCenter)
  -- RobotCenter.robotServiceAddressToRobot = {}
  RobotCenter.Robots = {}
  RobotCenter:LoadRobotsFromDB()
  return RobotCenter
end

function _RobotCenter.RegisterRobot(self, robotServiceAddress, robot)
  self.Robots[robotServiceAddress] = robot
end

function _RobotCenter.GetRobotByServiceAddress(self, robotServiceAddress)
  return self.Robots[robotServiceAddress]
end

function _RobotCenter.LoadRobotsFromDB(self)
    local sql = string.format([[
    select robot_id, data from b_robot
    ]])
    local rows = DB:Query(sql)
    local row = nil
    local robotCore
    while true do
      row = rows:FetchOne()
      if "table" ~= type(row) then
        break
      end
      robotCore = NewRobotCore()
      robotCore:Format(json.decode(row.data), row.robot_id)
      self:RegisterRobot(robotCore.Info.ServiceAddress, robotCore.Robot)
    end
    rows:Close()
end

function _RobotCenter.LoopEvent(self)
    for k, robot in pairs(GRobotCenter.Robots) do
        robot:LoopEvent()
    end
end
