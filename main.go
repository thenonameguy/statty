package main

import (
	"fmt"
  "github.com/jondot/gosigar"
  "encoding/json"
  "net/http"
  "flag"
)

type InfoHolder struct{
  Uptime sigar.Uptime
  Cpus sigar.CpuList
  FileSystems sigar.FileSystemList
  Swap sigar.Swap
  Mem sigar.Mem
}

func  NewInfoHolder() *InfoHolder{
  i:=new(InfoHolder)
  i.Uptime=sigar.Uptime{}
  i.Cpus=sigar.CpuList{}
  i.FileSystems=sigar.FileSystemList{}
  i.Swap=sigar.Swap{}
  i.Mem=sigar.Mem{}
  return i
}

func (i *InfoHolder)Update(){
  i.Uptime.Get()
  i.Cpus.Get()
  i.FileSystems.Get()
  i.Swap.Get()
  i.Mem.Get()
}

func (i *InfoHolder) String() string{
  b,err:=json.Marshal(i)
  if err!=nil{
    return ""
  }
  return string(b)
}

func DataServer(w http.ResponseWriter,req *http.Request){
    w.Header().Set("Content-Type", "application/json")
    info.Update()
    fmt.Fprint(w,info)
}

var info=NewInfoHolder()
var port string

func main() {
  flag.StringVar(&port,"port",":8080","Statty's port address")
  flag.Parse()
  fmt.Println("Listening on port",port)
  http.Handle("/", http.FileServer(http.Dir("./web/")))
  http.HandleFunc("/data",DataServer)
  http.ListenAndServe(port,nil)
}
