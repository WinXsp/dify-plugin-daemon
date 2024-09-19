package app

import "golang.org/x/exp/constraints"

func (config *Config) SetDefault() {
	setDefaultInt(&config.ServerPort, 5002)
	setDefaultInt(&config.RoutinePoolSize, 1000)
	setDefaultInt(&config.LifetimeCollectionGCInterval, 60)
	setDefaultInt(&config.LifetimeCollectionHeartbeatInterval, 5)
	setDefaultInt(&config.LifetimeStateGCInterval, 300)
	setDefaultInt(&config.DifyInvocationConnectionIdleTimeout, 120)
	setDefaultInt(&config.PluginRemoteInstallServerEventLoopNums, 8)
	setDefaultInt(&config.PluginRemoteInstallingMaxConn, 128)
	setDefaultInt(&config.MaxPluginPackageSize, 52428800)
	setDefaultInt(&config.MaxAWSLambdaTransactionTimeout, 150)
	setDefaultInt(&config.PluginMaxExecutionTimeout, 240)
	setDefaultInt(&config.PluginMediaCacheSize, 1024)
	setDefaultBool(&config.PluginRemoteInstallingEnabled, true)
	setDefaultBool(&config.PluginEndpointEnabled, true)
	setDefaultString(&config.DBSslMode, "disable")
	setDefaultString(&config.PluginMediaCachePath, "./storage/assets")
	setDefaultString(&config.PersistenceStorageLocalPath, "./storage/persistence")
	setDefaultString(&config.ProcessCachingPath, "./storage/subprocesses")
}

func setDefaultInt[T constraints.Integer](value *T, defaultValue T) {
	if *value == 0 {
		*value = defaultValue
	}
}

func setDefaultString(value *string, defaultValue string) {
	if *value == "" {
		*value = defaultValue
	}
}

func setDefaultBool(value *bool, defaultValue bool) {
	if !*value {
		*value = defaultValue
	}
}