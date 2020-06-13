package utils

// called endpoint service in node base on type of job as following

import (
	"Beq/dispurser/model"
	"fmt"
)

//Dispurse used for dispurse jobs to Node
func Dispurse(job *model.Job) {
	switch job.Type {
	case model.TypeAddRule:
		// add rule
		addRule()
		break
	case model.TypeRemoveRule:
		// remove rule
		removeRule()
		break
	case 3:
		// add flow
		break
	case 4:
		// remove flow
		break
	}
}

func removeRule() {
	fmt.Println("Remove RULE!")
}

func addRule() {
	fmt.Println("Add RULE!")
}
