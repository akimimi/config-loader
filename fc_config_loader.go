package config_loader

type FcServiceConfig struct {
	Endpoint        string `json:"endpoint" yaml:"endpoint"`
	AccessKeyId     string `json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" yaml:"access_key_secret"`
}

func (c *FcServiceConfig) LoadByFile(filename string) {
	LoadByFile(filename, c)
}

func (c *FcServiceConfig) LoadByBytes(content []byte) {
	LoadByBytes(content, c)
}

func (c *FcServiceConfig) SetDefault() {
}

type FunctionConfig struct {
	ServiceName  string `json:"service_name" yaml:"service_name"`
	FunctionName string `json:"function_name" yaml:"function_name"`
}

func (c *FunctionConfig) LoadByFile(filename string) {
	LoadByFile(filename, c)
}

func (c *FunctionConfig) LoadByBytes(content []byte) {
	LoadByBytes(content, c)
}

func (c *FunctionConfig) SetDefault() {
}
