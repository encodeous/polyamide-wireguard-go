//go:build !linux

package device

import (
	"github.com/encodeous/polyamide-wireguard-go/conn"
	"github.com/encodeous/polyamide-wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(_ conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
