package enum

type Scheme uint

const (
	HTTP Scheme = iota
	HTTPS
	SOCKS5
)

func (p Scheme) String() string {
	return []string{"http", "https", "socks5"}[p]
}

func SchemeList() []string {
	return []string{"http", "https", "socks5"}
}

type Protocol uint

const (
	PROTOCOL_UDP Protocol = iota
	PROTOCOL_TCP
)

func (p Protocol) String() string {
	return []string{"udp", "tcp"}[p]
}

func ProtocolList() []string {
	return []string{"udp", "tcp"}
}
