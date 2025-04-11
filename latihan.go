package main
import "fmt"

func bola(klubA, klubB *string){
	var a, b, c int
	var cek bool
	c = 1
	fmt.Print("Pertandingan ", c , " : ")
	fmt.Scan(&a, &b)
	if a >= 0 && b >= 0{
		cek = true
	}
	for cek == true{
		if a > b {
			fmt.Println("Hasil ", c, " : ", *klubA)
		}else if a < b{
			fmt.Println("Hasil ", c, " : ", *klubB)
		}else if a == b{
			fmt.Println("Hasil ", c, " : Draw")
		}
		c++
		fmt.Print("Pertandingan ", c , " : ")
		fmt.Scan(&a, &b)
		if a >= 0 && b >= 0{
			cek = true
		}else {
			cek = false
		}
	}
}
func main(){
	var klubA, klubB string
	fmt.Scan(&klubA, &klubB)
	bola(&klubA, &klubB)
}