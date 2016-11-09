local _Node = {}
local _mt = {__index = _Node}

function Node(id)
    local nodePointer = base.GetNodePointer(id)
    if nil == nodePointer then
        return nil
    end

    local ret = setmetatable({}, _mt)
    ret.nodePointer = nodePointer
    return ret
end

function _Node.SetActive(self)
    return base.NodeSetActive(self.nodePointer)
end

function _Node.GetHtmlData(self)
    return base.NodeGetHtmlData(self.nodePointer)
end

function _Node.SetText(self, text)
    return base.NodeSetText(self.nodePointer, text)
end

function _Node.GetValue(self)
    return base.NodeGetValue(self.nodePointer)
end

function _Node.OnKeyPressEnter(self, callback)
    return base.NodeOnKeyPressEnter(self.nodePointer, callback)
end

function _Node.Remove(self)
    return base.NodeRemove(self.nodePointer)
end

function WindowConfirm(title)
    content = string.format([[
    <table>
        <tr>
            <td offset=4 cols=4><par height=6></par></td>
        </tr>
        <tr>
            <td offset=4 cols=4><par>%s</par></td>
        </tr>
        <tr>
            <td offset=5 cols=2>
                <select id="SelectConfirm">
                    <option value="cancel">取消</option>
                    <option value="confirm">确定</option>
                </select>
            </td>
        </tr>
    </table>
    ]], title)
    return base.WindowConfirm(content)
end

function Log(content)
    base.Log(content)
end
