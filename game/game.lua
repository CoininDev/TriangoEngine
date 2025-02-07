local game = {
    Entities = {},
    Systems = {},
}
game.__index = game

function game:new()
    local instance = {
        Entities = {},
        Systems = {},
    }
    setmetatable(instance, game)
    return instance
end

function game:Init()
    for _, system in ipairs(game.Systems) do
        system:Refresh(game)
    end
end

function game:Update(dt)
    for _, system in ipairs(game.Systems) do
        system:Update(game, dt)
    end
end

function game:Draw()
    for _, system in ipairs(game.Systems) do
        system:Draw(game)
    end
end

function game:AddEntity(entity)
    table.insert(game.Entities, entity)
    for _, system in ipairs(game.Systems) do
        system:Refresh(game)
    end
end

function game:RemoveEntity(entity)
    for i, e in ipairs(game.Entities) do
        if e == entity then
            table.remove(game.Entities, i)
            for _, system in ipairs(game.Systems) do
                system:Refresh(game)
            end
            break
        end
    end
end


return game