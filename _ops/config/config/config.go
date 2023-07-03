package config

import (
	"time"
)

/**
 * Files
 **/

// Logging is a configuration struct used when generating logging.lua
type Logging struct {
	/*
	   # General pattern to be used by spdlog

	   ## Additional custom flags:
	   %& : the exe type name (login, map, search, world)
	   %_ : the first letter of the exe type name (L, M, S, W)
	   %* : a fixed-length (32) and right-truncated function name + file line number
	   %q : a fixed-length (32) and right-truncated file name + file line number

	   "Classic" pattern : "[%D %T:%e][%&]%^[%n]%$ %v (%!:%#)"
	   [12/03/22 19:25:12:812][map][info] do_init: connecting to database (do_init:221)
	   [12/03/22 19:25:12:815][map][info] database client version: 3.2.8 (do_init:224)
	   [12/03/22 19:25:12:815][map][info] database server version: 10.6.10-MariaDB (do_init:225)
	   [12/03/22 19:25:12:815][map][info] luautils: Lua initializing (luautils::init:113)

	   "New" pattern : "%Y/%m/%d %T:%e | %_ | %^%-4!l%$ | %* | %v"
	   2022/12/03 19:24:08:984 | M | info |                      do_init:221 | do_init: connecting to database
	   2022/12/03 19:24:08:987 | M | info |                      do_init:224 | database client version: 3.2.8
	   2022/12/03 19:24:08:987 | M | info |                      do_init:225 | database server version: 10.6.10-MariaDB
	   2022/12/03 19:24:08:987 | M | info |               luautils::init:113 | luautils: Lua initializing

	   ## Docs:
	   https://spdlog.docsforge.com/v1.x/3.custom-formatting/
	*/
	Pattern string `split_words:"true" default:"[%D %T:%e][%&]%^[%n]%$ %v (%!:%#)"`

	// Enable/Disable these logging types globally
	Log Log

	// Specific Debug loggers
	Debug Debug
}

// Login is a configuration struct used when generating login.lua
type Login struct {
	Client Client

	// 0 - disabled (normal operation)
	// 1 - enabled (only GM characters allowed online, no new character creation)
	MaintenanceMode int `split_word:"true" default:"0"`

	// Logging of user IP address to database (true/false)
	LogUserIp bool `split_words:"true" default:"false"`

	// Allow account creation via the loader (true/false)
	AccountCreation bool `split_words:"true" default:"true"`

	// Allow character deletion through the lobby (true/false)
	CharacterDeletion bool `split_words:"true" default:"true"`

	// Number of simultaneous game sessions per IP (0 for no limit)
	LoginLimit int `split_words:"true" default:"0"`

	// If true, blocks character creation with names of NPCs and Mobs in the database (Fafnir, Shantotto, etc.)
	DisableMobNpcCharNames bool `split_words:"true" default:"false"`

	// Character names with any of these words in, in any position, will be rejected
	//
	// Examples that will be rejected (using "badword"):
	// "badword"
	// "imbadword"
	// "badwordisme"
	// "lolbadwordlol"
	//
	// WARNING:
	// Be aware of the "Scunthorpe problem"!
	//
	// NOTE:
	// You can Google for "bad word list txt" to find lists of words to populate this table, if you'd like
	BannedWordsList []string `split_words:"true" default:""`
}

// Main is a configuration struct used when generating main.lua
type Main struct {
	Server Server

	Character        Character        // CHARACTER CONFIG
	Trust            Trust            // Trust Config
	NotoriousMonster NotoriousMonster `envconfig:"nm"` // NM Config

	Spell       Spell
	WeaponSkill WeaponSkill

	Exp       Exp       // Experience points
	Currency  Currency  // Currency Caps (Change at your own risk!)
	GobbieBox GobbieBox // Daily points / Gobbie mystery box.
	Gathering Gathering
	Treasure  Treasure

	// Quest/Missions/Tasks
	MagianTrials      MagianTrials      // Magian Trials
	ArtifactQuests    ArtifactQuests    `envconfig:"af"`     // QUEST/MISSION SPECIFIC SETTINGS
	LimitBreakQuests  LimitBreakQuests  `envconfig:"genkai"` // QUEST/MISSION SPECIFIC SETTINGS
	RecordsOfEminence RecordsOfEminence `envconfig:"roe"`    // Records of Eminence

	// Expansions
	// Setting to lock content more accurately to the expansions defined below.
	// This generally results in a more accurate presentation of your selected expansions,
	// as well as a less confusing player experience for things that are disabled (things that are disabled are not loaded).
	// This feature correlates to the content_tag column in the Sql files.
	RestrictContent int `split_words:"true" default:"0"`

	ChainsOfPromathia    ChainsOfPromathia    `envconfig:"cop"`
	TreasuresOfArtUrghan TreasuresOfArtUrghan `envconfig:"toau"`
	WingsOfTheGoddess    WingsOfTheGoddess    `envconfig:"wotg"`
	ACrystalineProphecy  ACrystalineProphecy  `envconfig:"acp"`
	AMoogleKupodEtat     AMoogleKupodEtat     `envconfig:"amk"`
	AShantottoAscension  AShantottoAscension  `envconfig:"asa"`
	Abyssea              Abyssea
	SeekersOfAdoulin     SeekersOfAdoulin     `envconfig:"soa"`
	RhapsodiesOfVanadiel RhapsodiesOfVanadiel `envconfig:"rov"`
	Voidwatch            Voidwatch            // Not an expansion, but has its own storyline. (Not Implemented)

	// Combat Systems
	Garrison              Garrison              // Garrison
	Valor                 Valor                 // Fields of Valor/Grounds of Valor settings
	Dynamis               Dynamis               // Dynamis
	Limbus                Limbus                // Limbus
	Nyzul                 Nyzul                 // Nyzul
	EmptyNotoriousMonster EmptyNotoriousMonster `envconfig:"enm"` // ENM configuration
	Voidwalker            Voidwalker            // VoidWalker

	// Events
	Celebration   Celebration   // CELEBRATIONS
	LoginCampaign LoginCampaign // Login Campaign

	// MISC
	ShopPrice                   float64       `split_words:"true" default:"1.000"`            // Multiplies prices in NPC shops.
	HealingTPChange             int           `split_words:"true" default:"-100"`             // Change in TP for each healing tick. Default is -100
	LanternsStayLit             time.Duration `split_words:"true" default:"20m"`              // time in seconds that lanterns in the Den of Rancor stay lit.
	NumberOfDivineMightEarrings int           `envconfig:"number_of_dm_earrings" default:"1"` // Number of earrings players can simultaneously own from Divine Might before scripts start blocking them (Default: 1)
	HomepointTeleport           int           `split_words:"true" default:"1"`                // Enables the homepoint teleport system
	EquipFromOtherContainers    bool          `split_words:"true" default:"false"`            // true/false. Allows equipping items from Mog Satchel, Sack, and Case. Only possible with the use of client addons.
}

