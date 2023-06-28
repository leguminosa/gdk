package envx

import "os"

const (
	envKey = "ENV"

	DevelopmentEnv = "development"
	StagingEnv     = "staging"
	ProductionEnv  = "production"
)

func ServiceEnv() string {
	return serviceEnv()
}

func IsDevelopment() bool {
	return serviceEnv() == DevelopmentEnv
}

func IsStaging() bool {
	return serviceEnv() == StagingEnv
}

func IsProduction() bool {
	return serviceEnv() == ProductionEnv
}

func serviceEnv() string {
	env := os.Getenv(envKey)
	if env == "" {
		env = DevelopmentEnv
	}
	return env
}
