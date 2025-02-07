local entity = {}

function entity:new(game, comps)
    local instance = {
        game = game,
        components = comps
    }
    setmetatable(instance, entity)
    return instance
end

function entity:AddComponent(comp)
    entity.components[comp.name] = comp
end

function entity:RemoveComponent(comp)
    entity.components[comp.name] = nil
end

return entity