package main

import (
    "log"
    "net/http"
    "github.com/hellstein/ssh-rpc-agent/jobmgr"
)

// our main function
func main() {
    mgr := &jobmgr.Mgr{}
    log.Fatal(http.ListenAndServe(":8000", CreateRouter(mgr)))
}