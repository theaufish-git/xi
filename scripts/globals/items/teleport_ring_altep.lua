-----------------------------------
-- ID: 14666
-- Teleport ring: Altep
-- Enchantment: "Teleport-Altep"
-----------------------------------
require("scripts/globals/teleports")
-----------------------------------
local itemObject = {}

itemObject.onItemCheck = function(target)
    local result = 0
    if not target:hasKeyItem(xi.ki.ALTEPA_GATE_CRYSTAL) then
        result = 445
    end

    return result
end

itemObject.onItemUse = function(target)
    target:addStatusEffectEx(xi.effect.TELEPORT, 0, xi.teleport.id.ALTEP, 0, 1)
end

return itemObject