// Map is a configuration struct used when generating map.lua
type Map struct {
	AntiCheat    AntiCheat
	AuctionHouse AuctionHouse
	Audit        Audit
	Craft        Craft
	Drops        Drops
	Exp          Exp
	Fishing      Fishing
	Garden       Garden
	Skillup      Skillup
	Speed        Speed
	Subjob       Subjob
	WeaponSkill  WeaponSkill

	// Acts as a multiplier, so default is 1.
	// Adjust base stats (str/vit/etc.) for NMs, regular mobs, players, and trusts/fellows. Acts as a multiplier, so default is 1.
	// Adjust max HP pool for NMs, regular mobs, players, and trusts/fellows. Acts as a multiplier, so default is 1.
	// Adjust max MP pool for NMs, regular mobs, players, and trusts/fellows. Acts as a multiplier, so default is 1.
	NotoriousMonster Mob `envconfig:"nm"`
	Mob              Mob
	Player           Entity
	AlterEgo         AlterEgo
	Pet              Entity
	Trust            Trust
	Fellow           Entity

	// --------------------------------
	// Packet settings
	// --------------------------------
	MaxTimeLastUpdate time.Duration `envconfig:"max_time_lastupdate" default:"60s"`

	// --------------------------------
	// SQL settings
	// --------------------------------

	// Used by serverutils::PersistServerVar() for the maximum attempts to retry verification
	// of a written Server Variable.
	SetvarRetryMax int `split_words:"true" default:"3"`

	// --------------------------------
	// Game settings
	// --------------------------------

	// Adjust the recast time for abilities. Acts as a multiplier, so default is 1
	AbilityRecastMultiplier float64 `split_words:"true" default:"1.0"`

	// Enable or disable Recycle Bin (Set to false for items to be dropped immediately)
	EnableItemRecycleBin bool `split_words:"true" default:"true"`

	// Determines Vana'diel time epoch (886/1/1 Firesday)
	// current timestamp - vanadiel_time_epoch = vana'diel time
	// 0 defaults to SE epoch 1009810800 (JP midnight 1/1/2002)
	// safe range is 1 - current timestamp
	VanadielTimeEpoch int `split_words:"true" default:"0"`

	// For old fame calculation use .25
	FameMultiplier float64 `split_words:"true" default:"1.0"`

	// Minimum level at which regional influence is lost in conquest when a player dies
	// Level 5 and below don't lose influence: http://wiki.ffo.jp/html/498.html
	MinimumLevelConquestInfluenceLoss int `split_words:"true" default:"6"`

	// Enable/disable Level Sync
	LevelSyncEnable bool `split_words:"true" default:"true"`

	// Disables ability to equip higher level gear when level cap/sync effect is on player.
	DisableGearScaling bool `split_words:"true" default:"false"`

	// Enable/disable jobs other than BST and RNG having widescan
	AllJobsWidescan bool `split_words:"true" default:"true"`

	// Enable/disable shared blood pact timer
	BloodPactSharedTimer bool `split_words:"true" default:"false"`

	// Globally adjusts ALL battlefield level caps by this many levels.
	BattleCapTweak int `split_words:"true" default:"0"`

	// Enable/disable level cap of mission battlefields stored in database.
	LevelCapMissionBCNM int `envconfig:"level_cap_mission_bcnm" default:"0"`

	// Minimum time between uses of yell command (in seconds).
	YellCooldown time.Duration `split_words:"true" default:"30s"`

	// Prevent players from sending tells to hidden GMs. You will still receive them from other GMs.
	BlockTellToHiddenGameMaster bool `envconfig:"block_tell_to_hidden_gm" default:"false"`

	// Todo = other logging including anti-cheat messages

	// Seconds between healing ticks. Default is 10
	HealingTickDelay time.Duration `split_words:"true" default:"10s"`

	// Enable/disable keeping jug pets through zoning
	KeepJugPetThroughZoning bool `envconfig:"keep_jugpet_through_zoning" default:"false"`

	// Send stack traces to the client after caught Lua errors if
	// their GM level is the same or higher than this number.
	// The max GM level is 5, so setting this to 6 disables it
	// for everone. Setting it to 0 enables for everyone.
	ReportLuaErrorsToPlayerLevel int `split_words:"true" default:"6"`
}

// Network is a configuration struct used when generating network.lua
type Network struct {
	Sql Sql

	Login Connect

	Map NetworkMap

	Search NetworkSearch

	Http Http

	Zmq Zmq

	// ===========================
	// NOTE: The settings that follow will not necessarily need to be modified
	//       in any way for the server to work out of the box.  This should only
	//       be modified by those who understand networking.  Modifying these
	//       values could potentially make it so that you can not log in to your
	//       server.
	// ===========================

	Udp Udp

	Tcp Tcp
}

