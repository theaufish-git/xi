package templates

import (
	"log"
	"text/template"
)

const (
	mapConfig string = `-----------------------------------
-- MAP SERVER SETTINGS
-----------------------------------
-- All settings are attached to the 'xi.settings' object. This is published globally, and be accessed from C++ and any script.
--
-- This file is concerned mainly with game administration and configuring the map executable
-----------------------------------

xi          = xi or {}
xi.settings = xi.settings or {}

xi.settings.map =
{
	-- --------------------------------
	-- Packet settings
	-- --------------------------------

	MAX_TIME_LASTUPDATE = {{ .MaxTimeLastUpdate.Seconds }},

	-- --------------------------------
	-- SQL settings
	-- --------------------------------

	-- Used by serverutils::PersistServerVar() for the maximum attempts to retry verification
	-- of a written Server Variable.
	SETVAR_RETRY_MAX = {{ .SetvarRetryMax }},

	-- --------------------------------
	-- Game settings
	-- --------------------------------

	-- PacketGuard will block and report any packets that aren't in the allow-list for a
	-- player's current state.
	PACKETGUARD_ENABLED = {{ .AntiCheat.PacketGuardEnabled }},

	-- Minimal number of 0x3A packets which uses for detect lightluggage (set 0 for disable)
	LIGHTLUGGAGE_BLOCK = {{ .AntiCheat.LightLuggageBlock }},

	-- Enable or disable Recycle Bin (Set to false for items to be dropped immediately)
	ENABLE_ITEM_RECYCLE_BIN = {{ .EnableItemRecycleBin }},

	-- AH fee structure, defaults are retail.
	AH_BASE_FEE_SINGLE = {{ .AuctionHouse.BaseFeeSingle }},
	AH_BASE_FEE_STACKS = {{ .AuctionHouse.BaseFeeStacks }},
	AH_TAX_RATE_SINGLE = {{ .AuctionHouse.TaxRateSingle }},
	AH_TAX_RATE_STACKS = {{ .AuctionHouse.TaxRateStacks }},
	AH_MAX_FEE         = {{ .AuctionHouse.MaxFee }},

	-- Max open listings per player, 0 = no limit. (Default 7)
	-- Note = Settings over 7 may need client-side plugin to work under all circumstances.
	-- If this is the case, consider using the ah_pagination module
	AH_LIST_LIMIT = {{ .AuctionHouse.ListLimit }},

	-- Misc EXP related settings
	EXP_RATE                = {{ .Exp.Rate }},
	EXP_LOSS_RATE           = {{ .Exp.LossRate }},
	EXP_PARTY_GAP_PENALTIES = {{ .Exp.PartyGapPenalties }},

	-- Capacity Point Settings
	CAPACITY_RATE = {{ .Exp.CapacityRate }},

	-- Determines Vana'diel time epoch (886/1/1 Firesday)
	-- current timestamp - vanadiel_time_epoch = vana'diel time
	-- 0 defaults to SE epoch 1009810800 (JP midnight 1/1/2002)
	-- safe range is 1 - current timestamp
	VANADIEL_TIME_EPOCH = {{ .VanadielTimeEpoch }},

	-- For old fame calculation use .25
	FAME_MULTIPLIER = {{ .FameMultiplier }},

	-- Percentage of experience normally lost to keep upon death. 0 means full loss, where 1 means no loss.
	EXP_RETAIN = {{ .Exp.Retain }},

	-- Minimum level at which experience points can be lost
	EXP_LOSS_LEVEL = {{ .Exp.LossLevel }},

	-- Minimum level at which regional influence is lost in conquest when a player dies
	-- Level 5 and below don't lose influence: http://wiki.ffo.jp/html/498.html
	MINIMUM_LEVEL_CONQUEST_INFUENCE_LOSS = {{ .MinimumLevelConquestInfluenceLoss }},

	-- Enable/disable Level Sync
	LEVEL_SYNC_ENABLE = {{ .LevelSyncEnable }},

	-- Disables ability to equip higher level gear when level cap/sync effect is on player.
	DISABLE_GEAR_SCALING = {{ .DisableGearScaling }},

	-- Weaponskill point base (before skillchain) for breaking latent - whole numbers only. retail is 1.
	WS_POINTS_BASE = {{ .WeaponSkill.PointsBase }},

	-- Weaponskill points per skillchain element - whole numbers only, retail is 1
	-- (tier 3 sc's have 4 elements, plus 1 for the ws itself, giving 5 points to the closer).
	WS_POINTS_SKILLCHAIN = {{ .WeaponSkill.PointsSkillchain }},

	-- Enable/disable jobs other than BST and RNG having widescan
	ALL_JOBS_WIDESCAN = {{ .AllJobsWidescan }},

	-- Modifier to apply to player speed. 0 is the retail accurate default. Negative numbers will reduce it.
	SPEED_MOD = {{ .Speed.PlayerMod }},

	-- Modifier to apply to mount speed. 0 is the retail accurate default. Negative numbers will reduce it.
	-- Note retail treats the mounted speed as double what it actually is.
	MOUNT_SPEED_MOD = {{ .Speed.MountMod }},

	-- Modifier to apply to agro'd monster speed. 0 is the retail accurate default. Negative numbers will reduce it.
	MOB_SPEED_MOD = {{ .Speed.MobMod }},

	-- Allows you to manipulate the constant multiplier in the skill-up rate formulas, having a potent effect on skill-up rates.
	SKILLUP_CHANCE_MULTIPLIER = {{ .Skillup.ChanceMultiplier }},
	CRAFT_CHANCE_MULTIPLIER   = {{ .Craft.ChanceMultiplier }},

	-- Multiplier for skillup amounts. Using anything above 1 will break the 0.5 cap, the cap will become 0.9 (For maximum, set to 5)
	SKILLUP_AMOUNT_MULTIPLIER = {{ .Skillup.AmountMultiplier }},
	CRAFT_AMOUNT_MULTIPLIER   = {{ .Craft.AmountMultiplier }},

	-- Gardening Factors. DO NOT change defaults without verifiable proof that your change IS how retail does it. Myths need to be optional.
	GARDEN_DAY_MATTERS       = {{ .Garden.DayMatters }},
	GARDEN_MOONPHASE_MATTERS = {{ .Garden.MoonphaseMatters }},
	GARDEN_POT_MATTERS       = {{ .Garden.PotMatters }},
	GARDEN_MH_AURA_MATTERS   = {{ .Garden.MogHouseAuraMatters }},

	-- Use current retail skill up rates and margins (Retail = High Skill-Up rate; Skill-Up when at or under 10 levels above synth recipe level.)
	CRAFT_MODERN_SYSTEM = {{ .Craft.ModernSystem }},

	-- Craft level limit from witch specialization points beginning to count. (Retail = 700; Level 75 era:600)
	CRAFT_COMMON_CAP = {{ .Craft.CommonCap }},

	-- Amount of points allowed in crafts over the level defined above. Points are shared across all crafting skills. (Retail = 400; All skills can go to max = 3200)
	CRAFT_SPECIALIZATION_POINTS = {{ .Craft.SpecializationPoints }},

	-- Enables fishing. 0 = Disabled. 1 = Enable. ENABLE AT YOUR OWN RISK.
	FISHING_ENABLE = {{ .Fishing.Enable }},

	-- Multiplier for fishing skill-up chance. Default = 1.0, very hard.
	FISHING_SKILL_MULTIPLIER = {{ .Fishing.SkillMultiplier }},

	-- Enable/disable skill-ups from bloodpacts
	SKILLUP_BLOODPACT = {{ .Skillup.BloodPact }},

	-- Adjust rate of TP gain for mobs, pets (includes charmed pets), fellows, trusts and players.
	-- Acts as a multiplier, so default is 1.
	MOB_TP_MULTIPLIER    = {{ .Mob.TPMultiplier }},
	PET_TP_MULTIPLIER    = {{ .Pet.TPMultiplier }},
	PLAYER_TP_MULTIPLIER = {{ .Player.TPMultiplier }},
	TRUST_TP_MULTIPLIER  = {{ .Trust.TPMultiplier }},
	FELLOW_TP_MULTIPLIER = {{ .Fellow.TPMultiplier }},

	-- Adjust max HP pool for NMs, regular mobs, players, and trusts/fellows. Acts as a multiplier, so default is 1.
	NM_HP_MULTIPLIER        = {{ .NotoriousMonster.HPMultiplier }},
	MOB_HP_MULTIPLIER       = {{ .Mob.HPMultiplier }},
	PLAYER_HP_MULTIPLIER    = {{ .Player.HPMultiplier }},
	ALTER_EGO_HP_MULTIPLIER = {{ .AlterEgo.HPMultiplier }},

	-- Adjust max MP pool for NMs, regular mobs, players, and trusts/fellows. Acts as a multiplier, so default is 1.
	NM_MP_MULTIPLIER        = {{ .NotoriousMonster.MPMultiplier }},
	MOB_MP_MULTIPLIER       = {{ .Mob.MPMultiplier }},
	PLAYER_MP_MULTIPLIER    = {{ .Player.MPMultiplier }},
	ALTER_EGO_MP_MULTIPLIER = {{ .AlterEgo.MPMultiplier }},

	-- Sets the fraction of MP a subjob provides to the main job. Retail is half and this acts as a divisor so default is 2
	SJ_MP_DIVISOR = {{ .Subjob.MPDivisor }},

	-- Modify ratio of subjob-to-mainjob
	-- 0            = no subjobs
	-- 1            = 1/2   (default, 75/37, 99/49)
	-- 2            = 2/3   (75/50, 99/66)
	-- 3            = equal (75/75, 99/99)
	SUBJOB_RATIO = {{ .Subjob.Ratio }},

	-- Also adjust monsters subjob in ratio adjustments? 1 = true / 0 = false
	INCLUDE_MOB_SJ = {{ .Mob.IncludeSubjob }},

	-- Adjust base stats (str/vit/etc.) for NMs, regular mobs, players, and trusts/fellows. Acts as a multiplier, so default is 1.
	NM_STAT_MULTIPLIER        = {{ .NotoriousMonster.StatMultiplier }},
	MOB_STAT_MULTIPLIER       = {{ .Mob.StatMultiplier }},
	PLAYER_STAT_MULTIPLIER    = {{ .Player.StatMultiplier }},
	ALTER_EGO_STAT_MULTIPLIER = {{ .AlterEgo.StatMultiplier }},

	-- Adjust skill caps for trusts/fellows. Acts as a multiplier, so default is 1.
	ALTER_EGO_SKILL_MULTIPLIER = {{ .AlterEgo.SkillMultiplier }},

	-- Adjust the recast time for abilities. Acts as a multiplier, so default is 1
	ABILITY_RECAST_MULTIPLIER = {{ .AbilityRecastMultiplier }},

	-- Enable/disable shared blood pact timer
	BLOOD_PACT_SHARED_TIMER = {{ .BloodPactSharedTimer }},

	-- Adjust mob drop rate. Acts as a multiplier, so default is 1.
	DROP_RATE_MULTIPLIER = {{ .Drops.RateMultiplier }},

	-- Multiplier for gil naturally dropped by mobs. Does not apply to the bonus gil from all_mobs_gil_bonus. Default is 1.0.
	MOB_GIL_MULTIPLIER = {{ .Drops.GilMultiplier }},

	-- All mobs drop this much extra gil per mob LV even if they normally drop zero.
	ALL_MOBS_GIL_BONUS = {{ .Drops.GilBonus }},

	-- Maximum total bonus gil that can be dropped. Default 9999 gil.
	MAX_GIL_BONUS = {{ .Drops.MaxGilBonus }},

	-- Allow mobs to walk back home instead of despawning
	MOB_NO_DESPAWN = {{ .Mob.NoDespawn }},

	-- Adds extra time to mob despawn in seconds. Base time is 25s, so a setting of 5 here would be a total of 30 seconds.
	MOB_ADDITIONAL_TIME_TO_DEAGGRO = {{ .Mob.AdditionalTimeToDeaggro.Seconds }},

	-- Allows parry, block, and guard to skill up regardless of the action occuring.
	-- This did not happen in previous eras
	PARRY_OLD_SKILLUP_STYLE = {{ .Skillup.OldParryStyle }},
	BLOCK_OLD_SKILLUP_STYLE = {{ .Skillup.OldBlockStyle }},
	GUARD_OLD_SKILLUP_STYLE = {{ .Skillup.OldGuardStyle }},

	-- Globally adjusts ALL battlefield level caps by this many levels.
	BATTLE_CAP_TWEAK = {{ .BattleCapTweak }},

	-- Enable/disable level cap of mission battlefields stored in database.
	LV_CAP_MISSION_BCNM = {{ .LevelCapMissionBCNM }},

	-- Max allowed merits points players can hold
	-- 10 classic
	-- 30 abyssea
	MAX_MERIT_POINTS = {{ .Exp.MaxMeritPoints }},

	-- Minimum time between uses of yell command (in seconds).
	YELL_COOLDOWN = {{ .YellCooldown.Seconds }},

	-- Prevent players from sending tells to hidden GMs. You will still receive them from other GMs.
	BLOCK_TELL_TO_HIDDEN_GM = {{ .BlockTellToHiddenGameMaster }},

	-- Command Audit [logging] commands with lower permission than this will not be logged.
	-- Zero for no logging at all. Commands given to non GMs are not logged.
	AUDIT_GM_CMD = {{ .Audit.GameMasterCommand }},

	-- Todo = other logging including anti-cheat messages

	-- Chat Audit[logging] settings
	AUDIT_CHAT      = {{ .Audit.Chat }},
	AUDIT_SAY       = {{ .Audit.Say }},
	AUDIT_SHOUT     = {{ .Audit.Shout }},
	AUDIT_TELL      = {{ .Audit.Tell }},
	AUDIT_YELL      = {{ .Audit.Yell }},
	AUDIT_LINKSHELL = {{ .Audit.Linkshell }},
	AUDIT_UNITY     = {{ .Audit.Unity }},
	AUDIT_PARTY     = {{ .Audit.Party }},

	-- Seconds between healing ticks. Default is 10
	HEALING_TICK_DELAY = {{ .HealingTickDelay.Seconds }},

	-- Set to 1 to enable server side anti-cheating measurements
	ANTICHEAT_ENABLED = {{ .AntiCheat.Enabled }},

	-- Set to 1 to completely disable auto-jailing offenders
	ANTICHEAT_JAIL_DISABLE = {{ .AntiCheat.JailDisable }},

	-- Enable/disable keeping jug pets through zoning
	KEEP_JUGPET_THROUGH_ZONING = {{ .KeepJugPetThroughZoning }},

	-- Send stack traces to the client after caught Lua errors if
	-- their GM level is the same or higher than this number.
	-- The max GM level is 5, so setting this to 6 disables it
	-- for everone. Setting it to 0 enables for everyone.
	REPORT_LUA_ERRORS_TO_PLAYER_LEVEL = {{ .ReportLuaErrorsToPlayerLevel }},
}
`
)

var (
	MapConfig *template.Template
)

func init() {
	tmpl, err := template.New("mapConfig").
		Funcs(template.FuncMap{}).
		Parse(mapConfig)
	if err != nil {
		log.Fatalln(err)
	}
	MapConfig = tmpl
}
