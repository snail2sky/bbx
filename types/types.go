package types

type CrtData struct {
	CertPath       string
	WarningDays    int
	TargetSuffixes string
	WebhookURL     string
}

type HTTPEchoData struct {
	Host string
	Port int
}
