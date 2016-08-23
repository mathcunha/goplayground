package main

import (
	"fmt"
	"github.com/mathcunha/CloudCapacitor/capacitor"
)

func main() {
	vms, _ := capacitor.LoadTypes("/home/vagrant/go/src/github.com/mathcunha/CloudCapacitor/config/dspace.yml")

	e := capacitor.MockExecutor{"/home/vagrant/go/src/github.com/mathcunha/CloudCapacitor/config/wordpress_cpu_mem.csv", nil}
	e.Load()
	slos := []float32{10000, 20000, 30000, 40000, 50000}
	modes := []string{"Strict", "Mem", "Mem", "CPU", "Price"}
	labels := []string{"Strict-Type", "Strict-Prop", "Mem", "CPU", "Price"}
	equi := []bool{false, true, true, true, true}
	fmt.Println("mode,slo,property,value")
	for i, mode := range modes {
		if i == 2 {
			vms, _ = capacitor.LoadTypes("/home/vagrant/go/src/github.com/mathcunha/CloudCapacitor/config/dspace_nocat.yml")
		}
		for _, slo := range slos {
			dspace := capacitor.NewDeploymentSpace(&vms, 7.0, 4)
			destDS := dspace.CapacityBy(mode)
			_ = *dspace.CalcMaxSLO(e, []string{"100", "200", "300", "400", "500", "600", "700", "800", "900", "1000"}, []float32{slo})

			c, a, d := capacitor.VerifyReflexionModel2(dspace.Configurations(), destDS, equi[i])
			fmt.Printf("\"%s\",%.0f,\"%s\",%d\n", labels[i], slo/1000, "convergence", c)
			fmt.Printf("\"%s\",%.0f,\"%s\",%d\n", labels[i], slo/1000, "abscence", d)
			fmt.Printf("\"%s\",%.0f,\"%s\",%d\n", labels[i], slo/1000, "divergence", a)
			fmt.Printf("\"%s\",%.0f,\"%s\",%.4f\n", labels[i], slo/1000, "metric", capacitor.CalcCapacityAccuracy(c, d, a))
			//fmt.Printf("\"%s\",%.0f,%d,%d,%d,%d\n", mode, slo/1000, c, a, d, c+a)
			//fmt.Printf("%s\tslo[%.0f]-\tconvergence:%d, absence:%d, divergence:%d\n", mode, slo/1000, c, a, d)
		}
	}
}
