package plexpac

import (
	"sync"
	"time"

	"github.com/cihub/seelog"
	"github.com/fanyang1988/proxy-services/winproxy"
)

// PacDataGetter getter for pac
type PacDataGetter struct {
	pacDatas []byte
	mutex    sync.RWMutex

	ip string
}

// NewPacDataGetter create a new PacDataGetter and then init
func NewPacDataGetter(ip string) *PacDataGetter {
	res := &PacDataGetter{
		ip: ip,
	}

	res.init()

	go func() {
		res.loop()
	}()

	return res
}

func (p *PacDataGetter) init() {
	reTryTimes := 0
	for {
		if err := p.updatePAC(); err == nil {
			return
		}

		reTryTimes++

		if reTryTimes >= 3 {
			seelog.Errorf("init pac failed!!!")
			return
		}

		seelog.Infof("update pac error re try %d times, wait 3 second", reTryTimes)
		time.Sleep(3 * time.Second)
	}
}

func (p *PacDataGetter) updatePAC() error {
	datas, err := winproxy.GetPAC()
	if err != nil {
		seelog.Errorf("get pac error %s", err.Error())
		return err
	}

	p.pacDatas = changePAC(p.ip, datas)
	return nil
}

func (p *PacDataGetter) loop() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case t := <-ticker.C:
			seelog.Debugf("update pac in %s", t)
			func() {
				p.mutex.Lock()
				defer p.mutex.Unlock()

				p.updatePAC()
			}()
		}
	}
}

// GetPAC get pac datas
func (p *PacDataGetter) GetPAC() []byte {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	res := make([]byte, len(p.pacDatas))
	copy(res, p.pacDatas)

	return res[:]
}
