-----------------------------------
-- Obfuscate
-- Description: A blinding wave hits players in an area of effect.
-- Type: Enfeebling
-- Utsusemi/Blink absorb: Ignores shadows
-- Range: 10' radial
-----------------------------------
require("scripts/globals/mobskills")
-----------------------------------
local mobskillObject = {}

mobskillObject.onMobSkillCheck = function(target, mob, skill)
    return 0
end

mobskillObject.onMobWeaponSkill = function(target, mob, skill)
    local typeEffect = xi.effect.BLINDNESS

    skill:setMsg(xi.mobskills.mobStatusEffectMove(mob, target, typeEffect, 20, 0, 120))

    return typeEffect
end

return mobskillObject
