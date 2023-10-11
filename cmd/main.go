package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/elliotwutingfeng/ccwc"
	"github.com/urfave/cli/v3"
)

var cmd = &cli.Command{
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "c",
			Usage: "Count number of bytes in a file",
		},
		&cli.BoolFlag{
			Name:  "l",
			Usage: "Count number of lines in a file",
		},
		&cli.BoolFlag{
			Name:  "w",
			Usage: "Count number of words in a file",
		},
		&cli.BoolFlag{
			Name:  "m",
			Usage: "Count number of characters in a file",
		},
	},
	Name:                   "ccwc",
	Usage:                  "Count bytes, lines, words, and characters in a file",
	UseShortOptionHandling: true,
	Action: func(cCtx *cli.Context) error {
		var filename string
		var content []byte
		var err error

		if cCtx.NArg() > 0 {
			filename = cCtx.Args().Get(0)
			content, err = os.ReadFile(filename)
			if err != nil {
				fmt.Println(fmt.Errorf("ccwc: %s: No such file or directory", filename).Error())
				return err
			}
		} else {
			content, err = io.ReadAll(os.Stdin)
			if err != nil {
				panic(err)
			}
		}
		c, l, w, m := ccwc.Counts(content)
		var output []string
		if cCtx.Bool("c") {
			output = append(output, c)
		}
		if cCtx.Bool("l") {
			output = append(output, l)
		}
		if cCtx.Bool("w") {
			output = append(output, w)
		}
		if cCtx.Bool("m") {
			output = append(output, m)
		}

		if len(filename) > 0 {
			output = append(output, filename)
			if len(output) == 1 {
				output = []string{l, w, c, filename}
			}
		} else {
			if len(output) == 0 {
				output = []string{l, w, c}
			}
		}
		fmt.Printf("   %s\n", strings.Join(output, "   "))
		return nil
	},
}

func main() {
	cmd.Run(context.Background(), os.Args)
}
