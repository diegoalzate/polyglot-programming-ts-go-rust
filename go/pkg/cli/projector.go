package projector

import (
	"path"
)

type Lookup = map[string]map[string]string
type ProjectorData struct {
	Projector Lookup `json:"projector"`
}

type Projector struct {
	Config ProjectorConfig
	// { pwd -> k,v } 
	Data ProjectorData
}

func (p *Projector) GetValue(key string) (string, bool) {
	found := false
	out := ""

	prev := ""
	cur := p.Config.Pwd

	for i := len(p.Data.Projector) - 1; i >= 0; i-- {
		if dir, ok := p.Data.Projector[cur]; ok {
			if value, ok := dir[key]; ok {
				found = true
				out = value
				break
			}
		}

		// stops infinite loop of going back at root '/' dir
		if (prev == cur) {
			break
		}

		// step back in dir
		prev = cur
		cur = path.Dir(cur)
	}

	return out, found
}

func (p *Projector) GetValueAll() map[string]string {
	out := map[string]string{}

	prev := ""
	cur := p.Config.Pwd

	paths := []string{}

	for ; cur != prev  ; {
		paths = append(paths, cur)
		// step back in dir
		prev = cur
		cur = path.Dir(cur)
	}

	// reverse paths so it goes from less specific to more specific
	for i := len(paths) - 1; i >= 0; i-- {
		for key, value := range(p.Data.Projector[paths[i]]) {
			out[key] = value
		}	
	}

	return out
}

func (p *Projector) SetValue(key, value string) {
	// create pwd data if it does not exist
	if _, ok := p.Data.Projector[p.Config.Pwd]; !ok {
		p.Data.Projector[p.Config.Pwd] = map[string]string{}
	}

	p.Data.Projector[p.Config.Pwd][key] = value
}

func (p *Projector) RemoveValue(key string) {
	if _, ok := p.Data.Projector[p.Config.Pwd]; ok {
		delete(p.Data.Projector[p.Config.Pwd], key)
	}
}