package main

import (
	"fmt"
	"github.com/mathcunha/CloudCapacitor/capacitor"
)

func main() {
	vms, _ := capacitor.LoadTypes("/home/vagrant/go/src/github.com/mathcunha/CloudCapacitor/config/dspace_nocat.yml")

	//e := capacitor.MockExecutor{"/home/vagrant/go/src/github.com/mathcunha/CloudCapacitor/config/wordpress_cpu_mem.csv", nil}
	e := capacitor.MockExecutor{"/home/vagrant/go/src/github.com/mathcunha/CloudCapacitor/config/terasort_cpu_mem.csv", nil}
	e.Load()
	//slos := []float32{10000, 20000, 30000, 40000, 50000}
	slos := []float32{0.01, 0.02, 0.03, 0.04, 0.05}
	//slos := []float32{30000}
	//modes := []string{"Strict", "Mem"}
	modes := []string{"Strict", "Mem", "CPU", "Price"}
	equi := []bool{false, true, true, true}
	fmt.Println("mode,slo,property,value")
	for i, mode := range modes {
		for _, slo := range slos {
			dspace := capacitor.NewDeploymentSpace(&vms, 7.0, 4)
			destDS := dspace.CapacityBy(mode)
			//_ = *dspace.CalcMaxSLO(e, []string{"100", "200", "300", "400", "500", "600", "700", "800", "900", "1000"}, []float32{slo})
			_ = *dspace.CalcMaxSLO(e, []string{"10000000", "20000000", "30000000", "40000000", "50000000"}, []float32{slo})

			c, a, d := capacitor.VerifyReflexionModel2(dspace.Configurations(), destDS, equi[i])
			fmt.Printf("\"%s\",%.2f,\"%s\",%d\n", mode, slo, "convergence", c)
			fmt.Printf("\"%s\",%.2f,\"%s\",%d\n", mode, slo, "abscence", a)
			fmt.Printf("\"%s\",%.2f,\"%s\",%d\n", mode, slo, "divergence", d)
			fmt.Printf("\"%s\",%.2f,\"%s\",%d\n", mode, slo, "soma", c+a)
			//fmt.Printf("\"%s\",%.0f,%d,%d,%d,%d\n", mode, slo/1000, c, a, d, c+a)
			//fmt.Printf("%s\tslo[%.0f]-\tconvergence:%d, absence:%d, divergence:%d\n", mode, slo/1000, c, a, d)
		}
	}
}
