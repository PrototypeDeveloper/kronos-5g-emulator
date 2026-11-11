package scenario

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"sync"

	libnas "libnas"
	libngap "libngap"

	"kronos/models"
)

// Creating signals in the editor: an NGAP message of a type, or a NAS
// message for a NAS-PDU, with its mandatory fields at their smallest
// values. The user edits them from there.

// NewSignal returns the octets of a new NGAP message of the named type.
func NewSignal(name string) ([]byte, error) {
	info, ok := LookupNGAP(name)
	if !ok {
		return nil, fmt.Errorf("there is no NGAP message %q", name)
	}
	v := reflect.New(info.Type)
	if err := fill(v.Elem(), reflect.StructField{}); err != nil {
		return nil, fmt.Errorf("cannot make a %s: %v", info.Name, err)
	}
	b, err := libngap.Encode(v.Interface())
	if err != nil {
		return nil, fmt.Errorf("cannot make a %s: %v", info.Name, err)
	}
	return b, nil
}

// nasLengthErr reads the length libnas wants in an encoding error.
var nasLengthErr = regexp.MustCompile(`outside (\d+)\.\.|expected (\d+) bytes, got 0`)

// NewNAS returns the octets of a new plain NAS message of the named type.
// libnas does not take an empty mandatory IE, so each one it refuses gets
// zero octets of the smallest length it asks for.
func NewNAS(name string) ([]byte, error) {
	var typ reflect.Type
	for _, m := range nasMessageTypes {
		if t := reflect.TypeOf(m).Elem(); t.Name() == name {
			typ = t
		}
	}
	if typ == nil {
		return nil, fmt.Errorf("libnas cannot make the NAS message %q", name)
	}
	v := reflect.New(typ)
	for try := 0; try <= typ.NumField(); try++ {
		b, err := libnas.Encode(v.Interface(), nil)
		if err == nil {
			return b, nil
		}
		m := nasLengthErr.FindStringSubmatch(err.Error())
		if m == nil {
			return nil, fmt.Errorf("cannot make a %s: %v", name, err)
		}
		n, _ := strconv.Atoi(m[1] + m[2])
		if !fillEmptyMandatory(v.Elem(), n) {
			return nil, fmt.Errorf("cannot make a %s: %v", name, err)
		}
	}
	return nil, errors.New("cannot make a " + name)
}

// fillEmptyMandatory gives the first empty mandatory octets of a NAS
// message (direct, or the Value of an IE type) n zero octets.
func fillEmptyMandatory(v reflect.Value, n int) bool {
	typ := v.Type()
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if !f.IsExported() || isHasFlag(typ, f) {
			continue
		}
		if _, optional := typ.FieldByName("Has" + f.Name); optional {
			continue
		}
		fv := v.Field(i)
		if fv.Kind() == reflect.Struct && fv.NumField() == 1 {
			fv = fv.Field(0)
		}
		if fv.Kind() == reflect.Slice && fv.Type().Elem().Kind() == reflect.Uint8 && fv.Len() == 0 {
			fv.SetBytes(make([]byte, n))
			return true
		}
	}
	return false
}

// The lists of SignalTypes, made once.
var (
	typesOnce           sync.Once
	ngapTypes, gmm, gsm []string
)

// SignalTypes lists what the editor can create: the NGAP messages, and the
// 5GMM and 5GSM messages for NAS-PDUs and NAS containers.
func SignalTypes() (ngap, gmmNames, gsmNames []string) {
	typesOnce.Do(func() {
		for _, m := range ngapMessageTypes {
			name := reflect.TypeOf(m).Elem().Name()
			if _, err := NewSignal(name); err == nil {
				ngapTypes = append(ngapTypes, name)
			}
		}
		for _, m := range nasMessageTypes {
			name := reflect.TypeOf(m).Elem().Name()
			info, ok := LookupNAS(name)
			if _, err := NewNAS(name); err != nil || !ok {
				continue
			}
			if info.EPD == models.EPD5GSM {
				gsm = append(gsm, name)
			} else {
				gmm = append(gmm, name)
			}
		}
		sort.Strings(ngapTypes)
		sort.Strings(gmm)
		sort.Strings(gsm)
	})
	return ngapTypes, gmm, gsm
}
