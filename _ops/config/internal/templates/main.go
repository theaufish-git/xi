package templates

import (
	"log"
	"text/template"
)

const (
	mainConfig string = `-----------------------------------
-- MAIN SETTINGS
-----------------------------------
-- All settings are attached to the 'xi.settings' object. This is published globally, and be accessed from C++ and any script.
--
-- This file is concerned mainly with content, balance, and gameplay tweaking.
-----------------------------------

xi = xi or {}
xi.settings = xi.settings or {}

xi.settings.main =
{
	-- Server name (not longer than 15 characters)
	SERVER_NAME = "{{ .Server.Name }}",

	SERVER_MESSAGE = "{{ .Server.Message }}",

	-- Setting to lock content more accurately to the expansions defined below.
	-- This generally results in a more accurate presentation of your selected expansions,
	-- as well as a less confusing player experience for things that are disabled (things that are disabled are not loaded).
	-- This feature correlates to the content_tag column in the SQL files.
	RESTRICT_CONTENT = {{ .RestrictContent }},

	-- Enable Expansion (1 = Enabled, 0 = Disabled)
	ENABLE_COP       = {{ .ChainsOfPromathia.Enable }},
	ENABLE_TOAU      = {{ .TreasuresOfArtUrghan.Enable }},
	ENABLE_WOTG      = {{ .WingsOfTheGoddess.Enable }},
	ENABLE_ACP       = {{ .ACrystalineProphecy.Enable }},
	ENABLE_AMK       = {{ .AMoogleKupodEtat.Enable }},
	ENABLE_ASA       = {{ .AShantottoAscension.Enable }},
	ENABLE_ABYSSEA   = {{ .Abyssea.Enable }},
	ENABLE_SOA       = {{ .SeekersOfAdoulin.Enable }},
	ENABLE_ROV       = {{ .RhapsodiesOfVanadiel.Enable }},
	ENABLE_VOIDWATCH = {{ .Voidwatch.Enable }}, -- Not an expansion, but has its own storyline. (Not Implemented)

	-- FIELDS OF VALOR/Grounds of Valor settings
	ENABLE_FIELD_MANUALS  = {{ .Valor.EnableFieldManuals }}, -- Enables Fields of Valor
	ENABLE_GROUNDS_TOMES  = {{ .Valor.EnableGroundsTomes }}, -- Enables Grounds of Valor
	ENABLE_SURVIVAL_GUIDE = {{ .Valor.EnableSurvivalGuide }}, -- Enables Survival Guides (Not Implemented)
	REGIME_WAIT           = {{ .Valor.RegimeWait }}, -- Make people wait till 00:00 game time as in retail. If it's 0, there is no wait time.
	FOV_REWARD_ALLIANCE   = {{ .Valor.FieldsRewardAlliance }}, -- Allow Fields of Valor rewards while being a member of an alliance. (default retail behavior: 0)
	GOV_REWARD_ALLIANCE   = {{ .Valor.GroundsRewardAlliance }}, -- Allow Grounds of Valor rewards while being a member of an alliance. (default retail behavior: {{ . }})

	-- Daily points / Gobbie mystery box.
	ENABLE_DAILY_TALLY = {{ .GobbieBox.Enable }},  -- Allows acquisition of daily points for gobbie mystery box.
	DAILY_TALLY_AMOUNT = {{ .GobbieBox.Amount }},
	DAILY_TALLY_LIMIT  = {{ .GobbieBox.Limit }},
	GOBBIE_BOX_MIN_AGE = {{ .GobbieBox.MinAge }}, -- Minimum character age in days before a character can sign up for Gobbie Mystery Box

	-- Records of Eminence
	ENABLE_ROE            = {{ .RecordsOfEminence.Enable }}, -- Enable Records of Eminence
	ENABLE_ROE_TIMED      = {{ .RecordsOfEminence.EnableTimed }}, -- Enable 4-hour timed records
	ENABLE_EXCHANGE_LIMIT = {{ .RecordsOfEminence.EnableExchangeLimit }}, -- Enable Maximum limit of sparks spent per Week (default retail behavior: 1)

	WEEKLY_EXCHANGE_LIMIT = {{ .RecordsOfEminence.WeeklyExchangeLimit }}, -- Maximum amount of sparks/accolades that can be spent per week (default retail value: 100000)

	-- Currency Caps (Change at your own risk!)
	CAP_CURRENCY_ACCOLADES = {{ .Currency.CapAccolades }},
	CAP_CURRENCY_BALLISTA  = {{ .Currency.CapBallista }},
	CAP_CURRENCY_SPARKS    = {{ .Currency.CapSparks }},
	CAP_CURRENCY_VALOR     = {{ .Currency.CapValor }},

	-- Magian Trials
	ENABLE_MAGIAN_TRIALS             = {{ .MagianTrials.Enable }},
	MAGIAN_TRIALS_MOBKILL_MULTIPLIER = {{ .MagianTrials.MobKillMultiplier }},
	MAGIAN_TRIALS_TRADE_MULTIPLIER   = {{ .MagianTrials.TradeMultiplier }},

	-- VoidWalker
	ENABLE_VOIDWALKER = {{ .Voidwalker.Enable }},

	-- TREASURE CASKETS
	-- Retail droprate = 0.1 (10%) with no other effects active
	-- Set to 0 to disable caskets.
	-- max is clamped to 1.0 (100%)
	CASKET_DROP_RATE = {{ .Abyssea.CasketDropRate }},

	-- Abyssea lights
	-- certain mobs that reduces the drop rate automatically depending on the light.
	-- pearl light is a dramaticly lower drop rate.
	-- min is 0 max is 100 (1 = 1%)
	ABYSSEA_LIGHTS_DROP_RATE = {{ .Abyssea.LightsDropRate }},

	-- This bonus will be added to players lights apon entering abyssea, it is mainly used during events
	-- recomended amount 0 - 100, some lights will cap at 255 while others are less, these are capped automatically
	ABYSSEA_BONUSLIGHT_AMOUNT = {{ .Abyssea.BonusLightAmount }},

	-- CHARACTER CONFIG
	INITIAL_LEVEL_CAP              = {{ .Character.InitialLevelCap }}, -- The initial level cap for new players.  There seems to be a hardcap of 255.
	MAX_LEVEL                      = {{ .Character.MaxLevel }}, -- Level max of the server, lowers the attainable cap by disabling Limit Break quests.
	NORMAL_MOB_MAX_LEVEL_RANGE_MIN = {{ .Character.NormalMobMaxLevelRangeMin }},  -- Lower Bound of Max Level Range for Normal Mobs (0 = Uncapped)
	NORMAL_MOB_MAX_LEVEL_RANGE_MAX = {{ .Character.NormalMobMaxLevelRangeMax }},  -- Upper Bound of Max Level Range for Normal Mobs (0 = Uncapped)
	START_GIL                      = {{ .Character.StartGil }}, -- Amount of gil given to newly created characters.
	START_INVENTORY                = {{ .Character.StartInventory }}, -- Starting inventory and satchel size.  Ignores values < 30.  Do not set above 80!
	NEW_CHARACTER_CUTSCENE         = {{ .Character.NewCharacterCutscene }},  -- Set to 1 to enable opening cutscenes, 0 to disable.
	SUBJOB_QUEST_LEVEL             = {{ .Character.SubjobQuestLevel }}, -- Minimum level to accept either subjob quest.  Set to 0 to start the game with subjobs unlocked.
	ADVANCED_JOB_LEVEL             = {{ .Character.AdvancedJobLevel }}, -- Minimum level to accept advanced job quests.  Set to 0 to start the game with advanced jobs.
	ALL_MAPS                       = {{ .Character.AllMaps }},  -- Set to 1 to give starting characters all the maps.
	UNLOCK_OUTPOST_WARPS           = {{ .Character.UnlockOutpostWarps }},  -- Set to 1 to give starting characters all outpost warps.  2 to add Tu'Lia and Tavnazia.

	SHOP_PRICE      = {{ .ShopPrice }}, -- Multiplies prices in NPC shops.
	GIL_RATE        = {{ .Currency.GilRate }}, -- Multiplies gil earned from quests.  Won't always display in game.
	BAYLD_RATE      = {{ .Currency.BayldRate }}, -- Multiples bayld earned from quests.
	-- Note: EXP rates are also influenced by conf setting
	EXP_RATE        = {{ .Exp.Rate }}, -- Multiplies exp from script (except FoV/GoV).
	CAPACITY_RATE   = {{ .Exp.CapacityRate }}, -- Multiplies capacy points gained.
	BOOK_EXP_RATE   = {{ .Exp.BookRate }}, -- Multiplies exp from FoV/GoV book pages.
	TABS_RATE       = {{ .Valor.TabsRate }}, -- Multiplies tabs earned from fov.
	ROE_EXP_RATE    = {{ .Exp.RecordsOfEminenceRate }}, -- Multiplies exp earned from records of eminence.
	SPARKS_RATE     = {{ .Currency.SparksRate }}, -- Multiplies sparks earned from records of eminence.
	CURE_POWER      = {{ .Spell.CurePower }}, -- Multiplies amount healed from Healing Magic, including the relevant Blue Magic.
	ELEMENTAL_POWER = {{ .Spell.ElementalPower }}, -- Multiplies damage dealt by Elemental and non-drain Dark Magic.
	DIVINE_POWER    = {{ .Spell.DivinePower }}, -- Multiplies damage dealt by Divine Magic.
	NINJUTSU_POWER  = {{ .Spell.NinjitsuPower }}, -- Multiplies damage dealt by Ninjutsu Magic.
	BLUE_POWER      = {{ .Spell.BluePower }}, -- Multiplies damage dealt by Blue Magic.
	DARK_POWER      = {{ .Spell.DarkPower }}, -- Multiplies amount drained by Dark Magic.
	ITEM_POWER      = {{ .Spell.ItemPower }}, -- Multiplies the effect of items such as Potions and Ethers.
	WEAPON_SKILL_POWER  = {{ .WeaponSkill.Power }}, -- Multiplies damage dealt by Weapon Skills.

	USE_ADOULIN_WEAPON_SKILL_CHANGES = {{ .SeekersOfAdoulin.UseWeaponSkillChanges }}, -- true/false. Change to toggle new Adoulin weapon skill damage calculations
	DISABLE_PARTY_EXP_PENALTY        = {{ .Exp.DisablePartyPenalty }}, -- true/false.

	-- TRUSTS
	ENABLE_TRUST_CASTING           = {{ .Trust.EnableCasting }},
	ENABLE_TRUST_QUESTS            = {{ .Trust.EnableQuests }},
	ENABLE_TRUST_CUSTOM_ENGAGEMENT = {{ .Trust.EnableCustomEngagement }},

	ENABLE_TRUST_ALTER_EGO_EXTRAVAGANZA          = {{ .Trust.EnableAlterEgoExtravaganza }}, -- 0 = disabled, 1 = summer/ny, 2 = spring/autumn, 3 = both
	ENABLE_TRUST_ALTER_EGO_EXTRAVAGANZA_ANNOUNCE = {{ .Trust.EnableAlterEgoExtravaganzaAnnounce }}, -- 0 = disabled, 1 = add announcement to player login
	ENABLE_TRUST_ALTER_EGO_EXPO                  = {{ .Trust.EnableAlterEgoExpo }}, -- 0 = disabled, 1 = expo - HPP/MPP/Status Resistance, 2 = expo plus (not implemented)
	ENABLE_TRUST_ALTER_EGO_EXPO_ANNOUNCE         = {{ .Trust.EnableAlterEgoExpoAnnounce }}, -- 0 = disabled, 1 = add announcement to player login

	TRUST_ALTER_EGO_EXTRAVAGANZA_MESSAGE = "{{ .Trust.AlterEgoExtravaganzaMessage }}",

	TRUST_ALTER_EGO_EXPO_MESSAGE = "{{ .Trust.AlterEgoExpoMessage }}",

	HARVESTING_BREAK_CHANCE = {{ .Gathering.HarvestingBreakChance }}, -- % chance for the sickle to break during harvesting.  Set between 0 and 100.
	EXCAVATION_BREAK_CHANCE = {{ .Gathering.ExcavationBreakChance }}, -- % chance for the pickaxe to break during excavation.  Set between 0 and 100.
	LOGGING_BREAK_CHANCE    = {{ .Gathering.LoggingBreakChance }}, -- % chance for the hatchet to break during logging.  Set between 0 and 100.
	MINING_BREAK_CHANCE     = {{ .Gathering.MiningBreakChance }}, -- % chance for the pickaxe to break during mining.  Set between 0 and 100.
	HARVESTING_RATE         = {{ .Gathering.HarvestingRate }}, -- % chance to recieve an item from haresting.  Set between 0 and 100.
	EXCAVATION_RATE         = {{ .Gathering.ExcavationRate }}, -- % chance to recieve an item from excavation.  Set between 0 and 100.
	LOGGING_RATE            = {{ .Gathering.LoggingRate }}, -- % chance to recieve an item from logging.  Set between 0 and 100.
	MINING_RATE             = {{ .Gathering.MiningRate }}, -- % chance to recieve an item from mining.  Set between 0 and 100.
	DIGGING_RATE            = {{ .Gathering.Digging.Rate }}, -- % chance to receive an item from chocbo digging during favorable weather.  Set between 0 and 100.

	HEALING_TP_CHANGE       = {{ .HealingTPChange }}, -- Change in TP for each healing tick. Default is -100

	-- SE implemented coffer/chest illusion time in order to prevent coffer farming. No-one in the same area can open a chest or coffer for loot (gil, gems & items)
	-- till a random time between MIN_ILLSION_TIME and MAX_ILLUSION_TIME. During this time players can loot keyitem and item related to quests (AF, maps... etc.)
	COFFER_MAX_ILLUSION_TIME = {{ .Treasure.CofferIllusionTime.Max.Seconds }},  -- 1 hour
	COFFER_MIN_ILLUSION_TIME = {{ .Treasure.CofferIllusionTime.Min.Seconds }},  -- 30 minutes
	CHEST_MAX_ILLUSION_TIME  = {{ .Treasure.ChestIllusionTime.Max.Seconds }},  -- 1 hour
	CHEST_MIN_ILLUSION_TIME  = {{ .Treasure.ChestIllusionTime.Min.Seconds }},  -- 30 minutes

	-- Multiplier to NM lottery spawn chance. (Default 1.0) eg. 0 = disable lottery spawns. -1 for always 100% chance.
	NM_LOTTERY_CHANCE = {{ .NotoriousMonster.LotteryChance }},
	-- Multiplier to NM lottery cooldown time (Default 1.0) eg. 2.0 = twice as long. 0 = no cooldowns.
	NM_LOTTERY_COOLDOWN = {{ .NotoriousMonster.LotteryCooldown }},

	-- DYNAMIS SETTINGS
	BETWEEN_2DYNA_WAIT_TIME     = {{ .Dynamis.ReentryWaitTime.Hours }},       -- Hours before player can re-enter Dynamis. Default is 1 Earthday (24 hours).
	DYNA_MIDNIGHT_RESET         = {{ .Dynamis.MidnightReset }},     -- If true, makes the wait time count by number of server midnights instead of full 24 hour intervals
	DYNA_LEVEL_MIN              = {{ .Dynamis.LevelMin }},       -- Level min for entering in Dynamis
	TIMELESS_HOURGLASS_COST     = {{ .Dynamis.TimelessHourglassCost }},   -- Refund for the timeless hourglass for Dynamis.
	PRISMATIC_HOURGLASS_COST    = {{ .Dynamis.PrismaticHourglassCost }},    -- Cost of the prismatic hourglass for Dynamis.
	CURRENCY_EXCHANGE_RATE      = {{ .Dynamis.CurrencyExchangeRate }},      -- X Tier 1 ancient currency -> 1 Tier 2, and so on. Certain values may conflict with shop items. Not designed to exceed 198.
	RELIC_2ND_UPGRADE_WAIT_TIME = {{ .Dynamis.Relic2ndUpgradeWaitTime.Seconds }},     -- Wait time for 2nd relic upgrade (stage 2 -> stage 3) in seconds. 7200s = 2 hours.
	RELIC_3RD_UPGRADE_WAIT_TIME = {{ .Dynamis.Relic3rdUpgradeWaitTime.Seconds }},     -- Wait time for 3rd relic upgrade (stage 3 -> stage 4) in seconds. 3600s = 1 hour.
	FREE_COP_DYNAMIS            = {{ .Dynamis.FreeChainsOfPromathia }},        -- Authorize player to entering inside COP Dynamis without completing COP mission (1 = enable 0 = disable)

	-- LIMBUS SETTINGS
	COSMO_CLEANSE_BASE_COST     = {{ .Limbus.CosmoCleanseBaseCost }},    -- Base gil cost for a Cosmo Cleanse from Sagheera

	-- QUEST/MISSION SPECIFIC SETTINGS
	AF1_QUEST_LEVEL = {{ .ArtifactQuests.QuestLevel1 }},    -- Minimum level to start AF1 quest
	AF2_QUEST_LEVEL = {{ .ArtifactQuests.QuestLevel2 }},    -- Minimum level to start AF2 quest
	AF3_QUEST_LEVEL = {{ .ArtifactQuests.QuestLevel3 }},    -- Minimum level to start AF3 quest
	OLDSCHOOL_G1    = {{ .LimitBreakQuests.OldSchoolLimitBreak1 }}, -- Set to true to require farming Exoray Mold, Bombd Coal, and Ancient Papyrus drops instead of allowing key item method.
	OLDSCHOOL_G2    = {{ .LimitBreakQuests.OldSchoolLimitBreak2 }}, -- Set true to require the NMs for "Atop the Highest Mountains" be dead to get KI like before SE changed it.
	FRIGICITE_TIME  = {{ .LimitBreakQuests.FrigiciteTime.Seconds }},    -- When OLDSCHOOL_G2 is enabled, this is the time (in seconds) you have from killing Boreal NMs to click the "???" target.
	ASSAULT_MINIMUM = {{ .TreasuresOfArtUrghan.AssaultMinimum }},     -- Minimum amount of people needed to start an assault mission. TOAU era is 3, Default is 1.

	-- SPELL SPECIFIC SETTINGS
	DIA_OVERWRITE                   = {{ .Spell.DiaOverwrite }},     -- Set to 1 to allow Bio to overwrite same tier Dia.  Default is 1.
	BIO_OVERWRITE                   = {{ .Spell.BioOverwrite }},     -- Set to 1 to allow Dia to overwrite same tier Bio.  Default is 0.
	STONESKIN_CAP                   = {{ .Spell.StoneskinCap }},   -- Soft cap for hp absorbed by stoneskin
	BLINK_SHADOWS                   = {{ .Spell.BlinkShadows }},     -- Number of shadows supplied by Blink spell
	SPIKE_EFFECT_DURATION           = {{ .Spell.SpikeEffectDuration.Seconds }},   -- the duration of RDM, BLM spikes effects (not Reprisal)
	ELEMENTAL_DEBUFF_DURATION       = {{ .Spell.ElementalDebuffDuration.Seconds }},   -- base duration of elemental debuffs
	AQUAVEIL_COUNTER                = {{ .Spell.AquaveilCounter }},     -- Base amount of hits Aquaveil absorbs to prevent spell interrupts. Retail is 1.
	ABSORB_SPELL_AMOUNT             = {{ .Spell.AbsorbSpellAmount }},     -- how much of a stat gets absorbed by DRK absorb spells - expected to be a multiple of 8.
	ABSORB_SPELL_TICK               = {{ .Spell.AbsorbSpellTick.Seconds }},     -- duration of 1 absorb spell tick
	SNEAK_INVIS_DURATION_MULTIPLIER = {{ .Spell.SneakInvisDurationMultiplier }},     -- multiplies duration of sneak, invis, deodorize to reduce player torture. 1 = retail behavior.
	USE_OLD_CURE_FORMULA            = {{ .Spell.UseOldCureFormula }}, -- true/false. if true, uses older cure formula (3*MND + VIT + 3*(healing skill/5)) // cure 6 will use the newer formula
	USE_OLD_MAGIC_DAMAGE            = {{ .Spell.UseOldMagicDamage }}, -- true/false. if true, uses older magic damage formulas

	-- CELEBRATIONS
	EXPLORER_MOOGLE_LV              = {{ .Celebration.ExplorerMoogleLevel }}, -- Enables Explorer Moogle teleports and sets required level. Zero to disable.
	HALLOWEEN_2005                  = {{ .Celebration.Halloween2005 }},  -- Set to 1 to Enable the 2005 version of Harvest Festival, will start on Oct. 20 and end Nov. 1.
	HALLOWEEN_YEAR_ROUND            = {{ .Celebration.HalloweenYearRound }},  -- Set to 1 to have Harvest Festival initialize outside of normal times.

	-- Login Campaign (Set to 0 if you don't want to run a Login Campaign)
	-- Please visit scripts/globals/events/login_campaign.lua for assigning the correct campaign dates.
	ENABLE_LOGIN_CAMPAIGN = {{ .LoginCampaign.Enable }},

	-- GARRISON
	GARRISON_LOCKOUT             = {{ .Garrison.Lockout.Seconds }},  -- Time in seconds before a new garrison can be started (default: 1800)
	GARRISON_TIME_LIMIT          = {{ .Garrison.TimeLimit.Seconds }},  -- Time in seconds before lose ongoing garrison (default: 1800)
	GARRISON_ONCE_PER_WEEK       = {{ .Garrison.OncePerWeek }},     -- Set to 1 to bypass the limit of one garrison per Conquest Tally Week.
	GARRISON_PARTY_LIMIT         = {{ .Garrison.PartyLimit }},    -- Set to max party members you want to do garrison (default: 18).
	GARRISON_NATION_BYPASS       = {{ .Garrison.NationBypass }},     -- Set to 1 to bypass the nation requirement.
	GARRISON_RANK                = {{ .Garrison.Rank }},     -- Set to minumum Nation Rank to start Garrison (default: 2).

	-- NYZUL
	RUNIC_DISK_SAVE      = {{ .Nyzul.RunicDiskSave }}, -- Allow anyone participating in Nyzul to save progress. Set to false so only initiator can save progress.
	ENABLE_NYZUL_CASKETS = {{ .Nyzul.EnableCaskets }}, -- Enable Treasure casket pops from NMs.
	ENABLE_VIGIL_DROPS   = {{ .Nyzul.EnableVigilDrops }}, -- Enable Vigil Weapon drops from NMs.
	ACTIVATE_LAMP_TIME   = {{ .Nyzul.ActivateLampTime.Milliseconds }}, -- Time in miliseconds for lamps to stay lit. TODO: Get retail confirmation.

	-- MISC
	RIVERNE_PORTERS              = {{ .ChainsOfPromathia.RivernePorters.Seconds }},   -- Time in seconds that Unstable Displacements in Cape Riverne stay open after trading a scale.
	LANTERNS_STAY_LIT            = {{ .LanternsStayLit.Seconds }},  -- time in seconds that lanterns in the Den of Rancor stay lit.
	ENABLE_COP_ZONE_CAP          = {{ .ChainsOfPromathia.EnableZoneCap }},     -- Enable or disable lvl cap
	ALLOW_MULTIPLE_EXP_RINGS     = {{ .Exp.RingsAllowMultiple }},     -- Set to 1 to remove ownership restrictions on the Chariot/Empress/Emperor Band trio.
	BYPASS_EXP_RING_ONE_PER_WEEK = {{ .Exp.RingsBypassOnePerWeek }},     -- Set to 1 to bypass the limit of one ring per Conquest Tally Week.
	NUMBER_OF_DM_EARRINGS        = {{ .NumberOfDivineMightEarrings }},     -- Number of earrings players can simultaneously own from Divine Might before scripts start blocking them (Default: 1)
	HOMEPOINT_TELEPORT           = {{ .HomepointTeleport }},     -- Enables the homepoint teleport system
	DIG_ABUNDANCE_BONUS          = {{ .Gathering.Digging.AbundanceBonus }},     -- Increase chance of digging up an item (450  = item digup chance +45)
	DIG_FATIGUE                  = {{ .Gathering.Digging.Fatigue }},     -- Set to 0 to disable Dig Fatigue
	DIG_GRANT_BURROW             = {{ .Gathering.Digging.GrantBurrow }},     -- Set to 1 to grant burrow ability
	DIG_GRANT_BORE               = {{ .Gathering.Digging.GrantBore }},     -- Set to 1 to grant bore ability
	ENM_COOLDOWN                 = {{ .EmptyNotoriousMonster.Cooldown.Hours }},   -- Number of hours before a player can obtain same KI for ENMs (default: 5 days)
	FORCE_SPAWN_QM_RESET_TIME    = {{ .EmptyNotoriousMonster.ForceSpawnQMResetTime.Seconds }},   -- Number of seconds the ??? remains hidden for after the despawning of the mob it force spawns.
	EQUIP_FROM_OTHER_CONTAINERS  = {{ .EquipFromOtherContainers }}, -- true/false. Allows equipping items from Mog Satchel, Sack, and Case. Only possible with the use of client addons.
}
`
)

var (
	MainConfig *template.Template
)

func init() {
	tmpl, err := template.New("mainConfig").
		Funcs(template.FuncMap{}).
		Parse(mainConfig)
	if err != nil {
		log.Fatalln(err)
	}
	MainConfig = tmpl
}
