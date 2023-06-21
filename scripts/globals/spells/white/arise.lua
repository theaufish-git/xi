-----------------------------------
-- Spell: Arise
-----------------------------------
require("scripts/globals/msg")
-----------------------------------
local spellObject = {}

spellObject.onMagicCastingCheck = function(caster, target, spell)
    return 0
end

spellObject.onSpellCast = function(caster, target, spell)
    if target:isPC() then
        target:sendRaise(4)
    else
        if target:getName() == "Prishe" then
            -- CoP 8-4 Prishe
            target:setLocalVar("Raise", 1)
            target:entityAnimationPacket("sp00")
            target:addHP(target:getMaxHP())
            target:addMP(target:getMaxMP())
        end
    end

    spell:setMsg(xi.msg.basic.MAGIC_CASTS_ON)

    return 4
end

return spellObject
