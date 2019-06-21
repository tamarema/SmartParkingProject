package DB_Parking

// Parking info Singeltone
type singelton map[string]string

var (
	instance singelton
)

func NewClass() singelton {

	if instance == nil {

		instance = make(singelton) // <-- not thread safe
	}

	return instance
}


// The function return list of all free parking
func GetFreeParking() [] int {
	listFreeParking := make([]int, 50)

	return listFreeParking
}

// The function get location of parking and update state
func UpdateParking(location int, state int) bool  {

	return true
}
