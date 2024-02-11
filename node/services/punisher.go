package services

import "sync"

var punisherService *PunisherService
var punisherServiceOnce sync.Once

type PunisherService struct {
	sync.Mutex
	blacklist map[string]struct{}
}

func GetPunisherService() *PunisherService {
	punisherServiceOnce.Do(func() {
		punisherService = &PunisherService{
			blacklist: make(map[string]struct{}),
		}
	})

	return punisherService
}

func (p *PunisherService) AddToBlacklist(ip string) {
	p.Lock()
	defer p.Unlock()
	p.blacklist[ip] = struct{}{}
}

func (p *PunisherService) IsBlacklisted(ip string) bool {
	p.Lock()
	defer p.Unlock()
	_, exists := p.blacklist[ip]
	return exists
}

func (p *PunisherService) RemoveFromBlacklist(ip string) {
	p.Lock()
	defer p.Unlock()
	delete(p.blacklist, ip)
}

func (p *PunisherService) GetBlacklist() []string {
	p.Lock()
	defer p.Unlock()

	blacklist := make([]string, 0, len(p.blacklist))
	for ip := range p.blacklist {
		blacklist = append(blacklist, ip)
	}

	return blacklist
}

func (p *PunisherService) ClearBlacklist() {
	p.Lock()
	defer p.Unlock()
	p.blacklist = make(map[string]struct{})
}
