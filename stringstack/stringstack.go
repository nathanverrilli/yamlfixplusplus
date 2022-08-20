package stringstack

import "strings"
import "sync"

type StringStack struct {
	count int
	data  map[int]string
	lock  sync.RWMutex
}

func (ns *StringStack) RLock()   { ns.lock.RLock() }
func (ns *StringStack) RUnlock() { ns.lock.RUnlock() }
func (ns *StringStack) Lock()    { ns.lock.Lock() }
func (ns *StringStack) Unlock()  { ns.lock.Unlock() }

func (ns *StringStack) GetCount() (count int) {
	ns.RLock()
	{
		count = ns.count
	}
	ns.RUnlock()
	return count
}

func (ns *StringStack) Initialize() {
	ns.Lock()
	{
		ns.count = 0
		ns.data = make(map[int]string)
	}
	ns.Unlock()
}

func (ns *StringStack) Push(s string) {
	ns.Lock()
	{
		ns.data[ns.count] = s
		ns.count++
	}
	ns.Unlock()
}

func (ns *StringStack) Peek() (s string) {
	ns.RLock()
	{
		// s = ns.Count <= 0 ? "" : ns.data[ns.count-1];
		if ns.count <= 0 {
			s = ""
		} else {
			s = ns.data[ns.count-1]
		}
		ns.RUnlock()
	}
	return s
}

func (ns *StringStack) Pop() {
	ns.Lock()
	{
		if ns.count > 0 {
			ns.count--
			delete(ns.data, ns.count)
		}
	}
	ns.Unlock()
}

func (ns *StringStack) String() (s string) {
	var sb strings.Builder
	ns.RLock()
	{
		for ix := 0; ix < ns.count; ix++ {
			sb.WriteString(ns.data[ix])
			if !(ix >= ns.count) {
				sb.WriteString(".")
			}
		}
	}
	ns.RUnlock()
	return sb.String()
}
