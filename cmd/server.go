package cmd

import (
	"fmt"
	"github.com/henrylee2cn/erpc/v6"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server_start",
	Short: "A brief description of your command",
	Long:  `A longer description.`,
	Run: func(cmd *cobra.Command, args []string) {
		ServerStart()
	},
}

func ServerStart() {
	defer erpc.FlushLogger()
	// graceful
	go erpc.GraceSignal()

	// server peer
	srv := erpc.NewPeer(erpc.PeerConfig{
		//CountTime:   true,
		ListenPort:  9090,
		PrintDetail: false,
	})
	srv.SetTLSConfig(erpc.GenerateTLSConfigForServer())

	// router
	srv.RoutePush(new(Server))

	go func() {
		for {
			fmt.Println("向客户端发送信息：")
			var content string
			fmt.Scanln(&content)
			srv.RangeSession(func(sess erpc.Session) bool {
				sess.Push(
					CLIENT_TOPIC,
					content,
				)
				return true
			})
		}
	}()

	// listen and serve
	srv.ListenAndServe()
}

const SERVER_TOPIC = "/server/msg"

type Server struct {
	erpc.PushCtx
}

// Add handles msg request
func (m *Server) Msg(arg *string) *erpc.Status {
	erpc.Printf("------------收到客户端的信息：%s\n", *arg)
	return nil
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
