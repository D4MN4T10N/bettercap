package modules

import (
	"fmt"
	"net"

	"github.com/bettercap/bettercap/log"
	"github.com/bettercap/bettercap/packets"
)

func (p *Prober) sendProbeWSD(from net.IP, from_hw net.HardwareAddr) {
	name := fmt.Sprintf("%s:%d", packets.WSDDestIP, packets.WSDPort)
	if addr, err := net.ResolveUDPAddr("udp", name); err != nil {
		log.Debug("could not resolve %s.", name)
	} else if con, err := net.DialUDP("udp", nil, addr); err != nil {
		log.Debug("could not dial %s.", name)
	} else {
		defer con.Close()
		if wrote, _ := con.Write(packets.WSDDiscoveryPayload); wrote > 0 {
			p.Session.Queue.TrackSent(uint64(wrote))
		} else {
			p.Session.Queue.TrackError()
		}
	}

}
