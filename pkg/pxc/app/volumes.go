package app

import (
	corev1 "k8s.io/api/core/v1"

	api "github.com/percona/percona-xtradb-cluster-operator/pkg/apis/pxc/v1alpha1"
)

func GetConfigVolumes(cvName string) corev1.Volume {
	vol1 := corev1.Volume{
		Name: "config-volume",
	}

	vol1.ConfigMap = &corev1.ConfigMapVolumeSource{}
	vol1.ConfigMap.Name = cvName
	t := true
	vol1.ConfigMap.Optional = &t
	return vol1
}

func Volumes(podSpec *api.PodSpec, dataVolumeName string) *api.Volume {
	var volume api.Volume

	if podSpec.VolumeSpec.PersistentVolumeClaim != nil {
		pvcs := PVCs(dataVolumeName, podSpec.VolumeSpec)
		volume.PVCs = pvcs
		return &volume
	}

	volume.Volumes = append(volume.Volumes, corev1.Volume{
		VolumeSource: corev1.VolumeSource{
			HostPath: podSpec.VolumeSpec.HostPath,
			EmptyDir: podSpec.VolumeSpec.EmptyDir,
		},
		Name: dataVolumeName,
	})

	return &volume
}
