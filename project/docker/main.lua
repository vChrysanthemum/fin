local unixsock = require("unixsock")
local json = require("json")

NodeSelectDockerImages = Node("SelectDockerImages")
NodeSelectDockerContainers = Node("SelectDockerContainers")

local dockerClient = unixsock.NewUnixSockClient("/var/run/docker.sock")
local ret

ret = dockerClient:Get("http://l/images/json?all=1")
ret = json.decode(ret)
NodeSelectDockerImages:SelectClearOptions()
for k, image in pairs(ret) do 
  NodeSelectDockerImages:SelectAppendOption(image.Id, image.RepoTags[1])
end

ret = dockerClient:Get("http://l/containers/json?all=1")
ret = json.decode(ret)
NodeSelectDockerContainers:SelectClearOptions()
for k, container in pairs(ret) do 
  NodeSelectDockerContainers:SelectAppendOption(container.Id,
  string.format("%-20s %-30s [%s]", container.Names[1], container.Image, container.State))
end
