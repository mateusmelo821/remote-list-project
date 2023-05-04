# README

Este é um projeto desenvolvido para a disciplina de Sistemas Distribuídos do curso de Mestrado em Tecnologia da Informação. Parte do código foi feito em cima da implementação feita pelo professor Ruan Delgado que pode ser conferida [aqui](https://github.com/ruandg/SD_PPGTI).

O objetivo do projeto era construir um componente denominado RemoteList que deveria ser capaz de gerenciar um conjunto de clientes e listas de valores inteiros. Cada cliente pode se comunicar com o servidor e executar as seguintes operações:

* Criar uma lista
* Inserir elemento em uma lista
* Remover elemento da lista
* Obter o tamanho da lista
* Persistir os dados da lista
* Obter item da lista

A comunicação dos clientes com o servidor é feita de forma síncrona e persistente. Para tal, o esquema de comunicação foi implementado usando a Remote Procedure Call (RPC).

A melhor forma de testar o projeto é executando-se em terminais distintos os seguintes comandos:


```
go run remote_rpc_server.go
```
```
go run remote_rpc_client2.go
```

Em seguida, é interessante criar uma lista de id 0, abrir um novo terminal e executar o comando:


```
go run remote_rpc_client.go
```

Desta forma, será possível observar que os clientes pode acessar as mesmas listas sem que haja inconsistência nos dados.

As listas são sempre armazenadas na pasta listas e elas são carregadas sempre que o comando de criação de lista é executado. Também se verifica a existência de uma lista tanto em memória como em disco, antes da execução de todas as operações. Desta forma, é garantida a consistência e confiabilidade dos dados entre todos os clientes.

Como há apenas um servidor, temos ele como principal ponto de falha, uma vez que se ele ficar indisponível, o serviço como um todo se torna indisponível.

