//Service

package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"

    //"os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    var err error
    readData("books.csv")
    switch request.Method {
    case "GET":
        err = handleGet(writer, request)
    case "POST":
        err = handlePost(writer, request)
    case "PUT":
        err = handlePut(writer, request)
    case "DELETE":
        err = handleDelete(writer, request)
    }
    writeData("books.csv")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
}

func marshalTest(){
    type Book struct {
        Id        string `json:"id"`
        Title     string `json:"title"`
        Edition   string `json:"edition"`
        Copyright string `json:"copyright"`
        Language  string `json:"language"`
        Pages     string `json:"pages"`
        Author    string `json:"author"`
        Publisher string `json:"publisher"`
    }

    b1 := Book{
        Id:"99",
        Title:"t",
        Edition:"1st",
        Copyright:"2020",
        Language:"SPANISH",
        Pages:"111",
        Author:"JJ",
        Publisher:"Pub",
    }
    books := []Book {b1}
    bs, err := json.Marshal(books)
    if(err!=nil){
             fmt.Println(err)
    }
    fmt.Println(string(bs))
}

func findN(x string) int {
    letters:= []string{"a","b","c","d","e","f"}

    for i, letter := range letters {
        if x == letter {
            return i
        }
    }
    return -1
}

func arrayTest(){
        s := make([]string, 3)
        fmt.Println("emp:", s)

        s[0] = "a"
        s[1] = "b"
        s[2] = "c"
        fmt.Println("set:", s)
        fmt.Println("get:", s[2])

        fmt.Println("len:", len(s))

        s = append(s, "d")
        s = append(s, "e", "f")
        fmt.Println("apd:", s)

        c := make([]string, len(s))
        copy(c, s)
        fmt.Println("cpy:", c)

        l := s[2:5]
        fmt.Println("sl1:", l)

        l = s[:5]
        fmt.Println("sl2:", l)

        l = s[2:]
        fmt.Println("sl3:", l)

        t := []string{"g", "h", "i"}
        fmt.Println("dcl:", t)

        twoD := make([][]int, 3)
        for i := 0; i < 3; i++ {
            innerLen := i + 1
            twoD[i] = make([]int, innerLen)
            for j := 0; j < innerLen; j++ {
                twoD[i][j] = i + j
            }
        }
        fmt.Println("2d: ", twoD)
        i := findN("d")
        fmt.Println(i)
        fmt.Println(s[:i])
        fmt.Println(s[i:])
        fmt.Println(append(s[:i], s[i+1:]...))
}

func main() {
    http.HandleFunc("/book/", handler)
    http.ListenAndServe(":"+os.Getenv("PORT"), nil)
    //arrayTest()
}