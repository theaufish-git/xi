package templates

import (
	"log"
	"text/template"
)

const (
	loggingConfig string = `-----------------------------------
-- LOGGING SETTINGS
-----------------------------------
-- All settings are attached to the 'xi.settings' object. This is published globally, and be accessed from C++ and any script.
--
-- This file is concerned mainly with configuring the logging across all server executables.
-----------------------------------

xi = xi or {}
xi.settings = xi.settings or {}

xi.settings.logging =
{
    --[[
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
    --]]
    PATTERN = "{{ .Pattern }}",

    -- Enable/Disable these logging types globally
    LOG_DEBUG   = {{ .Log.Debug }},
    LOG_INFO    = {{ .Log.Info }},
    LOG_WARNING = {{ .Log.Warning }},
    LOG_LUA     = {{ .Log.Lua }}, -- Prints from Lua using 'print()'

    -- Specific Debug loggers
    -- NOTE: None of these will print unless you also have the above LOG_DEBUG setting set to true!
    DEBUG_SOCKETS        = {{ .Debug.Sockets }}, -- Calls in C++: DebugSockets(...)
    DEBUG_NAVMESH        = {{ .Debug.Navmesh }}, -- Calls in C++: DebugNavmesh(...)
    DEBUG_PACKETS        = {{ .Debug.Packets }}, -- Calls in C++: DebugPackets(...)
    DEBUG_ACTIONS        = {{ .Debug.Actions }}, -- Calls in C++: DebugActions(...)
    DEBUG_SQL            = {{ .Debug.Sql }}, -- Calls in C++: DebugSQL(...)
    DEBUG_ID_LOOKUP      = {{ .Debug.IdLookup }}, -- Calls in C++: DebugIDLookup(...)
    DEBUG_MODULES        = {{ .Debug.Modules }}, -- Calls in C++: DebugModules(...)
    DEBUG_PACKET_BACKLOG = {{ .Debug.PacketBacklog }}, -- Special logic in map.cpp::send_parse
    DEBUG_DELIVERY_BOX   = {{ .Debug.DeliveryBox }}, -- Special logic in packet_system.cpp::SmallPacket0x04D
}
`
)

var (
	LoggingConfig *template.Template
)

func init() {
	tmpl, err := template.New("loggingConfig").
		Funcs(template.FuncMap{}).
		Parse(loggingConfig)
	if err != nil {
		log.Fatalln(err)
	}
	LoggingConfig = tmpl
}
