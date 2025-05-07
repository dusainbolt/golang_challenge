# Challenge 01 - Type Assertions in Go

In this challenge, we are going to become familiar with the concept of **Type Assertions** in Go!

---

## Preface

If you are new to the language, then type assertions are a concept that can sometimes trip you up and appear a little tricky at first, but after overcoming the syntax it becomes far easier to understand.

Through using type assertions, we can retrieve the dynamic value of an interface. For example:

```go
var myName interface{} = "Elliot"

name := myName.(string)
fmt.Println(name)