// Search is a configuration struct used when generating search.lua
type Search struct {
	// Omit items with no buy history from auction house results
	OmitNoHistory bool `split_words:"true" default:"false"`

	// After EXPIRE_DAYS, will listed auctions expire?
	ExpireAuctions bool `split_words:"true" default:"true"`

	// Expire items older than this number of days
	ExpireDays int `split_words:"true" default:"3"`

	// Interval is in seconds, default is one hour
	ExpireInterval time.Duration `split_words:"true" default:"1h"`
}

/**
 * Expansions
 **/
type ChainsOfPromathia struct {
	Enable         int           `split_words:"true" default:"1"`
	RivernePorters time.Duration `split_words:"true" default:"2m"` // Time in seconds that Unstable Displacements in Cape Riverne stay open after trading a scale.
	EnableZoneCap  int           `split_words:"true" default:"0"`  // Enable or disable lvle ecap
}

type TreasuresOfArtUrghan struct {
	Enable         int `split_words:"true" default:"1"`
	AssaultMinimum int `split_words:"true" default:"1"` // Minimum amount of people needed to start an assault mission. TOAU era is 3, Default is 1.
}

type WingsOfTheGoddess struct {
	Enable int `split_words:"true" default:"1"`
}

type ACrystalineProphecy struct {
	Enable int `split_words:"true" default:"1"`
}

type AMoogleKupodEtat struct {
	Enable int `split_words:"true" default:"1"`
}

type AShantottoAscension struct {
	Enable int `split_words:"true" default:"1"`
}

type Abyssea struct {
	Enable int `split_words:"true" default:"1"`

	// TREASURE CASKETS
	// Retail droprate = 0.1 (10%) with no other effects active
	// Set to 0 to disable caskets.
	// max is clamped to 1.0 (100%)
	CasketDropRate float64 `split_words:"true" default:"0.1"`

	// Abyssea lights
	// certain mobs that reduces the drop rate automatically depending on the light.
	// pearl light is a dramaticly lower drop rate.
	// min is 0 max is 100 (1 = 1%)
	LightsDropRate int `split_words:"true" default:"80"`

	// This bonus will be added to players lights apon entering abyssea, it is mainly used during events
	// recomended amount 0 - 100, some lights will cap at 255 while others are less, these are capped automatically
	BonusLightAmount int `split_words:"true" default:"0"`
}

type SeekersOfAdoulin struct {
	Enable                int  `split_words:"true" default:"1"`
	UseWeaponSkillChanges bool `split_words:"true" default:"true"` // true/false. Change to toggle new Adoulin weapon skill damage calculations
}

/**
 * Components
 **/
type AlterEgo struct {
	Entity

	// Adjust skill caps for trusts/fellows. Acts as a multiplier, so default is 1.
	SkillMultiplier string `split_words:"true" default:"1.0"`
}

type AntiCheat struct {
	// PacketGuard will block and report any packets that aren't in the allow-list for a
	// player's current state.
	PacketGuardEnabled bool `envconfig:"packetguard_enabled" default:"true"`

	// Minimal number of 0x3A packets which uses for detect lightluggage (set 0 for disable)
	LightLuggageBlock int `envconfig:"lightluggage_block" default:"4"`

	// Set to 1 to enable server side anti-cheating measurements
	Enabled bool `split_words:"true" default:"true"`

	// Set to 1 to completely disable auto-jailing offenders
	JailDisable bool `split_words:"true" default:"false"`
}

type ArtifactQuests struct {
	QuestLevel1 int `envconfig:"quest_level_1" default:"40"` // Minimum level to start AF1 quest
	QuestLevel2 int `envconfig:"quest_level_2" default:"50"` // Minimum level to start AF2 quest
	QuestLevel3 int `envconfig:"quest_level_3" default:"50"` // Minimum level to start AF3 quest
}

type AuctionHouse struct {
	// AH fee structure, defaults are retail.
	BaseFeeSingle int     `split_words:"true" default:"1"`
	BaseFeeStacks int     `split_words:"true" default:"4"`
	TaxRateSingle float64 `split_words:"true" default:"1.0"`
	TaxRateStacks float64 `split_words:"true" default:"0.5"`
	MaxFee        int     `split_words:"true" default:"10000"`

	// Max open listings per player, 0 = no limit. (Default 7)
	// Note = Settings over 7 may need client-side plugin to work under all circumstances.
	// If this is the case, consider using the ah_pagination module
	ListLimit int `split_words:"true" default:"7"`
}

type Audit struct {
	// Command Audit [logging] commands with lower permission than this will not be logged.
	// Zero for no logging at all. Commands given to non GMs are not logged.
	GameMasterCommand bool `envconfig:"gm_cmd" default:"false"`

	// Chat Audit[logging] settings
	Chat      bool `split_words:"true" default:"false"`
	Say       bool `split_words:"true" default:"false"`
	Shout     bool `split_words:"true" default:"false"`
	Tell      bool `split_words:"true" default:"false"`
	Yell      bool `split_words:"true" default:"false"`
	Linkshell bool `split_words:"true" default:"false"`
	Unity     bool `split_words:"true" default:"false"`
	Party     bool `split_words:"true" default:"false"`
}

type Celebration struct {
	ExplorerMoogleLevel int `split_words:"true" default:"10"` // Enables Explorer Moogle teleports and sets required level. Zero to disable.
	Halloween2005       int `split_words:"true" default:"0"`  // Set to 1 to Enable the 2005 version of Harvest Festival, will start on Oct. 20 and end Nov. 1.
	HalloweenYearRound  int `split_words:"true" default:"0"`  // Set to 1 to have Harvest Festival initialize outside of normal times.
}

