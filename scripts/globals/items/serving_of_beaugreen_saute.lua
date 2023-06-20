-----------------------------------
-- ID: 4572
-- Item: serving_of_beaugreen_sautee
-- Food Effect: 180Min, All Races
-----------------------------------
-- Agility 1
-- Vitality -1
-- Ranged Acc % 7
-- Ranged Acc Cap 15
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
    target:addStatusEffect(xi.effect.FOOD, 0, 0, 10800, 4572)
end

itemObject.onEffectGain = function(target, effect)
    target:addMod(xi.mod.AGI, 1)
    target:addMod(xi.mod.VIT, -1)
    target:addMod(xi.mod.FOOD_RACCP, 7)
    target:addMod(xi.mod.FOOD_RACC_CAP, 15)
end

itemObject.onEffectLose = function(target, effect)
    target:delMod(xi.mod.AGI, 1)
    target:delMod(xi.mod.VIT, -1)
    target:delMod(xi.mod.FOOD_RACCP, 7)
    target:delMod(xi.mod.FOOD_RACC_CAP, 15)
end

return itemObject
