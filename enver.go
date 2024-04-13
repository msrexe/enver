package enver

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	DefaultSources = []Source{
		&Source_Env{},
		&Source_DotEnv{},
	}
)

// Fill fills the given struct with environment variables from the sources.
// The given struct has struct tags to specify the environment variable key and default value.
// The struct tags have the following format:
// `env:"KEY,DEFAULT"`
// The KEY is the environment variable key.
// The DEFAULT is the default value to use if the environment variable is not set.
// The sources are used in the order they are provided. Precedence is given to sources that are provided first.
func Fill(envs any, sources ...Source) error {
	if reflect.TypeOf(envs).Kind() != reflect.Ptr {
		return fmt.Errorf("envs parameter must be a pointer to a struct")
	}

	for i := len(sources) - 1; i >= 0; i-- {
		envMap, err := sources[i].Get()
		if err != nil {
			return err
		}

		err = fill(envs, envMap)
		if err != nil {
			return err
		}
	}

	return nil
}

func fill(obj interface{}, envMap map[string]any) error {
	objType := reflect.TypeOf(obj)

	if objType.Kind() != reflect.Ptr {
		return fmt.Errorf("setField: obj parameter must be a pointer to a struct")
	}

	objValue := reflect.ValueOf(obj).Elem()

	for i := 0; i < objType.Elem().NumField(); i++ {
		field := objType.Elem().Field(i)

		envTag := field.Tag.Get("env")
		if envTag != "" {
			parts := strings.Split(envTag, ",")
			envKey := parts[0]
			defaultValue := parts[1]

			envValue, ok := envMap[envKey]
			if !ok {
				envValue = defaultValue
			}

			fieldValue := objValue.FieldByName(field.Name)
			convertedValue := reflect.ValueOf(envValue).Convert(fieldValue.Type())
			fieldValue.Set(convertedValue)
		}
	}

	return nil
}
