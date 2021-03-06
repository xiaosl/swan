package state

import (
	"time"

	rafttypes "github.com/Dataman-Cloud/swan/src/manager/raft/types"
	"github.com/Dataman-Cloud/swan/src/types"
)

func AppToRaft(app *App) *rafttypes.Application {
	raftApp := &rafttypes.Application{
		ID:        app.ID,
		Name:      app.Name,
		CreatedAt: app.Created.UnixNano(),
		UpdatedAt: app.Updated.UnixNano(),
		State:     app.State,
	}

	if app.CurrentVersion != nil {
		raftApp.Version = VersionToRaft(app.CurrentVersion, app.ID)
	}

	if app.ProposedVersion != nil {
		raftApp.ProposedVersion = VersionToRaft(app.ProposedVersion, app.ID)
	}

	return raftApp
}

func VersionToRaft(version *types.Version, appID string) *rafttypes.Version {
	raftVersion := &rafttypes.Version{
		ID:          version.ID,
		Command:     version.Command,
		Args:        version.Args,
		Cpus:        version.CPUs,
		Mem:         version.Mem,
		Disk:        version.Disk,
		Instances:   version.Instances,
		RunAs:       version.RunAs,
		Priority:    version.Priority,
		Labels:      version.Labels,
		Env:         version.Env,
		Constraints: version.Constraints,
		Uris:        version.URIs,
		Ip:          version.IP,
		Mode:        version.Mode,
		AppName:     version.AppName,
		AppID:       appID,
	}

	if version.Container != nil {
		raftVersion.Container = ContainerToRaft(version.Container)
	}

	if version.KillPolicy != nil {
		raftVersion.KillPolicy = KillPolicyToRaft(version.KillPolicy)
	}

	if version.UpdatePolicy != nil {
		raftVersion.UpdatePolicy = UpdatePolicyToRaft(version.UpdatePolicy)
	}

	if version.HealthCheck != nil {
		raftVersion.HealthCheck = HealthCheckToRaft(version.HealthCheck)
	}

	return raftVersion
}

func VersionFromRaft(raftVersion *rafttypes.Version) *types.Version {
	version := &types.Version{
		ID:          raftVersion.ID,
		AppName:     raftVersion.AppName,
		Command:     raftVersion.Command,
		Args:        raftVersion.Args,
		CPUs:        raftVersion.Cpus,
		Mem:         raftVersion.Mem,
		Disk:        raftVersion.Disk,
		Instances:   raftVersion.Instances,
		RunAs:       raftVersion.RunAs,
		Priority:    raftVersion.Priority,
		Labels:      raftVersion.Labels,
		Env:         raftVersion.Env,
		Constraints: raftVersion.Constraints,
		URIs:        raftVersion.Uris,
		IP:          raftVersion.Ip,
		Mode:        raftVersion.Mode,
	}

	if raftVersion.Container != nil {
		version.Container = ContainerFromRaft(raftVersion.Container)
	}

	if raftVersion.KillPolicy != nil {
		version.KillPolicy = KillPolicyFromRaft(raftVersion.KillPolicy)
	}

	if raftVersion.UpdatePolicy != nil {
		version.UpdatePolicy = UpdatePolicyFromRaft(raftVersion.UpdatePolicy)
	}

	if raftVersion.HealthCheck != nil {
		version.HealthCheck = HealthCheckFromRaft(raftVersion.HealthCheck)
	}

	return version
}

func ContainerToRaft(container *types.Container) *rafttypes.Container {
	raftContainer := &rafttypes.Container{
		Type: container.Type,
	}

	if container.Docker != nil {
		raftContainer.Docker = DockerToRaft(container.Docker)
	}

	if container.Volumes != nil {
		var volumes []*rafttypes.Volume

		for _, volume := range container.Volumes {
			volumes = append(volumes, VolumeToRaft(volume))
		}

		raftContainer.Volumes = volumes
	}

	return raftContainer
}

func ContainerFromRaft(raftContainer *rafttypes.Container) *types.Container {
	container := &types.Container{
		Type: raftContainer.Type,
	}

	if raftContainer.Docker != nil {
		container.Docker = DockerFromRaft(raftContainer.Docker)
	}

	if raftContainer.Volumes != nil {
		var volumes []*types.Volume

		for _, volume := range raftContainer.Volumes {
			volumes = append(volumes, VolumeFromFaft(volume))
		}

		container.Volumes = volumes
	}

	return container
}

