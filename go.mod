go 1.23

module github.com/aws/aws-dax-go-v2

require (
	github.com/antlr4-go/antlr/v4 v4.0.0-00010101000000-000000000000
	github.com/aws/aws-sdk-go-v2 v0.0.0-00010101000000-000000000000
	github.com/aws/smithy-go v0.0.0-00010101000000-000000000000
	github.com/davecgh/go-spew v0.0.0-00010101000000-000000000000
	github.com/gofrs/uuid v0.0.0-00010101000000-000000000000
	github.com/google/go-cmp/cmp v0.0.0-00010101000000-000000000000
	github.com/jmespath/go-jmespath v0.0.0-00010101000000-000000000000
	github.com/pmezard/go-difflib v0.0.0-00010101000000-000000000000
	github.com/stretchr/objx v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v3 v3.0.0-00010101000000-000000000000
)

replace github.com/antlr4-go/antlr/v4 v4.0.0-00010101000000-000000000000 => ./build/private/bgospace/Go3p-Github-Antlr4Go-Antlr-V4/src/github.com/antlr4-go/antlr/v4

replace github.com/aws/aws-sdk-go-v2 v0.0.0-00010101000000-000000000000 => ./build/private/bgospace/GoAmzn-Github-Aws-AwsSdkGoV2/src/github.com/aws/aws-sdk-go-v2

replace github.com/aws/smithy-go v0.0.0-00010101000000-000000000000 => ./build/private/bgospace/GoAmzn-Github-Aws-SmithyGo/src/github.com/aws/smithy-go

replace github.com/davecgh/go-spew v0.0.0-00010101000000-000000000000 => ./build/private/bgospace/Go3p-Github-Stretchr-Testify/src/github.com/davecgh/go-spew

replace github.com/gofrs/uuid v0.0.0-00010101000000-000000000000 => ./build/private/bgospace/Go3p-Github-Gofrs-Uuid/src/github.com/gofrs/uuid

replace github.com/google/go-cmp/cmp v0.0.0-00010101000000-000000000000 => ./build/private/bgospace/Go3p-Github-Google-GoCmp/src/github.com/google/go-cmp/cmp

replace github.com/jmespath/go-jmespath v0.0.0-00010101000000-000000000000 => ./build/private/bgospace/GoAmzn-Github-Aws-AwsSdkGoV2/src/github.com/jmespath/go-jmespath

replace github.com/pmezard/go-difflib v0.0.0-00010101000000-000000000000 => ./build/private/bgospace/Go3p-Github-Stretchr-Testify/src/github.com/pmezard/go-difflib

replace github.com/stretchr/objx v0.0.0-00010101000000-000000000000 => ./build/private/bgospace/Go3p-Github-Stretchr-Testify/src/github.com/stretchr/objx

replace github.com/stretchr/testify v0.0.0-00010101000000-000000000000 => ./build/private/bgospace/Go3p-Github-Stretchr-Testify/src/github.com/stretchr/testify

replace gopkg.in/yaml.v3 v3.0.0-00010101000000-000000000000 => ./build/private/bgospace/Go3p-Github-Stretchr-Testify/src/gopkg.in/yaml.v3
