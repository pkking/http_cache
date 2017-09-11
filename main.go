package main

import (
	"fmt"
	"strings"
)

const (
	HTTPS_POS = 0
	HTTP_POS  = 1
)

type service struct {
	//	ts int64
	port string
	Type string
	name string
}

type serviceindex struct {
	port  string
	index int
}

type servicelist struct {
	l          map[string][]*service
	index      map[string][]*serviceindex
	needUpdate bool
}

func InitService() *servicelist {
	return &servicelist{
		l:          make(map[string][]*service),
		index:      make(map[string][]*serviceindex),
		needUpdate: false,
	}
}

func (sl *servicelist) insertService(s *service) {
	ports := strings.Split(s.port, "|")
	for _, port := range ports {
		newindex := len(sl.l[port])
		sl.index[s.name] = append(sl.index[s.name], &serviceindex{
			port:  port,
			index: newindex,
		})
		sl.l[port] = append(sl.l[port], s)
		if newindex == 0 {
			sl.needUpdate = true
		}
	}
}

func (sl *servicelist) updateService(s *service, index []*serviceindex) {
	ports := strings.Split(s.port, "|")
	for i, port := range ports {
		if sl.l[index[i].port][index[i].index].name != s.name {
			fmt.Println("add service name not same with indexed service!")
		} else {
			if port != index[i].port {
				// remove old entry
				sl.l[index[i].port] = append(sl.l[index[i].port][:index[i].index], sl.l[index[i].port][index[i].index+1:]...)
				newindex := len(sl.l[port])
				// add new entry
				sl.index[s.name][i].index = newindex
				sl.index[s.name][i].port = port
				sl.l[port] = append(sl.l[port], s)
				if newindex == 0 {
					sl.needUpdate = true
				}

			} else {
				//oldts := sl.l[index.port][index.index].ts
				sl.l[index[i].port][index[i].index] = s
				//sl.l[index.port][index.index].ts = oldts
			}
			if index[i].index == 0 {
				sl.needUpdate = true
			}
		}
	}
}

func (sl *servicelist) AddService(s *service) {
	if i, ok := sl.index[s.name]; ok {
		sl.updateService(s, i)
	} else {
		sl.insertService(s)
	}
}

func (sl *servicelist) DelService(s *service) {
	if indexes, ok := sl.index[s.name]; ok {
		delete(sl.index, s.name)
		for _, index := range indexes {
			sl.l[index.port] = append(sl.l[index.port][:index.index], sl.l[index.port][index.index+1:]...)
			if len(sl.l[index.port]) == 0 {
				delete(sl.l, index.port)
			}
			if index.index == 0 {
				sl.needUpdate = true
			}
		}
	}
}

func (sl servicelist) Print() {
	fmt.Printf("[\n")
	for k, v := range sl.index {
		for _, index := range v {
			fmt.Printf("index:\tname:%s index:%d port:%s\n", k, index.index, index.port)
		}
	}
	fmt.Printf("]\n")

	fmt.Printf("[\n")
	for port, services := range sl.l {
		fmt.Printf("\t{port :%s, ", port)
		for _, d := range services {
			fmt.Printf("%+v  ", *d)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("]\n")
}

/*
func (sl *servicelist) UpdateService(s []*services) {

}

func (sl *servicelist) DelService(s []*service) {

}
*/
func main() {
	sl := InitService()
	s1 := &service{
		port: "1200|200",
		Type: "http",
		name: "test1",
	}
	s2 := &service{
		port: "1201",
		Type: "https",
		name: "test2",
	}
	s3 := &service{
		port: "1205",
		Type: "http-server",
		name: "test4",
	}
	s4 := &service{
		port: "100|200",
		Type: "https",
		name: "xxx",
	}

	sl.AddService(s4)
	sl.Print()
	sl.AddService(s1)
	sl.Print()
	sl.AddService(s3)
	sl.Print()
	sl.AddService(s2)
	sl.Print()
	sl.AddService(s1)
	sl.Print()
	sl.DelService(s4)
	sl.Print()
}
