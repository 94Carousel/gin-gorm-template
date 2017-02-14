package config

import ini "gopkg.in/ini.v1"

// GetSection return config.ini settings
func GetSection(key string) *ini.Section {
	iniCfg := ini.LoadOptions{AllowBooleanKeys: true, Insensitive: true}
	cfg, _ := ini.LoadSources(iniCfg, "config/config.ini")
	section, _ := cfg.GetSection(key)
	return section
}

// Get return section and value
func Get(key string, value string) *ini.Key {
	return GetSection(key).Key(value)
}
