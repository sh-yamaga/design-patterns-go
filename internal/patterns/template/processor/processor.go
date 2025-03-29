package processor

type process interface {
	load() error
	handle() error
	save() error
}

type processor struct {
	process process
}

func (p processor) Execute() error {
	if err := p.process.load(); err != nil {
		return err
	}
	if err := p.process.handle(); err != nil {
		return err
	}
	if err := p.process.save(); err != nil {
		return err
	}

	return nil
}
