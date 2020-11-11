package configuration

//ICredentialsProvider provider for credentials
type ICredentialsProvider interface {
	ProvideCredentials() (*ServiceCredentials, error)
}
