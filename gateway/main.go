package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/stack-labs/stack-rpc"
	"github.com/stack-labs/stack-rpc/cli"
	"github.com/stack-labs/stack-rpc/cmd"
	"github.com/stack-labs/stack-rpc/registry/etcd"

	"github.com/stack-labs/starter-kit/pkg/gateway/api"
	"github.com/stack-labs/starter-kit/pkg/gateway/platform"
	"github.com/stack-labs/starter-kit/pkg/gateway/plugin"
)

//App Info Vars
var (
	GitCommit string
	GitTag    string
	BuildDate string

	name        = "gateway"
	description = "A micro api gateway"
	version     = "0.0.1"
)

// var defination
var (
	args        []string //original parameters from os.Args
	mainFlags   []string //mainFlags parsed from os.Args
	subCmd      []string //subCmd parsed from os.Args
	subCmdFlags []string //subCmdFlags for subCmd parsed from os.Args

	flagNameSets = map[string]int{} //flags name cache for quick query

	ErrHelp                    = errors.New("flag: help requested")
	ErrParsedOver              = errors.New("arguments: parsed over")
	ErrParsedDoubleStrike      = errors.New("warning: unexpacted flag --")
	ErrParsedNoMainFlagValue   = errors.New("warning: no value found in Mian flag")
	ErrParsedNoSubCmdFlagValue = errors.New("warning: no value found in SubCmd flag")
	errorLastMainFlag          error
	errorLastSubCmdFlag        error

	defaultAPICmd   = "api"
	supportedWEBCmd = "web"
)

func main() {
	Init(
		stack.Registry(etcd.NewRegistry()),
	)
}

func Init(options ...stack.Option) {
	app := cmd.App()
	Setup(app, options...)

	regularArguments(cmd.App())

	cmd.Init(
		cmd.Name(name),
		cmd.Description(description),
		cmd.Version(buildVersion()),
	)
}

// Setup sets up a cli.App
func Setup(app *cli.App, options ...stack.Option) {
	// Add the various commands
	app.Commands = append(app.Commands, api.Commands(options...)...)

	// add the init command for our internal operator
	app.Commands = append(app.Commands, cli.Command{
		Name:  "init",
		Usage: "Run the micro operator",
		Action: func(c *cli.Context) {
			platform.Init(c)
		},
		Flags: []cli.Flag{},
	})

	// boot micro runtime
	app.Action = platform.Run

	setup(app)
}

func buildVersion() string {
	microVersion := version

	if GitTag != "" {
		microVersion = GitTag
	}

	if GitCommit != "" {
		microVersion += fmt.Sprintf("-%s", GitCommit)
	}

	if BuildDate != "" {
		microVersion += fmt.Sprintf("-%s", BuildDate)
	}

	return microVersion
}

func setup(app *cli.App) {
	app.Flags = append(app.Flags,
		cli.BoolFlag{
			Name:  "local",
			Usage: "Enable local only development",
		},
		cli.BoolFlag{
			Name:   "enable_acme",
			Usage:  "Enables ACME support via Let's Encrypt. ACME hosts should also be specified.",
			EnvVar: "MICRO_ENABLE_ACME",
		},
		cli.StringFlag{
			Name:   "acme_hosts",
			Usage:  "Comma separated list of hostnames to manage ACME certs for",
			EnvVar: "MICRO_ACME_HOSTS",
		},
		cli.StringFlag{
			Name:   "acme_provider",
			Usage:  "The provider that will be used to communicate with Let's Encrypt. Valid options: autocert, certmagic",
			EnvVar: "MICRO_ACME_PROVIDER",
		},
		cli.BoolFlag{
			Name:   "enable_tls",
			Usage:  "Enable TLS support. Expects cert and key file to be specified",
			EnvVar: "MICRO_ENABLE_TLS",
		},
		cli.StringFlag{
			Name:   "tls_cert_file",
			Usage:  "Path to the TLS Certificate file",
			EnvVar: "MICRO_TLS_CERT_FILE",
		},
		cli.StringFlag{
			Name:   "tls_key_file",
			Usage:  "Path to the TLS Key file",
			EnvVar: "MICRO_TLS_KEY_FILE",
		},
		cli.StringFlag{
			Name:   "tls_client_ca_file",
			Usage:  "Path to the TLS CA file to verify clients against",
			EnvVar: "MICRO_TLS_CLIENT_CA_FILE",
		},
		cli.StringFlag{
			Name:   "api_address",
			Usage:  "Set the api address e.g 0.0.0.0:8080",
			EnvVar: "MICRO_API_ADDRESS",
		},
		cli.StringFlag{
			Name:   "gateway_address",
			Usage:  "Set the micro default gateway address e.g. :9094",
			EnvVar: "MICRO_GATEWAY_ADDRESS",
		},
		cli.StringFlag{
			Name:   "api_handler",
			Usage:  "Specify the request handler to be used for mapping HTTP requests to services; {api, proxy, rpc}",
			EnvVar: "MICRO_API_HANDLER",
		},
		cli.StringFlag{
			Name:   "api_namespace",
			Usage:  "Set the namespace used by the API e.g. com.example.api",
			EnvVar: "MICRO_API_NAMESPACE",
		},
		cli.BoolFlag{
			Name:   "auto_update",
			Usage:  "Enable automatic updates",
			EnvVar: "MICRO_AUTO_UPDATE",
		},
		cli.BoolTFlag{
			Name:   "report_usage",
			Usage:  "Report usage statistics",
			EnvVar: "MICRO_REPORT_USAGE",
		},
		cli.StringFlag{
			Name:   "namespace",
			Usage:  "Set the micro service namespace",
			EnvVar: "MICRO_NAMESPACE",
			Value:  "stack.rpc",
		},
	)

	plugins := plugin.Plugins()

	for _, p := range plugins {
		if flags := p.Flags(); len(flags) > 0 {
			app.Flags = append(app.Flags, flags...)
		}

		if cmds := p.Commands(); len(cmds) > 0 {
			app.Commands = append(app.Commands, cmds...)
		}
	}

	before := app.Before

	app.Before = func(ctx *cli.Context) error {
		if len(ctx.String("api_handler")) > 0 {
			api.Handler = ctx.String("api_handler")
		}
		if len(ctx.String("api_address")) > 0 {
			api.Address = ctx.String("api_address")
		}
		if len(ctx.String("api_namespace")) > 0 {
			api.Namespace = ctx.String("api_namespace")
		}
		for _, p := range plugins {
			if err := p.Init(ctx); err != nil {
				return err
			}
		}

		// now do previous before
		return before(ctx)
	}
}

