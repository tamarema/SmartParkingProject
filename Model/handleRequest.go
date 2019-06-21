package Model

import (
	"SmartParkingProject/DB_Parking"
	"SmartParkingProject/Model/ConnectionToServices"
	"bufio"
	"fmt"
	"net"
	"strings"
)

var db = DB_Parking.NewClass()

func handleRequest(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	UserName := make([]byte, 64)
	Crippled := make([]byte, 64)
	scanner := bufio.NewScanner(conn)

	// get user name
	userName := getArgsFromUser(conn, UserName)

	// get if the user is crippled or not
	crippled := getArgsFromUser(conn, Crippled)

	if !DB_Parking.UserExists(userName){
		newUser := DB_Parking.User{userName, crippled}
		DB_Parking.AddNewUser(newUser)
	}

	for {
		ok := scanner.Scan()

		if !ok {
			break
		}

		getParkingRequest(scanner.Text(), conn)
	}

	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func getParkingRequest(message string, conn net.Conn) string {
	parking := DB_Parking.GetFreeParking()
	var path  = ConnectionToServices.BuildPath(parking[0])
	return path
}

func getArgsFromUser(conn net.Conn, args []byte) string {

	_, err := conn.Read(args)
	if err != nil {
		println("Error reading user name from user:")
		conn.Close()
		return ""
	}

	argument := strings.Split(string(args), "\n")[0]
	return argument
}