type Character struct {
	InitialLevelCap           int `split_words:"true" default:"50"` // The initial level cap for new players.  There seems to be a hardcap of 255.
	MaxLevel                  int `split_words:"true" default:"99"` // Level max of the server, lowers the attainable cap by disabling Limit Break quests.
	NormalMobMaxLevelRangeMin int `split_words:"true" default:"0"`  // Lower Bound of Max Level Range for Normal Mobs (0 = Uncapped)
	NormalMobMaxLevelRangeMax int `split_words:"true" default:"0"`  // Upper Bound of Max Level Range for Normal Mobs (0 = Uncapped)
	StartGil                  int `split_words:"true" default:"10"` // Amount of gil given to newly created characters.
	StartInventory            int `split_words:"true" default:"30"` // Starting inventory and satchel size.  Ignores values < 30.  Do not set above 80!
	NewCharacterCutscene      int `split_words:"true" default:"1"`  // Set to 1 to enable opening cutscenes, 0 to disable.
	SubjobQuestLevel          int `split_words:"true" default:"18"` // Minimum level to accept either subjob quest.  Set to 0 to start the game with subjobs unlocked.
	AdvancedJobLevel          int `split_words:"true" default:"30"` // Minimum level to accept advanced job quests.  Set to 0 to start the game with advanced jobs.
	AllMaps                   int `split_words:"true" default:"0"`  // Set to 1 to give starting characters all the maps.
	UnlockOutpostWarps        int `split_words:"true" default:"0"`  // Set to 1 to give starting characters all outpost warps.  2 to add Tu'Lia and Tavnazia.
}

type Client struct {
	// Expected Client version (wrong version cannot log in)
	Version string `split_words:"true" default:"30230415_1"`

	// 0 - disabled (every version allowed)
	// 1 - enabled - strict (only exact CLIENT_VER allowed)
	// 2 - enabled - greater than or equal  (matching or greater than CLIENT_VER allowed, default)
	//
	// WE STRONGLY ADVISE AGAINST LOCKING THE SERVER TO OLDER VERSIONS. IT IS A UNIVERSALLY BAD IDEA.
	VersionLock int `split_words:"true" default:"2"`
}

type Connect struct {
	DataIp   string `split_words:"true" default:"0.0.0.0"`
	DataPort int    `split_words:"true" default:"54230"`
	ViewIp   string `split_words:"true" default:"0.0.0.0"`
	ViewPort int    `split_words:"true" default:"54001"`
	AuthIp   string `split_words:"true" default:"0.0.0.0"`
	AuthPort int    `split_words:"true" default:"54231"`
	ConfIp   string `split_words:"true" default:"0.0.0.0"`
	ConfPort int    `split_words:"true" default:"51220"`
}

type Craft struct {
	// Use current retail skill up rates and margins (Retail = High Skill-Up rate; Skill-Up when at or under 10 levels above synth recipe level.)
	ModernSystem bool `split_words:"true" default:"true"`

	// Craft level limit from witch specialization points beginning to count. (Retail = 700; Level 75 era:600)
	CommonCap int `split_words:"true" default:"700"`

	// Amount of points allowed in crafts over the level defined above. Points are shared across all crafting skills. (Retail = 400; All skills can go to max = 3200)
	SpecializationPoints int `split_words:"true" default:"400"`

	// Allows you to manipulate the constant multiplier in the skill-up rate formulas, having a potent effect on skill-up rates.
	ChanceMultiplier float64 `split_words:"true" default:"1.0"`

	// Multiplier for skillup amounts. Using anything above 1 will break the 0.5 cap, the cap will become 0.9 (For maximum, set to 5)
	AmountMultiplier int `split_words:"true" default:"1"`
}

type Currency struct {
	CapAccolades int `split_words:"true" default:"99999"`
	CapBallista  int `split_words:"true" default:"2000"`
	CapSparks    int `split_words:"true" default:"99999"`
	CapValor     int `split_words:"true" default:"50000"`

	GilRate    float64 `split_words:"true" default:"1.000"` // Multiplies gil earned from quests.  Won't always display in game.
	BayldRate  float64 `split_words:"true" default:"1.000"` // Multiples bayld earned from quests.
	SparksRate float64 `split_words:"true" default:"1.000"` // Multiplies sparks earned from records of eminence.
}

type Debug struct {
	// NOTE: None of these will print unless you also have the above LOG_DEBUG setting set to true!
	Sockets       bool `split_words:"true" default:"false"` // Calls in C++: DebugSockets(...)
	Navmesh       bool `split_words:"true" default:"false"` // Calls in C++: DebugNavmesh(...)
	Packets       bool `split_words:"true" default:"false"` // Calls in C++: DebugPackets(...)
	Actions       bool `split_words:"true" default:"false"` // Calls in C++: DebugActions(...)
	Sql           bool `split_words:"true" default:"false"` // Calls in C++: DebugSql(...)
	IdLookup      bool `split_words:"true" default:"false"` // Calls in C++: DebugIDLookup(...)
	Modules       bool `split_words:"true" default:"false"` // Calls in C++: DebugModules(...)
	PacketBacklog bool `split_words:"true" default:"false"` // Special logic in map.cpp::send_parse
	DeliveryBox   bool `split_words:"true" default:"false"` // Special logic in packet_system.cpp::SmallPacket0x04D
}

type Digging struct {
	Rate           int `split_words:"true" default:"85"` // % chance to receive an item from chocbo digging during favorable weather.  Set between 0 and 100.
	AbundanceBonus int `split_words:"true" default:"0"`  // Increase chance of digging up an item (450  = item digup chance +45)
	Fatigue        int `split_words:"true" default:"1"`  // Set to 0 to disable Dig Fatigue
	GrantBurrow    int `split_words:"true" default:"0"`  // Set to 1 to grant burrow ability
	GrantBore      int `split_words:"true" default:"0"`  // Set to 1 to grant bore ability
}

