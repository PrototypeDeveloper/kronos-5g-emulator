package engine

import "net/netip"

// TUNDevice is a TUN device as the engine uses it.
type TUNDevice = tunDevice

// The tests never make host TUN devices, even with CAP_NET_ADMIN.
func init() { hasNetAdmin = func() bool { return false } }

// UseTUN gives the subscribers of e devices from open, as CAP_NET_ADMIN does
// with host TUN devices. Call it before linkup.
func UseTUN(e *Engine, open func(index int, ueIP netip.Addr) (TUNDevice, error)) { e.openTUN = open }
