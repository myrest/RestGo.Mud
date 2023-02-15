package Telnet

import (
	"fmt"
	"net"

	uu "github.com/satori/go.uuid"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/ScheduleService"
	"rest.com.tw/tinymud/src/RestGo.MUD.Core.Command/CommonServices/StructCollection"
	"rest.com.tw/tinymud/src/RestGo.Util/utility"
)

// TelnetServer struct holds the server's information
type TelnetServer struct {
	listener net.Listener
	clients  []*StructCollection.MudClient
	encoder  *encoding.Encoder
}

// NewTelnetServer returns a new TelnetServer
func Listen(portNumber int) (*TelnetServer, error) {
	ip := utility.MachingIPAddress
	listenIPAddr := fmt.Sprintf(":%d", portNumber)
	fmt.Printf("服務已啟動於： %s port:%d\n", ip, portNumber)

	listener, err := net.Listen("tcp", listenIPAddr)
	if err != nil {
		return nil, err
	}
	telnetServer := &TelnetServer{listener: listener, encoder: unicode.UTF8.NewEncoder()}
	return telnetServer, nil
}

// Start starts the telnet server
func (s *TelnetServer) Start() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Generate a new UUID and print out the connection information
		guid := uu.NewV4()
		fmt.Printf("New connection from: %s - %s\n", conn.RemoteAddr(), guid.String())

		// Create a new MudClient struct and assign the connection and connection ID
		mudClient := &StructCollection.MudClient{
			Conn:         conn,
			ConnectionID: guid.String(),
		}

		// Append the new client to the list of clients
		s.clients = append(s.clients, mudClient)

		// Handle the connection in a goroutine
		go s.handleConnection(mudClient)
	}
}

func (s *TelnetServer) Shutdown() {
	//停止所有的Job
	ScheduleService.Shutdown()

	s.listener.Close()

	// SendToAll sends a message to all clients
	for _, mudconn := range s.clients {
		fmt.Fprintln(mudconn.Conn, "Server is shutting down.")
		mudconn.Conn.Close()
	}
}
