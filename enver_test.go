package enver

import (
	"os"
	"testing"
)

func TestFill_ProcessEnv(t *testing.T) {
	type envs struct {
		Env string `env:"ENV,default"`
	}

	if err := os.Setenv("ENV", "test"); err != nil {
		t.Errorf("Setenv failed: %v", err)
	}

	testEnvs := envs{}
	err := Fill(&testEnvs, &Source_Env{})
	if err != nil {
		t.Errorf("Fill failed: %v", err)
	}

	if testEnvs.Env != "test" {
		t.Errorf("Expected test, got %s", testEnvs.Env)
	}
}

func TestFill_ProcessDotEnv(t *testing.T) {
	type envs struct {
		Env string `env:"ENV,default"`
	}

	if err := os.WriteFile(".env", []byte("ENV=test"), 0644); err != nil {
		t.Errorf("WriteFile failed: %v", err)
	}

	testEnvs := envs{}
	err := Fill(&testEnvs, &Source_DotEnv{})
	if err != nil {
		t.Errorf("Fill failed: %v", err)
	}

	if testEnvs.Env != "test" {
		t.Errorf("Expected test, got %s", testEnvs.Env)
	}

	if err := os.Remove(".env"); err != nil {
		t.Errorf("Remove failed: %v", err)
	}
}

func TestFill_ProcessEnvAndDotEnv(t *testing.T) {
	type envs struct {
		Env string `env:"ENV,default"`
	}

	if err := os.Setenv("ENV", "test"); err != nil {
		t.Errorf("Setenv failed: %v", err)
	}

	if err := os.WriteFile(".env", []byte("ENV=env"), 0644); err != nil {
		t.Errorf("WriteFile failed: %v", err)
	}

	testEnvs := envs{}
	err := Fill(&testEnvs, &Source_Env{}, &Source_DotEnv{})
	if err != nil {
		t.Errorf("Fill failed: %v", err)
	}

	if testEnvs.Env != "test" {
		t.Errorf("Expected test, got %s", testEnvs.Env)
	}

	if err := os.Remove(".env"); err != nil {
		t.Errorf("Remove failed: %v", err)
	}
}

func TestFill_DefaultValue(t *testing.T) {
	type envs struct {
		Env string `env:"ENV,default"`
	}

	testEnvs := envs{}
	err := Fill(&testEnvs, &Source_Env{})
	if err != nil {
		t.Errorf("Fill failed: %v", err)
	}

	if testEnvs.Env != "default" {
		t.Errorf("Expected default, got %s", testEnvs.Env)
	}
}

func TestFill_InvalidStruct(t *testing.T) {
	type envs struct {
		Env string `env:"ENVTEST,default"`
	}

	testEnvs := envs{}
	err := Fill(testEnvs, &Source_Env{})
	if err == nil {
		t.Errorf("Fill should have failed")
	}
}
