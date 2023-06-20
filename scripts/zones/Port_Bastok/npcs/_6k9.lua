-----------------------------------
-- Area: Port Bastok
--  NPC: Door: Arrivals Entrance
-- !pos -80 1 -26 236
-----------------------------------
local entity = {}

entity.onTrade = function(player, npc, trade)
end

entity.onTrigger = function(player, npc)
    player:startEvent(140)
    return 1
end

entity.onEventUpdate = function(player, csid, option)
end

entity.onEventFinish = function(player, csid, option)
end

return entity
