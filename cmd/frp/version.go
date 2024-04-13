package frp

import "github.com/fatedier/frp/pkg/util/version"

func Version() string {
	return version.Full()
}
