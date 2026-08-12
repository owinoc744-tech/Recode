package main
import(
        "strconv"
        "fmt"
       "reflect"
)
func main(){
// AtoI
agestr := "24"
fmt.Println("age this year",agestr)
age,_ := strconv.Atoi(agestr)
fmt.Println("nextyear",age+1)
//Atoi2
id := "423163"
fmt.Println("id number")
id1,_ := strconv.Atoi(id)
fmt.Println(id1)
fmt.Println("reflection of both id and id1")
fmt.Println(reflect.TypeOf(id))
fmt.Println(reflect.TypeOf(id1))
//Itoa 
score := 100
msg := "score:"+strconv.Itoa(score)
fmt.Println(msg)
fmt.Println("to provetha we converted Itoa we used package reflect")
fmt.Println(reflect.TypeOf(msg))
//Itoa2
on := 1
fmt.Println("reflection or var on")
fmt.Println(reflect.TypeOf(on))
off := strconv.Itoa(on)
fmt.Println("reflection of var off conv on(int) to off (str)")
fmt.Println(reflect.TypeOf(off))
//ParseFloat string to float
pricestr :=  "59.99"
price,_ :=strconv.ParseFloat(pricestr,32)
fmt.Println(price)
// ParseFloat2
w := "54.76"
we,_ :=strconv.ParseFloat(w,64)
fmt.Println(we + 31)
	fmt.Println(reflect.TypeOf(we))
// Parse bool
fmt.Println(strconv.ParseBool("true"))
fmt.Println(strconv.ParseBool("false"))
fmt.Println(strconv.ParseBool("1"))
fmt.Println(strconv.ParseBool("0"))

 b,_:=strconv.ParseBool("true")
fmt.Println(b)
}
