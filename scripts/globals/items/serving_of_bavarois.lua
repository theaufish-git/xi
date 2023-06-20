-----------------------------------
-- ID: 5729
-- Item: serving_of_bavarois
-- Food Effect: 3Hrs, All Races
-----------------------------------
-- HP 20
-- Intelligence 3
-- hHP +3
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
    target:addStatusEffect(xi.effect.FOOD, 0, 0, 10800, 5729)
end

itemObject.onEffectGain = function(target, effect)
    target:addMod(xi.mod.HP, 20)
    target:addMod(xi.mod.INT, 3)
    target:addMod(xi.mod.HPHEAL, 3)
end

itemObject.onEffectLose = function(target, effect)
    target:delMod(xi.mod.HP, 20)
    target:delMod(xi.mod.INT, 3)
    target:delMod(xi.mod.HPHEAL, 3)
end

return itemObject
