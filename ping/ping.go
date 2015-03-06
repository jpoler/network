package ping

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/internal/iana"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
	"time"
)

const net_type = "ip4"

func ICMPEcho(target string) (err error) {
	dst, err := net.ResolveIPAddr(net_type, target)
	if err != nil {
		return
	}
	sanityCheck := fmt.Sprintf("%s:%d", net_type, iana.ProtocolICMP)
	fmt.Println(sanityCheck)
	c, err := net.ListenPacket(sanityCheck, "0.0.0.0")
	if err != nil {
		return err
	}
	defer c.Close()

	fmt.Printf("pinging address: %v\n", dst.IP)

	//get a packetconn
	pc := ipv4.NewPacketConn(c)

	// set control message for FlagTTL, FlagSrc, FlagDst, FlagInterface
	if err = pc.SetControlMessage(
		ipv4.FlagTTL|ipv4.FlagSrc|ipv4.FlagDst|ipv4.FlagInterface, true); err != nil {
		return
	}

	//make icmp message
	ICMPMessage := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Data: []byte("Foo"),
		},
	}

	// make a buffer
	readBuffer := make([]byte, 1500)
	// iterate
	for i := 1; i < 64; i++ {
		//marshall message
		writeBuffer, err := ICMPMessage.Marshal(nil)

		// setTTL

		err = pc.SetTTL(i)

		if err != nil {
			return err
		}
		fmt.Println("foo")
		// set seq in ICMP header
		ICMPMessage.Body.(*icmp.Echo).Seq = i
		// take note of begin
		begin := time.Now()
		// write to packetconn
		_, err = pc.WriteTo(writeBuffer, nil, dst)
		if err != nil {
			return err
		}

		// set read deadline using an input timout param
		if err = pc.SetReadDeadline(time.Now().Add(3 * time.Second)); err != nil {
			return err
		}
		fmt.Println("bar")
		// try to read
		n, cm, peer, err := pc.ReadFrom(readBuffer)

		// error check for timeout and continue if so
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				fmt.Printf("%v\t*\n", i)
			}
		}
		rm, err := icmp.ParseMessage(iana.ProtocolICMP, readBuffer[:n])
		if err != nil {
			return err
		}

		rtt := time.Since(begin)

		// otherwise parse message using icmp package
		// use time.Since to find elapsed
		switch rm.Type {
		case ipv4.ICMPTypeTimeExceeded:
			names, _ := net.LookupAddr(peer.String())
			fmt.Printf("%d\t%v %+v %v\n%+v\n", i, peer, names, rtt, cm)
		case ipv4.ICMPTypeEchoReply:
			names, _ := net.LookupAddr(peer.String())
			fmt.Printf("%d\t%v %+v %v\n%+v\n", i, peer, names, rtt, cm)
			return err
		default:
			log.Printf("unknown ICMP message: %+v\n", rm)
		}

		// readfrom returns peer, so we need to use net.LookupAddr to find the peer

	}

	return
}
