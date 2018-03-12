package cfip

import (
	"net"

	"github.com/golang/glog"
	"github.com/pkg/errors"
)

func init() {
	cidrs := []string{
		"103.21.244.0/22",
		"103.22.200.0/22",
		"103.31.4.0/22",
		"104.16.0.0/12",
		"108.162.192.0/18",
		"131.0.72.0/22",
		"141.101.64.0/18",
		"162.158.0.0/15",
		"172.64.0.0/13",
		"173.245.48.0/20",
		"188.114.96.0/20",
		"190.93.240.0/20",
		"197.234.240.0/22",
		"198.41.128.0/17",
		"2400:cb00::/32",
		"2405:8100::/32",
		"2405:b500::/32",
		"2606:4700::/32",
		"2803:f800::/32",
		"2c0f:f248::/32",
		"2a06:98c0::/29",
	}

	nets = []*net.IPNet{}
	for _, v := range cidrs {
		_, n, e := net.ParseCIDR(v)
		if e != nil {
			glog.Errorf("%+v", e)
			continue
		}
		nets = append(nets, n)
	}
}

var (
	nets []*net.IPNet
)

func Contains(ip string) (bool, error) {
	if len(ip) == 0 {
		return false, errors.New("Empty ip.")
	}

	for _, v := range nets {
		if v.Contains(net.ParseIP(ip)) {
			return true, nil
		}
	}

	return false, nil
}
