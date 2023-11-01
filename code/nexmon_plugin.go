package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
)

import (
	"github.com/gorilla/mux"
)

var UNIX_PLUGIN_LISTENER = "/state/plugins/nexmon/socket"

func pluginTest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", 400)
}

func changeChannel(w http.ResponseWriter, r *http.Request) {
	channel := r.URL.Query().Get("channel")

	// Use regexp.MatchString to check if the input matches the pattern
	matches, err := regexp.MatchString("^[0-9/]*$", channel)
	if err != nil || !matches {
		http.Error(w, "Invalid channel string", 400)
		return
	}

	err = exec.Command("/nexmon/nexutil", "-k"+channel).Run()
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	unix_plugin_router := mux.NewRouter().StrictSlash(true)

	unix_plugin_router.HandleFunc("/change_channel", changeChannel).Methods("PUT")

	os.Remove(UNIX_PLUGIN_LISTENER)
	unixPluginListener, err := net.Listen("unix", UNIX_PLUGIN_LISTENER)
	if err != nil {
		panic(err)
	}

	pluginServer := http.Server{Handler: logRequest(unix_plugin_router)}

	pluginServer.Serve(unixPluginListener)
}
