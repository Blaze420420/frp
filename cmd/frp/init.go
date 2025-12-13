package frp

import (
	_ "github.com/fatedier/frp/assets/frpc"
	_ "github.com/fatedier/frp/assets/frps"
	_ "github.com/fatedier/frp/client"
	_ "github.com/fatedier/frp/client/proxy"
	_ "github.com/fatedier/frp/pkg/config"
	_ "github.com/fatedier/frp/pkg/config/legacy"
	_ "github.com/fatedier/frp/pkg/metrics/aggregate"
	_ "github.com/fatedier/frp/pkg/metrics/mem"
	_ "github.com/fatedier/frp/pkg/msg"
	_ "github.com/fatedier/frp/pkg/plugin/client"
	_ "github.com/fatedier/frp/pkg/util/log"

	_ "github.com/fatedier/frp/server"
	_ "github.com/fatedier/frp/server/proxy"
)
