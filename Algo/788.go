package algo

func RotatedDigits(n int) int {
	if n == 10000 {
		n = 9999
	}
	_1 := n / 1000
	hundred, ten, one := 1, 1, 1
	__1 := _1
	_hundred, _ten, _one := 1, 1, 1
	switch {
	case _1 < 3:
	case _1 == 3:
		hundred, _hundred, ten, _ten, one, _one = 0, 0, 0, 0, 0, 0
	case _1 == 4:
		_1, hundred, _hundred, ten, _ten, one, _one = 3, 0, 0, 0, 0, 0, 0
	case _1 == 5:
		_1 = 3
	case _1 == 6:
		_1 = 4
	case _1 == 7:
		_1, hundred, _hundred, ten, _ten, one, _one = 5, 0, 0, 0, 0, 0, 0
	case _1 <= 9:
		_1 -= 3
	}
	switch {
	case __1 < 2:
	case __1 == 2:
		_hundred, _ten, _one = 0, 0, 0
	case __1 <= 4:
		__1 = 2
	case __1 == 5 || __1 == 6:
		__1, _hundred, _ten, _one = 2, 0, 0, 0
	case __1 <= 8:
		__1 = 2
	case __1 == 9:
		__1, _hundred, _ten, _one = 3, 0, 0, 0
	}
	n %= 1000
	_2 := n / 100
	__2 := _2
	switch {
	case _2 < 3:
	case _2 == 3:
		ten, _ten, one, _one = 0, 0, 0, 0
	case _2 == 4:
		_2, ten, _ten, one, _one = 3, 0, 0, 0, 0
	case _2 == 5:
		_2 = 3
	case _2 == 6:
		_2 = 4
	case _2 == 7:
		_2, ten, _ten, one, _one = 5, 0, 0, 0, 0
	case _2 <= 9:
		_2 -= 3
	}
	switch {
	case __2 < 2:
	case __2 == 2:
		_ten, _one = 0, 0
	case __2 <= 4:
		__2 = 2
	case __2 == 5 || __2 == 6:
		__2, _ten, _one = 2, 0, 0
	case __2 <= 8:
		__2 = 2
	case __2 == 9:
		__2, _ten, _one = 3, 0, 0
	}
	n %= 100
	_3 := n / 10
	__3 := _3
	switch {
	case _3 < 3:
	case _3 == 3:
		one, _one = 0, 0
	case _3 == 4:
		_3, one, _one = 3, 0, 0
	case _3 == 5:
		_3 = 3
	case _3 == 6:
		_3 = 4
	case _3 == 7:
		_3, one, _one = 5, 0, 0
	case _3 <= 9:
		_3 -= 3
	}
	switch {
	case __3 < 2:
	case __3 == 2:
		_one = 0
	case __3 <= 4:
		__3 = 2
	case __3 == 5 || __3 == 6:
		__3, _one = 2, 0
	case __3 <= 8:
		__3 = 2
	case __3 == 9:
		__3, _one = 3, 0
	}
	_4 := n % 10
	__4 := _4
	switch {
	case _4 <= 2:
		_4 += 1
	case _4 <= 4:
		_4 = 3
	case _4 <= 6:
		_4 -= 1
	case _4 <= 9:
		_4 -= 2
	}
	switch {
	case __4 <= 1:
		__4 += 1
	case __4 <= 7:
		__4 = 2
	case __4 <= 9:
		__4 = 3
	}
	allNum := _1*7*7*7 + hundred*_2*7*7 + ten*_3*7 + one*_4
	allInvalid := __1*3*3*3 + _hundred*__2*3*3 + _ten*__3*3 + _one*__4
	return allNum - allInvalid
}
