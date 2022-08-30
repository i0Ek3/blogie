package version

const (
	Version = "1.0.0"
	AppName = "blogie"
	Address = "127.0.0.1:6831"
)

func GetVersion() string {
	return Version
}

func GetAppName() string {
	return AppName
}

func GetAddress() string {
	return Address
}
