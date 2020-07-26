package ice

type Agent struct {
}

func NewAgent() (*Agent, error) {
	agent := &Agent{}
	return agent, nil
}

func (a *Agent) GatherCandidates() error {
	return nil
}
