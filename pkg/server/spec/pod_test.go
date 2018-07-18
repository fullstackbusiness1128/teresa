package spec

import (
	"testing"

	"github.com/luizalabs/teresa/pkg/server/app"
	"github.com/luizalabs/teresa/pkg/server/storage"
)

func TestPodBuilder(t *testing.T) {
	cn := NewContainerBuilder("container", "image:v1").
		ExposePort("http", DefaultPort).
		Build()

	expectedName := "test"
	expectedNamespace := "ns"

	init := NewInitContainer("image:v3", "slug.tgz", storage.NewFake())
	mso := MountSecretInInitContainer("vl", "/s", "s")
	svio := ShareVolumeBetweenAppAndInitContainer("vol", "/app")

	ng := NewNginxContainer("image:v2", &app.App{})
	svo := ShareVolumeBetweenAppAndSideCar("vol", "/app")
	mcm := MountConfigMapInSideCar("cm", "/cm", "cm")

	ps := NewPodBuilder(expectedName, expectedNamespace).
		WithAppContainer(cn).
		WithInitContainer(init, mso, svio).
		WithSideCar(ng, svo, mcm, SwitchPortWithAppContainer).
		Build()

	if ps.Name != expectedName {
		t.Errorf("expected %s, got %s", expectedName, ps.Name)
	}
	if ps.Namespace != expectedNamespace {
		t.Errorf("expected %s, got %s", expectedNamespace, ps.Namespace)
	}
	if actual := len(ps.Containers); actual != 2 {
		t.Fatalf("expected 2 containers, got %d", actual)
	}
	if actual := ps.Containers[0].Ports[0].ContainerPort; actual != secondaryPort {
		t.Errorf("expected %d, got %d", secondaryPort, actual)
	}
	if actual := ps.Containers[1].Ports[0].ContainerPort; actual != DefaultPort {
		t.Errorf("expected %d, got %d", DefaultPort, actual)
	}
	if actual := len(ps.InitContainers); actual != 1 {
		t.Errorf("expected 1, got %d", actual)
	}
	// init secret, init shared with app, nginx config map, nginx shared with app
	if actual := len(ps.Volumes); actual != 4 {
		t.Errorf("expected 4, got %d", actual)
	}
	if actual := len(ps.Containers[0].VolumeMounts); actual != 2 {
		t.Errorf("expected 2, got %d", actual)
	}
}
