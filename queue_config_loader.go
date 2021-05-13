package config_loader

const DefaultDelaySeconds = 0
const DefaultMaxDequeueCount = 1
const DefaultMaxMessageSize = 65536
const DefaultMessageRetentionPeriod = 7 * 24 * 3600 // 7days
const DefaultVisibilityTimeout = 60
const DefaultPollingWaitSeconds = 10
const DefaultConsumeTime = -1
const DefaultMaxProcessingMessage = 200
const DefaultOverloadBreakSeconds = 120
const DefaultRecvMessageBatchSize = 1

type QueueConfig struct {
	Url                    string `json:"url" yaml:"url"`
	AccessKeyId            string `json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret        string `json:"access_key_secret" yaml:"access_key_secret"`
	QueueName              string `json:"queue_name" yaml:"queue_name"`
	MaxDequeueCount        int    `json:"max_dequeue_count" yaml:"max_dequeue_count"`
	DelaySeconds           int    `json:"delay_seconds" yaml:"delay_seconds"`
	MaxMessageSize         int    `json:"max_message_size" yaml:"max_message_size"`
	MessageRetentionPeriod int    `json:"message_retention_period" yaml:"message_retention_period"`
	VisibilityTimeout      int    `json:"visibility_timeout" yaml:"visibility_timeout"`
	PollingWaitSeconds     int    `json:"polling_wait_seconds" yaml:"polling_wait_seconds"`
	Verbose                bool   `json:"verbose" yaml:"verbose"`
	ConsumeTimeout         int    `json:"consume_timeout" yaml:"consume_timeout"`
	MaxProcessingMessage   int    `json:"max_processing_message" yaml:"max_processing_message"`
	OverloadBreakSeconds   int    `json:"overload_break_seconds" yaml:"overload_break_seconds"`
	RecvMessageBatchSize   int    `json:"recv_message_batch_size" yaml:"recv_message_batch_size"`
}

func (c *QueueConfig) LoadByFile(filename string) {
	LoadByFile(filename, c)
	c.SetDefault()
}

func (c *QueueConfig) LoadByBytes(content []byte) {
	LoadByBytes(content, c)
	c.SetDefault()
}

func (c *QueueConfig) SetDefault() {
	if c.DelaySeconds == 0 {
		c.DelaySeconds = DefaultDelaySeconds
	}
	if c.MaxDequeueCount == 0 {
		c.MaxDequeueCount = DefaultMaxDequeueCount
	}
	if c.MaxMessageSize == 0 {
		c.MaxDequeueCount = DefaultMaxMessageSize
	}
	if c.MessageRetentionPeriod == 0 {
		c.MessageRetentionPeriod = DefaultMessageRetentionPeriod
	}
	if c.VisibilityTimeout == 0 {
		c.VisibilityTimeout = DefaultVisibilityTimeout
	}
	if c.PollingWaitSeconds == 0 {
		c.PollingWaitSeconds = DefaultPollingWaitSeconds
	}
	if c.ConsumeTimeout == 0 {
		c.ConsumeTimeout = DefaultConsumeTime
	}
	if c.MaxProcessingMessage == 0 {
		c.MaxProcessingMessage = DefaultMaxProcessingMessage
	}
	if c.OverloadBreakSeconds == 0 {
		c.OverloadBreakSeconds = DefaultOverloadBreakSeconds
	}
	if c.RecvMessageBatchSize == 0 {
		c.RecvMessageBatchSize = DefaultRecvMessageBatchSize
	}
}
