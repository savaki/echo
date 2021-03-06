package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func echo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "<html>")
	fmt.Fprintln(w, "<head>")
	fmt.Fprintln(w, `<style>

	h1, td {
		font-family: helvetica neue, helvetica, arial, sans-serif;
	}

	h1 {
		font-size: 20px;
	}

	table {
		margin-top: 30px;
		border-spacing: 0;
    border-collapse: collapse;
	}

	tr.header {
		background-color: #42a5f5;
	}

	tr {
		border-bottom: 1px solid rgba(66, 165, 245, .25);
	}

	tr:nth-child(2n) {
		background-color: rgba(66, 165, 245, .1);
	}

	td {
		padding: 8px 16px;
	}

</style>`)
	fmt.Fprintln(w, "</head>")
	fmt.Fprintln(w, "<body>")

	fmt.Fprintf(w, "<h1>%v %v</h1>\n", req.Method, req.RequestURI)

	// -- Env ------------------------------------------------

	fmt.Fprintln(w, "<table>")
	fmt.Fprintln(w, `<tr class="header"><td>Env Var</td><td>Value</td></tr>`)
	kvs := os.Environ()
	sort.Strings(kvs)
	for _, kv := range kvs {
		segments := strings.Split(kv, "=")
		fmt.Fprintf(w, "<tr><td>%v</td><td>%v</td></tr>\n", segments[0], segments[1])
	}
	fmt.Fprintln(w, "</table>")

	// -- Headers --------------------------------------------

	keys := make([]string, 0, len(req.Header))
	for key := range req.Header {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	fmt.Fprintln(w, "<table>")
	fmt.Fprintln(w, `<tr class="header"><td>Header</td><td>Value</td></tr>`)
	for _, key := range keys {
		values := req.Header[key]
		sort.Strings(values)
		fmt.Fprintf(w, "<tr><td>%v</td><td>%v</td></tr>\n", key, strings.Join(values, ", "))
	}
	fmt.Fprintln(w, "</table>")

	fmt.Fprintln(w, "</body>")
	fmt.Fprintln(w, "</html>")

	json.NewEncoder(os.Stdout).Encode(map[string]interface{}{
		"timestamp": time.Now().Format("2006-01-02T15:04:05.000000Z"),
		"method":    req.Method,
		"path":      req.URL.Path,
	})
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	port := 3000
	if v, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		port = v
	}

	http.HandleFunc("/", echo)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	check(err)
}
