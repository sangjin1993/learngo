package ifwithatwist

// variable expression
func CanIDrink(age int) bool {
	// koreanAge = age + 2 를 안쓰고 if문에 넣어 명확히 사용할 변수를 알려줄수 있음
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}
