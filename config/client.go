package config

import (
	"api/db" // Replace with the actual import path
)

// Client is a package-level variable to be used across files
var Client *db.PrismaClient

func init() {
	Client = db.NewClient()
	if err := Client.Prisma.Connect(); err != nil {
		panic(err)
	}
}

// CloseClient is a function to be called when you want to disconnect the client
func CloseClient() {
	if err := Client.Prisma.Disconnect(); err != nil {
		panic(err)
	}
}
