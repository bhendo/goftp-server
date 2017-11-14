package server

import "net"

// Auth povides a mechinism for checking a username and password
type Auth interface {
	// params - username, password, connection
	// returns - true if the user is authenticated
	CheckPasswd(string, string, net.Conn) (bool, error)
}
