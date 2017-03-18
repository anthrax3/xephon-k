package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xephonhq/xephon-k/pkg/storage/cassandra"
)

var SchemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Create schema",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("create schema for cassandra using default setting")
		cassandra.CreateSchema()
		log.Info("schema created!")
	},
}

func init() {
	RootCmd.AddCommand(SchemaCmd)
}