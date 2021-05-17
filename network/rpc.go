package network

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"

	"github.com/EclesioMeloJunior/distributed/node"
)

const (
	PING      = "PING"
	STORE     = "STORE"
	FINDNODE  = "FIND-NODE"
	FINDVALUE = "FIND-VALUE"
)

type (
	EmptyRequest byte

	Server struct {
		node *node.RPC
		done chan bool
	}
)

func (s *Server) Start(ip string, port int) error {
	if err := rpc.Register(s.node); err != nil {
		return err
	}

	go func() {
		rpc.HandleHTTP()
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "server alive")
		})

		err := http.ListenAndServe(fmt.Sprintf("%s:%v", ip, port), nil)
		if err != nil {
			log.Println(err)
			s.done <- true
		}
	}()

	<-s.done
	log.Println("Shutdown server")

	return nil
}

func NewServer(noderpc *node.RPC) (*Server, error) {
	doneCh := make(chan bool, 1)

	return &Server{
		node: noderpc,
		done: doneCh,
	}, nil
}
