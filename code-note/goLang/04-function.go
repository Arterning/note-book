/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int
 
	if (num1 > num2) {
	   result = num1
	} else {
	   result = num2
	}
	return result 
 }


//  Go 函数可以返回多个值，例如：
 func swap(x, y string) (string, string) {
	return y, x
 }