package network

type NetAddr string


type RPC struct{//Message which is going to transfer over transport layer
	From NetAddr
	Payload []byte
}

type Transport interface{
	Consume() <-chan RPC //return channel RPC
	Connect(Transport) error
	SendMessage(NetAddr,[]byte) error
	Addr() NetAddr
}