type Drops struct {
	// Adjust mob drop rate. Acts as a multiplier, so default is 1.
	RateMultiplier float64 `split_words:"true" default:"1.0"`

	// Multiplier for gil naturally dropped by mobs. Does not apply to the bonus gil from all_mobs_gil_bonus. Default is 1.0.
	GilMultiplier float64 `split_words:"true" default:"1.0"`

	// All mobs drop this much extra gil per mob LV even if they normally drop zero.
	GilBonus int `split_words:"true" default:"0"`

	// Maximum total bonus gil that can be dropped. Default 9999 gil.
	MaxGilBonus int `split_words:"true" default:"9999"`
}

type Dynamis struct {
	ReentryWaitTime         time.Duration `split_words:"true" default:"24h"`                     // Hours before player can re-enter Dynamis. Default is 1 Earthday (24 hours).
	MidnightReset           bool          `split_words:"true" default:"true"`                    // If true, makes the wait time count by number of server midnights instead of full 24 hour intervals
	LevelMin                int           `split_words:"true" default:"65"`                      // Level min for entering in Dynamis
	TimelessHourglassCost   int           `split_words:"true" default:"500000"`                  // Refund for the timeless hourglass for Dynamis.
	PrismaticHourglassCost  int           `split_words:"true" default:"50000"`                   // Cost of the prismatic hourglass for Dynamis.
	CurrencyExchangeRate    int           `split_words:"true" default:"100"`                     // X Tier 1 ancient currency -> 1 Tier 2, and so on. Certain values may conflict with shop items. Not designed to exceed 198.
	Relic2ndUpgradeWaitTime time.Duration `envconfig:"relic_2nd_upgrade_wait_time" default:"2h"` // Wait time for 2nd relic upgrade (stage 2 -> stage 3) in seconds. 7200s = 2 hours.
	Relic3rdUpgradeWaitTime time.Duration `envconfig:"relic_3rd_upgrade_wait_time" default:"1h"` // Wait time for 3rd relic upgrade (stage 3 -> stage 4) in seconds. 3600s = 1 hour.
	FreeChainsOfPromathia   int           `envconfig:"free_cop" default:"0"`                     // Authorize player to entering inside COP Dynamis without completing COP mission (1 = enable 0 = disable)
}

type EmptyNotoriousMonster struct {
	Cooldown              time.Duration `split_words:"true" default:"120h"` // Number of hours before a player can obtain same KI for ENMs (default: 5 days)
	ForceSpawnQMResetTime time.Duration `split_words:"true" default:"5m"`   // Number of seconds the ??? remains hidden for after the despawning of the mob it force spawns.
}

type Entity struct {
	// Not all multipliers are used by all uses
	HPMultiplier   float64 `envconfig:"hp_multiplier" default:"1.0"`
	MPMultiplier   float64 `envconfig:"mp_multiplier" default:"1.0"`
	TPMultiplier   float64 `envconfig:"tp_multiplier" default:"1.0"`
	StatMultiplier float64 `split_words:"true" default:"1.0"`
}

type Exp struct {
	// Misc EXP related settings
	Rate                  float64 `split_words:"true" default:"1.000"`
	CapacityRate          float64 `split_words:"true" default:"1.000"`   // Multiplies capacy points gained.
	BookRate              float64 `split_words:"true" default:"1.000"`   // Multiplies exp from FoV/GoV book pages.
	RecordsOfEminenceRate float64 `envconfig:"roe_rate" default:"1.000"` // Multiplies exp earned from records of eminence.
	ScriptRate            float64 `split_words:"true" default:"1.000"`   // Multiplies exp from script (except FoV/GoV).

	LossRate            string `split_words:"true" default:"1.0"`
	PartyGapPenalties   bool   `split_words:"true" default:"true"`
	DisablePartyPenalty bool   `split_words:"true" default:"false"`

	// Percentage of experience normally lost to keep upon death. 0 means full loss, where 1 means no loss.
	Retain int `split_words:"true" default:"0"`

	// Minimum level at which experience points can be lost
	LossLevel int `split_words:"true" default:"31"`

	// Max allowed merits points players can hold
	// 10 classic
	// 30 abyssea
	MaxMeritPoints string `split_words:"true" default:"30"`

	RingsAllowMultiple    int `split_words:"true" default:"0"` // Set to 1 to remove ownership restrictions on the Chariot/Empress/Emperor Band trio.
	RingsBypassOnePerWeek int `split_words:"true" default:"0"` // Set to 1 to bypass the limit of one ring per Conquest Tally Week.
}

type Fishing struct {
	// Enables fishing. 0 = Disabled. 1 = Enable. ENABLE AT YOUR OWN RISK.
	Enable bool `split_words:"true" default:"false"`

	// Multiplier for fishing skill-up chance. Default = 1.0, very hard.
	SkillMultiplier float64 `split_words:"true" default:"1.0"`
}

type Garden struct {
	// Gardening Factors. DO NOT change defaults without verifiable proof that your change IS how retail does it. Myths need to be optional.
	DayMatters          bool `split_words:"true" default:"false"`
	MoonphaseMatters    bool `split_words:"true" default:"false"`
	PotMatters          bool `split_words:"true" default:"false"`
	MogHouseAuraMatters bool `envconfig:"mh_aura_matters" default:"false"`
}

type Garrison struct {
	Lockout      time.Duration `split_words:"true" default:"30m"` // Time in seconds before a new garrison can be started (default: 1800)
	TimeLimit    time.Duration `split_words:"true" default:"30m"` // Time in seconds before lose ongoing garrison (default: 1800)
	OncePerWeek  int           `split_words:"true" default:"0"`   // Set to 1 to bypass the limit of one garrison per Conquest Tally Week.
	PartyLimit   int           `split_words:"true" default:"18"`  // Set to max party members you want to do garrison (default: 18).
	NationBypass int           `split_words:"true" default:"0"`   // Set to 1 to bypass the nation requirement.
	Rank         int           `split_words:"true" default:"2"`   // Set to minumum Nation Rank to start Garrison (default: 2).
}

