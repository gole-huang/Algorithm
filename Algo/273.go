package algo

func toNumeric(num int) (str string) {
    switch num {
    case 1: str="One"
    case 2: str="Two"
    case 3: str="Three"
    case 4: str="Four"
    case 5: str="Five"
    case 6: str="Six"
    case 7: str="Seven"
    case 8: str="Eight"
    case 9: str="Nine"
    }
    return
}

func transform(num int) string {
	if num==0{
		return ""
	}
	hundred:=toNumeric(num/100)
	if hundred != "" {
		hundred+=" Hundred"
	}
    num%=100
    if num==0{
        return hundred
    } else if  hundred != "" {
        hundred+=" "
    }
    switch num {
    case 1: return hundred+"One"
    case 2: return hundred+"Two"
    case 3: return hundred+"Three"
    case 4: return hundred+"Four"
    case 5: return hundred+"Five"
    case 6: return hundred+"Six"
    case 7: return hundred+"Seven"
    case 8: return hundred+"Eight"
    case 9: return hundred+"Nine"
    case 10: return hundred+"Ten"
    case 11: return hundred+"Eleven"
    case 12: return hundred+"Twelve"
    case 13: return hundred+"Thirteen"
    case 14: return hundred+"Fourteen"
    case 15: return hundred+"Fifteen"
	case 16: return hundred+"Sixteen"
	case 17: return hundred+"Seventeen"
	case 18: return hundred+"Eighteen"
	case 19: return hundred+"Nineteen"
    }
    var str string
	ten:=num/10
	switch ten {
	case 2: str=hundred+"Twenty"
	case 3: str=hundred+"Thirty"
	case 4: str=hundred+"Forty"
	case 5: str=hundred+"Fifty"
	case 6: str=hundred+"Sixty"
	case 7: str=hundred+"Seventy"
	case 8: str=hundred+"Eighty"
	case 9: str=hundred+"Ninety"
	}
    one:=toNumeric(num%10)
    var conn string
    if str!="" && one!=""{
        conn=" "
    }
	return str + conn + one
}

func NumberToWords(num int) string {
    if num==0{
        return "Zero"
    }
    nums:=make([]int,0)
    for num!=0 {
        nums=append(nums,num%1000)
        num/=1000
    }
	str:=transform(nums[0])
	thousand:=[3]string{" Thousand", " Million", " Billion"}
	for i:=1;i<len(nums);i++{
        tmp:=transform(nums[i])
        var newstr,conn string
        if tmp!=""{
            newstr=tmp+thousand[i-1]
        }
        if newstr!="" && str!="" {
            conn=" "
        }
        str=newstr+conn+str
	}
	return str
}