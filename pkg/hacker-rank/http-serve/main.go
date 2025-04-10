package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
 * Complete the 'postHandler', 'deleteHandler' and 'getHandler' functions below.
 *
 * All functions are expected to be void.
 * All functions accept http.ResponseWriter w and *http.Request req as parameters.
 */

var lakes []Lake

func postHandler(w http.ResponseWriter, req *http.Request) {
	// read lake from body
	var lake Lake
	_ = json.NewDecoder(req.Body).Decode(&lake)

	lakes = append(lakes, lake)
	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, req *http.Request) {

}

func getHandler(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	for _, lake := range lakes {
		if lake.Id == id {
			json.NewEncoder(w).Encode(lake)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	stdout, err := os.Create("./env/log")
	checkError(err)

	lakes = make([]Lake, 0)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/delete", deleteHandler)
	go http.ListenAndServe(portSuffix, nil)
	time.Sleep(100 * time.Millisecond)

	actions := []string{
		`{"type":"post", "payload":"{\"id\":\"id0000\", \"name\":\"Malawi\", \"area\":29500}"}`,
		`{"type":"post", "payload":"{\"id\":\"id0001\", \"name\":\"Great Bear Lake\", \"area\":31000}"}`,
		`{"type":"get", "payload":"id0001"}`,
		`{"type":"get", "payload":"id0000"}`,
	}

	for _, actionStr := range actions {
		var action Action
		err := json.Unmarshal([]byte(actionStr), &action)
		checkError(err)
		switch action.Type {
		case "post":
			_, err := http.Post(address+"/post", "application/json", strings.NewReader(action.Payload))
			checkError(err)
		case "delete":
			client := &http.Client{}
			req, err := http.NewRequest("DELETE", address+"/delete?id="+action.Payload, nil)
			checkError(err)
			resp, err := client.Do(req)
			checkError(err)
			if resp.StatusCode != 200 {
				fmt.Fprintf(writer, "%s\n", resp.Status)
				continue
			}
		case "get":
			resp, err := http.Get(address + "/get?id=" + action.Payload)
			checkError(err)
			if resp.StatusCode != 200 {
				fmt.Fprintf(writer, "%s\n", resp.Status)
				continue
			}
			var lake Lake
			err = json.NewDecoder(resp.Body).Decode(&lake)
			checkError(err)
			fmt.Fprintf(writer, "%s\n", lake.Name)
			fmt.Fprintf(writer, "%d\n", lake.Area)
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

const portSuffix = ":3333"

var address = "http://127.0.0.1" + portSuffix

type Lake struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Area int32  `json:"area"`
}

type Action struct {
	Type    string
	Payload string
}

var store = map[string]Lake{}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
