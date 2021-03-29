# msqueue_config_loader


Define Message Queue Config struct and implement LoadByFile and LoadByBytes interfaces.

```
type QueueConfig struct {
	Url             string `json:"url" yaml:"url"`
	AccessKeyId     string `json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret" yaml:"access_key_secret"`
	QueueName       string `json:"queue_name" yaml:"queue_name"`
	MaxDequeueCount int    `json:"max_dequeue_count" yaml:"max_dequeue_count"`
}
```

