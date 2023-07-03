package templates

import (
	"log"
	"text/template"
)

const (
	searchConfig string = `-----------------------------------
-- SEARCH SERVER SETTINGS
-----------------------------------
-- All settings are attached to the 'xi.settings' object. This is published globally, and be accessed from C++ and any script.
--
-- This file is concerned mainly with /sea, searching, and the auction house.
-----------------------------------

xi = xi or {}
xi.settings = xi.settings or {}

xi.settings.search =
{
	-- Omit items with no buy history from auction house results
	OMIT_NO_HISTORY = {{ .OmitNoHistory }},

	-- After EXPIRE_DAYS, will listed auctions expire?
	EXPIRE_AUCTIONS = {{ .ExpireAuctions }},

	-- Expire items older than this number of days
	EXPIRE_DAYS = {{ .ExpireDays }},

	-- Interval is in seconds, default is one hour
	EXPIRE_INTERVAL = {{ .ExpireInterval.Seconds }},
}
`
)

var (
	SearchConfig *template.Template
)

func init() {
	tmpl, err := template.New("searchConfig").
		Funcs(template.FuncMap{}).
		Parse(searchConfig)
	if err != nil {
		log.Fatalln(err)
	}
	SearchConfig = tmpl
}
