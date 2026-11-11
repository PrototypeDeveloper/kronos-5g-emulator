package models

import "strings"

// MapItem is one item of the mapping (the auto-fix of sent signals).
type MapItem uint8

// MapOff is the set of mapping items that are off; the zero value has all on.
type MapOff uint32

// mapNames are the names of the items in mapon / mapoff / mapget.
var mapNames = [mapItemCount]string{
	MapPLMN: "plmn", MapTAC: "tac", MapNCI: "nci", MapGNodeBID: "gnodebid", MapRANName: "ranname",
	MapSNSSAI: "snssai", MapTimeStamp: "timestamp", MapNGAPID: "ngapid", MapIdentity: "identity",
	MapNgKSI: "ngksi", MapRES: "res", MapSecCap: "seccap", MapIMEISV: "imeisv", MapSession: "session",
	MapStatus: "status", MapTunnel: "tunnel", MapContainer: "container",
}

// String returns the name of the item.
func (i MapItem) String() string {
	if int(i) < len(mapNames) {
		return mapNames[i]
	}
	return "?"
}

// MapItems returns every item, in order.
func MapItems() []MapItem {
	out := make([]MapItem, mapItemCount)
	for i := range out {
		out[i] = MapItem(i)
	}
	return out
}

// MapNames returns the names of every item, in order.
func MapNames() []string { return append([]string(nil), mapNames[:]...) }

// MapItemByName returns the item named name (ignoring case).
func MapItemByName(name string) (MapItem, bool) {
	for i, n := range mapNames {
		if strings.EqualFold(n, name) {
			return MapItem(i), true
		}
	}
	return 0, false
}

// On reports whether the mapping of i is on.
func (o MapOff) On(i MapItem) bool { return o&(1<<i) == 0 }

// Of returns the set holding i.
func (i MapItem) Of() MapOff { return 1 << i }
