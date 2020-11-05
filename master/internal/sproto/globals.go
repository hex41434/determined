package sproto

import (
	"github.com/pkg/errors"

	"github.com/determined-ai/determined/master/pkg/actor"
)

var globalRM *actor.Ref

// SetRM sets the global resource manager.
func SetRM(ref *actor.Ref) {
	globalRM = ref
}

// GetRM returns the global resource manager.
func GetRM() *actor.Ref {
	return globalRM
}

// UseAgentRM returns if using the agent resource manager.
func UseAgentRM() bool {
	return globalRM.System().Get(actor.Addr("agents")) != nil
}

// UseK8sRM returns if using the kubernetes resource manager.
func UseK8sRM() bool {
	return globalRM.System().Get(actor.Addr("pods")) != nil
}

// GetRP returns the resource pool.
func GetRP(name string) *actor.Ref {
	return globalRM.Child(name)
}

// ValidateRP validates if the resource pool exists when using the agent resource manager.
func ValidateRP(name string) error {
	if UseAgentRM() && GetRP(name) != nil {
		return nil
	}
	return errors.Errorf("cannot find resource pool: %s", name)
}
