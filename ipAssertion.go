package arg

import (
	"fmt"
	"net"

	"github.com/Bofry/arg/internal"
)

var (
	_IPAssertion = IPAssertion("")

	_ IPValidator = _IPAssertion.IsValid
	_ IPValidator = _IPAssertion.GlobalUnicast
	_ IPValidator = _IPAssertion.InterfaceLocalMulticast
	_ IPValidator = _IPAssertion.LinkLocalMulticast
	_ IPValidator = _IPAssertion.LinkLocalUnicast
	_ IPValidator = _IPAssertion.Loopback
	_ IPValidator = _IPAssertion.Multicast
	_ IPValidator = _IPAssertion.Private
	_ IPValidator = _IPAssertion.Unspecified
	_ IPValidator = _IPAssertion.NotGlobalUnicast
	_ IPValidator = _IPAssertion.NotInterfaceLocalMulticast
	_ IPValidator = _IPAssertion.NotLinkLocalMulticast
	_ IPValidator = _IPAssertion.NotLinkLocalUnicast
	_ IPValidator = _IPAssertion.NotLoopback
	_ IPValidator = _IPAssertion.NotMulticast
	_ IPValidator = _IPAssertion.NotPrivate
	_ IPValidator = _IPAssertion.NotUnspecified
)

type IPAssertion string

func (IPAssertion) Assert(v net.IP, name string, validators ...IPValidator) error {
	for _, validate := range validators {
		if err := validate(v, name); err != nil {
			return err
		}
	}
	return nil
}

func (IPAssertion) Assertor(v net.IP, name string) *IPAssertor {
	return &IPAssertor{v, name}
}

func (IPAssertion) IsValid(v net.IP, name string) error {
	if v.To16() == nil {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

// GlobalUnicast checks whether ip is a global unicast address.
//
// The identification of global unicast addresses uses address type
// identification as defined in RFC 1122, RFC 4632 and RFC 4291 with
// the exception of IPv4 directed broadcast addresses.
// It returns true even if ip is in IPv4 private address space or
// local IPv6 unicast address space.
func (IPAssertion) GlobalUnicast(v net.IP, name string) error {
	if !v.IsGlobalUnicast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

// InterfaceLocalMulticast checks whether ip is an interface-local multicast address.
func (IPAssertion) InterfaceLocalMulticast(v net.IP, name string) error {
	if !v.IsInterfaceLocalMulticast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

// LinkLocalMulticast checks whether ip is a link-local multicast address.
func (IPAssertion) LinkLocalMulticast(v net.IP, name string) error {
	if !v.IsLinkLocalMulticast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

// LinkLocalUnicast checks whether ip is a link-local unicast address.
func (IPAssertion) LinkLocalUnicast(v net.IP, name string) error {
	if !v.IsLinkLocalUnicast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

// Loopback checks whether ip is a loopback address.
func (IPAssertion) Loopback(v net.IP, name string) error {
	if !v.IsLoopback() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

// Multicast checks whether ip is a multicast address.
func (IPAssertion) Multicast(v net.IP, name string) error {
	if !v.IsMulticast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

// Private checks whether ip is a private address, according to RFC 1918 (IPv4 addresses) and RFC 4193 (IPv6 addresses).
func (IPAssertion) Private(v net.IP, name string) error {
	if !v.IsPrivate() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

// Unspecified checks whether ip is an unspecified address, either the IPv4 address "0.0.0.0" or the IPv6 address "::".
func (IPAssertion) Unspecified(v net.IP, name string) error {
	if !v.IsUnspecified() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

func (IPAssertion) NotGlobalUnicast(v net.IP, name string) error {
	if v.IsGlobalUnicast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

func (IPAssertion) NotInterfaceLocalMulticast(v net.IP, name string) error {
	if v.IsInterfaceLocalMulticast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

func (IPAssertion) NotLinkLocalMulticast(v net.IP, name string) error {
	if v.IsLinkLocalMulticast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

func (IPAssertion) NotLinkLocalUnicast(v net.IP, name string) error {
	if v.IsLinkLocalUnicast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

func (IPAssertion) NotLoopback(v net.IP, name string) error {
	if v.IsLoopback() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

func (IPAssertion) NotMulticast(v net.IP, name string) error {
	if v.IsMulticast() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

func (IPAssertion) NotPrivate(v net.IP, name string) error {
	if v.IsPrivate() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

func (IPAssertion) NotUnspecified(v net.IP, name string) error {
	if v.IsUnspecified() {
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
	return nil
}

func (IPAssertion) Must(fn IPPredicate) IPValidator {
	return func(v net.IP, name string) error {
		if !fn(v) {
			return &InvalidArgumentError{
				Name:   name,
				Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
			}
		}
		return nil
	}
}

// BelongTo checks whether the ip is belong to the specified networks.
func (IPAssertion) BelongToAny(cidrs ...string) IPValidator {
	var ipnets []*net.IPNet
	for _, cidr := range cidrs {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(
				fmt.Sprintf("specified CIDR address %q is invalid", cidr),
			)
		}
		ipnets = append(ipnets, ipnet)
	}
	return func(v net.IP, name string) error {
		for _, ipnet := range ipnets {
			if ipnet.Contains(v) {
				return nil
			}
		}
		return &InvalidArgumentError{
			Name:   name,
			Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
		}
	}
}

func (IPAssertion) NotBelongToAny(cidrs ...string) IPValidator {
	var ipnets []*net.IPNet
	for _, cidr := range cidrs {
		_, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(
				fmt.Sprintf("specified CIDR address %q is invalid", cidr),
			)
		}
		ipnets = append(ipnets, ipnet)
	}
	return func(v net.IP, name string) error {
		for _, ipnet := range ipnets {
			if ipnet.Contains(v) {
				return &InvalidArgumentError{
					Name:   name,
					Reason: fmt.Sprintf(internal.ERR_INVALID_IP_ADDR, v),
				}
			}
		}
		return nil
	}
}
