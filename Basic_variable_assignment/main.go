/*
This code should store "codewa.rs" as a variable called name but it's not working.
Can you figure out why?

package kata
func Namevar() string {
  var a string == "code"
  var b string == "wa.rs"
  var name string == a + b
  return name
}

*/

package kata

func Namevar() string {
	var a string = "code"
	var b string = "wa.rs"
	var name string = a + b
	return name
}
