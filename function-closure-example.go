import (
	"fmt"
)

func main() {

         num := 3
         multiup := func(n int) int {
            num *= n
            return num
           }
	fmt.Println(multiup(5))
	fmt.Println(multiup(2))
}
