package config

// type EnvVar string

const (
  Host = "HOST"
  Username = "USERNAME"
  Password = "PASSWORD"
  DatabaseName = "DATABASE"
  Flavor = "FLAVOR"
  Port = "PORT"
)

// TODO:
// What are the configurables for this program?
/////// EXTERNAL CONFIGURABLES /////////////////
// • Database Flavor: MySQL, Postgres, or Aurora
// • Username
// • Password
// • Host
// • Port
// • Database name
/////// INTERNAL CONFIGURABLES /////////////////
// • How often do we monitor?
// • How long is the ctx timeout between monitoring intervals?
// • What is the API URI for the OT API Gateway?

func init() {
  viper.SetEnvPrefix("OT")
  viper.BindEnv(Host)
  viper.BindEnv(Username)
  viper.BindEnv(Password)
  viper.BindEnv(DatabaseName)
  viper.BindEnv(Flavor)
  viper.BindEnv(Port)
}
