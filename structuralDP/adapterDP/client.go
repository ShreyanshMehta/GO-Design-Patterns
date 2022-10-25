package main

import "fmt"

type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com IComputer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
