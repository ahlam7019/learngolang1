package main
import "fmt"
type student struct{
    name string
	regno float32
	dept string
}
func main(){
	st :=student{name:"student 1",regno: 12.2, dept: "cs"}
	fmt.Println("Name:",st.name,"\nRegister num:",st.regno,"\nDepartment:",st.dept)
}