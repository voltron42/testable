package icli

import (
	"flag"
	"io"
	"time"

	"github.com/codegangsta/cli"
)

var (
	AppHelpTemplate        = cli.AppHelpTemplate
	BashCompletionFlag     = cli.BashCompletionFlag
	CommandHelpTemplate    = cli.CommandHelpTemplate
	HelpFlag               = cli.HelpFlag
	HelpPrinter            = cli.HelpPrinter
	SubcommandHelpTemplate = cli.SubcommandHelpTemplate
	VersionFlag            = cli.VersionFlag
	VersionPrinter         = cli.VersionPrinter
)

type Factory interface {
	DefaultAppComplete(c *Context)
	ShowAppHelp(c *Context)
	ShowCommandCompletions(ctx *Context, command string)
	ShowCommandHelp(ctx *Context, command string)
	ShowCompletions(c *Context)
	ShowVersion(c *Context)
	ShowSubcommandHelp(c *Context)

	NewApp() *App
	NewContext(app *App, set *flag.FlagSet, parentCtx *Context) *Context
	NewMultiError(err ...error) MultiError
}

type App interface {
	Name() string
	SetName(name string)
	Usage() string
	SetUsage(usage string)
	Version() string
	SetVersion(version string)
	Commands() []Command
	SetCommands(commands []Command)
	Flags() []flag.Flag
	SetFlags(flags []flag.Flag)
	EnableBashCompletion() bool
	SetEnableBashCompletion(enablebashcompletion bool)
	HideHelp() bool
	SetHideHelp(hidehelp bool)
	HideVersion() bool
	SetHideVersion(hideversion bool)
	BashComplete() func(context Context)
	SetBashComplete(bashcomplete func(context Context))
	Before() func(context Context) error
	SetBefore(before func(context Context) error)
	After() func(context Context) error
	SetAfter(after func(context Context) error)
	Action() func(context Context)
	SetAction(action func(context Context))
	CommandNotFound() func(context Context, command string)
	SetCommandNotFound(commandnotfound func(context Context, command string))
	Compiled() time.Time
	SetCompiled(compiled time.Time)
	Authors() []Author
	SetAuthors(authors []Author)
	Author() string
	SetAuthor(author string)
	Email() string
	SetEmail(email string)
	Writer() io.Writer
	SetWriter(writer io.Writer)

	Command(name string) *Command
	Run(arguments []string) (err error)
	RunAndExitOnError()
	RunAsSubcommand(ctx *Context) (err error)
}

type Args interface {
	AllArgs() []string

	First() string
	Get(n int) string
	Present() bool
	Swap(from, to int) error
	Tail() []string
}

type Author interface {
	Name() string
	SetName(name string)
	Email() string
	SetEmail(email string)

	String() string
}

type BoolFlag interface {
	Name() string
	SetName(name string)
	Usage() string
	SetUsage(usage string)
	EnvVar() string
	SetEnvVar(envvar string)

	Apply(set *flag.FlagSet)
	String() string
}

type BoolTFlag interface {
	Name() string
	SetName(name string)
	Usage() string
	SetUsage(usage string)
	EnvVar() string
	SetEnvVar(envvar string)

	Apply(set *flag.FlagSet)
	String() string
}

type Command interface {
	Name() string
	SetName(name string)
	ShortName() string
	SetShortName(shortname string)
	Aliases() []string
	SetAliases(aliases []string)
	Usage() string
	SetUsage(usage string)
	Description() string
	SetDescription(description string)
	BashComplete() func(context Context)
	SetBashComplete(bashcomplete func(context Context))
	Before() func(context Context) error
	SetBefore(before func(context Context) error)
	After() func(context Context) error
	SetAfter(after func(context Context) error)
	Action() func(context Context)
	SetAction(action func(context Context))
	Subcommands() []Command
	SetSubcommands(subcommands []Command)
	Flags() []flag.Flag
	SetFlags(flags []flag.Flag)
	SkipFlagParsing() bool
	SetSkipFlagParsing(skipflagparsing bool)
	HideHelp() bool
	SetHideHelp(hidehelp bool)

	HasName(name string) bool
	Names() []string
	Run(ctx *Context) error
}

type Context interface {
	App() *App
	SetApp(app *App)
	Command() Command
	SetCommand(command Command)

	Args() Args
	Bool(name string) bool
	BoolT(name string) bool
	Duration(name string) time.Duration
	FlagNames() (names []string)
	Float64(name string) float64
	Generic(name string) interface{}
	GlobalBool(name string) bool
	GlobalDuration(name string) time.Duration
	GlobalFlagNames() (names []string)
	GlobalGeneric(name string) interface{}
	GlobalInt(name string) int
	GlobalIntSlice(name string) []int
	GlobalIsSet(name string) bool
	GlobalString(name string) string
	GlobalStringSlice(name string) []string
	Int(name string) int
	IntSlice(name string) []int
	IsSet(name string) bool
	NumFlags() int
	String(name string) string
	StringSlice(name string) []string
}

type DurationFlag interface {
	Name() string
	SetName(name string)
	Value() time.Duration
	SetValue(value time.Duration)
	Usage() string
	SetUsage(usage string)
	EnvVar() string
	SetEnvVar(envvar string)

	Apply(set *flag.FlagSet)
	String() string
}

type Float64Flag interface {
	Name() string
	SetName(name string)
	Value() float64
	SetValue(value float64)
	Usage() string
	SetUsage(usage string)
	EnvVar() string
	SetEnvVar(envvar string)

	Apply(set *flag.FlagSet)
	String() string
}

type GenericFlag interface {
	Name() string
	SetName(name string)
	Value() cli.Generic
	SetValue(value cli.Generic)
	Usage() string
	SetUsage(usage string)
	EnvVar() string
	SetEnvVar(envvar string)

	Apply(set *flag.FlagSet)
	String() string
}

type IntFlag interface {
	Name() string
	SetName(name string)
	Value() int
	SetValue(value int)
	Usage() string
	SetUsage(usage string)
	EnvVar() string
	SetEnvVar(envvar string)

	Apply(set *flag.FlagSet)
	String() string
}

type IntSlice interface {
	AllInts() []int
	SetAllInts(allints []int)

	Set(value string) error
	String() string
	Value() []int
}

type IntSliceFlag interface {
	Name() string
	SetName(name string)
	Value() *IntSlice
	SetValue(value *IntSlice)
	Usage() string
	SetUsage(usage string)
	EnvVar() string
	SetEnvVar(envvar string)

	Apply(set *flag.FlagSet)
	String() string
}

type MultiError interface {
	Errors() []error
	SetErrors(errors []error)

	Error() string
}

type StringFlag interface {
	Name() string
	SetName(name string)
	Value() string
	SetValue(value string)
	Usage() string
	SetUsage(usage string)
	EnvVar() string
	SetEnvVar(envvar string)

	Apply(set *flag.FlagSet)
	String() string
}

type StringSlice interface {
	AllStrings() []string
	SetAllStrings(allstrings []string)

	Set(value string) error
	String() string
	Value() []string
}

type StringSliceFlag interface {
	Name() string
	SetName(name string)
	Value() *StringSlice
	SetValue(value *StringSlice)
	Usage() string
	SetUsage(usage string)
	EnvVar() string
	SetEnvVar(envvar string)

	Apply(set *flag.FlagSet)
	String() string
}
