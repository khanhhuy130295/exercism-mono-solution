package raindrops

import "fmt"

func Convert(number int) string {
    var result string = ""
    
	if (number % 3 == 0){
        result = "Pling"
    }
    
    if (number % 5 == 0){
        result += "Plang"
    }
    
    if (number % 7 == 0){
       result += "Plong"
    }
    if (len(result) == 0){
        result =  fmt.Sprintf("%d",number)
    }
	return result
}
