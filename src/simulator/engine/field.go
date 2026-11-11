package engine

import (
	"errors"

	"kronos/command"
	"kronos/scenario"
)

// newReceived keeps a received NGAP message and the plain text of its NAS-PDUs.
func newReceived(msg any, plain [][]byte) received {
	return received{ngap: msg, plain: plain}
}

// field reads an IE of the message for $ngap / $nas.
func (r *received) field(proto string, path []command.FieldSeg) (command.Value, error) {
	switch proto {
	case command.FieldNGAP:
		return scenario.FieldValue([]any{r.ngap}, path)
	case command.FieldNAS:
		if !r.done {
			var nas [][]byte
			for i, b := range scenario.NASPDUs(r.ngap) {
				if i < len(r.plain) && r.plain[i] != nil {
					b = r.plain[i]
				}
				nas = append(nas, b)
			}
			r.decoded, r.done = scenario.NASMessages(nas), true
		}
		return scenario.FieldValue(r.decoded, path)
	}
	return command.Value{}, errors.New("only $nas and $ngap name an IE")
}

// Field implements command.Env: the subscriber's last received message.
func (s *subscriber) Field(proto string, path []command.FieldSeg) (command.Value, error) {
	return s.rx.field(proto, path)
}

// Field implements command.Env: the node's last received message.
func (n *node) Field(proto string, path []command.FieldSeg) (command.Value, error) {
	return n.rx.field(proto, path)
}

// Field implements command.Env; the console receives nothing.
func (consoleEnv) Field(string, []command.FieldSeg) (command.Value, error) {
	return command.Value{}, errors.New("the console has received no message")
}
