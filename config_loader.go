package config_loader

type ConfigLoader interface {
	LoadByFile(filename string)
	LoadByBytes(content []byte)
	SetDefault()
}
