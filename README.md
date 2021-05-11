# config-loader


Define Message Queue Config struct and implement LoadByFile, LoadByBytes and SetDefault interfaces.

```
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
}
```

