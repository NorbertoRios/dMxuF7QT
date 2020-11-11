package configuration

//IReportConfigProvider represents provider
type IReportConfigProvider interface {
	Provide() ([]Field, error)
}
