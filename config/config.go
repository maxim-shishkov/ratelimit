package config

import (
	"flag"
	"github.com/go-ini/ini"
	"github.com/rs/zerolog/log"
)

type Sectioner interface {
	Init()
}

var (
	ConfigPath        = flag.String("config", "", "Path to file conf.ini")
	config            *ini.File
	ConfigSectionsPtr *map[string]Sectioner
)

var sections = map[string]Sectioner{
	"common": Common,
	"limit":  Limit,
	"redis":  Redis,
}

func Init() {
	flag.Parse()
	if *ConfigPath == "" {
		*ConfigPath = "./config.ini"
	}

	var err error
	config, err = ini.ShadowLoad(*ConfigPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Bad Config file: " + *ConfigPath)
	}
	config.NameMapper = ini.TitleUnderscore

	for sectionName, section := range sections {
		initSection(sectionName, section)
	}
	ConfigSectionsPtr = &sections
}

func initSection(sectionName string, section Sectioner) {
	config.Section(sectionName).MapTo(section)
	section.Init()
}
