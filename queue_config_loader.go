package config_loader

const DefaultDelaySeconds = 0
const DefaultMaxDequeueCount = 1
const DefaultMaxMessageSize = 65536
const DefaultMessageRententionPeriod = 7 * 24 * 3600 // 7days
const DefaultVisibilityTimeout = 60
const DefaultPollingWaitSeconds = 10

type QueueConfig struct {
	Url                     string `json:"url" yaml:"url"`
	AccessKeyId             string `json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret         string `json:"access_key_secret" yaml:"access_key_secret"`
	QueueName               string `json:"queue_name" yaml:"queue_name"`
	MaxDequeueCount         int    `json:"max_dequeue_count" yaml:"max_dequeue_count"`
	DelaySeconds            int    `json:"delay_seconds" yaml:"delay_seconds"`
	MaxMessageSize          int    `json:"max_message_size" yaml:"max_message_size"`
	MessageRententionPeriod int    `json:"message_rentention_period" yaml:"message_rentention_period"`
	VisibilityTimeout       int    `json:"visibility_timeout" yaml:"visibility_timeout"`
	PollingWaitSeconds      int    `json:"polling_wait_seconds" yaml:"polling_wait_seconds"`
	Verbose                 bool   `json:"verbose" yaml:"verbose"`
}

func (c *QueueConfig) LoadByFile(filename string) {
	LoadByFile(filename, c)
	c.setDefault()
}

func (c *QueueConfig) LoadByBytes(content []byte) {
	LoadByBytes(content, c)
	c.setDefault()
}

func (c *QueueConfig) setDefault() {
	if c.DelaySeconds == 0 {
		c.DelaySeconds = DefaultDelaySeconds
	}
	if c.MaxDequeueCount == 0 {
		c.MaxDequeueCount = DefaultMaxDequeueCount
	}
	if c.MaxMessageSize == 0 {
		c.MaxDequeueCount = DefaultMaxMessageSize
	}
	if c.MessageRententionPeriod == 0 {
		c.MessageRententionPeriod = DefaultMessageRententionPeriod
	}
	if c.VisibilityTimeout == 0 {
		c.VisibilityTimeout = DefaultVisibilityTimeout
	}
	if c.PollingWaitSeconds == 0 {
		c.PollingWaitSeconds = DefaultPollingWaitSeconds
	}
}
