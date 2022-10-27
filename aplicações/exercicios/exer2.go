package exercicios

import (
	"encoding/json"
	"fmt"
)

func Exer2() {

	//	usando recursos de noRep.go
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	var recebe []User
	err := json.Unmarshal([]byte(s), &recebe)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(recebe)
}
