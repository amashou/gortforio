package main

// "GortfoRio/yahoo"

import (
	// "database/sql"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	// "io/ioutil"
	// "net/http"
	// _ "github.com/mattn/go-sqlite3"
)

// type JsonTestData struct {
// 	Name string
// 	Age  int
// }

// func getXemHandler(w http.ResponseWriter, r *http.Request) {
// 	apiClient := bitflyer.New(os.Getenv("BITFLYER_API_KEY"), os.Getenv("BITFLYER_API_SECRET"))
// 	fmt.Println(apiClient.GetBalance())
// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	// w.Header().Set("Access-Control-Allow-Headers","")
// 	testPerson1 := JsonTestData{"amano", 3}
// 	res, _ := json.Marshal(testPerson1)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(res)
// }

// func getHelloHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	w.Header().Set("Content-Type", "application/json")
// 	hello := "konnnitiha"
// 	fmt.Println("文字列は: %s", hello)
// 	helloRes, _ := json.Marshal(hello)
// 	fmt.Println("返却されるMarshalされた文字列は: %s", helloRes)
// 	w.Write(helloRes)
// }

// func main() {

// 	http.HandleFunc("/xem", getXemHandler)
// 	http.HandleFunc("/hello", getHelloHandler)
// 	// http.HandlerFunc("/hello", getHelloHandler)
// 	http.ListenAndServe(":8080", nil)

// 	// fmt.Println(yahoo.GetYahooFinance)
// }

// var DbConnection *sql.DB

// type Person struct {
// 	Name string
// 	Age  int
// }

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h2>%s</h2><div>%s</div>", p.Title, p.Body)
}

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}

func main() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)

	Server(apiKey, sign, data)

	// http.HandleFunc("/view/", viewHandler)
	// http.HandleFunc("/test/", testHandler)
	// log.Fatalln(http.ListenAndServe(":8080", nil))

	// p1 := &Page{Title: "test", Body: []byte("This is a first page.")}
	// p1.save()

	// p2, _ := loadPage(p1.Title)
	// fmt.Println(string(p2.Body))

	// DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	// defer DbConnection.Close()
	// cmd := `CREATE TABLE IF NOT EXISTS person(
	// 		name STRING,
	// 		age INT)`
	// _, err := DbConnection.Exec(cmd)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = DbConnection.Exec(cmd, "Nancy", 12)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, 32, "Mike")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// cmd = "SELECT * FROM person"
	// rows, _ := DbConnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// cmd = "SELECT * FROM person where age = ?"
	// row := DbConnection.QueryRow(cmd, 32)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// cmd = "DELETE FROM person WHERE name = ?"
	// _, err = DbConnection.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

}
