package frp

import (
	"context"
	"fmt"
	"os"

	"github.com/fatedier/frp/pkg/config"
	v1 "github.com/fatedier/frp/pkg/config/v1"
	"github.com/fatedier/frp/pkg/config/v1/validation"
	"github.com/fatedier/frp/pkg/util/log"
	"github.com/fatedier/frp/server"
)

var (
	gService *server.Service

	serverCh = make(chan bool)
)

func RunFrps(cfgFile string, strictConfigMode bool) {
	var (
		svrCfg         *v1.ServerConfig
		isLegacyFormat bool
		err            error
	)
	if cfgFile != "" {
		svrCfg, isLegacyFormat, err = config.LoadServerConfig(cfgFile, strictConfigMode)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if isLegacyFormat {
			fmt.Printf("WARNING: ini format is deprecated and the support will be removed in the future, " +
				"please use yaml/json/toml format instead!\n")
		}
	} else {
		fmt.Println("no cfgFile")
		os.Exit(1)
	}

	warning, err := validation.ValidateServerConfig(svrCfg)
	if warning != nil {
		fmt.Printf("WARNING: %v\n", warning)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := runServer(svrCfg, cfgFile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func StopFrps() {
	if gService != nil {
		gService.Close()
		gService = nil

		close(serverCh)
	}
}

func runServer(cfg *v1.ServerConfig, cfgFile string) (err error) {
	log.InitLogger(cfg.Log.To, cfg.Log.Level, int(cfg.Log.MaxDays), cfg.Log.DisablePrintColor)

	if cfgFile != "" {
		log.Infof("frps uses config file: %s", cfgFile)
	} else {
		log.Infof("frps uses command line arguments for config")
	}

	svr, err := server.NewService(cfg)
	if err != nil {
		return err
	}
	gService = svr
	log.Infof("frps started successfully")

	go svr.Run(context.Background())
	<-serverCh
	return
}
