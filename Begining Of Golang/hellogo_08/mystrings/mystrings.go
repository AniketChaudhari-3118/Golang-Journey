package mystrings

func Reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

//we need to capatilize the first letter of the function because if we dont we will be not able to use the function outside the package
