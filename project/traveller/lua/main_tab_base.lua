NodeTabBaseSelectEnterMain = Node("TabBaseSelectEnterMain")

NodeTabBaseSelectEnterMainKeyPressEnterKeySig = 
NodeTabBaseSelectEnterMain:RegisterKeyPressEnterHandler(function(nodePointer)
  NodeTabpaneMain:TabpaneSetActiveTab("main")
end)
