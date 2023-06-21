-----------------------------------
-- ID: 5747
-- Item: Jar of Marinara Sauce
-- Food Effect: 5Min, All Races
-----------------------------------
-- Mind 2
-- Intelligence 1
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
    target:addStatusEffect(xi.effect.FOOD, 0, 0, 300, 5747)
end

itemObject.onEffectGain = function(target, effect)
    target:addMod(xi.mod.MND, 2)
    target:addMod(xi.mod.INT, 1)
end

itemObject.onEffectLose = function(target, effect)
    target:delMod(xi.mod.MND, 2)
    target:delMod(xi.mod.INT, 1)
end

return itemObject
