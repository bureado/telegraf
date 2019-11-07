package psi

import (
	"github.com/prometheus/procfs"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

type PSIStats struct {
	Log telegraf.Logger
}

func (*PSIStats) Description() string {
	return "Read pressure stall information (PSI) for CPU, Memory and I/O"
}

func (*PSIStats) SampleConfig() string {
	return ``
}

func (s *PSIStats) Gather(acc telegraf.Accumulator) error {

	fs, _ := procfs.NewFS("/proc")

	fields := make(map[string]interface{})

	vals, _ := fs.PSIStatsForResource("io")

        fields["Avg10"] = vals.Some.Avg10
        fields["Avg60"] = vals.Some.Avg60
        fields["Total"] = vals.Some.Total

	acc.AddFields("psiIO", fields, nil)

	return nil
}

func init() {
	inputs.Add("psi", func() telegraf.Input {
		return &PSIStats{}
	})
}
