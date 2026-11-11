package gnb

import (
	"encoding/hex"
	"fmt"
	"sort"

	"kronos/models"
)

// Status implements role.UE: it fills st with the UE's registration,
// connection, NAS security and PDU session state.
func (u *UE) Status(st *models.UEStatus) {
	// Registration and connection management states, with the NGAP IDs.
	st.RM, st.CM = models.RMDeregistered, models.CMIdle
	if u.registered {
		st.RM = models.RMRegistered
	}
	if u.hasRAN {
		st.CM = models.CMConnected
		id := u.ranID
		st.RANID = &id
	}
	if u.hasAMF {
		id := u.amfID
		st.AMFID = &id
	}
	if u.guti != nil {
		st.GUTI = hex.EncodeToString(u.guti)
	}
	// The NAS security context, once there is one.
	if u.sec != nil {
		st.Security = &models.NASSecurity{
			NgKSI:     u.ngKSI & ksiMask,
			Integrity: fmt.Sprintf("NIA%d", u.sec.IntegrityAlgorithm),
			Ciphering: fmt.Sprintf("NEA%d", u.sec.CipheringAlgorithm),
			ULCount:   uint32(u.sec.UplinkCount),
			DLCount:   uint32(u.sec.DownlinkCount),
		}
	}
	// The PDU sessions, by ID.
	ids := make([]int, 0, len(u.sessions))
	for id := range u.sessions {
		ids = append(ids, int(id))
	}
	sort.Ints(ids)
	for _, id := range ids {
		s := u.sessions[uint8(id)]
		ss := models.SessionStatus{ID: uint8(id), ULTEID: s.ulTEID, DLTEID: s.dlTEID}
		for _, q := range s.qfis {
			ss.QFIs = append(ss.QFIs, int(q))
		}
		if s.ueIP.IsValid() {
			ss.UEIP = s.ueIP.String()
		}
		if s.upf.IsValid() {
			ss.UPF = s.upf.String()
		}
		st.Sessions = append(st.Sessions, ss)
	}
}
