package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	var (
		Interval string
		Ratio    string
	)

	cmd.Flags().StringVarP(&Interval, "interval", "i", "1000", "ログを吐く間隔(MilliSeconds)")
	cmd.Flags().StringVarP(&Ratio, "ratio", "r", "0/1/0/0", "ログの割合(DEBUG / INFO / WARN / ERROR)")
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type LogRatio struct {
	Debug float64
	Info  float64
	Warn  float64
	Error float64
}
type CmdValues struct {
	Interval float64
	Ratio    LogRatio
}

func parseCmdValues(cmd *cobra.Command) (*CmdValues, error) {
	intervalStr := cmd.Flags().Lookup("interval").Value.String()
	interval, err := strconv.ParseFloat(intervalStr, 64)
	if err != nil {
		return &CmdValues{}, err
	}

	ratioStr := cmd.Flags().Lookup("ratio").Value.String()
	ratioFloat := []float64{}
	for _, ratio := range strings.Split(ratioStr, "/") {
		r, err := strconv.ParseFloat(ratio, 64)
		if err != nil {
			return &CmdValues{}, err
		}
		ratioFloat = append(ratioFloat, r)
	}
	if len(ratioFloat) != 4 {
		return &CmdValues{}, fmt.Errorf("ratioは DEBUG / INFO / WARN / ERROR の値を4つ設定してください。")
	}

	return &CmdValues{
		Interval: interval,
		Ratio: LogRatio{
			Debug: ratioFloat[0],
			Info:  ratioFloat[1],
			Warn:  ratioFloat[2],
			Error: ratioFloat[3],
		},
	}, nil
}

var cmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		cmdValues, err := parseCmdValues(cmd)
		if err != nil {
			panic(err)
		}

		totalRatio := cmdValues.Ratio.Debug + cmdValues.Ratio.Info + cmdValues.Ratio.Warn + cmdValues.Ratio.Error
		debugRatio := cmdValues.Ratio.Debug / totalRatio
		infoRatio := cmdValues.Ratio.Info / totalRatio
		warnRatio := cmdValues.Ratio.Warn / totalRatio

		for range time.Tick(time.Duration(cmdValues.Interval) * time.Millisecond) {
			var logLevel = ""
			randomFloat := rand.Float64()
			if randomFloat <= debugRatio {
				logLevel = "DEBUG"
			} else if randomFloat <= debugRatio+infoRatio {
				logLevel = "INFO"
			} else if randomFloat <= debugRatio+infoRatio+warnRatio {
				logLevel = "WARN"
			} else {
				logLevel = "ERROR"
			}

			log.Printf("%s: xxx\n", logLevel)
		}
	},
}
