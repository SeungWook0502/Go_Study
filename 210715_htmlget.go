package main
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
)
 
func main() {
    // GET ȣ��
    resp, err := http.Get("https://news.naver.com/")
    if err != nil {
        panic(err)
    }
 
    defer resp.Body.Close()
 
    // ��� ���
    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    fmt.Printf("%s\n", string(data))
}