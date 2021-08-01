package simplenet

type IP []byte

const (
	IPv4len = 4
	IPv6len = 16
	big     = 0xFFFFFF
)

var v4InV6Prefix = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff}

func ParseIP(s string) IP {
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return parseIPv4(s)
		}
	}
	return nil
}

func parseIPv4(s string) IP {
	var p [IPv4len]byte
	for i := 0; i < IPv4len; i++ {
		if len(s) == 0 {
			return nil
		}

		if i > 0 {
			if s[0] != '.' {
				return nil
			}
			s = s[1:]
		}

		n, c, ok := dtoi(s)

		if !ok || n > 0xFF {
			return nil
		}

		s = s[c:]

		p[i] = byte(n)
	}

	if len(s) != 0 {
		return nil
	}
	return IPv4(p[0], p[1], p[2], p[3])
}

func IPv4(a, b, c, d byte) IP {
	p := make(IP, IPv6len)
	copy(p, v4InV6Prefix)
	p[12] = a
	p[13] = b
	p[14] = c
	p[15] = d
	return p
}
func dtoi(s string) (n int, i int, ok bool) {
	n = 0
	for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
		n = n*10 + int(s[i]-'0')
		if n >= big {
			return big, i, false
		}
	}
	if i == 0 {
		return 0, 0, false
	}
	return n, i, true
}
