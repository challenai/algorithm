package solution

func devide(devidend int, devisor int) int {
	needMinusSymbol := devidend^devisor < 0
	devidend = abs(devidend)
	devisor = abs(devisor)
	counter := 0
	for devidend >= devisor {
		counter++
		devidend -= devisor
	}
	if needMinusSymbol {
		return -counter
	}
	return counter
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func devide2(devidend int, devisor int) int {
   //  needMinusSymbol := devidend ^ devisor < 0
    return 0
}
