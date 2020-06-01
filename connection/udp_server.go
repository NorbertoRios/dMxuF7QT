package connection

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"
)

//ConstructUDPServer returns new UDP server
func ConstructUDPServer(port int) *UDPServer {
	addr := fmt.Sprintf(":%v", port)
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
	onNewPacket func(channel *net.UDPAddr, packet []byte)
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
		server.onNewPacket(addr, buf[0:n])
	}
}

func (server *UDPServer) sendBytes(channel IChannel, packet []byte) {
	n, err := server.connection.WriteToUDP(packet, channel.RemoteAddr())
	if err != nil {
		log.Println("[UDPServer] Error while sending bytes. ", err)
		return
	}
	channel.AddTransmitted(int64(n))
}

//Send send message to device
func (server *UDPServer) Send(channel IChannel, message interface{}) {
	switch message.(type) {
	case string:
		{
			server.sendBytes(channel, []byte(message.(string)))
			break
		}
	case []byte:
		{
			server.sendBytes(channel, message.([]byte))
			break
		}
	default:
		{
			server.sendBytes(channel, message.([]byte))
			break
		}
	}
}
