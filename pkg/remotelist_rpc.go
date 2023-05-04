package remotelist

import (
	"errors"
	"fmt"
	"sync"
	"os"
	"strconv"
	"io"
)

type RemoteList struct {
	mu   sync.Mutex
	list_map map[int][]int
}

const dir = "./listas/"

func (l *RemoteList)ReadFile (id int, filePath string) error {
	l.list_map[id]=[]int{}
    fd, err := os.Open(filePath)
    if err != nil {
        return err
    }
    var line int
    for {

        _, err := fmt.Fscanf(fd, "%d\n", &line)

        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err

        }
        l.list_map[id] = append(l.list_map[id], line)
    }
}

func (l *RemoteList)GetLists() error {
	f, err := os.Open(dir)
    if err != nil {
        return err
    }
	files, err := f.Readdir(0)
    if err != nil {
        return err
    }

    for _, v := range files {
		file := v.Name()
		id, err := strconv.Atoi(file[0:len(file)-4])
		if err != nil {
			return err
		}
		filePath := dir+file
		l.ReadFile(id, filePath)
    }
	return nil
}

func (l *RemoteList) Exists() bool {
	if len(l.list_map) == 0 {
		return false
	}
	
	return true
}

func (l *RemoteList) ListExists(id int) bool {
	_, check := l.list_map[id]
	if !check{
		filePath := dir+strconv.Itoa(id)+".txt"
		err := l.ReadFile(id, filePath)
		if err!=nil {
			return false
		}
	}
	return true
}


func (l *RemoteList) Create(id int, reply *bool) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.Exists() {
		l.list_map = make(map[int][]int)
		l.GetLists()
	}
	if l.ListExists(id){
		*reply=false
		return nil
	} 
	l.list_map[id]=[]int{}
	*reply = true
	return nil
}

func (l *RemoteList) Append(arg []int, reply *bool) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	id := arg[0]
	value := arg[1]
	if !l.ListExists(id){
		*reply=false
		return errors.New("A lista não existe")
	} 
	l.list_map[id] = append(l.list_map[id], value)
	*reply = true
	return nil
}

func (l *RemoteList) Remove(id int, reply *int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.ListExists(id){
		return errors.New("A lista não existe")
	} 

	if len(l.list_map[id]) > 0 {
		*reply = l.list_map[id][len(l.list_map[id])-1]
		l.list_map[id] = l.list_map[id][:len(l.list_map[id])-1]
		fmt.Println(l.list_map[id])
	} else {
		return errors.New("empty list")
	}
	return nil
}

func (l *RemoteList) SaveFile(id int, reply *int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.ListExists(id){
		return errors.New("A lista não existe")
	} 
	filePath := "listas/"+strconv.Itoa(id)+".txt"
    f, err := os.Create(filePath)
    if err != nil {
		*reply=-1
        return err
    }
    defer f.Close()
    for _, value := range l.list_map[id] {
       fmt.Fprintln(f, value)  
    }
    return nil
}

func (l *RemoteList) Get(arg []int, reply *int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	id := arg[0]
	pos := arg[1]
	if !l.ListExists(id){
		return errors.New("A lista não existe")
	} 
	if pos >= len(l.list_map[id]){
		return errors.New("Posicao invalida")
	}
	*reply = l.list_map[id][pos]
	return nil
}

func (l *RemoteList) Size(id int, reply *int) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.ListExists(id){
		return errors.New("A lista não existe")
	} 
	*reply = len(l.list_map[id])
	return nil
}

func NewRemoteList() *RemoteList {
	return new(RemoteList)
}

