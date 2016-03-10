package redismq

import (
  "net"
  "net/rpc"
)
type Worker {
  l net.Listener
  WorkerAddress string
}

func InitWorker(WorkerAddress string, nRPC int) *Worker{
  w := &Worker{}
  w.nRPC = nRPC
  w.WorkerAddress = WorkerAddress
  return w
}

func RunWorker(MasterAddress, WorkerAddress string, nRPC int) {
  w := InitWorker(WorkerAddress, nRPC)
  w.StartRpcServer()
  w.Register(MasterAddress, WorkerAddress string)
}

func (w *Worker) Register(MasterAddress, WorkerAddress string) {
  /*
  args := &RegisterArgs{}
	args.Worker = me
	var reply RegisterReply
	ok := call(master, "MapReduce.Register", args, &reply)
	if ok == false {
		fmt.Printf("Register: RPC %s register error\n", master)
	}
  */
  args := &RegisterArgs{}
  args.worker = WorkerAddress
  var reply RegisterReply
  ok := call(MasterAddress, "Master.Register", args, &reply)
}

func (w *Worker) StartRpcServer() {
  rpcs := rpc.NewServer()
  rpcs.Register(w)
  l, e := net.Listen("unix", w.WorkerAddress)
  if e != nil {
		log.Fatal("RunWorker: worker ", me, " error: ", e)
	}
	wk.l = l
  //add your code here

  for wk.nRPC != 0 {
		conn, err := wk.l.Accept()
		if err == nil {
			wk.nRPC -= 1
			go rpcs.ServeConn(conn)
			wk.nJobs += 1
		} else {
			break
		}
	}
	wk.l.Close()
	DPrintf("RunWorker %s exit\n", me)
}