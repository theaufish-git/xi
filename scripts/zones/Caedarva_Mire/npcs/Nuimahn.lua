-----------------------------------
-- Area: Caedarva Mire
--  NPC: Nuimahn
-- Type: Alzadaal Undersea Ruins
-- !pos  -380 0 -381 79
-----------------------------------
-----------------------------------
local entity = {}

entity.onTrade = function(player, npc, trade)
    if trade:getItemCount() == 1 and trade:hasItemQty(2185, 1) then
        player:tradeComplete()
        player:startEvent(203)
    end
end

entity.onTrigger = function(player, npc)
    if player:getZPos() < -281 then
        player:startEvent(204) -- leaving
    else
        player:startEvent(202) -- entering
    end
end

entity.onEventUpdate = function(player, csid, option)
end

entity.onEventFinish = function(player, csid, option)
    if csid == 203 then
        player:setPos(-515, -6.5, 740, 0, 72)
    end
end

return entity
