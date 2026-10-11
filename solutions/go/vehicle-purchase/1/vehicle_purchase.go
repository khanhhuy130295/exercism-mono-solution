package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	var is_need_license bool = false
	switch kind {
	case "car", "truck":
		is_need_license = true
	default:
		is_need_license = false
	}

	return is_need_license
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	var comparison string = ""

	if option1 < option2 {
		comparison = option1
	} else {
		comparison = option2
	}

	return fmt.Sprintf("%s is clearly the better choice.", comparison)
}

var vehicle_ages = []struct {
	value   string
	percent float64
}{
	{value: "like_new", percent: 0.8},
	{value: "second_hand", percent: 0.7},
	{value: "bad_condition", percent: 0.5},
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {

	if age == 0 || originalPrice == 0 {
		return 0
	}
	var resell_price float64 = 0

	if age < 3 {
		resell_price = vehicle_ages[0].percent * originalPrice
	}
	if age >= 3 && age < 10 {
		resell_price = vehicle_ages[1].percent * originalPrice
	}
	if age >= 10 {
		resell_price = vehicle_ages[2].percent * originalPrice
	}

	return resell_price

}
