local json = require("json")

local _RobotCenter = {}
local _mtRobotCenter = {__index = _RobotCenter} 

function NewRobotCenter()
  local RobotCenter = setmetatable({}, _mtRobotCenter)
  RobotCenter.robotServiceAddressToRobot = {}
  RobotCenter:LoadRobotsFromDB()
  return RobotCenter
end

function _RobotCenter.RegisterRobot(self, robotServiceAddress, robot)
  self.robotServiceAddressToRobot[robotServiceAddress] = robot
end

function _RobotCenter.GetRobotByServiceAddress(self, robotServiceAddress)
  return self.robotServiceAddressToRobot[robotServiceAddress]
end

function _RobotCenter.LoadRobotsFromDB(self)
    sql = string.format([[
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
