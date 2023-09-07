// 1ファイル 1パッケージ
package main

import(
  "fmt"
  "time"
)

func outer() {
  var s4 string = "outer"
  fmt.Println(s4)
}

// エントリポイント
func main() {
  fmt.Printf("Hello World\n")
  // time.Now()
  fmt.Println(time.Now())

  var n int = 100
  fmt.Println(n+10)

  var s string = "Hello Go"
  fmt.Println(s)


  var t, f bool = true, false

  fmt.Println(t,f)

  var (
    i2 int = 100
    s2 string = "Golang"
  )

  fmt.Println(i2, s2)

  var i3 int 
  var s3 string
  var t2 bool

  i2 = 5000

  fmt.Println(i3, s3, t2)
  
  i := 150

  fmt.Println(i,i2)

  outer()

  var i6 int16 = 100

  fmt.Println(i6)

  fmt.Printf("%T\n", i6)

  fmt.Printf("%T\n", int32(i2))

  fmt.Printf(`
  test
      test  
          test
  `)

  fmt.Printf("\"")
  fmt.Println(string(s[0]))


  // 型
  // byte(uint8)型

  byteA := []byte{72, 73}
  fmt.Printf(byteA)  

}
