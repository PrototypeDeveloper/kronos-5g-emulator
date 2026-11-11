package engine

import "net/netip"

// TUNDevice is a TUN device as the engine uses it.
type TUNDevice = tunDevice

// UseTUN gives the subscribers of e devices from open, as ue.tun does with
// host TUN devices. Call it before linkup.
func UseTUN(e *Engine, open func(index int, ueIP netip.Addr) (TUNDevice, error)) { e.openTUN = open }
