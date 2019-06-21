package View

import "SmartParkingProject/Controller"

func main()  {
	newParking()
}

func newParking()  {
	var path = Controller.FindfreeParking()
	showPath(path)
}

func showPath(path string)  {

}

