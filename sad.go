package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Info struct {
	Alias string
	ID    string
	Sex   string
	Quote string
	Phone string
}

const (
	cook = "jsaknfjfjkasdflkajffadsjkladsf"
)

var baseHtml = `

{{ define "base" }}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title> DarkEye </title>


    <style>
        body {
        	background-color:black;
        	color:red;
        	text-align:center; 
        	padding: 50px;
        	font-family:monospace;
            font-size: 15px;
        }

        input[type=text] {
        	color:red;
        	outline:off;
        	padding: 1px solid red;
        	padding: 5px;
        }
     </style>
</head>
<body>
    
</body>
</html>

{{ end }}

`

var start = `

{{ template "base" }}

welcome

<h3> if a giraffe has two eyes, a monkey has two eyes,<br> and an elephant has two eyes, how many eyes do we have? </h3>
<br><br>
<h3> What's the Answer </h3>
<form action="index", method="post" autocomplete="off">
<input type="text", name="res" placeholder="answer">
<input type="submit" value="submit">
</form>

`

var page2 = `

{{ template "base" }}

    <div>

        <h1> EFCC DarkEye </h1><hr><br>
        <h4> Welcome To EFCC DarkEye </h1>

        <p> You are seen this page, Because you 
            solve The previous puzzle or you are a
            member of sadsec(a sec org started by
            a group of researchers and students)
        </p><br>
        
        <p> DarkEye is a secret program 
            created by the Nigeria EFCC to
            Employ 10 Random people
            Around The World.
         </p>
        <br>

        <p>
            Member of the Sadsec org are giving the 
            privilege to this secret program
            without having to solve a puzzle 
            was due to their phenomenal contribution
            to the Country National Security in the
            past.   
        </p><br><br>
            

        <h3> ANSWER IS IN HERE </H3>
		<p>
		<p>Hint</p>
		the HTTP headers are used to pass additional infomation between the clients and the server
		through the request and response header
		</p>

		<h3> Solution </h3>
		<form action="sadsec" method="post" autocomplete="off">

		<input type="text"  name="enter" placeholdee="flag">
		<input type="submit"  value="submit">

		</form>

    </div>

`

// secret data rendered
var tmpl = `

{{ $r := . }}


    {{ $r.cne3rd.Alias}}
    {{ $r.cne3rd.ID}}
    {{ $r.cne3rd.Sex}}
    {{ $r.cne3rd.Quote}}
    {{ $r.cne3rd.Phone}}



    {{ $r.abdulconsol3_.Alias}}
    {{ $r.abdulconsol3_.ID}}
    {{ $r.abdulconsol3_.Sex}}
    {{ $r.abdulconsol3_.Quote}}
    {{ $r.abdulconsol3_.Phone}}



    {{ $r.devsammy.Alias}}
    {{ $r.devsammy.ID}}
    {{ $r.devsammy.Sex}}
    {{ $r.devsammy.Quote}}
    {{ $r.devsammy.Phone}}



    {{ $r.devfemibadmus.Alias}}
    {{ $r.devfemibadmus.ID}}
    {{ $r.devfemibadmus.Sex}}
    {{ $r.devfemibadmus.Quote}}
    {{ $r.devfemibadmus.Phone}}



	{{ $r.md.Alias}}
    {{ $r.md.ID}}
    {{ $r.md.Sex}}
    {{ $r.md.Quote}}
    {{ $r.md.Phone}}



`

// secret data
var cloud9 = map[string]Info{
	"cne3rd": Info{
		Alias: "Cne3Rd",
		ID:    "007",
		Sex:   "Male",
		Quote: "Talk is cheap show me the code",
		Phone: "01077777",
	},

	"abdulconsol3_": Info{
		Alias: "Abdul Console",
		ID:    "006",
		Sex:   "Male",
		Quote: "There is no security Just levels of Difficultues to avoid",
		Phone: "01066666",
	},

	"devsammy": Info{
		Alias: "dev sammy",
		ID:    "005",
		Sex:   "Male",
		Quote: "some of the best programming is done on paper, really. putting it into the computer is just a minor detail",
		Phone: "01055555",
	},

	"devfemibadmus": Info{
		Alias: "dev FemiBadmus",
		ID:    "004",
		Sex:   "Male",
		Quote: "Testing can only prove the presence of bugs not their absence",
		Phone: "01044444",
	},

	"md": Info{
		Alias: "md",
		ID:    "003",
		Sex:   "Male",
		Quote: "laughing at your security",
		Phone: "01033333",
	},
}

