package Algo

type OrderedStream struct {
	steam []string
	cur   int
}

func Constructor(n int) OrderedStream {
	os := new(OrderedStream)
	os.steam = make([]string, n)
	return *os
}

func (os *OrderedStream) Insert(idKey int, value string) []string {
	os.steam[idKey-1] = value
	if os.steam[os.cur] == "" {
		return nil
	} else {
		start := os.cur
		for os.cur < len(os.steam) && os.steam[os.cur] != "" {
			os.cur++
		}
		return os.steam[start:os.cur]
	}
}
