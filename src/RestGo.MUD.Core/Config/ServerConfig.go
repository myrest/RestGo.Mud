package Config

var ServiceConfig serviceConfig

type serviceConfig struct {
	ListenOnPortNumber int `json:"ListenOnPortNumber"`
	Firebase           struct {
		ProjectID      string `json:"ProjectID"`
		CredentialFile string `json:"CredentialFile"`
	} `json:"Firebase"`
	MaxLength int `json:"MaxLength"`
}