func constrainSubCmd(subCmd []string) string {

	lenth := len(subCmd)

	switch lenth {
	case 0:
		subCmd = append(subCmd, defaultAPICmd)
		break
	case 1:
		if !(strings.EqualFold(subCmd[0], defaultAPICmd) ||
			strings.EqualFold(subCmd[0], supportedWEBCmd) ||
			strings.EqualFold(subCmd[0], "version") ||
			strings.EqualFold(subCmd[0], "v") ||
			strings.EqualFold(subCmd[0], "help") ||
			strings.EqualFold(subCmd[0], "h")) {
			subCmd[0] = defaultAPICmd
		}
		break
	default:
		subCmd = []string{defaultAPICmd}
	}

	return subCmd[0]
}

func regularArguments(app *cli.App) {

	//point to original parameters
	args = os.Args[1:]

	//cache the flag name for quick searching
	for idx, f := range app.Flags {
		flagNameSets[f.GetName()] = idx
	}

	for _, item := range args {
		seen, err := parseOne(item, app.Flags)
		if seen {
			args = args[1:]
			continue
		}
		if !seen && err == nil {
			break
		}
	}

	var newArgs []string
	newArgs = append(newArgs, os.Args[0])
	newArgs = append(newArgs, mainFlags...)
	newArgs = append(newArgs, constrainSubCmd(subCmd))
	newArgs = append(newArgs, subCmdFlags...)

	os.Args = newArgs
}

// parseOne parses one flag. It reports whether a flag was seen.
func parseOne(s string, appFlags []cli.Flag) (bool, error) {
	if len(args) == 0 {
		return false, ErrParsedOver
	}

	//check last errorflags is novalue
	if s[0] != '-' {
		if errors.Is(errorLastMainFlag, ErrParsedNoMainFlagValue) {
			mainFlags = append(mainFlags, s)
			errorLastMainFlag = nil
			return true, nil
		}
		if errors.Is(errorLastSubCmdFlag, ErrParsedNoSubCmdFlagValue) {
			subCmdFlags = append(subCmdFlags, s)
			errorLastSubCmdFlag = nil
			return true, nil
		}

	}
	errorLastMainFlag = nil
	errorLastSubCmdFlag = nil

	// merge subcmd
	if len(s) < 2 || s[0] != '-' {
		subCmd = append(subCmd, s)
		return true, nil
	}
	numMinuses := 1
	if s[1] == '-' {
		numMinuses++
		if len(s) == 2 { // "--" terminates the flags
			//	args = args[1:]
			return true, ErrParsedDoubleStrike
		}
	}
	name := s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return true, errors.New("warning :unexpacted bad flag syntax: " + s)
	}

	hasValue := false
	//value := ""
	for i := 1; i < len(name); i++ { // equals cannot be first
		if name[i] == '=' {
			//		value = name[i+1:]
			hasValue = true // equals found
			name = name[0:i]
			break
		}
	}

	if _, found := flagNameSets[name]; found {
		mainFlags = append(mainFlags, s) //add to mainflag set
		//check the next value
		if !hasValue {
			errorLastMainFlag = ErrParsedNoMainFlagValue
			return true, ErrParsedNoMainFlagValue
		}
		return true, nil
	}

	// clearly we did not find s in appFlags which must be subcmdflag
	//check the next value
	subCmdFlags = append(subCmdFlags, s) //add to flag set
	if !hasValue {
		errorLastMainFlag = ErrParsedNoSubCmdFlagValue
		return true, ErrParsedNoSubCmdFlagValue
	}
	//find a value
	return true, nil

}
