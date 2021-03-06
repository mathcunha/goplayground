package main

import (
	"fmt"
	"github.com/mathcunha/CloudCapacitor/capacitor"
)

func main() {
	vms, _ := capacitor.LoadTypes("/home/vagrant/go/src/github.com/mathcunha/CloudCapacitor/config/dspace.yml")
	vmsNCat, _ := capacitor.LoadTypes("/home/vagrant/go/src/github.com/mathcunha/CloudCapacitor/config/dspace_nocat.yml")

	e := capacitor.MockExecutor{"/home/vagrant/go/src/github.com/mathcunha/CloudCapacitor/config/terasort_cpu_mem.csv", nil}
	e.Load()
	slos := []float32{100000, 200000, 300000, 400000}

	modes := []string{"Strict", "CPU", "Mem", "Price", "CPU"}
	equi := []bool{false, false, false, false, false}
	bypass := []bool{true, false, false, false, true}

	fmt.Println("mode,slo,property,value,length")
	for i, mode := range modes {
		for _, slo := range slos {
			var dspace capacitor.DeploymentSpace
			if bypass[i] {
				dspace = capacitor.NewDeploymentSpace(&vms, 7.0, 4)
			} else {
				dspace = capacitor.NewDeploymentSpace(&vmsNCat, 7.0, 4)
			}
			destDS := dspace.CapacityBy(mode)
			_ = *dspace.CalcMaxSLO(e, []string{"10000000", "20000000", "30000000", "40000000", "50000000"}, []float32{slo})

			c, a, d := capacitor.VerifyReflexionModel2(dspace.Configurations(), destDS, equi[i])
			fmt.Printf("\"Prop:%s Cat:%t\",%.0f,%q,%d\n", modes[i], bypass[i], slo/1000, "convergence", c)
			fmt.Printf("\"Prop:%s Cat:%t\",%.0f,%q,%d\n", modes[i], bypass[i], slo/1000, "abscence", d)
			fmt.Printf("\"Prop:%s Cat:%t\",%.0f,%q,%d\n", modes[i], bypass[i], slo/1000, "divergence", a)
			fmt.Printf("\"Prop:%s Cat:%t\",%.0f,%q,%.4f\n", modes[i], bypass[i], slo/1000, "metric", capacitor.CalcCapacityAccuracy(c, d, a))
			fmt.Printf("\"Prop:%s Cat:%t\",%.0f,%q,%.4f\n", modes[i], bypass[i], slo/1000, "v1", capacitor.CalcCapacityAccuracyV1(c, d, a))
			fmt.Printf("\"Prop:%s Cat:%t\",%.0f,%q,%.4f\n", modes[i], bypass[i], slo/1000, "v2", capacitor.CalcCapacityAccuracyV2(c, d, a))
			fmt.Printf("\"Prop:%s Cat:%t\",%.0f,%q,%.4f,%d\n", modes[i], bypass[i], slo/1000, "v3", capacitor.CalcCapacityAccuracyV3(c, d, a), c+d)
			//fmt.Printf("\"%s\",%.0f,\"%s\",%d\n", labels[i], slo/10000, "soma", c+a)
			//fmt.Printf("\"%s\",%.0f,%d,%d,%d,%d\n", mode, slo/1000, c, a, d, c+a)
			//fmt.Printf("%s\tslo[%.0f]-\tconvergence:%d, absence:%d, divergence:%d\n", mode, slo/1000, c, a, d)
		}
	}
}
