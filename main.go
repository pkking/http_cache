package main

import (
	"fmt"
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
	index      map[string]*serviceindex
	needUpdate bool
}

func InitService() *servicelist {
	return &servicelist{
		l:          make(map[string][]*service),
		index:      make(map[string]*serviceindex),
		needUpdate: false,
	}
}

func (sl *servicelist) insertService(s *service) {
	newindex := len(sl.l[s.port])
	sl.index[s.name] = &serviceindex{
		port:  s.port,
		index: newindex,
	}
	sl.l[s.port] = append(sl.l[s.port], s)
	if newindex == 0 {
		sl.needUpdate = true
	}
}

func (sl *servicelist) updateService(s *service, index *serviceindex) {
	if sl.l[index.port][index.index].name != s.name {
		fmt.Println("add service name not same with indexed service!")
	} else {
		if s.port != index.port {
			// remove old entry
			sl.l[index.port] = append(sl.l[index.port][:index.index], sl.l[index.port][index.index+1:]...)
			newindex := len(sl.l[s.port])
			// add new entry
			sl.index[s.name].index = newindex
			sl.index[s.name].port = s.port
			sl.l[s.port] = append(sl.l[s.port], s)
			if newindex == 0 {
				sl.needUpdate = true
			}

		} else {
			//oldts := sl.l[index.port][index.index].ts
			sl.l[index.port][index.index] = s
			//sl.l[index.port][index.index].ts = oldts
		}
		if index.index == 0 {
			sl.needUpdate = true
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
	if index, ok := sl.index[s.name]; ok {
		sl.l[index.port] = append(sl.l[index.port][:index.index], sl.l[index.port][index.index+1:]...)
		delete(sl.index, s.name)
		if index.index == 0 {
			sl.needUpdate = true
		}
	}
}

func (sl servicelist) Print() {
	for k, v := range sl.index {
		fmt.Printf("[\n\tname:%s index:%d port:%s\n]\n", k, v.index, v.port)
	}

	for k, v := range sl.l {
		fmt.Printf("[\n\t")
		for _, d := range v {
			fmt.Printf("\t{port :%s, service:%+v}\n", k, *d)
		}
		fmt.Printf("]\n")
	}
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
		port: "1200",
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
		port: "100",
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
}
