-----------------------------------
-- Zone: Chamber_of_Oracles (168)
-----------------------------------
local ID = require('scripts/zones/Chamber_of_Oracles/IDs')
require('scripts/globals/conquest')
require('scripts/globals/missions')
require('scripts/globals/zone')
-----------------------------------
local zoneObject = {}

zoneObject.onInitialize = function(zone)
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
        player:setPos(-177.804, -2.765, -37.893, 179)
    end

    return cs
end

zoneObject.onTriggerAreaEnter = function(player, triggerArea)
end

zoneObject.onEventUpdate = function(player, csid, option)
end

zoneObject.onEventFinish = function(player, csid, option)
end

return zoneObject
