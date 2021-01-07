package languangeinsite

//给定两个字符串，从hayStack中找出 needle第一次出现的位置
func StrStr(hayStack,needle string)int {
	if len(hayStack)==0 || len(needle)==0 {
		return -1
	}
	var index int = -1
	for i:=0 ; i <len(hayStack);i++{
		if hayStack[i] == needle[0]{
			index = i
			for j:=0; j < len(needle) ;j++{
				if hayStack[i]==needle[j]{
					i++
				}else{
						index = -1
						break
				}
			}
		}else{
			continue
		}
	}
	return index
}
//范例