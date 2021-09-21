package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	pingProto "github.com/shubhamagarwal19/protobufWithRest/proto"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", Ping)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}

func Ping(resp http.ResponseWriter, req *http.Request) {
	log.Println("Received PING from client")
	request := &pingProto.Ping{}
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	req.Body.Close()
	if err = proto.Unmarshal(data, request); err != nil {
		log.Fatal(err)
	}

	response := &pingProto.Pong{Response: "PONG"}
	res, err := proto.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	resp.Write(res)
	return
}
