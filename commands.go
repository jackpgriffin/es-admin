package main

import (
	"fmt"
	elasticgo "github.com/mattbaird/elastigo/lib"
	"os"
	"text/tabwriter"
)

/*{
  "cluster_name" : "elasticsearch",
  "status" : "green",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 0,
  "active_shards" : 0,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 0
}
*/
func PrintHealth(health elasticgo.ClusterHealthResponse) error {
	w := tabwriter.NewWriter(os.Stdout, 2, 1, 3, ' ', 0)
	fmt.Fprint(w, "cluster\tstatus\ttimed_out\tnodes\tdata\tprimary\tactive\trelo\tinit\tunassign\n")
	fmt.Fprintf(w, "%s\t%s\t%t\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n",
		health.ClusterName, health.Status, health.TimedOut, health.NumberOfNodes, health.NumberOfDataNodes, health.ActivePrimaryShards, health.ActiveShards,
		health.RelocatingShards, health.InitializingShards, health.UnassignedShards)

	w.Flush()
	return nil
}
