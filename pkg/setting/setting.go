package setting

type Config struct {
	PostgreSql PostgreSQLSetting `mapstructure:"postgresql"`
	Logger     LoggerSetting     `mapstructure:"logger"`
	Redis      RedisSetting      `mapstructure:"redis"`
	Server     ServerSetting     `mapstructure:"server"`
	Jwt        JwtSetting        `mapstructure:"token"`
	CoinMarket CoinMarketSetting `mapstructure:"coinmarket"`
	SMTP       SMTPSetting       `mapstructure:"smtp"`
	Kafka      KafkaSetting      `mapstructure:"kafka"`
}

type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"database"`
}

type PostgreSQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	DBName          string `mapstructure:"dbname"`
	SslMode         string `mapstructure:"sslmode"`
	Timezone        string `mapstructure:"timezone"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

type LoggerSetting struct {
	Log_level     string `mapstructure:"log_level"`
	File_log_name string `mapstructure:"file_log_name"`
	Max_backups   int    `mapstructure:"max_backups"`
	Max_age       int    `mapstructure:"max_age"`
	Max_size      int    `mapstructure:"max_size"`
	Compress      bool   `mapstructure:"compress"`
}

type CoinMarketSetting struct {
	URLApi      string `mapstructure:"url_api"`
	Host        string `mapstructure:"host"`
	CurrencyAPI string `mapstructure:"currency_api"`
	APIKey      string `mapstructure:"api_key"`
}

type JwtSetting struct {
	AccessToken       string `mapstructure:"access_token"`
	RefreshToken      string `mapstructure:"refresh_token"`
	Expiration        int    `mapstructure:"expiration"`
	RefreshExpiration int    `mapstructure:"refresh_expiration"`
}

type SMTPSetting struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
}

type KafkaSetting struct {
	Brokers      []string          `mapstructure:"brokers"`
	GroupID      string            `mapstructure:"group_id"`
	Topics       KafkaTopics       `mapstructure:"topics"`
	Producer     KafkaProducer     `mapstructure:"producer"`
	Consumer     KafkaConsumer     `mapstructure:"consumer"`
	Notification KafkaNotification `mapstructure:"notification"`
}

type KafkaTopics struct {
	Notification KafkaTopic `mapstructure:"notification"`
}

type KafkaTopic struct {
	Name              string `mapstructure:"name"`
	Partitions        int    `mapstructure:"partitions"`
	ReplicationFactor int    `mapstructure:"replicationFactor"`
}

type KafkaProducer struct {
	Retries         int    `mapstructure:"retries"`
	Acks            string `mapstructure:"acks"`
	BatchSize       int    `mapstructure:"batchSize"`
	LingerMs        int    `mapstructure:"lingerMs"`
	BufferMemory    int    `mapstructure:"bufferMemory"`
	KeySerializer   string `mapstructure:"keySerializer"`
	ValueSerializer string `mapstructure:"valueSerializer"`
}

type KafkaConsumer struct {
	GroupID           string `mapstructure:"groupId"`
	EnableAutoCommit  bool   `mapstructure:"enableAutoCommit"`
	AutoOffsetReset   string `mapstructure:"autoOffsetReset"`
	MaxPollRecords    int    `mapstructure:"maxPollRecords"`
	KeyDeserializer   string `mapstructure:"keyDeserializer"`
	ValueDeserializer string `mapstructure:"valueDeserializer"`
}

type KafkaNotification struct {
	Retry KafkaRetry `mapstructure:"retry"`
}

type KafkaRetry struct {
	MaxRetries int `mapstructure:"maxRetries"`
	DelayMs    int `mapstructure:"delayMs"`
}
