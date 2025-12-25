package config

var TomlConfig = []byte(`
# This is a TOML document.
hostname = "127.0.0.1"
hostport = "auto"
ssccprefix = "1462709"
ssccstartnumber = 21
perpallet = 24
entirely = false
statuskm = true

[layouts]
timelayout = "2006-01-02T15:04:05-0700"
timelayoutclear = "2006.01.02 15:04:05"
timelayoutday = "2006.01.02"
timelayoututc = "2006-01-02T15:04:05"

`)

type Configuration struct {
	Hostname        string              `json:"hostname"`
	HostPort        string              `json:"hostport"`
	Output          string              `json:"output"`
	SsccPrefix      string              `json:"ssccprefix"`
	SsccStartNumber int                 `json:"ssccstartnumber"`
	PerPallet       int                 `json:"perpallet"`
	Entirely        bool                `json:"entirely"`
	StatusKM        bool                `json:"statuskm"`
	Layouts         LayoutConfiguration `json:"layouts"`
}

type LayoutConfiguration struct {
	TimeLayout      string `json:"timelayout"`
	TimeLayoutClear string `json:"timelayoutclear"`
	TimeLayoutDay   string `json:"timelayoutday"`
	TimeLayoutUTC   string `json:"timelayoututc"`
}
