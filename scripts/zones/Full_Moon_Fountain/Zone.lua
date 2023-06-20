-----------------------------------
-- Zone: Full_Moon_Fountain (170)
-----------------------------------
local ID = require('scripts/zones/Full_Moon_Fountain/IDs')
require('scripts/globals/conquest')
-----------------------------------
local zoneObject = {}

zoneObject.onInitialize = function(zone)
end

zoneObject.onZoneIn = function(player, prevZone)
    local cs = -1

    if
        player:getXPos() == 0 and
        player:getYPos() == 0 and
        player:getZPos() == 0
    then
        player:setPos(-260.136, 2.09, -325.702, 188)
    end

    return cs
end

zoneObject.afterZoneIn = function(player)
    player:entityVisualPacket("kilk")
    player:entityVisualPacket("izum")
    player:entityVisualPacket("hast")
end

zoneObject.onConquestUpdate = function(zone, updatetype, influence, owner, ranking, isConquestAlliance)
    xi.conq.onConquestUpdate(zone, updatetype, influence, owner, ranking, isConquestAlliance)
end

zoneObject.onTriggerAreaEnter = function(player, triggerArea)
end

zoneObject.onEventUpdate = function(player, csid, option)
end

zoneObject.onEventFinish = function(player, csid, option)
end

return zoneObject
