package main

type insuranceStatus struct {
	hasInsuranceFlag bool
	totaled          bool
	dented           bool
	bigDent          bool
}

func (s insuranceStatus) hasInsurance() bool {
	return s.hasInsuranceFlag
}

func (s insuranceStatus) isTotaled() bool {
	return s.totaled
}

func (s insuranceStatus) isDented() bool {
	return s.dented
}

func (s insuranceStatus) isBigDent() bool {
	return s.bigDent
}

func thisExplicit() (int, int) {
	var x int
	var y int

	return x, y // This is explicit return
}

func thisImplicit() (x, y int) {
	return // return // This is implicit return
}

func veryNestedLoop(status insuranceStatus) int {
	amount := 0
	if !status.hasInsurance() {
		amount = 1
	} else {
		if status.isTotaled() {
			amount = 10000
		} else {
			if status.isDented() {
				amount = 160
				if status.isBigDent() {
					amount = 270
				}
			} else {
				amount = 0
			}
		}
	}
	return amount
}

func loopWithGuardClause(status insuranceStatus) int {
	if !status.hasInsurance() {
		return 1
	}
	if status.isTotaled() {
		return 10000
	}
	if !status.isDented() {
		return 0
	}
	if status.isBigDent() {
		return 270
	}
	return 160
}
