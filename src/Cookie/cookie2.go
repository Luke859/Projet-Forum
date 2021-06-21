import (
    "io"
    "net/http"
    "log"
    "fmt"
    "time"
	"github.com/satori/go.uuid"
)

func DrawMenu(w http.ResponseWriter){
    w.Header().Set("Content-Type", "text/html")
    io.WriteString(w, "<a href='/writecookie'>Write Cookie <ba><br/>" + "\n")
}

func WriteCookieServer(w http.ResponseWriter, req *http.Request) {
    myuuid, err := uuid.NewV4()
    if err == nil{
        fmt.Println("Your UUID is:", myuuid)
    }
    expire := time.Now().AddDate(0, 0, 1)
    cookie := http.Cookie{Name: "testcookiename", Value: myuuid, Path: "/", Expires: expire, MaxAge: 86400}

    http.SetCookie(w, &cookie)
    DrawMenu(w)

}

func main() {

    http.HandleFunc("/writecookie", WriteCookieServer)
    fmt.Println("listen on 3000")
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}