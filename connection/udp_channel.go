package connection

import (
	"net"
	"time"
)

//IChannel channel interface
type IChannel interface {
	AddReceived(int64)
	AddTransmitted(int64)
	Received() int64
	Transmitted() int64
	RemoteAddr() *net.UDPAddr
	RemoteIP() string
}

//ConstructUDPChannel returns new channel
func ConstructUDPChannel(addr *net.UDPAddr) IChannel {
	return &UDPChannel{
		ConnectedAt: time.Now().UTC(),
		clientAddr:  addr,
	}
}

//UDPChannel cahnnel for device
type UDPChannel struct {
	ConnectedAt    time.Time
	LastActivityTs time.Time
	received       int64
	transmitted    int64
	clientAddr     *net.UDPAddr
}

//Received received bytes
func (c *UDPChannel) Received() int64 {
	return c.received
}

//Transmitted transmitted bytes
func (c *UDPChannel) Transmitted() int64 {
	return c.transmitted
}

//AddTransmitted to cahnnel
func (c *UDPChannel) AddTransmitted(count int64) {
	c.transmitted += count
}

//AddReceived to cahnnel
func (c *UDPChannel) AddReceived(count int64) {
	c.received += count
}

//RemoteAddr client address
func (c *UDPChannel) RemoteAddr() *net.UDPAddr {
	return c.clientAddr
}

//RemoteIP indicates device remote address
func (c *UDPChannel) RemoteIP() string {
	return c.clientAddr.String()
}
