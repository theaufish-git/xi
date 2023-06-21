-----------------------------------
-- Mix: Gold Needle - Removes Petrification.
-----------------------------------
local mobskillObject = {}

mobskillObject.onMobSkillCheck = function(target, mob, skill)
    return 0
end

mobskillObject.onMobWeaponSkill = function(target, mob, skill)
    target:delStatusEffect(xi.effect.PETRIFICATION)
    return 0
end

return mobskillObject
