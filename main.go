package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/jondot/gosigar"
	"log"
	"net/http"
)

type InfoHolder struct {
	Uptime sigar.Uptime
	Cpus   sigar.CpuList
	Swap   sigar.Swap
	Mem    sigar.Mem
}

func (i *InfoHolder) Update() {
	i.Uptime.Get()
	i.Cpus.Get()
	i.Swap.Get()
	i.Mem.Get()
}

func (i *InfoHolder) String() string {
	b, err := json.Marshal(i)
	if err != nil {
		return ""
	}
	return string(b)
}

func DataServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	info.Update()
	fmt.Fprint(w, info)
}

var info = new(InfoHolder)
var port string

func main() {
	flag.StringVar(&port, "port", ":8080", "Statty's port address")
	flag.Parse()
	fmt.Println("Listening on port", port)
	http.Handle("/", http.FileServer(http.Dir("./web/")))
	http.HandleFunc("/data", DataServer)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
