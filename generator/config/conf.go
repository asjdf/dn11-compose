package config

type Config struct {
	//这部分信息是不需要公开但要保存的
	keepInMind
	//这部分信息是对方需要告诉你的
	needToKnowFromPeer
	//这些部分是你需要告诉对方的
	needToTellPeer
}

type keepInMind struct {
	OwnTunnelPrivateKey string
	OwnIP               string
	OwnNet              string
	OwnNets             []string
	PeerAlias           string
}

type needToKnowFromPeer struct {
	OwnIPAtPeer    string
	PeerIPInTunnel string
	PeerEndPoint   string
	PeerPublicKey  string
	PeerAS         string
}

type needToTellPeer struct {
	OwnTunnelListenPort int
	OwnAS               string
}
