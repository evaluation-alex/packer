package shell

import (
	"runtime"

	sl "github.com/hashicorp/packer/common/shell-local"
	"github.com/hashicorp/packer/packer"
)

type Provisioner struct {
	config sl.Config
}

func (p *Provisioner) Prepare(raws ...interface{}) error {
	err := sl.Decode(&p.config, raws...)
	if err != nil {
		return err
	}
	if len(p.config.ExecuteCommand) == 0 && runtime.GOOS == "windows" {
		p.config.ExecuteCommand = []string{
			"bash",
			"-c",
			"\"{{.Vars}} {{.Script}}\"",
		}
	}

	return sl.Validate(&p.config)
}

func (p *Provisioner) Provision(ui packer.Ui, _ packer.Communicator) error {
	_, retErr := sl.Run(ui, &p.config)
	if retErr != nil {
		return retErr
	}

	return nil
}

func (p *Provisioner) Cancel() {
	// Just do nothing. When the process ends, so will our provisioner
}
