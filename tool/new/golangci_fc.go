package new

var (
	golangcifc = &FileContent{
		FileName: ".golangci.yml",
		Dir:      ".",
		Content: `linters-settings:
  depguard:
    list-type: blacklist
    packages:
      # logging is allowed only by logutils.Log, logrus
      # is allowed to use only in logutils package
      - github.com/sirupsen/logrus
    packages-with-error-message:
      - github.com/sirupsen/logrus: "logging is allowed only by logutils.Log"
  dupl:
    threshold: 100
  funlen:
    lines: 100
    statements: 50
  goconst:
    min-len: 2
    min-occurrences: 2
  gocritic:
    enabled-tags:
      - diagnostic
      - experimental
      - opinionated
      - performance
      - style
    disabled-checks:
      - dupImport # https://github.com/go-critic/go-critic/issues/845
      - ifElseChain
      - octalLiteral
      - whyNoLint
      - wrapperFunc
      - commentedOutCode
  gocyclo:
    min-complexity: 15
  goimports:
    local-prefixes: github.com/golangci/golangci-lint
  golint:
    min-confidence: 0
  gomnd:
    settings:
      mnd:
        # don't include the "operation" and "assign"
        checks: argument,case,condition,return
  govet:
    check-shadowing: true
    settings:
      printf:
        funcs:
          - (github.com/golangci/golangci-lint/pkg/logutils.Log).Infof
          - (github.com/golangci/golangci-lint/pkg/logutils.Log).Warnf
          - (github.com/golangci/golangci-lint/pkg/logutils.Log).Errorf
          - (github.com/golangci/golangci-lint/pkg/logutils.Log).Fatalf
  lll:
    line-length: 100
  maligned:
    suggest-new: true
  misspell:
    locale: US
  nolintlint:
    allow-leading-space: true # don't require machine-readable nolint directives (i.e. with no leading space)
    allow-unused: false # report any unused nolint directives
    require-explanation: false # don't require an explanation for nolint directives
    require-specific: false # don't require nolint directives to be specific about which linter is being skipped

linters:
  # please, do not use {{.SingleMark}}enable-all{{.SingleMark}}: it's deprecated and will be removed soon.
  # inverted configuration with {{.SingleMark}}enable-all{{.SingleMark}} and {{.SingleMark}}disable{{.SingleMark}} is not scalable during updates of golangci-lint
  disable-all: true
  enable:
    - bodyclose
    - deadcode
    - depguard
    - dogsled
    - dupl
    - errcheck
    - funlen
   # - gochecknoinits
   # - goconst
    - gocritic
    - gocyclo
    - gofmt
    - goimports
    - golint
   # - gomnd
    - goprintffuncname
    - gosec
    - gosimple
    - govet
    - ineffassign
    - interfacer
    - lll
    - misspell
    - nakedret
    - nolintlint
    - rowserrcheck
    - scopelint
    - staticcheck
    - structcheck
    - stylecheck
    - typecheck
    - unconvert
    - unparam
    - unused
    - varcheck
    - whitespace

  # don't enable:
  # - asciicheck
  # - gochecknoglobals
  # - gocognit
    - godot
  # - godox
  # - goerr113
  # - maligned
  # - nestif
  # - prealloc
  # - testpackage
  # - wsl

issues:
  # Excluding configuration per-path, per-linter, per-text and per-source
  # exclude-rules:
  #  - path: _test\.go
  #    linters:
  #      - gomnd

run:
  skip-dirs:
    - tool/new
  skip-files:
    - _test\.go

# golangci.com configuration
# https://github.com/golangci/golangci/wiki/Configuration
service:
  golangci-lint-version: 1.23.x # use the fixed version to not introduce new linters unexpectedly
  prepare:
    - echo "here I can run custom commands, but no preparation needed for this repo"`,
	}
)

func initGolangCiFiles() {
	Files = append(Files, golangcifc)
}
