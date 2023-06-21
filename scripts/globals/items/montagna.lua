-----------------------------------
-- ID: 5887
-- Item: montagna
-- Food Effect: 30Min, All Races
-----------------------------------
-- HP +8% (cap 140)
-- Increases rate of combat skill gains by 60%
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
    target:addStatusEffect(xi.effect.FOOD, 0, 0, 1800, 5887)
end

itemObject.onEffectGain = function(target, effect)
    target:addMod(xi.mod.FOOD_HPP, 8)
    target:addMod(xi.mod.FOOD_HP_CAP, 140)
    target:addMod(xi.mod.COMBAT_SKILLUP_RATE, 60)
end

itemObject.onEffectLose = function(target, effect)
    target:delMod(xi.mod.FOOD_HPP, 8)
    target:delMod(xi.mod.FOOD_HP_CAP, 140)
    target:delMod(xi.mod.COMBAT_SKILLUP_RATE, 60)
end

return itemObject
