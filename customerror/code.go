package customerror

// Code is type for error code.
type Code uint32

// UInt32 converts code to uint32.
func (c Code) UInt32() uint32 {
	return uint32(c)
}

const (
	// DefaultErrorCode is DefaultError code.
	DefaultErrorCode Code = 5000
	// PostgreSQLErrorCode is PostgreSQLError code.
	PostgreSQLErrorCode Code = 5001
	// RedisErrorCode is RedisError code.
	RedisErrorCode Code = 5002
	// ToolkitErrorCode is ToolkitError code.
	ToolkitErrorCode Code = 5003
	// CryptoGenerationErrorCode is error code for generation crypto.
	CryptoGenerationErrorCode Code = 5004
	// CryptoEncryptErrorCode is encrypt error code.
	CryptoEncryptErrorCode Code = 5005
	// CryptoDecryptErrorCode is decrypt error code.
	CryptoDecryptErrorCode Code = 5006
	// DefaultCryptoErrorCode is default crypto error code.
	DefaultCryptoErrorCode Code = 5007
	// RSAKeyFilesNotFoundErrorCode is RSAKeyFilesNotFoundError code
	RSAKeyFilesNotFoundErrorCode Code = 5008
	// NotImplementedErrorCode is NotImplementedError code.
	NotImplementedErrorCode Code = 5011
	// OAuthManagementAPIServiceUnavailableErrorCode is OAuthManagementAPIServiceUnavailableError code.
	OAuthManagementAPIServiceUnavailableErrorCode Code = 5031
	// UserManagementAPIServiceUnavailableErrorCode is UserManagementAPIServiceUnavailableError code.
	UserManagementAPIServiceUnavailableErrorCode Code = 5032
	// ColangAPIServiceUnavailableErrorCode is ColangAPIServiceUnavailableError code.
	ColangAPIServiceUnavailableErrorCode Code = 5033
	// WorldAPIServiceUnavailableErrorCode is WorldAPIServiceUnavailableError code.
	WorldAPIServiceUnavailableErrorCode Code = 5034
	// AdminAPIServiceUnavailableErrorCode is AdminAPIServiceUnavailableError code.
	AdminAPIServiceUnavailableErrorCode Code = 5035
	// SecretsManagementAPIServiceUnavailableErrorCode is SecretsManagementAPIServiceUnavailableError code.
	SecretsManagementAPIServiceUnavailableErrorCode Code = 5036
	// RequestValidationFailureCode is RequestValidationFailure code.
	RequestValidationFailureCode Code = 4001
	// ResponseValidationFailureCode is ResponseValidationFailure code.
	ResponseValidationFailureCode Code = 4002
	// OAuthManagementAPIBadRequestCode is OAuthManagementAPIBadRequest code.
	OAuthManagementAPIBadRequestCode Code = 4003
	// UserManagementAPIBadRequestCode is UserManagementAPIBadRequest code.
	UserManagementAPIBadRequestCode Code = 4004
	// ColangAPIServiceBadRequestCode is ColangAPIServiceBadRequest code.
	ColangAPIServiceBadRequestCode Code = 4005
	// WorldAPIServiceBadRequestCode is WorldAPIServiceBadRequest code.
	WorldAPIServiceBadRequestCode Code = 4006
	// AdminAPIServiceBadRequestCode is AdminAPIServiceBadRequest code.
	AdminAPIServiceBadRequestCode Code = 4007
	// SecretsManagementAPBadRequestCode is SecretsManagementAPBadRequest code.
	SecretsManagementAPBadRequestCode Code = 4008
	// TokenIsInvalidFailureCode is TokenIsInvalidFailure code.
	TokenIsInvalidFailureCode Code = 4011
	// TokenIsExpiredFailureCode is TokenIsExpiredFailure code.
	TokenIsExpiredFailureCode Code = 4012
	// IncorrectTokenFormatFailureCode is IncorrectTokenFormat code.
	IncorrectTokenFormatFailureCode Code = 4013
	// RoleForbiddenFailureCode is RoleForbiddenFailure code.
	RoleForbiddenFailureCode Code = 4031
	// EntityNotFoundFailureCode is EntityNotFoundFailure code.
	EntityNotFoundFailureCode Code = 4041
	// HTTPMethodNotAllowedFailureCode is HTTPMethodNotAllowedFailure code.
	HTTPMethodNotAllowedFailureCode Code = 4051
	// EntityAlreadyExistsFailureCode is EntityAlreadyExistsFailure code.
	EntityAlreadyExistsFailureCode Code = 4091
	// ExtractIPAddressFromRequestFailureCode is ExtractIPAddressFromRequestFailure code.
	ExtractIPAddressFromRequestFailureCode Code = 4003
)
