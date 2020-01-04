package lib

var (
	ConfPath = "config"
	ConfName = "task3"
)

type AppConfig struct {
	HttpAPI struct {
		Base string `mapstructure:"base"`
		Addr string `mapstructure:"addr"`
	} `mapstructure:"httpApi"`
	Task struct {
		Addr    string `mapstructure:"addr"`
		Clients int    `mapstructure:"clients"`
	} `mapstructure:"task"`
	Redis struct {
		Addr       string `mapstructure:"addr"`
		ScoreKey   string `mapstructure:"scoreKey"`
		WinnersKey string `mapstructure:"winnersKey"`
	} `mapstructure:"redis"`
}
