package main

import (
	"github.com/CrackedPoly/AES-implementation-in-Golang/src/action"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func main() {
	var app = &cli.App{
		Name:                 "AES encryption and decryption",
		Usage:                "encrypt / decrypt",
		EnableBashCompletion: true,
		Commands: cli.Commands{
			{
				Name:   "encrypt",
				Usage:  "AES Encrypt",
				Action: action.EncryptAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "mode",
            Usage:    "specify encryption mode: ECB, CBC, CFB, OFB, CTR, GCM",
						Aliases:  []string{"m"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "plainfile",
						Usage:    "path to plain text file",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "keyfile",
						Usage:    "path to key file",
						Aliases:  []string{"k"},
						Required: true,
					},
					&cli.StringFlag{
						Name:    "vifile",
						Usage:   "path to initial vector file",
						Aliases: []string{"v"},
					},
					&cli.StringFlag{
						Name:     "cipherfile",
						Usage:    "path to save cipher text file",
						Aliases:  []string{"c"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "authfile",
						Usage:    "path to authentication file in GCM mode",
						Aliases:  []string{"a"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "tagfile",
						Usage:    "path to tag file",
						Aliases:  []string{"tag"},
						Required: true,
					},
				},
			},
			{
				Name:   "decrypt",
				Usage:  "AES Decrypt",
				Action: action.DecryptAction,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "mode",
            Usage:    "specify decryption mode: ECB, CBC, CFB, OFB, CTR, GCM",
						Aliases:  []string{"m"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "plainfile",
						Usage:    "path to plain text file",
						Aliases:  []string{"p"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "keyfile",
						Usage:    "path to key file",
						Aliases:  []string{"k"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "vifile",
						Usage:    "path to initial vector file",
						Aliases:  []string{"v"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "cipherfile",
						Usage:    "path to save cipher text file",
						Aliases:  []string{"c"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "authfile",
						Usage:    "path to authentication file in GCM mode",
						Aliases:  []string{"a"},
						Required: true,
					},
					&cli.StringFlag{
						Name:     "tagfile",
						Usage:    "path to tag file",
						Aliases:  []string{"tag"},
						Required: true,
					},
				},
			},
		},
		Action: func(ctx *cli.Context) (err error) {
			_ = ctx.App.Command("help").Action(ctx)
			_ = action.AfterAction(ctx)
			return
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err)
	}

}