type Gathering struct {
	HarvestingBreakChance int `split_words:"true" default:"33"` // % chance for the sickle to break during harvesting.  Set between 0 and 100.
	ExcavationBreakChance int `split_words:"true" default:"33"` // % chance for the pickaxe to break during excavation.  Set between 0 and 100.
	LoggingBreakChance    int `split_words:"true" default:"33"` // % chance for the hatchet to break during logging.  Set between 0 and 100.
	MiningBreakChance     int `split_words:"true" default:"33"` // % chance for the pickaxe to break during mining.  Set between 0 and 100.
	HarvestingRate        int `split_words:"true" default:"50"` // % chance to recieve an item from haresting.  Set between 0 and 100.
	ExcavationRate        int `split_words:"true" default:"50"` // % chance to recieve an item from excavation.  Set between 0 and 100.
	LoggingRate           int `split_words:"true" default:"50"` // % chance to recieve an item from logging.  Set between 0 and 100.
	MiningRate            int `split_words:"true" default:"50"` // % chance to recieve an item from mining.  Set between 0 and 100.
	Digging               Digging
}

type GobbieBox struct {
	Enable int `split_words:"true" default:"1"` // Allows acquisition of daily points for gobbie mystery box.
	Amount int `split_words:"true" default:"10"`
	Limit  int `split_words:"true" default:"50000"`
	MinAge int `split_words:"true" default:"45"` // Minimum character age in days before a character can sign up for Gobbie Mystery Box
}

type Http struct {
	Enable bool   `split_words:"true" default:"false"`
	Host   string `split_words:"true" default:"localhost"`
	Port   int    `split_words:"true" default:"8088"`
}

type IllusionTime struct {
	Max time.Duration `split_words:"true" default:"1h"`  // 1 hour
	Min time.Duration `split_words:"true" default:"30m"` // 30 minutes
}

type Limbus struct {
	CosmoCleanseBaseCost int `split_words:"true" default:"15000"` // Base gil cost for a Cosmo Cleanse from Sagheera
}

type LimitBreakQuests struct {
	OldSchoolLimitBreak1 bool          `envconfig:"oldschool_1" default:"false"` // Set to true to require farming Exoray Mold, Bombd Coal, and Ancient Papyrus drops instead of allowing key item method.
	OldSchoolLimitBreak2 bool          `envconfig:"oldschool_2" default:"false"` // Set true to require the NMs for "Atop the Highest Mountains" be dead to get KI like before SE changed it.
	FrigiciteTime        time.Duration `split_words:"true" default:"30s"`        // When OLDSCHOOL_G2 is enabled, this is the time (in seconds) you have from killing Boreal NMs to click the "???" target.
}

type Log struct {
	Debug   bool `split_words:"true" default:"true"`
	Info    bool `split_words:"true" default:"true"`
	Warning bool `split_words:"true" default:"true"`
	Lua     bool `split_words:"true" default:"true"` // Prints from Lua using `print()`
}

type LoginCampaign struct {
	// Login Campaign (Set to 0 if you don't want to run a Login Campaign)
	// Please visit scripts/globals/events/login_campaign.lua for assigning the correct campaign dates.
	Enable int `split_words:"true" default:"0"`
}

type MagianTrials struct {
	Enable            int `split_words:"true" default:"1"`
	MobKillMultiplier int `split_words:"true" default:"1"`
	TradeMultiplier   int `split_words:"true" default:"1"`
}

type Mob struct {
	Entity

	// Also adjust monsters subjob in ratio adjustments? 1 = true / 0 = false
	IncludeSubjob bool `split_words:"true" default:"false"`

	// Allow mobs to walk back home instead of despawning
	NoDespawn bool `split_words:"true" default:"false"`

	// Adds extra time to mob despawn in seconds. Base time is 25s, so a setting of 5 here would be a total of 30 seconds.
	AdditionalTimeToDeaggro time.Duration `split_words:"true" default:"0s"`
}

type NetworkMap struct {
	Port int `split_words:"true" default:"54230"`
}

type NetworkSearch struct {
	Port int `split_words:"true" default:"54002"`
}

type NotoriousMonster struct {
	Entity

	// Multiplier to NM lottery spawn chance. (Default 1.0) eg. 0 = disable lottery spawns. -1 for always 100% chance.
	LotteryChance float64 `split_words:"true" default:"1.0"`
	// Multiplier to NM lottery cooldown time (Default 1.0) eg. 2.0 = twice as long. 0 = no cooldowns.
	LotteryCooldown float64 `split_words:"true" default:"1.0"`
}

type Nyzul struct {
	RunicDiskSave    bool          `split_words:"true" default:"true"` // Allow anyone participating in Nyzul to save progress. Set to false so only initiator can save progress.
	EnableCaskets    bool          `split_words:"true" default:"true"` // Enable Treasure casket pops from NMs.
	EnableVigilDrops bool          `split_words:"true" default:"true"` // Enable Vigil Weapon drops from NMs.
	ActivateLampTime time.Duration `split_words:"true" default:"6s"`   // Time in miliseconds for lamps to stay lit. TODO: Get retail confirmation.
}

type RecordsOfEminence struct {
	Enable              int `split_words:"true" default:"1"`      // Enable Records of Eminence
	EnableTimed         int `split_words:"true" default:"1"`      // Enable 4-hour timed records
	EnableExchangeLimit int `split_words:"true" default:"1"`      // Enable Maximum limit of sparks spent per Week (default retail behavior: 1)
	WeeklyExchangeLimit int `split_words:"true" default:"100000"` // Maximum amount of sparks/accolades that can be spent per week (default retail value: 100000)
}

