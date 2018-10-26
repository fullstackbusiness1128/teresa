package spec

import (
	"errors"
	"reflect"
	"testing"

	"github.com/luizalabs/teresa/pkg/server/app"
)

func TestNewCloudSQLProxy(t *testing.T) {
	img := "image"
	want := &CloudSQLProxy{
		Instances:      "instances",
		CredentialFile: "file",
		Image:          img,
	}
	fn := func(v interface{}) error {
		v.(*CloudSQLProxy).Instances = want.Instances
		v.(*CloudSQLProxy).CredentialFile = want.CredentialFile
		return nil
	}
	ty := &TeresaYaml{
		SideCars: map[string]RawData{
			"cloudsql-proxy": {fn},
		},
	}

	csp, err := NewCloudSQLProxy(img, ty, &app.App{})
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if !reflect.DeepEqual(csp, want) {
		t.Errorf("got %v; want %v", csp, want)
	}
}

func TestNewCloudSQLProxyFromEnvVar(t *testing.T) {
	a := &app.App{
		EnvVars: []*app.EnvVar{
			&app.EnvVar{Key: "DB_PROJECT", Value: "project"},
			&app.EnvVar{Key: "DB_ZONE", Value: "zone"},
			&app.EnvVar{Key: "DB_NAME", Value: "name"},
		},
	}
	img := "image"

	want := &CloudSQLProxy{
		Instances:      "$(DB_PROJECT):$(DB_ZONE):$(DB_NAME)=tcp:3306",
		CredentialFile: "file",
		Image:          img,
	}
	fn := func(v interface{}) error {
		v.(*CloudSQLProxy).CredentialFile = want.CredentialFile
		return nil
	}
	ty := &TeresaYaml{
		SideCars: map[string]RawData{
			"cloudsql-proxy": {fn},
		},
	}

	csp, err := NewCloudSQLProxy(img, ty, a)
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if !reflect.DeepEqual(csp, want) {
		t.Errorf("got %v; want %v", csp, want)
	}
	a = &app.App{
		EnvVars: []*app.EnvVar{
			&app.EnvVar{Key: "GCP_CLOUDSQL_INSTANCE_NAME", Value: "project:zone:name=tcp:3306"},
			&app.EnvVar{Key: "DB_PROJECT", Value: "a_project"},
			&app.EnvVar{Key: "DB_ZONE", Value: "a_zone"},
			&app.EnvVar{Key: "DB_NAME", Value: "a_name"},
		},
	}
	want.Instances = "project:zone:name=tcp:3306"
	csp, err = NewCloudSQLProxy(img, ty, a)
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if !reflect.DeepEqual(csp, want) {
		t.Errorf("got %v; want %v", csp, want)
	}
}

func TestNewCloudSQLProxyError(t *testing.T) {
	fn := func(v interface{}) error {
		return errors.New("test")
	}
	ty := &TeresaYaml{
		SideCars: map[string]RawData{
			"cloudsql-proxy": {fn},
		},
	}

	if _, err := NewCloudSQLProxy("test", ty, &app.App{}); err == nil {
		t.Error("got nil; want error")
	}
}

func TestNewCloudSQLProxyNil(t *testing.T) {
	ty := &TeresaYaml{
		SideCars: map[string]RawData{
			"test": RawData{},
		},
	}

	csp, err := NewCloudSQLProxy("test", ty, &app.App{})
	if err != nil {
		t.Fatal("got unexpected error:", err)
	}
	if csp != nil {
		t.Errorf("got %v; want nil", csp)
	}
}

func TestNewCloudSQLProxyContainerFromEnvVar(t *testing.T) {
	a := &app.App{
		EnvVars: []*app.EnvVar{
			&app.EnvVar{Key: "DB_PROJECT", Value: "project"},
			&app.EnvVar{Key: "DB_ZONE", Value: "zone"},
			&app.EnvVar{Key: "DB_NAME", Value: "name"},
		},
	}
	csp := &CloudSQLProxy{
		Instances:      "$(DB_PROJECT):$(DB_ZONE):$(DB_NAME)=tcp:3306",
		CredentialFile: "file",
		Image:          "image",
	}
	want := &Container{
		Name: "cloudsql-proxy",
		ContainerLimits: &ContainerLimits{
			CPU:    cloudSQLProxyDefaultCPULimit,
			Memory: cloudSQLProxyDefaultMemoryLimit,
		},
		Image:   csp.Image,
		Command: []string{"/cloud_sql_proxy"},
		Args: []string{
			"-instances=$(DB_PROJECT):$(DB_ZONE):$(DB_NAME)=tcp:3306",
			"-credential_file=/secrets/cloudsql/file",
		},
		VolumeMounts: []*VolumeMounts{
			{
				Name:      AppSecretName,
				MountPath: "/secrets/cloudsql/file",
				SubPath:   "file",
				ReadOnly:  true,
			},
		},
		Env: map[string]string{
			"DB_PROJECT": "project",
			"DB_ZONE":    "zone",
			"DB_NAME":    "name",
		},
		Ports:   []Port{},
		Secrets: []string{},
	}

	cn := NewCloudSQLProxyContainer(csp, a)
	if !reflect.DeepEqual(cn, want) {
		t.Errorf("got %v; want %v", cn, want)
	}
}

func TestNewCloudSQLProxyContainer(t *testing.T) {
	a := &app.App{
		EnvVars: []*app.EnvVar{&app.EnvVar{Key: "key1", Value: "value1"}},
	}
	csp := &CloudSQLProxy{
		Instances:      "instances",
		CredentialFile: "file",
		Image:          "image",
	}
	want := &Container{
		Name: "cloudsql-proxy",
		ContainerLimits: &ContainerLimits{
			CPU:    cloudSQLProxyDefaultCPULimit,
			Memory: cloudSQLProxyDefaultMemoryLimit,
		},
		Image:   csp.Image,
		Command: []string{"/cloud_sql_proxy"},
		Args: []string{
			"-instances=instances",
			"-credential_file=/secrets/cloudsql/file",
		},
		VolumeMounts: []*VolumeMounts{
			{
				Name:      AppSecretName,
				MountPath: "/secrets/cloudsql/file",
				SubPath:   "file",
				ReadOnly:  true,
			},
		},
		Env: map[string]string{
			"key1": "value1",
		},
		Ports:   []Port{},
		Secrets: []string{},
	}

	cn := NewCloudSQLProxyContainer(csp, a)
	if !reflect.DeepEqual(cn, want) {
		t.Errorf("got %v; want %v", cn, want)
	}
}
