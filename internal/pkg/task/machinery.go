package task

import (
	"common/internal/conf"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/backends/result"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

type TaskServer struct {
	server  *machinery.Server
	workers map[string]*machinery.Worker
}

func NewTaskServer(taskConf *conf.Task) (*TaskServer, error) {
	server, err := newMachineryServer(taskConf)
	if err != nil {
		return nil, err
	}

	return &TaskServer{
		server:  server,
		workers: make(map[string]*machinery.Worker),
	}, nil
}

func newMachineryServer(mc *conf.Task) (*machinery.Server, error) {
	cnf := &config.Config{
		Broker:        mc.Broker,
		DefaultQueue:  mc.Queue,
		ResultBackend: mc.ResultBackend,
		AMQP: &config.AMQPConfig{
			Exchange:     mc.Exchange,
			ExchangeType: mc.ExchangeType,
		},
	}

	server, err := machinery.NewServer(cnf)
	if err != nil {
		return nil, err
	}

	tasks := map[string]interface{}{
		"push":          push,
		"audit":         audit,
		"auditCallback": auditCallback,
	}

	return server, server.RegisterTasks(tasks)
}

func (t *TaskServer) NewWorker(customerTag string, concurrency int) {
	//fmt.Printf("%v", t.server)
	worker := t.server.NewWorker(customerTag, concurrency)
	t.workers[customerTag] = worker
	worker.Launch()
}

func (t *TaskServer) SendTask(task *tasks.Signature) (*result.AsyncResult, error) {
	result, err := t.server.SendTask(task)
	if err != nil {
		return nil, err
	}

	return result, nil
}
