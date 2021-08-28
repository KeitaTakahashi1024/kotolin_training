package main

import(
	//標準のやつ
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
	"io"

	// mysql関係のライブラリ
	 "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/* ユーザー情報の構造体 */
type User struct {
	Name string //名前
	Token string //トークン
}

var user User;

type AllUser struct {
	ID int //id
	Name string //名前
	Token string //トークン
}

var allUser AllUser;
var allUsers []AllUser

func main() {
	createUser()
	http.HandleFunc("/user/get", getUser)
	http.HandleFunc("/user/update", updateUser)
	http.ListenAndServe(":8080", nil)
}

/* /user/createにおける処理を行う関数 */
func createUser() {
	// POSTリクエストからユーザーの名前を抽出する
	http.HandleFunc("/user/create", getJsonRequest)

}

/* POSTリクエストのJson形式からユーザーの名前を抽出する関数 */
func getJsonRequest(w http.ResponseWriter, req *http.Request) {
	// リクエストがPOSTか確認
	if req.Method != "POST" { //POST以外だったら...
		w.WriteHeader(http.StatusBadRequest) //400:bad requestにする
		return
	}

	// リクエストのコンテンツタイプがjson確認
	if req.Header.Get("Content-Type") != "application/json" { //json以外だったら
		w.WriteHeader(http.StatusBadRequest) //400:bad requestにする
		return
	}

	// コンテンツの長さをstringからintに変換してlengthに格納する
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// bodyという配列を作って、リクエストされたjsonのデータを格納する
	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// jsonBodyに、bodyに格納されているjsonデータをパースして格納する
	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user.Name = jsonBody["name"].(string)
	user.Token = "0614";

	// ユーザーの名前をmysqlに追加する
	insertMysql()
	// MySQLに追加されたデータを出力する
	outputMysql()

	jsonData, err := json.Marshal(allUsers)
	if err != nil{
		panic(err.Error())
	}
	fmt.Fprintln(w, string(jsonData))

	w.WriteHeader(http.StatusOK)
}

/* MySQLにデータを挿入する関数 */
func insertMysql() {
	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql_users")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmtInsert, err := db.Prepare("INSERT INTO users(id, name, token) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtInsert.Close()

	result, err := stmtInsert.Exec(0, user.Name, user.Token)
	if err != nil {
		panic(err.Error())
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Print(lastInsertID)
}

func outputMysql() {

	db, err := sql.Open("mysql", "root:password@tcp(mysql:3306)/mysql_users")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, token FROM users")
	if err != nil{
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next(){
		if err := rows.Scan(&allUser.ID,&allUser.Name, &allUser.Token);
		err != nil{
			panic(err.Error())
		}
		allUsers = append(allUsers, allUser)
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to user/get, from Docker container!")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to user/update, from Docker container!")
}
