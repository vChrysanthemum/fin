function PointToStr(point)
    return tostring(point.X) .. ":" .. tostring(point.Y)
end

function InitRandomPoints(max, rectangle)
    local point, points = {}, {}
    local i = 0
    local isConflict = false
    while i < max do
        isConflict = false
        point = {
            X = math.random(rectangle.Min.X, rectangle.Max.X), 
            Y = math.random(rectangle.Min.Y, rectangle.Max.Y)
        }

        for _, v in pairs(points) do
            if point.X == v.X and point.Y == v.Y then
                isConflict = true
                break
            end
        end

        if false == isConflict then
            i = i + 1
            table.insert(points, point)
        end
    end

    return points
end
