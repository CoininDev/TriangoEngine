local renderSystem = {
    name = "render"
}

function renderSystem:new()
    local instance = {
        notableEntities = {},
    }
    setmetatable(instance, renderSystem)
    return instance
end

function renderSystem:Update(game, dt) end

function renderSystem:Refresh(game)
    self.notableEntities = {}
    for _, entity in ipairs(game.Entities) do
        if entity.components["sprite"] and entity.components["transform"] then
            table.insert(self.notableEntities, entity)
        end
    end
    print("renderSystem refreshed")
end

function renderSystem:Draw(game)
    love.graphics.clear()
    for _, entity in ipairs(self.notableEntities) do
        local sprite = entity.components["sprite"]
        local transform = entity.components["transform"]
        love.graphics.draw(sprite.texture, transform.position.x, transform.position.y, transform.rotation, transform.scale.x, transform.scale.y)
    end
end

return renderSystem