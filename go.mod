//In Go, modules are used for dependency management and versioning.

//Dependency Management: Go modules make it easier to manage dependencies for your project. When you initialize a module, you create a go.mod file that keeps track of the dependencies your project needs and their versions. This makes it easier to ensure that everyone working on the project has the same versions of dependencies.

// Versioning: Go modules allow you to version your project's code. When you create a module, you specify a module path (e.g., github.com/your-username/project-name), which serves as a unique identifier for your project. This makes it easier to share and distribute your code, as well as to manage different versions of your project.

// Reproducible Builds: Go modules ensure that your project's dependencies are reproducible. When you run go mod tidy, Go will update the go.mod file to include the exact versions of the dependencies your project is using. This makes it easier to reproduce the same build environment, even if you switch machines or work on the project at a later time.

module go-aws-s3-cli

go 1.21.5

require github.com/aws/aws-sdk-go v1.51.25 // or the latest version

require github.com/spf13/cobra v1.8.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)
