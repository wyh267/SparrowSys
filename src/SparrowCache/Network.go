package SparrowCache

import (
	"bufio"
	"fmt"
	"net"
	"utils"
)

type Network struct {
	workerChans []chan *net.TCPConn
	workerNum   int
	port        int
    cache       *SCache
	Logger      *utils.Log4FE
}

func NewNetwork(workernum int, mport int, logger *utils.Log4FE) *Network {

	this := &Network{workerNum: workernum, port: mport, Logger: logger, workerChans: make([]chan *net.TCPConn, workernum)}

	for idx := range this.workerChans {
		this.workerChans[idx] = make(chan *net.TCPConn, 100)
	}
    
    this.cache=NewSCache()

	return this

}

// Start function description : 启动服务器
// params :
// return :
func (this *Network) Start() error {
    
    var tcpAddr *net.TCPAddr

    tcpAddr, _ = net.ResolveTCPAddr("tcp", fmt.Sprintf(":%v", this.port))

	//绑定端口
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		this.Logger.Error("[ERROR] listen port [%v] fail ... %v", this.port, err)
		return err
	}

	//启动处理协程
	for i := 0; i < this.workerNum; i++ {

		go this.runWorker(fmt.Sprintf("worker_%v", i),i)

	}

	this.Logger.Info("[INFO] listen port[%v] , waitting for connection ", this.port)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			this.Logger.Error("[ERROR] Accept port [%v] fail ... %v", this.port, err)
			return err
		}
        
        select {
            case this.workerChans[0] <- conn :
            default :
                this.Logger.Error("[ERROR] Not Resouce ")
            
        }
        

	}
}

func (this *Network) runWorker(name string, num int) {

	for {

		select {
		case conn := <-this.workerChans[num]:
			requestStr, err := this.readFromConn(conn)
            if err != nil {
                continue
            }
            this.Logger.Info("[INFO] request %v",requestStr)
			//this.processRequst(requestStr)
            conn.Write([]byte(requestStr))
            this.Logger.Info("[INFO]  disconnected ")
		    conn.Close()
		default:

		}

	}

}

func (this *Network) readFromConn(conn *net.TCPConn) (string, error) {

	//ipStr := conn.RemoteAddr().String()
	//defer func() {
	//	this.Logger.Info("[INFO]  disconnected %v", ipStr)
	//	conn.Close()
	//}()
	reader := bufio.NewReader(conn)

	message, err := reader.ReadString('\n')
	if err != nil {
		this.Logger.Error("[ERROR] Read Error %v", err)
		return "", err
	}
    
    this.Logger.Info("[INFO] Message : [ %v ]",message)

    return message,nil
}
