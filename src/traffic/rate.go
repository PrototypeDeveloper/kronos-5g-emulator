package traffic

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseRate reads a bit rate as iperf writes it: a number of bits per
// second with an optional K, M, G or T suffix (powers of 1000, either case),
// e.g. 100M, 1.5G, 500k.
func ParseRate(s string) (uint64, error) {
	t := strings.TrimSpace(s)
	mult := 1.0
	if n := len(t); n > 0 {
		switch t[n-1] {
		case 'k', 'K':
			mult = 1e3
		case 'm', 'M':
			mult = 1e6
		case 'g', 'G':
			mult = 1e9
		case 't', 'T':
			mult = 1e12
		}
		if mult != 1 {
			t = t[:n-1]
		}
	}
	v, err := strconv.ParseFloat(t, 64)
	if err != nil || v <= 0 || v*mult > maxParsedRate {
		return 0, fmt.Errorf("cannot read the rate %q (e.g. 100M, 1G, 500k)", s)
	}
	return uint64(v*mult + 0.5), nil
}

// FormatRate renders bits per second like "99.8 Mbit/s".
func FormatRate(bps float64) string {
	switch {
	case bps >= 1e9:
		return fmt.Sprintf("%.2f Gbit/s", bps/1e9)
	case bps >= 1e6:
		return fmt.Sprintf("%.1f Mbit/s", bps/1e6)
	case bps >= 1e3:
		return fmt.Sprintf("%.1f kbit/s", bps/1e3)
	}
	return fmt.Sprintf("%.0f bit/s", bps)
}

// FormatBytes renders a byte count like "11.9 MB" (powers of 1000).
func FormatBytes(n uint64) string {
	switch {
	case n >= 1e9:
		return fmt.Sprintf("%.2f GB", float64(n)/1e9)
	case n >= 1e6:
		return fmt.Sprintf("%.1f MB", float64(n)/1e6)
	case n >= 1e3:
		return fmt.Sprintf("%.1f kB", float64(n)/1e3)
	}
	return fmt.Sprintf("%d B", n)
}
