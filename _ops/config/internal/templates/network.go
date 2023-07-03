package templates

import (
	"log"
	"text/template"
)

const (
	networkConfig string = `-----------------------------------
-- NETWORK SETTINGS
-----------------------------------
-- All settings are attached to the 'xi.settings' object. This is published globally, and be accessed from C++ and any script.
--
-- This file is concerned mainly with networking between the database, client, and server executables.
-----------------------------------

xi = xi or {}
xi.settings = xi.settings or {}

xi.settings.network =
{
    SQL_HOST     = "{{ .Sql.Host }}",
    SQL_PORT     = {{ .Sql.Port }},
    SQL_LOGIN    = "{{ .Sql.Username }}",
    SQL_PASSWORD = "{{ .Sql.Password }}",
    SQL_DATABASE = "{{ .Sql.Database }}",

    LOGIN_DATA_IP   = "{{ .Login.DataIp }}",
    LOGIN_DATA_PORT = {{ .Login.DataPort }},
    LOGIN_VIEW_IP   = "{{ .Login.ViewIp }}",
    LOGIN_VIEW_PORT = {{ .Login.ViewPort }},
    LOGIN_AUTH_IP   = "{{ .Login.AuthIp }}",
    LOGIN_AUTH_PORT = {{ .Login.AuthPort }},
    LOGIN_CONF_IP   = "{{ .Login.ConfIp }}",
    LOGIN_CONF_PORT = {{ .Login.ConfPort }},

    MAP_PORT = {{ .Map.Port }},

    SEARCH_PORT = {{ .Search.Port }},

    ENABLE_HTTP = {{ .Http.Enable }},
    HTTP_HOST   = "{{ .Http.Host }}",
    HTTP_PORT   = {{ .Http.Port }},

    -- Central message server settings (ensure these are the same on both all map servers and the central (lobby) server
    ZMQ_IP   = "{{ .Zmq.Ip }}",
    ZMQ_PORT = {{ .Zmq.Port }},

    -- ===========================
    -- NOTE: The settings that follow will not necessarily need to be modified
    --       in any way for the server to work out of the box.  This should only
    --       be modified by those who understand networking.  Modifying these
    --       values could potentially make it so that you can not log in to your
    --       server.
    -- ===========================

    -- ===========================
    -- UDP Sockets Configuration
    -- ===========================

    -- Display debug reports (When something goes wrong during the report, the report is saved.)
    UDP_DEBUG = {{ .Udp.Debug }},

    -- ===========================
    -- TCP Sockets Configuration
    -- ===========================

    -- Display debug reports (When something goes wrong during the report, the report is saved.)
    TCP_DEBUG = {{ .Tcp.Debug }},

    -- How long can a socket stall before closing the connection (in seconds)
    TCP_STALL_TIME = {{ .Tcp.StallTime.Seconds }},

    -- ===========================
    -- IP Rules Settings
    -- ===========================

    -- If IP's are checked when connecting. This also enables DDoS protection.
    TCP_ENABLE_IP_RULES = {{ .Tcp.EnableIpRules }},

    -- Order of the checks
    --   deny,allow     : Checks deny rules, then allow rules. Allows if no rules match.
    --   allow,deny     : Checks allow rules, then deny rules. Allows if no rules match.
    --   mutual-failure : Allows only if an allow rule matches and no deny rules match.
    -- (default is deny,allow)
    --TCP_ORDER = "deny,allow",
    --TCP_ORDER = "allow,deny",
    --TCP_ORDER = "mutual-failure",
    TCP_ORDER = "{{ .Tcp.Order }}",

    -- ===========================
    -- IP rules
    -- ===========================
    --   allow : Accepts connections from the ip range (even if flagged as DDoS)
    --   deny  : Rejects connections from the ip range
    -- The rules are processed in order, the first matching rule of each list
    -- (allow and deny) is used

    --TCP_ALLOW = "",
    --TCP_ALLOW = "127.0.0.1,192.168.0.0/16",
    --TCP_ALLOW = "127.0.0.1"
    --TCP_ALLOW = "192.168.0.0/16"
    --TCP_ALLOW = "10.0.0.0/255.0.0.0"
    --TCP_ALLOW = "all"
    TCP_ALLOW = "{{ .Tcp.Allow }}",

    --TCP_DENY = "",
    --TCP_DENY = "10.0.0.0/8,192.168.0.0/16",
    --TCP_DENY = "127.0.0.1",
    --TCP_DENY = "10.0.0.0/255.0.0.0",
    TCP_DENY = "{{ .Tcp.Deny }}",

    -- ===========================
    -- Connection Limit Settings
    -- ===========================

    -- If connect_count connection request are made within connect_interval
    -- msec, it blocks it

    -- Consecutive attempts interval (msec)
    -- (default is 3000 msecs, 3 seconds)
    TCP_CONNECT_INTERVAL = {{ .Tcp.ConnectInterval.Milliseconds }},

    -- Consecutive attempts trigger
    -- (default is 10 attempts)
    TCP_CONNECT_COUNT = {{ .Tcp.ConnectCount }},

    -- The time interval after which the connection lockout is lifted. (msec)
    -- After this amount of time, the connection restrictions are lifted.
    -- (default is 600000 msecs, 10 minutes)
    TCP_CONNECT_LOCKOUT = {{ .Tcp.ConnectLockout.Milliseconds }}
}
`
)

var (
	NetworkConfig *template.Template
)

func init() {
	tmpl, err := template.New("networkConfig").
		Funcs(template.FuncMap{}).
		Parse(networkConfig)
	if err != nil {
		log.Fatalln(err)
	}
	NetworkConfig = tmpl
}
