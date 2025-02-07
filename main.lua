local game = require "game.game"
local entity = require "game.entity"

local spriteComp = require "comps.sprite"
local transformComp = require "comps.transform"

local renderSystem = require "systems.render"

function love.load()
    G = game.new(G)
    table.insert(G.Systems, renderSystem.new())
    game.AddEntity(entity.new(game, {
        transformComp.new({ x = 0, y = 0 }, 0, { x = 1, y = 1 }),
        spriteComp.new(love.graphics.newImage("assets/images/player.jpg")),
    }))

    game.Init(G)
end

function love.update(dt)
    game.Update(G, dt)
end

function love.draw()
    game.Draw(G)
end

