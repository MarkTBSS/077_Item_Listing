package main

import (
	"github.com/MarkTBSS/077_Item_Listing/config"
	"github.com/MarkTBSS/077_Item_Listing/databases"
	"github.com/MarkTBSS/077_Item_Listing/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
