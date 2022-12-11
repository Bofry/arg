package arg

import (
	"fmt"
	"net"

	"github.com/Bofry/arg/internal"
)

func (fn IPValidator) AssertString(v string, name string) error {
	return fn.AssertValue(v, name)
}

func (fn IPValidator) AssertValue(v interface{}, name string) error {
	var ip IP

	switch v.(type) {
	case IP:
		{
			ip = v.(IP)
		}
	case []byte:
		{
			ip = IP(v.([]byte))
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
