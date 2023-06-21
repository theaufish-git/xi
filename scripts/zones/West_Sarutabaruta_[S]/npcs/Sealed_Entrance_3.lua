-----------------------------------
-- Area: West Sarutabaruta [S]
--  NPC: Sealed Entrance (Sealed_Entrance_3)
-- !pos -340.000 1.825 -364.825 95
-----------------------------------
local ID = require("scripts/zones/West_Sarutabaruta_[S]/IDs")
require("scripts/globals/quests")
require("scripts/globals/utils")
-----------------------------------
local entity = {}

entity.onTrigger = function(player, npc)
    local snakeOnThePlains = player:getQuestStatus(xi.quest.log_id.CRYSTAL_WAR, xi.quest.id.crystalWar.SNAKE_ON_THE_PLAINS)
    local maskBit1 = utils.mask.getBit(player:getCharVar("SEALED_DOORS"), 0)
    local maskBit2 = utils.mask.getBit(player:getCharVar("SEALED_DOORS"), 1)
    local maskBit3 = utils.mask.getBit(player:getCharVar("SEALED_DOORS"), 2)

    if
        snakeOnThePlains == QUEST_ACCEPTED and
        player:hasKeyItem(xi.ki.ZONPA_ZIPPAS_ALL_PURPOSE_PUTTY)
    then
        if not maskBit3 then
            player:setCharVar("SEALED_DOORS", utils.mask.setBit(player:getCharVar("SEALED_DOORS"), 2, true))

            if not maskBit2 or not maskBit1 then
                player:messageSpecial(ID.text.DOOR_OFFSET + 1, xi.ki.ZONPA_ZIPPAS_ALL_PURPOSE_PUTTY)
            else
                player:messageSpecial(ID.text.DOOR_OFFSET + 4, xi.ki.ZONPA_ZIPPAS_ALL_PURPOSE_PUTTY)
                player:delKeyItem(xi.ki.ZONPA_ZIPPAS_ALL_PURPOSE_PUTTY)
            end
        else
            player:messageSpecial(ID.text.DOOR_OFFSET + 2, xi.ki.ZONPA_ZIPPAS_ALL_PURPOSE_PUTTY)
        end
    elseif snakeOnThePlains == QUEST_COMPLETED then
        player:messageSpecial(ID.text.DOOR_OFFSET + 2, xi.ki.ZONPA_ZIPPAS_ALL_PURPOSE_PUTTY)
    else
        player:messageSpecial(ID.text.DOOR_OFFSET + 3)
    end
end

entity.onEventUpdate = function(player, csid, option)
end

entity.onEventFinish = function(player, csid, option)
end

return entity
