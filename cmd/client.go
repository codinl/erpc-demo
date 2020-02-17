package cmd

import (
	"fmt"
	"github.com/henrylee2cn/erpc/v6"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client_start",
	Short: "A brief description of your command",
	Long:  `A longer description.`,
	Run: func(cmd *cobra.Command, args []string) {
		ClientStart()
	},
}

func ClientStart() {
	//var addr string
	//fmt.Println(`请输入服务器端地址(默认是":9090")：`)
	//if _, err := fmt.Scanln(&addr); err != nil {
	//	erpc.Fatalf("", err)
	//}
	//initClient(addr)
	initClient(":9090")

	for {
		fmt.Println("向服务端发送信息：")
		var content string
		fmt.Scanln(&content)
		sendMsg(content, SERVER_TOPIC)
	}
}

var clientSession erpc.Session
var clientPeer erpc.Peer

func initClient(addr string) {
	defer erpc.SetLoggerLevel("ERROR")()

	clientPeer = erpc.NewPeer(erpc.PeerConfig{RedialTimes: -1})
	//defer cli.Close()
	clientPeer.SetTLSConfig(erpc.GenerateTLSConfigForClient())

	clientPeer.RoutePush(new(Client))

	var stat *erpc.Status
	clientSession, stat = clientPeer.Dial(addr)
	if !stat.OK() {
		erpc.Fatalf("%v", stat)
	}
}

func sendMsg(content, topic string) {
	//defer erpc.SetLoggerLevel("ERROR")()
	//
	//cli := erpc.NewPeer(erpc.PeerConfig{RedialTimes: -1})
	//defer cli.Close()
	//cli.SetTLSConfig(erpc.GenerateTLSConfigForClient())
	//
	//cli.RoutePush(new(Client))
	//
	//var stat *erpc.Status
	//clientSession, stat = cli.Dial(":9090")
	//if !stat.OK() {
	//	erpc.Fatalf("%v", stat)
	//}

	//var result int
	stat := clientSession.Push(
		topic,
		content,
	)
	if !stat.OK() {
		erpc.Fatalf("%v", stat)
	}

	//erpc.Printf("result: %d", result)
	//erpc.Printf("Wait 10 seconds to receive the push...")
	//time.Sleep(time.Second * 10)
}

const CLIENT_TOPIC = "/client/msg"

// Client handler
type Client struct {
	erpc.PushCtx
}

// Add handles msg request
func (m *Client) Msg(arg *string) *erpc.Status {
	erpc.Printf("-------------收到服务器端信息：%s\n", *arg)
	return nil
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
