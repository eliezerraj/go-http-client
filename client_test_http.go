package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "time"
   "bytes"
   "encoding/json"
   "flag"
   "os"
   "strconv"

)

type Balance struct {
	Id					string		`json:"balance_id"`
    Account 			string 		`json:"account"`
	Amount				int32 		`json:"amount"`
    DateBalance  		time.Time 	`json:"date_balance"`
	Description			string 		`json:"description"`
}

func main() {
	port := flag.String("port","","")
	flag.Parse()
	flag.VisitAll(func (f *flag.Flag) {
		if f.Value.String()=="" {
			log.Printf("A flag -%v não foi informado \n", f.Name )
			os.Exit(1)
		}
	})

	//var host = "a344da4888686401bab959ee6d0f98e7-304745981.us-east-2.elb.amazonaws.com" + ":" + *port 
	var host = "127.0.0.1:" + *port
	
	client := http.Client{}
	done := make(chan string)

	log.Println("-----------------------------")
	log.Println("Goroutine - POST data")
	go post(host, client, done)
	/*go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)

	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)

	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)

	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)

	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)
	go post(host, client, done)*/
	
	log.Println("End POST data")
	log.Println("-----------------------------")

	//log.Println("-----------------------------")
	//log.Println("Goroutine - GET Data")
	//go get(host, client, done)
	//log.Println("End GET Data")
	//log.Println("-----------------------------")

	log.Println(<-done)
}

func get(host string, client http.Client, done chan string){
	for i:=0; i < 3600; i++ {
		host_url := "http://" + host + "/list_balance"
		get_data(host_url, client)
		time.Sleep(time.Millisecond * time.Duration(2000))
	}
	done <- "END"
}

func post(host string, client http.Client, done chan string){
	for a:=0; a < 3600; a++ {
		for i:=0; i < 50; i++ {
			host_url := "http://" + host + "/balance/save"
			post_data(i, host_url, client)
			time.Sleep(time.Millisecond * time.Duration(2000))
		}
		//time.Sleep(time.Millisecond * time.Duration(1))
	}
	done <- "END"
}

func get_data(host_url string, client http.Client){
	log.Println("GET DATA.....")
	
	req_get , err := http.NewRequest("GET", host_url, nil)
	if err != nil {
		log.Println("Error http.NewRequest : ", err)
		panic(err)
	}

	req_get.Header = http.Header{
		"Accept_Language": []string{"pt-BR"},
		"jwt": []string{"cookie"},
	}

	req_get.Close = true
	resp, err := client.Do(req_get)
	defer resp.Body.Close()

	if err != nil {
		log.Println("Error doing GET : ", err)
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error : ", err)
		panic(err)
	}

	sb := string(body)
	log.Printf(sb)

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("Erro Close :",err)
		}
	}()

	log.Println("###########")
}

func post_data(i int ,host string, client http.Client){
	log.Println("POST DATA.....")

	balance := NewBalance(i)
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(balance)

	req_post , err := http.NewRequest("POST", host, payload)
	if err != nil {
		log.Println("Error http.NewRequest : ", err)
		panic(err)
	}

	req_post.Header = http.Header{
		"Accept_Language": []string{"pt-BR"},
		"jwt": []string{"cookie"},
		"Content-Type": []string{"application/json"},
	}

	resp, err := client.Do(req_post)
	defer resp.Body.Close()
	if err != nil {
	   log.Println("Error doing POST : ", err)
	   panic(err)
	}

	log.Println("StatusCode : ", resp.StatusCode )
	
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("Erro Close :",err)
		}
	}()
}

func NewBalance(i int) Balance{
	acc := "acc-" + strconv.Itoa(i)
	description := "COOKIE-"+ strconv.Itoa(i) + " - OK"
	
	balance := Balance{
		Id:    strconv.Itoa(i),
		Account: acc,
		Amount: 1,
		DateBalance: time.Now(),
		Description: description,
	}
	return balance
 }