package main

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/golang/glog"

	"github.com/gangwn/dvc/internal/pkg/ccs/config"
	"github.com/gangwn/dvc/pkg/util"

		"github.com/gangwn/dvc/internal/pkg/ccs/controler"
	"github.com/gangwn/dvc/pkg/net/basic"
)

type options struct {
	cfgPath *string
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func main() {
	flag.Set("logtostderr", "true")

	// Read arguments from command line
	opt := options{nil}

	opt.cfgPath = flag.String("config", path.Join(getCurrentDirectory(), "dvc.json"), "Configuration Path")
	flag.Parse()

	glog.Infof("Config path is %v", *opt.cfgPath)

	cfg, err := config.ReadConfig(*opt.cfgPath)
	if err != nil {
		glog.Errorf("Can't read the config: %v, err: %v", *opt.cfgPath, err)
		return
	}

	glog.Infof("Listen address is %v", cfg.Net.LisAddrs)

	addr := util.NewMultiaddr(cfg.Net.LisAddrs)
	_, priv, err := util.LoadKeyPair(cfg.Key.Dir)
	if err != nil {
		return
	}

	node := basic.NewBasicNode(addr, priv)
	if node == nil {
		return
	}

	glog.Infof("Local address: /ip4/%s/tcp/%d/ipfs/%s", cfg.Net.LisAddrs[0].IP, cfg.Net.LisAddrs[0].Port, node.ID().Pretty())

	controler := controler.NewControler(node)
	if controler == nil {
		return
	}

	select {}
}
