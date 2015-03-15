package main

import (
	// "encoding/json"
	"fmt"
	elasticgo "github.com/mattbaird/elastigo/lib"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var (
// host *string = flag.Stringt(name, value, usage)
)

func main() {

	var rootCmd = &cobra.Command{
		Use:   "es-admin",
		Short: "Elasticsearch admin command line interface",
	}

	flDebug := rootCmd.PersistentFlags().BoolP("debug", "D", false, "Enable debug mode.")
	flQuiet := rootCmd.PersistentFlags().BoolP("quiet", "q", true, "Disable output column headers.")

	flHost := rootCmd.PersistentFlags().StringP("host", "H", "", "ES_HOST to connect to. Default: localhost")
	flPort := rootCmd.PersistentFlags().StringP("port", "P", "", "ES_PORT to connect on. Defaulf: 9200")
	flAuth := rootCmd.PersistentFlags().StringP("auth", "A", "", "ES_AUTH username:password for authentication.")

	var healthCmd = &cobra.Command{
		Use:   "health",
		Short: "Print health of cluster",
		Run: func(cmd *cobra.Command, args []string) {
			parseDefaultFlags(*flDebug, *flQuiet, *flHost, *flPort, *flAuth)
			client := connect()

			resp, err := client.Health()
			if err != nil {
				fmt.Printf("ERR: %v", err)
				return
			}

			PrintHealth(resp)
		},
	}

	rootCmd.AddCommand(healthCmd)

	rootCmd.Execute()
}

func parseDefaultFlags(flDebug, flQuiet bool, flHost, flPort, flAuth string) {

	if flDebug {
		os.Setenv("DEBUG", "1")
	}

	if flQuiet {
		os.Setenv("QUIET", "1")
	}

	if flHost != "" {
		os.Setenv("ES_HOST", flHost)
	}

	if flPort != "" {
		os.Setenv("ES_PORT", flPort)
	}

	if flAuth != "" {
		os.Setenv("ES_AUTH", flAuth)
	}
}

func connect() *elasticgo.Conn {
	conn := elasticgo.NewConn()

	if os.Getenv("ES_HOST") != "" {
		conn.Domain = os.Getenv("ES_HOST")
	}

	if os.Getenv("ES_PORT") != "" {
		conn.Port = os.Getenv("ES_PORT")
	}

	if os.Getenv("ES_AUTH") != "" {
		userInfo := strings.Split(os.Getenv("ES_AUTH"), ":")
		if len(userInfo) == 2 {
			conn.Username = userInfo[0]
			conn.Password = userInfo[1]
		}
	}

	return conn
}
