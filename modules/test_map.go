package test_map

func getAllMathOperations(a, b int) (mathOperate map[string]int) {
	mathOperate = map[string]int{
		"sum": a + b,
		"dec": a - b,
		"mul": a * b,
		"div": a / b,
	}
	return
}
