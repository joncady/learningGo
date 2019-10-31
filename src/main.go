package main

import (
    "log"
    "net/http"
    "os"
)

//HelloHandler handles requests for the `/hello` resource
func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hi Thomas, Web 2!\n"))
}

func main() {
    //get the value of the ADDR environment variable
    addr := os.Getenv("ADDR")

    //if it's blank, default to ":80", which means
    //listen port 80 for requests addressed to any host
    if len(addr) == 0 {
        addr = ":80"
    }

    //create a new mux (router)
    //the mux calls different functions for
    //different resource paths
    mux := http.NewServeMux()

	methmux := NewMethodMux()
	methmux.HandlerFuncs["GET"] = HelloHandler

    //tell it to call the HelloHandler() function
    //when someone requests the resource path `/hello`
    mux.Handle("/", methmux)

    //start the web server using the mux as the root handler,
    //and report any errors that occur.
    //the ListenAndServe() function will block so
    //this program will continue to run until killed
    log.Printf("server is listening at %s...", addr)
    log.Fatal(http.ListenAndServe(addr, mux))
}

//MethodMux sends the request to the function
//associated with the HTTP request method
type MethodMux struct {
    //use a map where the  key is a string (method name) 
    //and the value is the associated handler function
    HandlerFuncs map[string]func(http.ResponseWriter, *http.Request)
}

//ServeHTTP sends the request to the appropriate handler based on
//the HTTP method in the request
func (mm *MethodMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    //r.Method will be the method used in the request (GET, PUT, PATCH, POST, etc.)
    fn := mm.HandlerFuncs[r.Method]
    if fn != nil {
        fn(w, r)
    } else {
        http.Error(w, "that method is not allowed", http.StatusMethodNotAllowed)
    }
}

//NewMethodMux constructs a new MethodMux
func NewMethodMux() *MethodMux {
    return &MethodMux{
        HandlerFuncs: map[string]func(http.ResponseWriter, *http.Request){},
    }
}