local spriteComp = {}
spriteComp.__index = spriteComp

function spriteComp:new(image)
    local instance = {
        active = true,
        texture = image,
        name = "sprite"
    }
    setmetatable(instance, spriteComp)
    return instance
end

return spriteComp