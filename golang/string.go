//string manipulation
package main
 import(
         "fmt"
         "strings"
)
func main(){
  name := "caro baby dorothy abigail"
  fish := "TILAPIA SALMON DAGAA SHARK"
fmt.Println(name)
// Contains is used to check for items and returns a true or false
// and returns false if the item is not in the string
fmt.Println(strings.Contains(name,"baby"))
//ReplaceAll is used to replace some items in a string
fmt.Println(strings.ReplaceAll(name,"baby","tiffany"))
fmt.Println(strings.Contains(name,"tif"))
// ToUpper and ToLower is used to tur strings into either upper case letters or lowercase
fmt.Println(strings.ToUpper(name))
fmt.Println(fish)
fmt.Println(strings.ToLower(fish))
// Index is used to tell the position of the string
fmt.Println("the index ofof SALMON is below ")
fmt.Println(strings.Index(fish,"SALMON"))
// Split can be used to remove strings
fmt.Println(strings.Split(name,"baby"))
// HasPrefix is used to show  what a string  starts with
// while HasSuffix is for what the string ends with
fmt.Println(strings.HasPrefix(name,"ca"))
fmt.Println(strings.HasSuffix(name,"gail"))
// Join is used to join two strings together
join := []string{ "mary","ann"}
fmt.Println(join)
//nam := strings.Join(nam,"+")
//fmt.Println(join)
parts := []string{"https:", "", "api.example.com", "v1", "users"}
url := strings.Join(parts, "/")
fmt.Println(url)
}
