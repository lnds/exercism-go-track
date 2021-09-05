package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill create a new bill
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem add item to customer bill
func AddItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	if value, ok := units[unit]; ok {
		bill[item] = value
		return true
	}
	return false
}

// RemoveItem remove item from customer bill
func RemoveItem(bill map[string]int, units map[string]int, item string, unit string) bool {
	if item_value, ok := bill[item]; !ok {
		return false
	} else if value, ok := units[unit]; !ok {
		return false
	} else {
		new_val := item_value - value
		switch {
		case new_val < 0:
			return false
		case new_val == 0:
			delete(bill, item)
			return true
		default:
			bill[item] = new_val
			return true
		}
	}
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	val, ok := bill[item]
	return val, ok
}
