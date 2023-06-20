-----------------------------------
-- Area: The Shrine of Ru'Avitau
--  NPC: ??? (Spawn Olla Pequena)
-- !pos 851 0.1 92 178
-----------------------------------
local ID = require("scripts/zones/The_Shrine_of_RuAvitau/IDs")
-----------------------------------
local entity = {}

entity.onTrade = function(player, npc, trade)
    if trade:hasItemQty(1195, 1) and trade:getItemCount() == 1 then -- Trade Ro'Maeve Water
        for i = ID.mob.OLLAS_OFFSET, ID.mob.OLLAS_OFFSET + 2 do
            if GetMobByID(i):isSpawned() then
                return
            end
        end

        player:tradeComplete()
        SpawnMob(ID.mob.OLLAS_OFFSET):updateClaim(player)
        npc:setStatus(xi.status.DISAPPEAR)
    end
end

entity.onTrigger = function(player, npc)
    player:messageSpecial(ID.text.SMALL_HOLE_HERE)
end

entity.onEventUpdate = function(player, csid, option)
end

entity.onEventFinish = function(player, csid, option)
end

return entity
