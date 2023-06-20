-----------------------------------
-- ID: 15842
-- Recall ring: Pashhow
-- Enchantment: "Recall-Pashhow"
-----------------------------------
require("scripts/globals/teleports")
-----------------------------------
local itemObject = {}

itemObject.onItemCheck = function(target)
    local result = 0
    if not target:hasKeyItem(xi.ki.PASHHOW_GATE_CRYSTAL) then
        result = 445
    end

    return result
end

itemObject.onItemUse = function(target)
    target:addStatusEffectEx(xi.effect.TELEPORT, 0, xi.teleport.id.PASHH, 0, 1)
end

return itemObject
