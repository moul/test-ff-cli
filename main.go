package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/peterbourgon/ff/ffcli"
	"golang.org/x/xerrors"
)

func main() {
	log.SetFlags(0)

	var (
		globalFlags   = flag.NewFlagSet("objectctl", flag.ExitOnError)
		globalVerbose = globalFlags.Bool("v", false, "verbose mode")
		aaaFlags      = flag.NewFlagSet("aaa", flag.ExitOnError)
		aaaFoo        = aaaFlags.Bool("foo", false, "foo")
		bbbFlags      = flag.NewFlagSet("bbb", flag.ExitOnError)
		bbbFoo        = bbbFlags.Bool("foo", false, "foo")
		bbbBar        = bbbFlags.Bool("bar", false, "bar")
		cccFlags      = flag.NewFlagSet("ccc", flag.ExitOnError)
		cccBaz        = cccFlags.Bool("baz", false, "baz")
	)

	aaa := &ffcli.Command{
		Name:    "aaa",
		FlagSet: aaaFlags,
		Exec: func(args []string) error {
			fmt.Printf("verbose=%v, args=%q, foo=%v\n", *globalVerbose, args, *aaaFoo)
			return nil
		},
	}

	bbb := &ffcli.Command{
		Name:    "bbb",
		FlagSet: bbbFlags,
		Exec: func(args []string) error {
			fmt.Printf("verbose=%v, args=%q, foo=%v, bar=%v\n", *globalVerbose, args, *bbbFoo, *bbbBar)
			return nil
		},
	}

	ccc := &ffcli.Command{
		Name:    "ccc",
		FlagSet: cccFlags,
		Exec: func(args []string) error {
			fmt.Printf("verbose=%v, args=%q, baz=%v\n", *globalVerbose, args, *cccBaz)
			return nil
		},
	}

	root := &ffcli.Command{
		Usage:       "just-a-test",
		FlagSet:     globalFlags,
		Subcommands: []*ffcli.Command{aaa, bbb, ccc},
		Exec: func([]string) error {
			return xerrors.New("specify a subcommand")
		},
	}

	if err := root.Run(os.Args[1:]); err != nil {
		log.Fatalf("error: %v", err)
	}
}
