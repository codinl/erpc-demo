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
		//PrintDetail: true,
	})
	srv.SetTLSConfig(erpc.GenerateTLSConfigForServer())

	// router
	srv.RoutePush(new(Server))

	// listen and serve
	srv.ListenAndServe()
}

// Server push handler
type Server struct {
	erpc.PushCtx
}

// Server handles '/server/msg' message
func (p *Server) Msg(arg *string) *erpc.Status {
	fmt.Printf("--------%s\n", *arg)
	return nil
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
