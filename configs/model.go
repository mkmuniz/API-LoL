package configs

type config struct {
	API     APIConfig
	RiotAPI APILoL
}

type APIConfig struct {
	Port string
	Host string
}

type APILoL struct {
	Key string
}
