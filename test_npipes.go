package main

import(
	"fmt"
    "net/http"
	"github.com/rpc"
	"errors"
	"log"
    "github.com/natefinch/npipe"
)
		type Args struct {
			A, B int
		}

		type Quotient struct {
			Quo, Rem int
		}

		type Arith int

		func (t *Arith) Multiply(args *Args, reply *int) error {
			*reply = args.A * args.B
			return nil
		}

		func (t *Arith) Divide(args *Args, quo *Quotient) error {
			if args.B == 0 {
				return errors.New("divide by zero")
			}
			quo.Quo = args.A / args.B
			quo.Rem = args.A % args.B
			return nil
		}

func main(){
        var pipe_path = "\\\\.\\pipe\\test"
	    
		arith := new(Arith)
		rpc.Register(arith)
		rpc.HandleHTTP()
		l, e := npipe.Listen(pipe_path)
		if e != nil {
			log.Fatal("listen error:", e)
		}
		go http.Serve(l, nil)

        client, err := rpc.DialHTTP("pipe", pipe_path)
		if err != nil {
			log.Fatal("dialing:", err)
		}
        
		// Synchronous call
		args := &Args{7,8}
		var reply int
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		}
		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}


