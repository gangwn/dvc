package main

import (
	"flag"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/golang/glog"

	"github.com/gangwn/dvc/internal/pkg/config"
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

	glog.Infof("Config path is %v", cfg.Net.LisAddrs)

	// var maddrs []ma.Multiaddr
	// if *bindIPs != "" {
	// 	maddrs = make([]ma.Multiaddr, 0)
	// 	mas := strings.Split(*bindIPs, ",")
	// 	i := 0
	// 	for _, m := range mas {
	// 		addr, err := ma.NewMultiaddr(m)
	// 		if err != nil {
	// 			glog.Errorf("Error creating bindIP %v to multiaddr: %v", m, err)
	// 			continue // nonfatal
	// 		}
	// 		maddrs = append(maddrs, addr)
	// 		i++
	// 	}
}
