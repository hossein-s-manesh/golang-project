package orented

import(
	"fmt"
)


type codeIde string


func cheekOry(Ide codeIde)bool{
	if len(Ide)==10{
		return true
	}
	return false
}

func Ory(){
	var Ide codeIde
	Ide="01_3456789"
	fmt.Println("code_ide:>",Ide,"cheek_code_id:>",cheekOry(Ide))
}