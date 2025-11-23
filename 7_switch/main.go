package main

import "fmt"

func main() {
// 	i := 1

// 	switch i {
// 	case 1:
// 		fmt.Println("one")
	
//     case 2:
// 	    fmt.Println("two")

// 	default:
// 		fmt.Println("no case")
// }

    //type switch

	whoAmi := func (i interface{})  {

		switch t := i.(type){
		case int:
			fmt.Println("its integer")
		case string:
			fmt.Println("its a string")

		default:{
			fmt.Println("other",t)
		}
		}
		
	}


	whoAmi("gola")
}



