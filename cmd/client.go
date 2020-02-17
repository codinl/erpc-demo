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
	for {
		fmt.Println("请输入信息：")
		var content string
		fmt.Scanln(&content)
		sendMsg(content)
	}
}

func sendMsg(content string) {
	//fmt.Println("sendMsg --",  content)
	defer erpc.SetLoggerLevel("ERROR")()

	cli := erpc.NewPeer(erpc.PeerConfig{RedialTimes: -1})
	defer cli.Close()
	cli.SetTLSConfig(erpc.GenerateTLSConfigForClient())

	cli.RoutePush(new(Server))

	sess, stat := cli.Dial(":9090")
	if !stat.OK() {
		erpc.Fatalf("%v", stat)
	}

	//var result int
	stat = sess.Push("/server/msg",
		content,
	)

	//if !stat.OK() {
	//	erpc.Fatalf("%v", stat)
	//}
	//erpc.Printf("result: %d", result)
	//erpc.Printf("Wait 10 seconds to receive the push...")
	//time.Sleep(time.Second * 10)
}

// Client handler
type Client struct {
	erpc.PushCtx
}

// Add handles msg request
func (m *Client) Msg(arg *string) *erpc.Status {
	erpc.Printf("%s\n", *arg)
	return nil
}

func init() {
	rootCmd.AddCommand(clientCmd)
}
