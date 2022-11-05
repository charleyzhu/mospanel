/*
 * @Author: Charley
 * @Date: 2022-11-04 14:05:41
 * @LastEditors: Charley
 * @LastEditTime: 2022-11-05 16:22:19
 * @FilePath: /mospanel/core/run.go
 * @Description: start command
 */
package cmd

import (
	"fmt"

	"github.com/charleyzhu/mospanel/configs"
	"github.com/charleyzhu/mospanel/web"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type serverFlags struct {
	c   string
	dir string
}

var sf *serverFlags

var startCmd = &cobra.Command{
	Use:   "start [-c config_file] [-d working_dir]",
	Short: "Start Service",
	RunE: func(cmd *cobra.Command, args []string) error {

		v := viper.New()
		v.SetConfigType("yaml")

		if len(sf.c) > 0 {
			v.SetConfigFile(sf.c)
		} else {
			v.SetConfigName("config")
			v.AddConfigPath(".")
		}

		if err := v.ReadInConfig(); err != nil {
			return fmt.Errorf("failed to read config file %w", err)
		}

		cfg := new(configs.Config)

		if err := v.Unmarshal(cfg); err != nil {
			return fmt.Errorf("failed to parse config file, %w", err)
		}

		server := web.NewWebServer(&cfg.Panel)
		server.Run()

		return nil
	},
	DisableFlagsInUseLine: true,
	SilenceUsage:          true,
}

func init() {
	sf = new(serverFlags)
	fs := startCmd.Flags()
	fs.StringVarP(&sf.c, "config", "c", "", "config file")
	fs.StringVarP(&sf.dir, "dir", "d", "", "working dir")
	rootCmd.AddCommand(startCmd)
}
