package version

const (
	Version = "1.0.0"
	AppName = "blogie"
)

func GetVersion() string {
	return Version
}

func GetAppName() string {
	return AppName
}
