package main

func addTen(b *int ) {

*b = *b+10
}

func main(){
a:= 10
addTen(&a)
println(a)
}