func main() {

	server()

}

func begin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
	}

	http.Redirect(w, r, "/index/", 301)
	return
}

// stage 1
func index(w http.ResponseWriter, r *http.Request) {

    id := strings.SplitN(r.URL.Path, "/", 3)[2]
    _, ok := cloud9[id]
    
    if ok{
    	cookie := &http.Cookie{Name: "token", Value: cook, Path: "/sadsec/"}
    	http.SetCookie(w, cookie)
        http.Redirect(w, r, "/sadsec/", 301)
        return
    }
    
	if r.Method == http.MethodPost {
		out := r.FormValue("res")
		//fmt.Println(out)
		if out == "4"{
			cookie := &http.Cookie{Name: "token", Value: cook, Path: "/sadsec/"}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/sadsec/", 301)
			return
		}

		http.Redirect(w, r, "/", 301)
		fmt.Println("wrong answer try again")

	}

	if r.Method == http.MethodGet {
		ts, err := template.New("start").Parse(start)
		ts, err = ts.Parse(baseHtml)
		if err != nil {
			fmt.Println(err)
		}

		ts.Execute(w, nil)
	}

}

// stage two
func home(w http.ResponseWriter, r *http.Request) {

	r.Header.Set("flag", "You Are 1337")
	val := r.FormValue("enter")
	if r.Method == http.MethodPost {
		flag := r.Header.Get("flag")

		if val == flag {
			http.Redirect(w, r, "/sec/", 301)
			return
		}

		http.Redirect(w, r, "/sadsec", 301)
		fmt.Println("i guess you are missing The header")

	}

		c, err := r.Cookie("token")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		if c.Value != cook {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		ts, err := template.New("page2").Parse(page2)
		ts, err = ts.Parse(baseHtml)
		if err != nil {
			fmt.Println(err)
			return
		}

		ts.Execute(w, nil)

}

// stage 3
func secret(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are The Flag")
	id := strings.SplitN(r.URL.Path, "/", 3)[2]
	_, ok := cloud9[id]
	if ok {

		if r.Header.Get("whoami") == id {

			ts, err := template.New("templ").Parse(tmpl)
			if err != nil {
				panic(err)
			}

			ts.Execute(w, cloud9)
			return

		}

	}

	w.Write([]byte("Oops, Only Real Sadsec Member can Acces this"))
}

var sect = `

{{ template "base" }}

<h2> bypass the captcha </h2> </br>

<h3> are you a robot </h3>


<h1> {{ . }} </h1><br>

<h4> Hint </h4>

<p> The HTTP 200 OK success status response code indicates that the request has succeeded </p>

<form action="sec" method="post" autocomplete="off">
    <input type="text" name="sec"><br>
	<input type="submit" value="send">
</form>


`

// bypass the captcha
func sec(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		fmt.Println("Try The Path")
		http.Redirect(w, r, "/awesome/", 301)
		return

	}

	if r.URL.Path == "/sec/200/" {
		cookie := &http.Cookie{Name: "awesome", Value: "end", Path: "/awesome/"}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/awesome/", 301)
		return
	}

	make := make([]byte, 32)

	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	rand.Read(make)

	capt := sha256.Sum256(make)

	var newrand []byte
	for _, v := range capt {
		newrand = append(newrand, v)
	}
	h := hex.EncodeToString(newrand)
	h = h[:5]

	ts, err := template.New("sec").Parse(sect)
	if err != nil {
		fmt.Println(err)
	}
	ts.Parse(baseHtml)

	err = ts.Execute(w, h)
	if err != nil {
		fmt.Println(err)
	}
}

func awesome(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("awesome")
	if err != nil {
		http.Redirect(w, r, "/sec", 301)
		return
	}

	if cookie.Value == "end" {
		w.Write([]byte("say, you are awesome"))
		return
	}

	http.Redirect(w, r, "/sec", 301)
	return

}

func server() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", begin)
	mux.HandleFunc("/index/", index)
	mux.HandleFunc("index/*", index)
	mux.HandleFunc("/sadsec/", home)
	mux.HandleFunc("/secret/", secret)
	mux.HandleFunc("/sec/", sec)
	mux.HandleFunc("/awesome/", awesome)

	fmt.Println("listening on port 4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		panic(err)
	}
}
