package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":5000")
	if err != nil {
		fmt.Print("dialing:", err)
	}

	var reply bool

	err = client.Call("RemoteList.Create", 0, &reply)
	if err != nil {
		fmt.Println("Error:", err)
	}

	if reply {
		fmt.Println("Lista criada com sucesso")
	} else {
		fmt.Println("A lista já existe")
	}

	



	// Synchronous call
	
	var reply_i int
	err = client.Call("RemoteList.Append", []int{0,10}, &reply)
	err = client.Call("RemoteList.Append", []int{0,20}, &reply)
	err = client.Call("RemoteList.Append", []int{0,30}, &reply)
	err = client.Call("RemoteList.Append", []int{0,40}, &reply)
	err = client.Call("RemoteList.Append", []int{0,50}, &reply)

	err = client.Call("RemoteList.Remove", 0, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", reply_i)
	}
	err = client.Call("RemoteList.Remove", 0, &reply_i)
	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", reply_i)
	}

	err = client.Call("RemoteList.SaveFile", 0, &reply_i)

	if err != nil {
		fmt.Print("Error:", err)
	} else {
		fmt.Println("Elemento retirado:", reply_i)
	}

	err = client.Call("RemoteList.Get", []int{0,0}, &reply_i)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Elemento obtido:", reply_i)
	}
	err = client.Call("RemoteList.Size", 0, &reply_i)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Tamanho da Lista:", reply_i)
	}

	err = client.Call("RemoteList.Create", 5, &reply)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
