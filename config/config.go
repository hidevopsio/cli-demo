package config

type Cluster struct {
	Cluster string 		`json:"cluster"`
	Username string 	`json:"username"`
	Token string    	`json:"token"`
}

type Hicli struct {
	Clusters  []Cluster `json:"clusters"`
	LastIndex string 	`json:"lastIndex"`
}

type Configuration struct {
	 Hicli Hicli  		 `mapstructure:"hicli"`
}