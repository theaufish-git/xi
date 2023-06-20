-----------------------------------
-- ID: 6498
-- Item: Bunch of Fortune Fruits
-- Food Effect: 30Min, All Races
-----------------------------------
-- Charisma +7
-- may have unknown hidden effects
-----------------------------------
require("scripts/globals/msg")
-----------------------------------
local itemObject = {}

itemObject.onItemCheck = function(target)
    local result = 0
    if target:hasStatusEffect(xi.effect.FOOD) then
        result = xi.msg.basic.IS_FULL
    end

    return result
end

itemObject.onItemUse = function(target)
    target:addStatusEffect(xi.effect.FOOD, 0, 0, 1800, 6498)
end

itemObject.onEffectGain = function(target, effect)
    target:addMod(xi.mod.CHR, 7)
end

itemObject.onEffectLose = function(target, effect)
    target:delMod(xi.mod.CHR, 7)
end

return itemObject
