package arg

import (
	"fmt"
	"net"

	"github.com/Bofry/arg/internal"
)

func (fn IPValidator) AssertString(v string, name string) error {
	return fn.Assert(v, name)
}

func (fn IPValidator) Assert(v interface{}, name string) error {
	var ip net.IP

	switch v.(type) {
	case net.IP:
		{
			ip = v.(net.IP)
		}
	case []byte:
		{
			ip = net.IP(v.([]byte))
		}
	case string:
		{
			str := v.(string)
			ip = net.ParseIP(str)
		}
	}

	// is valid ip?
	err := IPs.IsValid(ip, name)
	if err != nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_UNSUPPORTED_CAST_IP_ADDR, v),
		}
	}
	return fn(ip, name)
}
