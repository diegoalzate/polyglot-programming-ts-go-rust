package projector

import (
	"encoding/json"
	"errors"
	"os"
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
		if prev == cur {
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

	for cur != prev {
		paths = append(paths, cur)
		// step back in dir
		prev = cur
		cur = path.Dir(cur)
	}

	// reverse paths so it goes from less specific to more specific
	for i := len(paths) - 1; i >= 0; i-- {
		for key, value := range p.Data.Projector[paths[i]] {
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

func (p *Projector) Save() error {
	// write config file down if it does not exist
	if _, err := os.Stat(p.Config.Config); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(path.Dir(p.Config.Config), 0755); err != nil {
			bytes := []byte(" \"projector\": {}")
			err := os.WriteFile(p.Config.Config, bytes, 0755)
			if err != nil {
				return err
			}
		}
	}
	// save data to config
	bytesJson, err := json.Marshal(p.Data)

	if err != nil {
		return err
	}

	err = os.WriteFile(p.Config.Config, bytesJson, 0755)

	if err != nil {
		return err
	}

	return nil
}

func FromConfig(config *ProjectorConfig) (*Projector, error) {
	// load data from config
	// check if config exists
	if _, err := os.Stat(config.Config); errors.Is(err, os.ErrNotExist) {
		// create projector if it does not exist for default scenario
		return &Projector{
			Config: *config,
			Data: ProjectorData{
				Projector: map[string]map[string]string{},
			},
		}, nil
	}
	// load data from config
	bytes, err := os.ReadFile(config.Config)

	if err != nil {
		return &Projector{
			Config: *config,
			Data: ProjectorData{
				Projector: map[string]map[string]string{},
			},
		}, nil
	}

	var jsonData ProjectorData
	err = json.Unmarshal(bytes, &jsonData)

	if err != nil {
		return &Projector{
			Config: *config,
			Data: ProjectorData{
				Projector: map[string]map[string]string{},
			},
		}, nil
	}

	return &Projector{
		Config: *config,
		Data:   jsonData,
	}, nil
}
