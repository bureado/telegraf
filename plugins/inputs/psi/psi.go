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
	return `# d00d`
}

func (s *PSIStats) Gather(acc telegraf.Accumulator) error {

	// TODO: error handling, constant for /proc
	fs, _ := procfs.NewFS("/proc")

	fields := make(map[string]interface{})

	// TODO: iterator of some sort
	vals, _ := fs.PSIStatsForResource("io")

        fields["IOSome10"] = vals.Some.Avg10
        fields["IOSome60"] = vals.Some.Avg60
        fields["IOFull10"] = vals.Full.Avg10
        fields["IOFull60"] = vals.Full.Avg60
        fields["IOSomeTotal"] = vals.Some.Total
        fields["IOFullTotal"] = vals.Full.Total

	vals, _ = fs.PSIStatsForResource("cpu")

        fields["CPUSome10"] = vals.Some.Avg10
        fields["CPUSome60"] = vals.Some.Avg60
        fields["CPUSomeTotal"] = vals.Some.Total

	vals, _ = fs.PSIStatsForResource("memory")

        fields["MEMFull10"] = vals.Full.Avg10
        fields["MEMFull60"] = vals.Full.Avg60
        fields["MEMSome10"] = vals.Some.Avg10
        fields["MEMSome60"] = vals.Some.Avg60
        fields["MEMSomeTotal"] = vals.Some.Total
        fields["MEMFullTotal"] = vals.Full.Total

	// TODO: alternatively, have a full/some field or 10/60/total fields
	acc.AddFields("psi", fields, nil)

	return nil
}

func init() {
	inputs.Add("psi", func() telegraf.Input {
		return &PSIStats{}
	})
}