func DockerToRaft(docker *types.Docker) *rafttypes.Docker {
	raftDocker := &rafttypes.Docker{
		ForcePullImage: docker.ForcePullImage,
		Image:          docker.Image,
		Network:        docker.Network,
		Privileged:     docker.Privileged,
	}

	if docker.Parameters != nil {
		var parameters []*rafttypes.Parameter
		for _, parameter := range docker.Parameters {
			parameters = append(parameters, ParameterToRaft(parameter))
		}

		raftDocker.Parameters = parameters
	}

	if docker.PortMappings != nil {
		var portMappings []*rafttypes.PortMapping

		for _, portMapping := range docker.PortMappings {
			portMappings = append(portMappings, PortMappingToRaft(portMapping))
		}

		raftDocker.PortMappings = portMappings
	}

	return raftDocker
}

func DockerFromRaft(raftDocker *rafttypes.Docker) *types.Docker {
	docker := &types.Docker{
		ForcePullImage: raftDocker.ForcePullImage,
		Image:          raftDocker.Image,
		Network:        raftDocker.Network,
		Privileged:     raftDocker.Privileged,
	}

	if raftDocker.Parameters != nil {
		var parameters []*types.Parameter
		for _, parameter := range raftDocker.Parameters {
			parameters = append(parameters, ParameterFromRaft(parameter))
		}

		docker.Parameters = parameters
	}

	if raftDocker.PortMappings != nil {
		var portMappings []*types.PortMapping
		for _, portMapping := range raftDocker.PortMappings {
			portMappings = append(portMappings, PortMappingFromRaft(portMapping))
		}

		docker.PortMappings = portMappings
	}

	return docker
}

func ParameterToRaft(parameter *types.Parameter) *rafttypes.Parameter {
	return &rafttypes.Parameter{
		Key:   parameter.Key,
		Value: parameter.Value,
	}
}

func ParameterFromRaft(raftParameter *rafttypes.Parameter) *types.Parameter {
	return &types.Parameter{
		Key:   raftParameter.Key,
		Value: raftParameter.Value,
	}
}

func PortMappingToRaft(portMapping *types.PortMapping) *rafttypes.PortMapping {
	return &rafttypes.PortMapping{
		ContainerPort: portMapping.ContainerPort,
		Name:          portMapping.Name,
		Protocol:      portMapping.Protocol,
	}
}

func PortMappingFromRaft(raftPortMapping *rafttypes.PortMapping) *types.PortMapping {
	return &types.PortMapping{
		ContainerPort: raftPortMapping.ContainerPort,
		Name:          raftPortMapping.Name,
		Protocol:      raftPortMapping.Protocol,
	}
}

func VolumeToRaft(volume *types.Volume) *rafttypes.Volume {
	return &rafttypes.Volume{
		ContainerPath: volume.ContainerPath,
		HostPath:      volume.HostPath,
		Mode:          volume.Mode,
	}
}

func VolumeFromFaft(raftVolume *rafttypes.Volume) *types.Volume {
	return &types.Volume{
		ContainerPath: raftVolume.ContainerPath,
		HostPath:      raftVolume.HostPath,
		Mode:          raftVolume.Mode,
	}
}

func KillPolicyToRaft(killPolicy *types.KillPolicy) *rafttypes.KillPolicy {
	return &rafttypes.KillPolicy{
		Duration: killPolicy.Duration,
	}
}

func KillPolicyFromRaft(raftKillPolicy *rafttypes.KillPolicy) *types.KillPolicy {
	return &types.KillPolicy{
		Duration: raftKillPolicy.Duration,
	}
}

func UpdatePolicyToRaft(updatePolicy *types.UpdatePolicy) *rafttypes.UpdatePolicy {
	return &rafttypes.UpdatePolicy{
		UpdateDelay:  updatePolicy.UpdateDelay,
		MaxRetries:   updatePolicy.MaxRetries,
		MaxFailovers: updatePolicy.MaxFailovers,
		Action:       updatePolicy.Action,
	}
}

func UpdatePolicyFromRaft(raftUpdatePolicy *rafttypes.UpdatePolicy) *types.UpdatePolicy {
	return &types.UpdatePolicy{
		UpdateDelay:  raftUpdatePolicy.UpdateDelay,
		MaxRetries:   raftUpdatePolicy.MaxRetries,
		MaxFailovers: raftUpdatePolicy.MaxFailovers,
		Action:       raftUpdatePolicy.Action,
	}
}

func HealthCheckToRaft(healthCheck *types.HealthCheck) *rafttypes.HealthCheck {
	raftHealthCheck := &rafttypes.HealthCheck{
		ID:                  healthCheck.ID,
		Address:             healthCheck.Address,
		Protocol:            healthCheck.Protocol,
		PortName:            healthCheck.PortName,
		Path:                healthCheck.Path,
		Value:               healthCheck.Value,
		ConsecutiveFailures: healthCheck.ConsecutiveFailures,
		GracePeriodSeconds:  healthCheck.GracePeriodSeconds,
		IntervalSeconds:     healthCheck.IntervalSeconds,
		TimeoutSeconds:      healthCheck.TimeoutSeconds,
		DelaySeconds:        healthCheck.DelaySeconds,
	}

	return raftHealthCheck
}

