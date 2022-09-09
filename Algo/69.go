package algo

func MySqrt(x int) int {
	res := x
	for res*res > x {
		res = (res + x/res) >> 1
	}
	return res
}

/*	magic number: 0x5f375a86
float InvSqrt(float x)
{
	float xhalf = 0.5f * x;
	int i = *(int *)&x;
	i = 0x5f3759df - (i>>1);
	x = *(float *)&i;
	x = x * (1.5f - xhalf * x * x);
	return x;
}
*/
