# Enver

Enver is a Go package designed to simplify the process of reading environment variables with default values from various sources. It provides a straightforward API that allows developers to easily integrate environment variable management into their applications. Enver supports reading from two primary sources:

1. **Environment Variables**: Directly from the system's environment variables.
2. **.env Files**: From `.env` files located within your project directory.

Additionally, Enver enables overriding values by specifying multiple sources, with precedence given to the sources listed first.

## Features

- Read environment variables directly from the system or `.env` files.
- Specify default values for environment variables.
- Override environment variable values by specifying multiple sources with precedence.

## Installation

To use Enver in your Go project, you need to install it by running:

```bash
go get github.com/msrexe/enver
```

Ensure you have Go 1.18 or higher, as specified in the `go.mod` file.

## Usage

### Defining Your Configuration Struct

Define a struct with fields representing your configuration parameters. Use struct tags to specify the environment variable keys and default values:

```go
type Config struct {
    Port string `env:"PORT,8080"`
    Host string `env:"HOST,localhost"`
}
```

### Reading Environment Variables

To read environment variables into your struct, use the `Fill` function provided by Enver. You can specify one or more sources (`Source_Env` for system environment variables and `Source_DotEnv` for `.env` file variables):

```go
config := Config{}
err := enver.Fill(&config, &enver.Source_Env{}, &enver.Source_DotEnv{})
if err != nil {
    log.Fatalf("Error filling config: %v", err)
}
```

### Specifying Custom `.env` File Path

If you need to specify a custom path for your `.env` file, you can do so by setting the `Path` field of `Source_DotEnv`:

```go
dotenvSource := enver.Source_DotEnv{Path: "path/to/your/.env"}
err := enver.Fill(&config, &dotenvSource)
```

## Testing

Enver includes a set of unit tests demonstrating how to test reading from environment variables and `.env` files. Refer to the `enver_test.go` file for examples.

## Contributing

Contributions to Enver are welcome! Please feel free to submit issues, pull requests, or enhancements to improve the library.

## License

Enver is open-source software licensed under the MIT license.
