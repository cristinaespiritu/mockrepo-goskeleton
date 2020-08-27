
# go-skeleton-project
This is go skeleton project which will serve as template for all new go microservices.

# Setting up new go project

1. Clone `go-skeleton-project` using `git clone git@gitlab.eng.vmware.com:dell-iot/iotss/go-skeleton-project.git`

2. Run `rm -rf .git` from root folder to remove existing `.git` folder

3. Run `git init` to initialize new git folder

4. Create new project in GitLab under `iotss` Group.

5. Create new project URL and add it as your git remote origin into your local git repo

6. Add new `README.md` file

7. Commit your initial file changes and push your changes to remote new project

### Using pre-commit go tools
Make sure you install all these required go tools in your local machine first using below command line. If your go toolings are not installed properly, your `git commit` will fail.
1. Get the tooling on your local
    - go get -u golang.org/x/lint/golint
    - go get -u golang.org/x/tools/cmd/goimports
    - go get -u golang.org/x/tools
2. Clean the go.mod addition of the tooling 
    - go mod tidy
3. Get golang-ci
    - For linux and windows: `curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.30.0` (For others check: https://golangci-lint.run/usage/install/)
4. Make sure to define your GOPATH and PATH 
    - Linux: in your .zshrc or .bashrc or .profile
    `export GOPATH=$HOME/go`
    `export PATH=$HOME/bin:/usr/local/bin:$GOPATH/bin:$PATH`
    - For Windows to create and add GOPATH to PATH env variable 

Configuring pre-commit inside .git folder for each go repo
5. Go to your git repo root directory and run `cp githooks/pre-commit .git/hooks/pre-commit`

6. Provide correct permission to pre-commit file by running `chmod +x .git/hooks/pre-commit`  

### Testing and linting
* `golangci-lint`: golangci-lint is a linters aggregator. It's fast: on average 5 times faster. It's easy to integrate and use, has nice output and has a minimum number of false positives. It supports go modules. Check all avaiable linting tools and configurations here: https://github.com/golangci/golangci-lint
* `go test`: go test will run all unit tests from the project
* `go test ./... -tags=unit -v -coverprofile cover.out`: is the command used to run tests for all files
* For Unit Tests coverage will be checked whether it is greated than 80.0%
* Commits will fail either if unit test fails or coverage is less than 80.0%

### How to run, build, test and lint go files on your local machine
1. Run `go build` or `go mod tidy` to download all required packages.

2. Run by passing command line arguments as follow, 
    * `go run main.go -secret=TEST_SECRET -stage=dev`
    
3. Run `go build` and run `./go-skeleton-project -secret=IotssSaltSecrets -stage=dev`

4. To run specific unit test, `cd` into your `directory` where your test files live and run `go test --tags=unit -run=YOUR_TEST_FUNCTION_NAME -v`

5. To check code coverage, run follow commands:
    * `go test ./... -v -tags=unit -coverprofile cover.out`
    * `go tool cover -html=cover.out -o cover.html`
    * Open `cover.html` with any browser

6. As convention, unit test file name should be `*_unit_test.go` and place this line `// +build unit` at the top of unit test file.

7. As convention, integration test file name should be `*_integration_test.go` and place this line `// +build integration` at the top of integration test file.
    * Every `*_integration_test.go` should have `init` method which should call `test\service\setup.Setup` which will check export IOTSS_TEST_SECRET as global variable once we set it as environment variable

8. To run all unit and integration tests, 
    * `go test ./... -v -tags=unit`
    * `go test ./... -v -tags=integration`

9. To run golangci-lint linter on `go` files in your project, simply run `golangci-lint run`. Make sure you have already installed `golangci-lint` in your `$GOPATH`.

### How to Encrypt and Decrypt password

1. `util\crypto.go` has methods to encrypt and decrypt
2. `Encrypt`: pass plain text and key which will give the encrypted password
3. `Decrypt`: pass encrypted password and key to get the actual password 
4. Refer to `util\crypto_unit_test.go` which has sample method which will encrypt plaint text with key (saltstring) and decrypt


### Adding new command line flags
If you want to pass more command line flags or arguments to your go binary, you need to add those flags to `initFlag()` function inside `service > setup.go` file. Please, refer to existing implementation of `initFlag()` function. For now, there are only two flags available: `-secret` and `-stage`.


# Folder structures

```
├── Dockerfile
├── README.md
├── asset
│   ├── cert
│   └── config
│       └── dev_application.json
├── githooks
│   └── pre-commit
├── gitlab-ci.yml
├── go.mod
├── go.sum
├── main.go
├── main_integration_test.go
├── main_unit_test.go
├── model
│   ├── appConfigModel.go
│   └── appEnvModel.go
├── repository
├── service
│   ├── setup.go
│   ├── setup_integration_test.go
│   └── setup_unit_test.go
├── shared
│   └── global.go
├── test
│   ├── service
│   │   └── setup.go
│   ├── shared
│   │   └── global.go
│   └── testdata
└── util
    ├── crypto.go
    ├── crypto_unit_test.go
    ├── fileio.go
    └── logger.go
```

### Setting up Go’s new dependency management system using go mod in new project for first time

This is only applicable for go version v1.11 and above. Go modules are officially supported in v1.14 and we can start using it in production.

1. Run `go mod init YOUR_HTTPS_GITLAB_REMOTE_REPO_URL_WITHOUT_HTTPS_PREFIX`
2. Run `go mod tidy`. This will prune any no-longer-needed dependencies from `go.mod` and add any dependencies needed for project. You can also run `go build` which will download all required modules too.
3. Every time you want to import local folder in your `.go` file, you need to use module name that you use to initialize go module with `go mod init YOUR_HTTPS_GITLAB_REMOTE_REPO_URL_WITHOUT_HTTPS_PREFIX` followed by your local folder name. For example, if you want to import model folder to one of `.go` files, `import "gitlab.eng.vmware.com/dell-iot/iotss/api/iotss-api-server/model"`
4. Run `go mod vendor`. This resets the main module's vendor directory to include all packages needed to build and test all of the module's packages based on the state of the go.mod files and Go source code.
5. Run excutable generated by `go build` using 

## Adding license content to source files

In order to automatically license content for all sources in this project

*  Get the addlicense tool using:  **go get -u -v github.com/tpryan/addlicense**
*  From the parent folder of the project run the following command: **addlicense -v -f HEADER .**

