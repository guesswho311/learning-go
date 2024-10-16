package main
import (
    "encoding/json"
    "fmt"
    "os"
)

type response1 struct {
    Page   int
    Fruits []string
}

type response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

type Message struct {
    Name string
    Body string
    Time int64
}

func main() {
    bolB, _ := json.Marshal(true)
        fmt.Println(string(bolB))

        intB, _ := json.Marshal(1)
            fmt.Println(string(intB))
            fltB, _ := json.Marshal(2.34)
            fmt.Println(string(fltB))
            strB, _ := json.Marshal("gopher")
            fmt.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
        slcB, _ := json.Marshal(slcD)
        fmt.Println(string(slcB))
    mapD := map[string]int{"apple": 5, "lettuce": 7}
        mapB, _ := json.Marshal(mapD)
        fmt.Println(string(mapB))


    res1D := &response1{
            Page:   1,
            Fruits: []string{"banana", "peach", "pear"}}
        res1B, _ := json.Marshal(res1D)
        fmt.Println(string(res1B))

    res2D := &response2{
            Page:   1,
            Fruits: []string{"banana", "peach", "pear"}}
        res2B, _ := json.Marshal(res2D)
        fmt.Println(string(res2B))

    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    var dat map[string]interface{}

    if err := json.Unmarshal(byt, &dat); err != nil {
            panic(err)
        }
        fmt.Println(dat)

    num := dat["num"].(float64)
        fmt.Println(num)

    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    str2 := strs[1].(string)
    fmt.Println(str1)
    fmt.Println(str2)


    str := `{"page": 1, "fruits": ["apple", "peach"]}`
        res := response2{}
        json.Unmarshal([]byte(str), &res)
        fmt.Println(res)
        fmt.Println(res.Fruits[0])

    enc := json.NewEncoder(os.Stdout)
        d := map[string]int{"apple": 5, "lettuce": 7}
        enc.Encode(d)

    m := Message{"Alice", "Hello", 1294706395881547000}
    b, _ := json.Marshal(m)
    fmt.Println(b)

    var mm Message
    json.Unmarshal(b, &mm)
    fmt.Println(mm)

}