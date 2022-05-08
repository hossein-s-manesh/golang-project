package orented

import(
	"fmt"
)


type codeIde string


func cheekOry(Ide codeIde)bool{
	if len(Ide)==10{
		for _,val:=range Ide{
			if !(val>=48 && val<=57){
				return false
			}
		}
		return true
	}
	return false
}

func Ory(){
	var Ide codeIde
	Ide="01234567.9"
	fmt.Println("code_ide:>",Ide,"cheek_code_id:>",cheekOry(Ide))
}