func HealthCheckFromRaft(raftHealthCheck *rafttypes.HealthCheck) *types.HealthCheck {
	healthCheck := &types.HealthCheck{
		ID:                  raftHealthCheck.ID,
		Address:             raftHealthCheck.Address,
		Protocol:            raftHealthCheck.Protocol,
		PortName:            raftHealthCheck.PortName,
		Path:                raftHealthCheck.Path,
		Value:               raftHealthCheck.Value,
		ConsecutiveFailures: raftHealthCheck.ConsecutiveFailures,
		GracePeriodSeconds:  raftHealthCheck.GracePeriodSeconds,
		IntervalSeconds:     raftHealthCheck.IntervalSeconds,
		TimeoutSeconds:      raftHealthCheck.TimeoutSeconds,
		DelaySeconds:        raftHealthCheck.DelaySeconds,
	}

	return healthCheck
}

func SlotToRaft(slot *Slot) *rafttypes.Slot {
	raftSlot := &rafttypes.Slot{
		Index:                int32(slot.Index),
		ID:                   slot.ID,
		AppID:                slot.App.ID,
		VersionID:            slot.Version.ID,
		State:                slot.State,
		MarkForDeletion:      slot.MarkForDeletion(),
		MarkForRollingUpdate: slot.MarkForRollingUpdate(),
		Healthy:              slot.Healthy(),
	}

	if slot.CurrentTask != nil {
		raftSlot.CurrentTask = TaskToRaft(slot.CurrentTask)
	}

	// TODO store slot.restartPolicy

	return raftSlot
}

func SlotFromRaft(raftSlot *rafttypes.Slot) *Slot {
	slot := &Slot{
		Index:                int(raftSlot.Index),
		ID:                   raftSlot.ID,
		State:                raftSlot.State,
		CurrentTask:          TaskFromRaft(raftSlot.CurrentTask),
		OfferID:              raftSlot.CurrentTask.OfferID,
		AgentID:              raftSlot.CurrentTask.AgentID,
		Ip:                   raftSlot.CurrentTask.Ip,
		AgentHostName:        raftSlot.CurrentTask.AgentHostName,
		markForDeletion:      raftSlot.MarkForDeletion,
		markForRollingUpdate: raftSlot.MarkForRollingUpdate,
		healthy:              raftSlot.Healthy,
	}

	raftVersion, err := persistentStore.GetVersion(raftSlot.AppID, raftSlot.VersionID)
	if err == nil {
		slot.Version = VersionFromRaft(raftVersion)
	}

	return slot
}

func TaskToRaft(task *Task) *rafttypes.Task {
	return &rafttypes.Task{
		ID:            task.ID,
		AppID:         task.Slot.App.ID,
		VersionID:     task.Version.ID,
		SlotID:        task.Slot.ID,
		State:         task.State,
		Stdout:        task.Stdout,
		Stderr:        task.Stderr,
		HostPorts:     task.HostPorts,
		OfferID:       task.OfferID,
		AgentID:       task.AgentID,
		Ip:            task.Ip,
		AgentHostName: task.AgentHostName,
		Reason:        task.Reason,
		Message:       task.Message,
		CreatedAt:     task.Created.UnixNano(),
		ArchivedAt:    task.ArchivedAt.UnixNano(),
	}
}

func TaskFromRaft(raftTask *rafttypes.Task) *Task {
	task := &Task{
		ID:            raftTask.ID,
		State:         raftTask.State,
		Stdout:        raftTask.Stdout,
		Stderr:        raftTask.Stderr,
		HostPorts:     raftTask.HostPorts,
		OfferID:       raftTask.OfferID,
		AgentID:       raftTask.AgentID,
		Ip:            raftTask.Ip,
		AgentHostName: raftTask.AgentHostName,
		Reason:        raftTask.Reason,
		Message:       raftTask.Message,
		Created:       time.Unix(0, raftTask.CreatedAt),
	}

	raftVersion, err := persistentStore.GetVersion(raftTask.AppID, raftTask.VersionID)
	if err == nil {
		task.Version = VersionFromRaft(raftVersion)
	}

	return task
}

func OfferAllocatorItemToRaft(slotID string, offerInfo *OfferInfo) *rafttypes.OfferAllocatorItem {
	item := &rafttypes.OfferAllocatorItem{
		OfferID:  offerInfo.OfferID,
		SlotID:   slotID,
		Hostname: offerInfo.Hostname,
		AgentID:  offerInfo.AgentID,
	}

	return item
}

func OfferAllocatorItemFromRaft(item *rafttypes.OfferAllocatorItem) (slotID string, offerInfo *OfferInfo) {
	return item.SlotID, &OfferInfo{
		OfferID:  item.OfferID,
		Hostname: item.Hostname,
		AgentID:  item.AgentID,
	}
}