type RhapsodiesOfVanadiel struct {
	Enable int `split_words:"true" default:"1"`
}

type Server struct {
	// Server name (not longer than 15 characters)
	Name    string `split_words:"true" default:"nameless"`
	Message string `split_words:"true" default:"Welcome to the server."`
}

type Skillup struct {
	// Enable/disable skill-ups from bloodpacts
	BloodPact bool `split_words:"true" default:"true"`

	// Allows parry, block, and guard to skill up regardless of the action occuring.
	// This did not happen in previous eras
	OldParryStyle bool `split_words:"true" default:"false"`
	OldBlockStyle bool `split_words:"true" default:"false"`
	OldGuardStyle bool `split_words:"true" default:"false"`

	// Allows you to manipulate the constant multiplier in the skill-up rate formulas, having a potent effect on skill-up rates.
	ChanceMultiplier float64 `split_words:"true" default:"1.0"`

	// Multiplier for skillup amounts. Using anything above 1 will break the 0.5 cap, the cap will become 0.9 (For maximum, set to 5)
	AmountMultiplier int `split_words:"true" default:"1"`
}

type Speed struct {
	// Modifier to apply to player speed. 0 is the retail accurate default. Negative numbers will reduce it.
	PlayerMod int `split_words:"true" default:"0"`

	// Modifier to apply to mount speed. 0 is the retail accurate default. Negative numbers will reduce it.
	// Note retail treats the mounted speed as double what it actually is.
	MountMod int `split_words:"true" default:"0"`

	// Modifier to apply to agro'd monster speed. 0 is the retail accurate default. Negative numbers will reduce it.
	MobMod int `split_words:"true" default:"0"`
}

type Spell struct {
	CurePower      float64 `split_words:"true" default:"1.000"` // Multiplies amount healed from Healing Magic, including the relevant Blue Magic.
	ElementalPower float64 `split_words:"true" default:"1.000"` // Multiplies damage dealt by Elemental and non-drain Dark Magic.
	DivinePower    float64 `split_words:"true" default:"1.000"` // Multiplies damage dealt by Divine Magic.
	NinjitsuPower  float64 `split_words:"true" default:"1.000"` // Multiplies damage dealt by Ninjutsu Magic.
	BluePower      float64 `split_words:"true" default:"1.000"` // Multiplies damage dealt by Blue Magic.
	DarkPower      float64 `split_words:"true" default:"1.000"` // Multiplies amount drained by Dark Magic.
	ItemPower      float64 `split_words:"true" default:"1.000"` // Multiplies the effect of items such as Potions and Ethers.

	// SPELL SPECIFIC SETTINGS
	DiaOverwrite                 int           `split_words:"true" default:"1"`     // Set to 1 to allow Bio to overwrite same tier Dia.  Default is 1.
	BioOverwrite                 int           `split_words:"true" default:"0"`     // Set to 1 to allow Dia to overwrite same tier Bio.  Default is 0.
	StoneskinCap                 int           `split_words:"true" default:"350"`   // Soft cap for hp absorbed by stoneskin
	BlinkShadows                 int           `split_words:"true" default:"2"`     // Number of shadows supplied by Blink spell
	SpikeEffectDuration          time.Duration `split_words:"true" default:"3m"`    // the duration of RDM, BLM spikes effects (not Reprisal)
	ElementalDebuffDuration      time.Duration `split_words:"true" default:"2m"`    // base duration of elemental debuffs
	AquaveilCounter              int           `split_words:"true" default:"1"`     // Base amount of hits Aquaveil absorbs to prevent spell interrupts. Retail is 1.
	AbsorbSpellAmount            int           `split_words:"true" default:"8"`     // how much of a stat gets absorbed by DRK absorb spells - expected to be a multiple of 8.
	AbsorbSpellTick              time.Duration `split_words:"true" default:"9s"`    // duration of 1 absorb spell tick
	SneakInvisDurationMultiplier float64       `split_words:"true" default:"1.000"` // multiplies duration of sneak, invis, deodorize to reduce player torture. 1 = retail behavior.
	UseOldCureFormula            bool          `split_words:"true" default:"false"` // true/false. if true, uses older cure formula (3*MND + VIT + 3*(healing skill/5)) // cure 6 will use the newer formula
	UseOldMagicDamage            bool          `split_words:"true" default:"false"` // true/false. if true, uses older magic damage formulas
}

type Sql struct {
	Host     string `split_words:"true" default:"127.0.0.1"`
	Port     int    `split_words:"true" default:"3306"`
	Username string `split_words:"true" required:"true"`
	Password string `split_words:"true" required:"true"`
	Database string `split_words:"true" default:"xidb"`
}

type Subjob struct {
	// Sets the fraction of MP a subjob provides to the main job. Retail is half and this acts as a divisor so default is 2
	MPDivisor float64 `envconfig:"mp_divisor" default:"2.0"`

	// Modify ratio of subjob-to-mainjob
	// 0            = no subjobs
	// 1            = 1/2   (default, 75/37, 99/49)
	// 2            = 2/3   (75/50, 99/66)
	// 3            = equal (75/75, 99/99)
	Ratio int `split_words:"true" default:"1"`
}

