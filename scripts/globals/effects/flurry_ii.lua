-----------------------------------
-- xi.effect.FLURRY_II
-----------------------------------
local effectObject = {}

effectObject.onEffectGain = function(target, effect)
    target:addMod(xi.mod.SNAP_SHOT, effect:getPower())
end

effectObject.onEffectTick = function(target, effect)
end

effectObject.onEffectLose = function(target, effect)
    target:delMod(xi.mod.SNAP_SHOT, effect:getPower())
end

return effectObject
