package connection

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

//ConstructUDPServer returns new UDP server
func ConstructUDPServer(host string, port int) *UDPServer {
	addr := fmt.Sprintf("%v:%v", host, port)
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("[UDPServer] Wrong UDP Address:%v", addr)
		return nil
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)

	if err != nil {
		log.Fatalf("Create udp server error:%v", err.Error())
		return nil
	}

	server := &UDPServer{
		port:       port,
		connection: udpConn,
	}
	return server
}

//UDPServer for genx service
type UDPServer struct {
	port        int
	connection  *net.UDPConn
	onNewPacket func(channel IChannel, packet []byte)
}

//Listen incoming packet
func (server *UDPServer) Listen() {
	for {
		var buf [4096]byte
		n, addr, err := server.connection.ReadFromUDP(buf[0:])
		if err != nil {
			log.Fatalf("Error Reading from udp connection:%v", err.Error())
			return
		}
		log.Println("Received UDP packet:", hex.EncodeToString(buf[0:n]))
		channel := ConstructUDPChannel(addr, server)
		server.onNewPacket(channel, buf[0:n])
	}
}

//SendBytes send bytes
func (server *UDPServer) SendBytes(addr *net.UDPAddr, packet []byte) (int64, error) {
	n, err := server.connection.WriteToUDP(packet, addr)
	if err != nil {
		log.Println("[UDPServer] Error while sending bytes. ", err)
		return 0, err
	}
	return int64(n), nil
}
