package arg

import (
	"net"
	"testing"
)

func TestIPAssertion(t *testing.T) {
	var (
		invalidIP   net.IP = []byte{0}
		ip128_0_0_1 net.IP = net.ParseIP("128.0.0.1")
	)

	{
		err := _IPAssertion.Assert(invalidIP, "invalidIP",
			_IPAssertion.IsValid,
			_IPAssertion.Loopback,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"invalidIP\"; specified ip address ?00 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		err := _IPAssertion.Assert(ip128_0_0_1, "ip128_0_0_1",
			_IPAssertion.IsValid,
			_IPAssertion.Loopback,
		)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"ip128_0_0_1\"; specified ip address 128.0.0.1 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertor(t *testing.T) {
	var (
		invalidIP net.IP = []byte{0}
	)

	{
		err := _IPAssertion.Assertor(invalidIP, "invalidIP").
			Assert(
				_IPAssertion.IsValid,
				_IPAssertion.Loopback,
			)
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"invalidIP\"; specified ip address ?00 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}

}

func TestIPAssertion_IsValid(t *testing.T) {
	{
		var arg net.IP = []byte{1, 1, 1, 1}
		err := _IPAssertion.IsValid(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = []byte{3, 5, 6, 7, 8, 9}
		err := _IPAssertion.IsValid(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address ?030506070809 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_GlobalUnicast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("2000::")
		err := _IPAssertion.GlobalUnicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("FF00::")
		err := _IPAssertion.GlobalUnicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address ff00:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("10.255.0.0")
		err := _IPAssertion.GlobalUnicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("8.8.8.8")
		err := _IPAssertion.GlobalUnicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("255.255.255.255")
		err := _IPAssertion.GlobalUnicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 255.255.255.255 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_InterfaceLocalMulticast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("ff01::1")
		err := _IPAssertion.InterfaceLocalMulticast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("2000::")
		err := _IPAssertion.InterfaceLocalMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 2000:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("255.0.0.0")
		err := _IPAssertion.InterfaceLocalMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 255.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_LinkLocalMulticast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("ff02::2")
		err := _IPAssertion.LinkLocalMulticast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("fe80::")
		err := _IPAssertion.LinkLocalMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address fe80:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("224.0.0.0")
		err := _IPAssertion.LinkLocalMulticast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("169.254.0.0")
		err := _IPAssertion.LinkLocalMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 169.254.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_LinkLocalUnicast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("fe80::")
		err := _IPAssertion.LinkLocalUnicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("2000::")
		err := _IPAssertion.LinkLocalUnicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 2000:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("169.254.0.0")
		err := _IPAssertion.LinkLocalUnicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("224.0.0.0")
		err := _IPAssertion.LinkLocalUnicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 224.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_Loopback(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("::1")
		err := _IPAssertion.Loopback(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("ff02::1")
		err := _IPAssertion.Loopback(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address ff02::1 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("127.0.0.0")
		err := _IPAssertion.Loopback(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("128.0.0.0")
		err := _IPAssertion.Loopback(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 128.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_Multicast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("FF00::")
		err := _IPAssertion.Multicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("ff02::1")
		err := _IPAssertion.Multicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("::1")
		err := _IPAssertion.Multicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address ::1 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("239.0.0.0")
		err := _IPAssertion.Multicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("224.0.0.0")
		err := _IPAssertion.Multicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("127.0.0.0")
		err := _IPAssertion.Multicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 127.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_Private(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("fc00::")
		err := _IPAssertion.Private(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("fe00::")
		err := _IPAssertion.Private(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address fe00:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("10.255.0.0")
		err := _IPAssertion.Private(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("11.0.0.0")
		err := _IPAssertion.Private(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 11.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_Unspecified(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("::")
		err := _IPAssertion.Unspecified(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("fe00::")
		err := _IPAssertion.Unspecified(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address fe00:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("0.0.0.0")
		err := _IPAssertion.Unspecified(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("8.8.8.8")
		err := _IPAssertion.Unspecified(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 8.8.8.8 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_NonGlobalUnicast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("2000::")
		err := _IPAssertion.NotGlobalUnicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 2000:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("FF00::")
		err := _IPAssertion.NotGlobalUnicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("10.255.0.0")
		err := _IPAssertion.NotGlobalUnicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 10.255.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("8.8.8.8")
		err := _IPAssertion.NotGlobalUnicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 8.8.8.8 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("255.255.255.255")
		err := _IPAssertion.NotGlobalUnicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIPAssertion_NonInterfaceLocalMulticast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("ff01::1")
		err := _IPAssertion.NotInterfaceLocalMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address ff01::1 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("2000::")
		err := _IPAssertion.NotInterfaceLocalMulticast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("255.0.0.0")
		err := _IPAssertion.NotInterfaceLocalMulticast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIPAssertion_NotLinkLocalMulticast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("ff02::2")
		err := _IPAssertion.NotLinkLocalMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address ff02::2 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("fe80::")
		err := _IPAssertion.NotLinkLocalMulticast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("224.0.0.0")
		err := _IPAssertion.NotLinkLocalMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 224.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("169.254.0.0")
		err := _IPAssertion.NotLinkLocalMulticast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIPAssertion_NotLinkLocalUnicast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("fe80::")
		err := _IPAssertion.NotLinkLocalUnicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address fe80:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("2000::")
		err := _IPAssertion.NotLinkLocalUnicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("169.254.0.0")
		err := _IPAssertion.NotLinkLocalUnicast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 169.254.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("224.0.0.0")
		err := _IPAssertion.NotLinkLocalUnicast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIPAssertion_NotLoopback(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("::1")
		err := _IPAssertion.NotLoopback(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address ::1 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("ff02::1")
		err := _IPAssertion.NotLoopback(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("127.0.0.0")
		err := _IPAssertion.NotLoopback(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 127.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("128.0.0.0")
		err := _IPAssertion.NotLoopback(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIPAssertion_NotMulticast(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("FF00::")
		err := _IPAssertion.NotMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address ff00:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("ff02::1")
		err := _IPAssertion.NotMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address ff02::1 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("::1")
		err := _IPAssertion.NotMulticast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("239.0.0.0")
		err := _IPAssertion.NotMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 239.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("224.0.0.0")
		err := _IPAssertion.NotMulticast(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 224.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("127.0.0.0")
		err := _IPAssertion.NotMulticast(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIPAssertion_NotPrivate(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("fc00::")
		err := _IPAssertion.NotPrivate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address fc00:: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("fe00::")
		err := _IPAssertion.NotPrivate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("10.255.0.0")
		err := _IPAssertion.NotPrivate(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 10.255.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("11.0.0.0")
		err := _IPAssertion.NotPrivate(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIPAssertion_NotUnspecified(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("::")
		err := _IPAssertion.NotUnspecified(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address :: is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("fe00::")
		err := _IPAssertion.NotUnspecified(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("0.0.0.0")
		err := _IPAssertion.NotUnspecified(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 0.0.0.0 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("8.8.8.8")
		err := _IPAssertion.NotUnspecified(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}

func TestIPAssertion_BelongToAny(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("10.10.20.3")
		err := _IPAssertion.BelongToAny("10.10.20.0/30").Assert(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("10.10.20.5")
		err := _IPAssertion.BelongToAny("10.10.20.0/30").Assert(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 10.10.20.5 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("1::1")
		err := _IPAssertion.BelongToAny("1::/64").Assert(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("2::1")
		err := _IPAssertion.BelongToAny("1::/64").Assert(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 2::1 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
}

func TestIPAssertion_NotBelongToAny(t *testing.T) {
	{
		var arg net.IP = net.ParseIP("10.10.20.3")
		err := _IPAssertion.NotBelongToAny("10.10.20.0/30").Assert(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 10.10.20.3 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("10.10.20.5")
		err := _IPAssertion.NotBelongToAny("10.10.20.0/30").Assert(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
	{
		var arg net.IP = net.ParseIP("1::1")
		err := _IPAssertion.NotBelongToAny("1::/64").Assert(arg, "arg")
		if err == nil {
			t.Errorf("should get error")
		}
		exceptedErrorMsg := "invalid argument \"arg\"; specified ip address 1::1 is invalid"
		if err.Error() != exceptedErrorMsg {
			t.Errorf("except: %v\ngot: %v", exceptedErrorMsg, err.Error())
		}
	}
	{
		var arg net.IP = net.ParseIP("2::1")
		err := _IPAssertion.NotBelongToAny("1::/64").Assert(arg, "arg")
		if err != nil {
			t.Errorf("should not error")
		}
	}
}
