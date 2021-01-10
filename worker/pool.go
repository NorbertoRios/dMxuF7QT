package worker

import "genx-go/unitofwork"

//Newpool ...
func newPool(workersCount int, _uow unitofwork.IDeviceUnitOfWork) *pool {
	_workers := []*Worker{}
	for i := 0; i < workersCount; i++ {
		w := NewWorker(_uow)
		_workers = append(_workers, w)
	}
	return &pool{
		currentNum: 0,
		workers:    _workers,
	}
}

//pool ...
type pool struct {
	currentNum int
	workers    []*Worker
}

func (p *pool) all() []*Worker {
	return p.workers
}

func (p *pool) next() *Worker {
	defer func() { p.currentNum++ }()
	if p.currentNum == len(p.workers) {
		p.currentNum = 0
	}
	return p.workers[p.currentNum]
}
