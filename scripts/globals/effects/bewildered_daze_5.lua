-----------------------------------
-- xi.effect.BEWILDERED_DAZE_5
-----------------------------------
local effectObject = {}

effectObject.onEffectGain = function(target, effect)
    target:addMod(xi.mod.CEVA, -13)
end

effectObject.onEffectTick = function(target, effect)
end

effectObject.onEffectLose = function(target, effect)
    target:delMod(xi.mod.CEVA, -13)
end

return effectObject
