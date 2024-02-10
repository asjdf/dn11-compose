package main

import (
	"fmt"
	"generator/config"
	generatorTemplate "generator/template"
	"generator/wg"
	"github.com/pelletier/go-toml/v2"
	"io"
	"os"
	"strings"
	"text/template"
)

func main() {
	configData, err := os.ReadFile("config.toml")
	if err != nil {
		panic(err)
	}
	var conf = new(config.Config)
	if err := toml.Unmarshal(configData, conf); err != nil {
		panic(err)
	}
	if conf.OwnTunnelPrivateKey == "" {
		var wgKeyPair = wg.GenerateSelfKeyPair()
		conf.OwnTunnelPrivateKey = wgKeyPair.PrivateKeyB64
		fmt.Println("private key is empty, use generated key:", wgKeyPair.PrivateKeyB64)
		fmt.Println("tell peer to use this public key:", wgKeyPair.PublicKeyB64)
	}
	fmt.Println()
	printFileNameBanner(fmt.Sprintf("peer/wg-peers/%s.conf", conf.PeerAlias))
	var wgTunnelVars = &generatorTemplate.WGTunnelTemplateFileVars{
		OwnPrivateKey: conf.OwnTunnelPrivateKey,
		OwnListenPort: conf.OwnTunnelListenPort,
		OwnIPAtPeer:   conf.OwnIPAtPeer,
		PeerIP:        conf.PeerIPInTunnel,
		PeerEndPoint:  conf.PeerEndPoint,
		PeerPublicKey: conf.PeerPublicKey,
	}
	readTemplateFileAndExecute("wg-tunnel", "../template/peer/wg-peers/wg-template.conf.temp",
		wgTunnelVars,
		os.Stdout)
	fmt.Println()
	printFileNameBanner("peer/bird/variable.conf")
	var birdVariableVars = &generatorTemplate.BirdVariableTemplateFileVars{
		OwnAS:        conf.OwnAS,
		OwnIP:        conf.OwnIP,
		OwnNet:       conf.OwnNet,
		OwnNetSetStr: strings.Join(conf.OwnNets, ","),
	}
	readTemplateFileAndExecute("bird-vars", "../template/peer/bird/variable.conf.temp",
		birdVariableVars,
		os.Stdout)
	fmt.Println()
	printFileNameBanner(fmt.Sprintf("peer/bird/include/%s.conf", conf.PeerAlias))
	var birdEGBPTemplateVars = &generatorTemplate.BirdEGBPTemplateFileVars{
		PeerAlias: conf.PeerAlias,
		PeerIP:    conf.PeerIPInTunnel,
		PeerAS:    conf.PeerAS,
	}
	readTemplateFileAndExecute("bird-include", "../template/peer/bird/include/dn11-ebgp-template.conf.temp",
		birdEGBPTemplateVars,
		os.Stdout)
	fmt.Println()
	printFileNameBanner("docker-compose.yaml")
	var dockerComposeYamlVars = &generatorTemplate.DockerComposeYamlTemplateFileVars{
		OwnIP:         conf.OwnIP,
		OwnListenPort: conf.OwnTunnelListenPort,
		OwnNet:        conf.OwnNet,
	}
	readTemplateFileAndExecute("docker-compose-yml", "../template/docker-compose.yml.temp",
		dockerComposeYamlVars,
		os.Stdout)
}

func printFileNameBanner(fileName string) {
	fmt.Println()
	fmt.Printf("\n%s\n", fileName)
}

func readTemplateFileAndExecute(templateName string, templateFilePath string,
	vars any,
	out io.Writer) {
	templateData, err := os.ReadFile(templateFilePath)
	if err != nil {
		panic(err)
	}
	wgTunnelTemplate, err := template.New(templateName).Parse(string(templateData))
	if err != nil {
		panic(err)
	}
	if err := wgTunnelTemplate.Execute(out, vars); err != nil {
		panic(err)
	}
}
