package templates

import (
	"log"
	"text/template"
)

const (
	loginConfig string = `-----------------------------------
-- LOGIN SERVER SETTINGS
-----------------------------------
-- All settings are attached to the 'xi.settings' object. This is published globally, and be accessed from C++ and any script.
--
-- This file is concerned mainly with logging into the game, and managing login sessions.
-----------------------------------

xi = xi or {}
xi.settings = xi.settings or {}

xi.settings.login =
{
    -- Expected Client version (wrong version cannot log in)
    CLIENT_VER = "{{ .Client.Version }}",

    -- 0 - disabled (every version allowed)
    -- 1 - enabled - strict (only exact CLIENT_VER allowed)
    -- 2 - enabled - greater than or equal  (matching or greater than CLIENT_VER allowed, default)
    --
    -- WE STRONGLY ADVISE AGAINST LOCKING THE SERVER TO OLDER VERSIONS. IT IS A UNIVERSALLY BAD IDEA.
    VER_LOCK = {{ .Client.VersionLock }},

    -- 0 - disabled (normal operation)
    -- 1 - enabled (only GM characters allowed online, no new character creation)
    MAINT_MODE = {{ .MaintenanceMode }},

    -- Logging of user IP address to database (true/false)
    LOG_USER_IP = {{ .LogUserIp }},

    -- Allow account creation via the loader (true/false)
    ACCOUNT_CREATION = {{ .AccountCreation }},

    -- Allow character deletion through the lobby (true/false)
    CHARACTER_DELETION = {{ .CharacterDeletion }},

    -- Number of simultaneous game sessions per IP (0 for no limit)
    LOGIN_LIMIT = {{ .LoginLimit }},

    -- If true, blocks character creation with names of NPCs and Mobs in the database (Fafnir, Shantotto, etc.)
    DISABLE_MOB_NPC_CHAR_NAMES = {{ .DisableMobNpcCharNames }},

    -- Character names with any of these words in, in any position, will be rejected
    --
    -- Examples that will be rejected (using "badword"):
    -- "badword"
    -- "imbadword"
    -- "badwordisme"
    -- "lolbadwordlol"
    --
    -- WARNING:
    -- Be aware of the "Scunthorpe problem"!
    --
    -- NOTE:
    -- You can Google for "bad word list txt" to find lists of words to populate this table, if you'd like
    BANNED_WORDS_LIST =
    {
    {{- range $word := .BannedWordsList }}
        "{{ $word }}",
    {{- end }}
    }
}
`
)

var (
	LoginConfig *template.Template
)

func init() {
	tmpl, err := template.New("loginConfig").
		Funcs(template.FuncMap{}).
		Parse(loginConfig)
	if err != nil {
		log.Fatalln(err)
	}
	LoginConfig = tmpl
}
