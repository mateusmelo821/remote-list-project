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
	var reply_i, op, id, val int
	for {
		fmt.Println("\n\n1 - Criar lista")
		fmt.Println("2 - Inserir elemento na lista")
		fmt.Println("3 - Remover elemento da lista")
		fmt.Println("4 - Obter tamanho da lista")
		fmt.Println("5 - Persistir dados da lista")
		fmt.Println("6 - Obter item da lista")
		fmt.Println("0 - Encerrar aplicação")
		fmt.Println("Insira a opção desejada: ")
    	fmt.Scanln(&op)
		if op==0 {
			break
		}
		fmt.Println("Informe o id da lista:")
		fmt.Scanln(&id)
		if op==1 {
			err = client.Call("RemoteList.Create", id, &reply)
			if reply {
				fmt.Println("Lista criada com sucesso")
			} else {
				fmt.Println("A lista já existe")
			}
		}

		if op==2 {
			fmt.Println("Informe o valor a ser inserido:")
			fmt.Scanln(&val)
			err = client.Call("RemoteList.Append", []int{id,val}, &reply)
		}

		if op==3 {
			err = client.Call("RemoteList.Remove", 0, &reply_i)
			fmt.Println("Elemento retirado:", reply_i)
		}

		if op==4 {
			err = client.Call("RemoteList.Size", id, &reply_i)	
			fmt.Println("Tamanho da Lista:", reply_i)
		}

		if op==5 {
			err = client.Call("RemoteList.SaveFile", id, &reply_i)
		}

		if op==6 {
			fmt.Println("Informe a posição do valor:")
			fmt.Scanln(&val)
			err = client.Call("RemoteList.Get", []int{id, val}, &reply_i)
		}
		if err != nil {
			fmt.Print("Error:", err)
		}
	}
}
