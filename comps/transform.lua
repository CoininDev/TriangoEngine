local transformComp = {}

function transformComp:new(pos, rot, scale)
    local instance = {
        active = true,
        position = pos,
        rotation = rot,
        scale = scale,
        name = "transform"
    }
    setmetatable(instance, transformComp)
    return instance
end

return transformComp