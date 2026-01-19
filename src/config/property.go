package config

// PropertyName represents the name of a configuration key.
type PropertyName string

const (
	// PropertyNameCacheHost represents the cache host address.
	PropertyNameCacheHost PropertyName = "CACHE_HOST"

	// PropertyNameCacheIdentifier represents the cache identifier.
	PropertyNameCacheIdentifier PropertyName = "CACHE_IDENTIFIER"

	// PropertyNameCacheUsername represents the cache username.
	PropertyNameCacheUsername PropertyName = "CACHE_USERNAME"

	// PropertyNameCachePassword represents the cache password.
	PropertyNameCachePassword PropertyName = "CACHE_PASSWORD"

	// PropertyNameCachePasswordFile represents the file path from which to read the cache password.
	PropertyNameCachePasswordFile PropertyName = "CACHE_PASSWORD_FILE"

	// PropertyNameCachePort represents the cache port.
	PropertyNameCachePort PropertyName = "CACHE_PORT"

	// PropertyNameDatabaseHost represents the database host address.
	PropertyNameDatabaseHost PropertyName = "DATABASE_HOST"

	// PropertyNameDatabaseName represents the database name.
	PropertyNameDatabaseName PropertyName = "DATABASE_NAME"

	// PropertyNameDatabasePassword represents the database password.
	PropertyNameDatabasePassword PropertyName = "DATABASE_PASSWORD"

	// PropertyNameDatabasePasswordFile represents the file path from which to read the database password.
	PropertyNameDatabasePasswordFile PropertyName = "DATABASE_PASSWORD_FILE"

	// PropertyNameDatabasePort represents the database port.
	PropertyNameDatabasePort PropertyName = "DATABASE_PORT"

	// PropertyNameDatabaseUsername represents the database username.
	PropertyNameDatabaseUsername PropertyName = "DATABASE_USERNAME"

	// PropertyNameDatabaseSSLMode represents the database SSL mode configuration.
	PropertyNameDatabaseSSLMode PropertyName = "DATABASE_SSL_MODE"

	// PropertyNameDatabaseTimezone represents the database default timezone.
	PropertyNameDatabaseTimezone PropertyName = "DATABASE_TIMEZONE"

	// PropertyNameDebugMode represents whether debugging mode is turned on.
	PropertyNameDebugMode PropertyName = "DEBUG"

	// PropertyNameEnvironment represents the environment on which the service is running.
	PropertyNameEnvironment PropertyName = "ENVIRONMENT"

	// PropertyNameFeatureFlagSDKKey represents the SDK key for the feature flag service.
	PropertyNameFeatureFlagSDKKey PropertyName = "FEATURE_FLAG_SDK_KEY"

	// PropertyNameFeatureFlagSDKKeyFile represents the file path from which to read the feature flag SDK key.
	PropertyNameFeatureFlagSDKKeyFile PropertyName = "FEATURE_FLAG_SDK_KEY_FILE"

	// PropertyNameGRPCPort represents the port on which the gRPC server will be listening.
	PropertyNameGRPCPort PropertyName = "GRPC_PORT"

	// PropertyNameHttpPort represents the port on which the HTTP service will be listening.
	PropertyNameHTTPPort PropertyName = "PORT"

	// PropertyNameLoadEnvFromFile represents whether to load environment variables from a .env file.
	PropertyNameLoadEnvFromFile PropertyName = "LOAD_ENV_FROM_FILE"

	// PropertyNameMailHost represents the mail server host address.
	PropertyNameMailHost PropertyName = "MAIL_HOST"

	// PropertyNameMailPassword represents the mail server password.
	PropertyNameMailPassword PropertyName = "MAIL_PASSWORD"

	// PropertyNameMailPasswordFile represents the file path from which to read the mail server password.
	PropertyNameMailPasswordFile PropertyName = "MAIL_PASSWORD_FILE"

	// PropertyNameMailPort represents the mail server port.
	PropertyNameMailPort PropertyName = "MAIL_PORT"

	// PropertyNameMailSenderAddress represents the mail sender's email address.
	PropertyNameMailSenderAddress PropertyName = "MAIL_SENDER_ADDRESS"

	// PropertyNameMailSenderName represents the mail sender's human-readable name.
	PropertyNameMailSenderName PropertyName = "MAIL_SENDER_NAME"

	// PropertyNameMailUsername represents the mail server username.
	PropertyNameMailUsername PropertyName = "MAIL_USERNAME"

	// PropertyNameServiceName represents the human-readable name of the service that is running.
	PropertyNameServiceName PropertyName = "NAME"
)

// GetAvailableConfigurationKeys returns a slice of all available configuration property names.
func GetAvailableConfigurationKeys() []PropertyName {
	return []PropertyName{
		PropertyNameCacheHost,
		PropertyNameCacheIdentifier,
		PropertyNameCacheUsername,
		PropertyNameCachePassword,
		PropertyNameCachePasswordFile,
		PropertyNameCachePort,
		PropertyNameDatabaseHost,
		PropertyNameDatabaseName,
		PropertyNameDatabasePassword,
		PropertyNameDatabasePasswordFile,
		PropertyNameDatabasePort,
		PropertyNameDatabaseUsername,
		PropertyNameDatabaseSSLMode,
		PropertyNameDatabaseTimezone,
		PropertyNameDebugMode,
		PropertyNameEnvironment,
		PropertyNameFeatureFlagSDKKey,
		PropertyNameFeatureFlagSDKKeyFile,
		PropertyNameGRPCPort,
		PropertyNameHTTPPort,
		PropertyNameLoadEnvFromFile,
		PropertyNameMailHost,
		PropertyNameMailPassword,
		PropertyNameMailPasswordFile,
		PropertyNameMailPort,
		PropertyNameMailSenderAddress,
		PropertyNameMailSenderName,
		PropertyNameMailUsername,
		PropertyNameServiceName,
	}
}
