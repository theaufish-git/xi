-----------------------------------
-- ID: 5236
-- Item: Clump of Shungiku
-- Food Effect: 5Min, All Races
-----------------------------------
-- Agility +1
-- Vitality -1
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
    target:addStatusEffect(xi.effect.FOOD, 0, 0, 300, 5236)
end

itemObject.onEffectGain = function(target, effect)
    target:addMod(xi.mod.AGI, 1)
    target:addMod(xi.mod.VIT, -1)
end

itemObject.onEffectLose = function(target, effect)
    target:delMod(xi.mod.AGI, 1)
    target:delMod(xi.mod.VIT, -1)
end

return itemObject
