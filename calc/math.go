
package calc

// returns sum of the two integers

func Add(numbers ...int) int {
      
      	sum := 0
	for _, num := range numbers{

        sum = sum + num
        }

	return sum
}
