package config

import (
	"flag"
	"os"
	"path/filepath"
)

// IniParserOptName is the option name to indicate the path of the nin config file.
var IniParserOptName = "config-file"

// Conf is the global config manager.
var Conf = NewDefault()

// NewDefault returns a new default config manager.
func NewDefault() *Config {
	cli := NewFlagCliParser(filepath.Base(os.Args[0]), flag.ExitOnError)
	ini := NewSimpleIniParser(IniParserOptName)
	conf := NewConfig(cli).AddParser(ini)
	conf.RegisterCliOpt("", StrOpt("", IniParserOptName, nil, "The path of the ini config file."))
	return conf
}
