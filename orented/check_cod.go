package orented

import "strconv"

var palacealue = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
const residConst=11
const controldigitindex=9


func CheekOry(Ide codeIde) bool {
	if len(Ide) == 10 {
		for _, val := range Ide {
			if !(val >= 48 && val <= 57) {
				return false
			}
		}
		if Ide.cheek_hoviyat_cood_id(){
			return true

		}
	}
	return false
}

func (Ide codeIde) cheek_hoviyat_cood_id() bool {
	var (
		sum int
		resid int
	)
	for index := 0; index < 9; index++ {
		digit, err := strconv.Atoi(string(Ide[index]))
		if err != nil {
			return false
		}
		sum = sum + digit*palacealue[index]
	}

	resid = sum % residConst

	controlDigit,err :=strconv.Atoi(string(Ide[controldigitindex]))
	if err !=nil{
		return false
	}
	if resid<2{
		if resid==controlDigit{
			return true
		}
	}else{
		if controlDigit==(11-resid){
			return true
		}
	}

	return false
}