package frp

import "github.com/fatedier/frp/cmd/frpc/sub"

func RunFrpc(cfgDir, cfgFile string) {
	sub.Init(cfgDir, cfgFile)
	sub.Execute()
}

func StopFrpc() {
	sub.Stop()
}
