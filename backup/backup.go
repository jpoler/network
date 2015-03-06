package ping

import (
	"fmt"
	"os"
	"syscall"

	// "golang.org/x/net/icmp"
	// "golang.org/x/net/internal/iana"
	"golang.org/x/net/ipv4"
	// "net"
)

// PING Algorithm
// initialize echo request
// send echo request
// wait for echo reply (or time out)
// recieve reply
// report results
// loop back to 1 or stop

func ICMPEcho(addr string, ttl int) (err error) {

}

func EchoPacket(ttl int) (packet []byte, err error) {
	h := ipv4.Header{
	Version: iana.ProtocolICMP
	}
}

func ICMPListen(addr string) (err error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	f := os.NewFile(uintptr(fd), fmt.Sprintf("fd %d", fd))

	for {
		buf := make([]byte, 1024)
		_, err := f.Read(buf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		head, err := ipv4.ParseHeader(buf)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(head.String())

	}

}

// func Ping(addr string, ttl int) (err error) {
// 	fmt.Println(addr, ttl)
// 	protocol := "icmp"
// 	netaddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
// 	conn, _ := net.ListenIP("ip4:"+protocol, netaddr)
// 	defer conn.Close()
// 	buf := make([]byte, 1024)

// 	for {
// 		_, _, _ = conn.ReadFrom(buf)
// head, _ := ipv4.ParseHeader(buf)
// fmt.Println(head.String())

// 	}
// 	return
// }

// func Ping(addr string, ttl int) (err error) {
// 	fmt.Println("FOO")
// 	c, err := icmp.ListenPacket("ip4:icmp", "127.0.0.1")
// 	if err != nil {
// 		return err
// 	}
// 	defer c.Close()

// 	buf := make([]byte, 2000)
// 	for {
// 		n, addr, err := c.ReadFrom(buf)
// 		if err != nil {
// 			return errors.New("LSKDFJLSKDF")
// 		}
// 		msg, _ := icmp.ParseMessage(iana.ProtocolICMP, buf[:n])
// 		fmt.Printf("%v\n", *msg)
// 		fmt.Println(addr)
// 	}
// 	// icmpMsg := icmp.Message{
// 	// 	Type: ipv4.ICMPTypeEcho,
// 	// 	Code: 0,
// 	// 	Body: &icmp.Echo{
// 	// 		ID:   os.Getpid() & 0xffff,
// 	// 		Seq:  1,
// 	// 		Data: []byte("Weeeeeee"),
// 	// 	},
// 	// }

// 	// icmpMsgBytes := icmpMsg.Marshall()

// 	return
// 
}
