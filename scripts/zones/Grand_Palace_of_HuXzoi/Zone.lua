-----------------------------------
-- Zone: Grand_Palace_of_HuXzoi (34)
-----------------------------------
local huxzoiGlobal = require('scripts/zones/Grand_Palace_of_HuXzoi/globals')
local ID = require('scripts/zones/Grand_Palace_of_HuXzoi/IDs')
require('scripts/globals/conquest')
-----------------------------------
local zoneObject = {}

zoneObject.onInitialize = function(zone)
    zone:registerTriggerArea(1, -507, -4, 697, -501, 4, 702)
    zone:registerTriggerArea(2, -102, -4, 541,  -97, 4, 546)
    zone:registerTriggerArea(3, -178, -4,  97, -173, 4, 103)
    zone:registerTriggerArea(4, -497, -4,  97, -492, 4, 102)
    zone:registerTriggerArea(5, -742, -4, 372, -736, 4, 379)
    zone:registerTriggerArea(6,  332, -4, 696,  338, 4, 702)
    zone:registerTriggerArea(7,  737, -4, 541,  742, 4, 546)
    zone:registerTriggerArea(8,  661, -4,  87,  667, 4, 103)
    zone:registerTriggerArea(9,  340, -4,  97,  347, 4, 102)
    zone:registerTriggerArea(10,  97, -4, 372,  103, 4, 378)

    huxzoiGlobal.pickTemperancePH()
end

zoneObject.onConquestUpdate = function(zone, updatetype, influence, owner, ranking, isConquestAlliance)
    xi.conq.onConquestUpdate(zone, updatetype, influence, owner, ranking, isConquestAlliance)
end

zoneObject.onZoneIn = function(player, prevZone)
    local cs = -1

    if
        player:getXPos() == 0 and
        player:getYPos() == 0 and
        player:getZPos() == 0
    then
        player:setPos(-20, -1.5, -355.482, 192)
    end

    player:setCharVar("Hu-Xzoi-TP", 0)

    return cs
end

zoneObject.afterZoneIn = function(player)
    player:entityVisualPacket("door")
    player:entityVisualPacket("dtuk")
    player:entityVisualPacket("2dor")
    player:entityVisualPacket("cryq")
end

zoneObject.onTriggerAreaEnter = function(player, triggerArea)
    if
        player:getLocalVar("Hu-Xzoi-TP") == 0 and
        player:getAnimation() == xi.anim.NONE
    then
        -- prevent 2cs at same time
        player:startEvent(149 + triggerArea:GetTriggerAreaID())
    end
end

zoneObject.onTriggerAreaLeave = function(player, triggerArea)
end

zoneObject.onEventUpdate = function(player, csid, option)
    if csid >= 150 and csid <= 159 then
        player:setLocalVar("Hu-Xzoi-TP", 1)
    end
end

zoneObject.onEventFinish = function(player, csid, option)
    if csid >= 150 and csid <= 159 then
        player:setLocalVar("Hu-Xzoi-TP", 0)
    end
end

zoneObject.onGameHour = function(zone)
end

return zoneObject
