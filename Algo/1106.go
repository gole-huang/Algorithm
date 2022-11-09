package algo

type expr struct {
	exp []bool
	opr int
}

type exprs []expr

func (e expr) res() bool {
	switch e.opr {
	case 1: //neg
		return !e.exp[0]
	case 2: //and
		for _, v := range e.exp {
			if !v {
				return false
			}
		}
		return true
	case 3: //or
		for _, v := range e.exp {
			if v {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func parseBoolExpr(expression string) bool {
	var exps exprs = make([]expr, 0)
	var res bool
	var opt int
	for _, b := range expression {
		switch b {
		case 't':
			res = true
		case 'f':
			res = false
		case ',':
			exps[len(exps)-1].exp = append(exps[len(exps)-1].exp, res)
		case '!':
			opt = 1
		case '&':
			opt = 2
		case '|':
			opt = 3
		case '(':
			exps = append(exps, expr{make([]bool, 0), opt})
		case ')':
			exps[len(exps)-1].exp = append(exps[len(exps)-1].exp, res)
			res = exps[len(exps)-1].res()
			exps = exps[:len(exps)-1]
		}
	}
	return res
}
