package helper

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func LogFile(filename string) {
	if filename != "" {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
		if err != nil {
			IfPanic(err)
		}
		log.SetOutput(file)
	}
}

type LoggerStruct struct {
	http.ResponseWriter
	statusCode int
}

func (l *LoggerStruct) WriteHeader(code int) {
	l.ResponseWriter.WriteHeader(code)
	l.statusCode = code
}

func Logging(out io.Writer, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := log.New(out, "", 0)
		addr := r.RemoteAddr
		endpointResponse := &LoggerStruct{ResponseWriter: w}
		if i := strings.LastIndex(addr, ":"); i != -1 {
			addr = addr[:i]
		}
		h.ServeHTTP(endpointResponse, r)
		logger.Printf("%s -- [%s] %q %d %q",
			addr,
			time.Now().Format("02/Jan/2006:15:04:05"),
			fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
			endpointResponse.statusCode,
			r.UserAgent(),
		)
		log.Printf("%s - -  %q %d %q %q \n",
			addr,
			fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
			endpointResponse.statusCode,
			r.Referer(),
			r.UserAgent())

	})
}

// func Logger(out io.Writer, h http.Handler) http.Handler {
// 	logger := log.New(out, "", 0)
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		o := &responseObserver{ResponseWriter: w}
// 		h.ServeHTTP(o, r)
// 		addr := r.RemoteAddr
// 		if i := strings.LastIndex(addr, ":"); i != -1 {
// 			addr = addr[:i]
// 		}
// 		logger.Printf("%s - - [%s] %q %d %d %q %q",
// 			addr,
// 			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
// 			fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
// 			o.status,
// 			o.written,
// 			r.Referer(),
// 			r.UserAgent())

// 		log.Printf("%s - -  %q %d %d %q %q \n",
// 			addr,
// 			fmt.Sprintf("%s %s %s", r.Method, r.URL, r.Proto),
// 			o.status,
// 			o.written,
// 			r.Referer(),
// 			r.UserAgent())
// 	})

// }

// type responseObserver struct {
// 	http.ResponseWriter
// 	status      int
// 	written     int64
// 	wroteHeader bool
// }

// func (o *responseObserver) Write(p []byte) (n int, err error) {
// 	if !o.wroteHeader {
// 		o.WriteHeader(http.StatusOK)
// 	}
// 	n, err = o.ResponseWriter.Write(p)
// 	o.written += int64(n)
// 	return
// }

// func (o *responseObserver) WriteHeader(code int) {
// 	o.ResponseWriter.WriteHeader(code)
// 	if o.wroteHeader {
// 		return
// 	}
// 	o.wroteHeader = true
// 	o.status = code
// }
