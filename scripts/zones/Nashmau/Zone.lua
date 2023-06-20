-----------------------------------
-- Zone: Nashmau (53)
-----------------------------------
local ID = require('scripts/zones/Nashmau/IDs')
require('scripts/globals/missions')
require('scripts/globals/zone')
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
        if prevZone == xi.zone.SILVER_SEA_ROUTE_TO_NASHMAU then
            cs = 201
            player:setPos(11, 2, -102, 128)
        else
            player:setPos(40.658, -7.527, -24.001, 128)
        end
    end

    return cs
end

zoneObject.onTriggerAreaEnter = function(player, triggerArea)
end

zoneObject.onTriggerAreaLeave = function(player, triggerArea)
end

zoneObject.onTransportEvent = function(player, transport)
    if transport == 59 then
        player:startEvent(200)
    end
end

zoneObject.onEventUpdate = function(player, csid, option)
end

zoneObject.onEventFinish = function(player, csid, option)
    if csid == 200 then
        player:setPos(0, -2, 0, 0, 59)
    end
end

return zoneObject
