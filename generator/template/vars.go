package template

type ConfigVars struct {
	TunnelType string
	Tunnel     TemplateFileVars
}

type TemplateFileVars interface {
	//Generate() string
}

type WGTunnelTemplateFileVars struct {
	OwnPrivateKey string
	OwnListenPort int
	OwnIPAtPeer   string
	PeerIP        string
	PeerEndPoint  string
	PeerPublicKey string
}

type BirdVariableTemplateFileVars struct {
	OwnAS        string
	OwnIP        string
	OwnNet       string
	OwnNetSetStr string
}

type BirdEGBPTemplateFileVars struct {
	PeerAlias string
	PeerIP    string
	PeerAS    string
}

type DockerComposeYamlTemplateFileVars struct {
	OwnIP         string
	OwnListenPort int
	OwnNet        string
}