type Trust struct {
	Entity

	EnableCasting          int `split_words:"true" default:"1"`
	EnableQuests           int `split_words:"true" default:"1"`
	EnableCustomEngagement int `split_words:"true" default:"0"`

	EnableAlterEgoExtravaganza         int `split_words:"true" default:"0"` // 0 = disabled, 1 = summer/ny, 2 = spring/autumn, 3 = both
	EnableAlterEgoExtravaganzaAnnounce int `split_words:"true" default:"0"` // 0 = disabled, 1 = add announcement to player login
	EnableAlterEgoExpo                 int `split_words:"true" default:"0"` // 0 = disabled, 1 = expo - HPP/MPP/Status Resistance, 2 = expo plus (not implemented)
	EnableAlterEgoExpoAnnounce         int `split_words:"true" default:"0"` // 0 = disabled, 1 = add announcement to player login

	AlterEgoExtravaganzaMessage string `split_words:"true" default:"The Alter Ego Extravaganza Campaign is active!"`

	AlterEgoExpoMessage string `split_words:"true" default:"The Alter Ego Expo Campaign is active!"`
}

type Tcp struct {
	// ===========================
	// TCP Sockets Configuration
	// ===========================

	// Display debug reports (When something goes wrong during the report, the report is saved.)
	Debug bool `split_words:"true" default:"false"`

	// How long can a socket stall before closing the connection (in seconds)
	StallTime time.Duration `split_words:"true" default:"1m"`

	// ===========================
	// IP Rules Settings
	// ===========================

	// If IP's are checked when connecting. This also enables DDoS protection.
	EnableIpRules bool `split_words:"true" default:"true"`

	// Order of the checks
	//   deny,allow     : Checks deny rules, then allow rules. Allows if no rules match.
	//   allow,deny     : Checks allow rules, then deny rules. Allows if no rules match.
	//   mutual-failure : Allows only if an allow rule matches and no deny rules match.
	// (default is deny,allow)
	//TCP_ORDER = "deny,allow",
	//TCP_ORDER = "allow,deny",
	//TCP_ORDER = "mutual-failure",
	Order string `split_words:"true" default:"deny,allow"`

	// ===========================
	// IP rules
	// ===========================
	//   allow : Accepts connections from the ip range (even if flagged as DDoS)
	//   deny  : Rejects connections from the ip range
	// The rules are processed in order, the first matching rule of each list
	// (allow and deny) is used

	//TCP_ALLOW = "",
	//TCP_ALLOW = "127.0.0.1,192.168.0.0/16",
	//TCP_ALLOW = "127.0.0.1"
	//TCP_ALLOW = "192.168.0.0/16"
	//TCP_ALLOW = "10.0.0.0/255.0.0.0"
	//TCP_ALLOW = "all"
	Allow string `split_words:"true" default:""`

	//TCP_DENY = "",
	//TCP_DENY = "10.0.0.0/8,192.168.0.0/16",
	//TCP_DENY = "127.0.0.1",
	//TCP_DENY = "10.0.0.0/255.0.0.0",
	Deny string `split_words:"true" default:""`

	// ===========================
	// Connection Limit Settings
	// ===========================

	// If connect_count connection request are made within connect_interval
	// msec, it blocks it

	// Consecutive attempts interval (msec)
	// (default is 3000 msecs, 3 seconds)
	ConnectInterval time.Duration `split_words:"true" default:"3s"`

	// Consecutive attempts trigger
	// (default is 10 attempts)
	ConnectCount int `split_words:"true" default:"10"`

	// The time interval after which the connection lockout is lifted. (msec)
	// After this amount of time, the connection restrictions are lifted.
	// (default is 600000 msecs, 10 minutes)
	ConnectLockout time.Duration `split_words:"true" default:"10m"`
}

type Treasure struct {
	// SE implemented coffer/chest illusion time in order to prevent coffer farming. No-one in the same area can open a chest or coffer for loot (gil, gems & items)
	// till a random time between MIN_ILLSION_TIME and MAX_ILLUSION_TIME. During this time players can loot keyitem and item related to quests (AF, maps... etc.)
	CofferIllusionTime IllusionTime `split_words:"true"`
	ChestIllusionTime  IllusionTime `split_words:"true"`
}

type Udp struct {
	// ===========================
	// UDP Sockets Configuration
	// ===========================

	// Display debug reports (When something goes wrong during the report, the report is saved.)
	Debug bool `split_words:"true" default:"false"`
}

type Valor struct {
	EnableFieldManuals    int     `split_words:"true" default:"1"` // Enables Fields of Valor
	EnableGroundsTomes    int     `split_words:"true" default:"1"` // Enables Grounds of Valor
	EnableSurvivalGuide   int     `split_words:"true" default:"1"` // Enables Survival Guides (Not Implemented)
	RegimeWait            int     `split_words:"true" default:"1"` // Make people wait till 00:00 game time as in retail. If it's 0, there is no wait time.
	FieldsRewardAlliance  int     `split_words:"true" default:"0"` // Allow Fields of Valor rewards while being a member of an alliance. (default retail behavior: 0)
	GroundsRewardAlliance int     `split_words:"true" default:"1"` // Allow Grounds of Valor rewards while being a member of an alliance. (default retail behavior: 1)
	TabsRate              float64 `split_words:"true" default:"1.000"`
}

type Voidwalker struct {
	Enable int `split_words:"true" default:"1"`
}

type Voidwatch struct {
	Enable int `split_words:"true" default:"1"`
}

type WeaponSkill struct {
	// Multiplies damage dealt by Weapon Skills.
	Power float64 `split_words:"true" default:"1.000"`

	// Weaponskill point base (before skillchain) for breaking latent - whole numbers only. retail is 1.
	PointsBase int `envconfig:"points_base" default:"1"`

	// Weaponskill points per skillchain element - whole numbers only, retail is 1
	// (tier 3 sc's have 4 elements, plus 1 for the ws itself, giving 5 points to the closer).
	PointsSkillchain int `envconfig:"points_skillchain" default:"1"`
}

type Zmq struct {
	// Central message server settings (ensure these are the same on both all map servers and the central (lobby) server
	Ip   string `split_words:"true" default:"127.0.0.1"`
	Port int    `split_words:"true" default:"54003"`
}
