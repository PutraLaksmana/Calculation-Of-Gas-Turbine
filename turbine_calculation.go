//* @Author Putra Laksmana - @PutraLaksmana

// ██████╗ ██████╗ ███████╗
// ██╔══██╗██╔══██╗██╔════╝
// ██████╔╝██████╔╝█████╗
// ██╔══██╗██╔═══╝ ██╔══╝
// ██║  ██║██║     ███████╗
// ╚═╝  ╚═╝╚═╝     ╚══════╝

// 50 55 54 52 41  4C 41 4B 53 4D 41 4E 41

// ░█▀█░█░█░▀█▀░█▀▄░█▀█░░░█░░░█▀█░█░█░█▀▀░█▄█░█▀█░█▀█░█▀█
// ░█▀▀░█░█░░█░░█▀▄░█▀█░░░█░░░█▀█░█▀▄░▀▀█░█░█░█▀█░█░█░█▀█
// ░▀░░░▀▀▀░░▀░░▀░▀░▀░▀░░░▀▀▀░▀░▀░▀░▀░▀▀▀░▀░▀░▀░▀░▀░▀░▀░▀

package main

import (
	"fmt"
	"math"
)

func main() {

	watermark := `
   ░█▀█░█░█░▀█▀░█▀▄░█▀█░░░█░░░█▀█░█░█░█▀▀░█▄█░█▀█░█▀█░█▀█
   ░█▀▀░█░█░░█░░█▀▄░█▀█░░░█░░░█▀█░█▀▄░▀▀█░█░█░█▀█░█░█░█▀█
   ░▀░░░▀▀▀░░▀░░▀░▀░▀░▀░░░▀▀▀░▀░▀░▀░▀░▀▀▀░▀░▀░▀░▀░▀░▀░▀░▀
`
	fmt.Println(watermark)

	const beban = 35   //MW
	const Flow_Gas = 4 //kg/s
	const LHV = 10000  //Kj/kg
	const P2 = 13      //bar//(Pressure Outlet Comp.)
	const P3 = 12      //kg/cm2 //Pressur Combustor
	const T1 = 305     //kelvin //(Temperature Inlet Comp)
	const T2 = 700     //kelvin //(Temperature Outlet Comp.)
	const T4 = 653     //kelvin //(Temperature Outlet Turbine)
	const P4 = 1.015   //bar //(Pressure Outlet Turbin)
	const K = 1.4      // konstanta

	Temperature := []float64{200, 210, 220, 230, 240, 250, 260, 270, 280, 285, 290, 295, 298, 300, 305, 310, 315, 320, 325, 330, 340, 350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470, 480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580, 590, 600, 610, 620, 630, 640, 650, 660, 670, 680, 690, 700, 710, 720, 730, 740, 750, 760, 780, 800, 820, 840, 860, 880, 900, 920, 940, 960, 980, 1000, 1020, 1040, 1060, 1080, 1100, 1120, 1140, 1160, 1180, 1200, 1220, 1240, 1260, 1280, 1300, 1320, 1340, 1360, 1380, 1400, 1420, 1440, 1460, 1480, 1500, 1520, 1540, 1560, 1580, 1600, 1620, 1640, 1660, 1680, 1700, 1750, 1800, 1850, 1900, 1950, 2000, 2050, 2100, 2150, 2200, 2250}
	Entalphi := []float64{199.97, 209.97, 219.97, 230.02, 240.02, 250.05, 260.09, 270.11, 280.13, 285.14, 290.16, 295.17, 298.18, 300.19, 305.22, 310.24, 315.27, 320.29, 325.31, 330.34, 340.42, 350.49, 360.58, 370.67, 380.77, 390.88, 400.98, 411.12, 421.26, 431.43, 441.61, 451.80, 462.02, 472.24, 482.49, 492.74, 503.02, 513.32, 523.63, 533.98, 544.35, 555.74, 565.17, 575.59, 586.04, 596.52, 607.02, 617.53, 628.07, 638.63, 649.22, 659.84, 670.47, 681.14, 691.82, 702.52, 713.27, 724.04, 734.82, 745.62, 756.44, 767.29, 778.18, 800.03, 821.95, 843.98, 866.08, 888.27, 910.56, 932.93, 955.38, 977.92, 1000.55, 1023.25, 1046.04, 1068.89, 1091.85, 1114.86, 1137.89, 1161.07, 1184.28, 1207.57, 1230.92, 1254.34, 1277.79, 1301.31, 1324.93, 1348.55, 1372.24, 1395.97, 1419.76, 1443.60, 1467.49, 1491.44, 1515.42, 1539.44, 1563.51, 1587.63, 1611.79, 1635.97, 1660.23, 1684.51, 1708.82, 1733.17, 1757.57, 1782.00, 1806.46, 1830.96, 1855.50, 1880.1, 1941.6, 2003.3, 2065.3, 2127.4, 2189.7, 2252.1, 2314.6, 2377.7, 2440.3, 2503.2, 2566.4}
	relative_pressure := []float64{0.3363, 0.3987, 0.4690, 0.5477, 0.6355, 0.7329, 0.8405, 0.9590, 1.0889, 1.1584, 1.2311, 1.3068, 1.3543, 1.3860, 1.4686, 1.5546, 1.6442, 1.7375, 1.8345, 1.9352, 2.149, 2.379, 2.626, 2.892, 3.176, 3.481, 3.806, 4.153, 4.522, 4.915, 5.332, 5.775, 6.245, 6.742, 7.268, 7.824, 8.411, 9.031, 9.684, 10.37, 11.10, 11.86, 12.66, 13.50, 14.38, 15.31, 16.28, 17.30, 18.36, 19.84, 20.64, 21.86, 23.13, 24.46, 25.85, 27.29, 28.80, 30.38, 32.02, 33.72, 35.50, 37.35, 39.27, 43.35, 47.75, 52.59, 57.60, 63.09, 68.98, 75.29, 82.05, 89.28, 97.00, 105.2, 114.0, 123.4, 133.3, 143.9, 155.2, 167.1, 179.7, 193.1, 207.2, 222.2, 238.0, 254.7, 272.3, 290.8, 310.4, 330.9, 352.5, 375.3, 399.1, 424.2, 450.5, 478.0, 506.9, 537.1, 568.8, 601.9, 636.5, 672.8, 710.5, 750.0, 791.2, 834.1, 878.9, 925.6, 974.2, 1025, 1161, 1310, 1475, 1655, 1852, 2068, 2303, 2559, 2837, 3138, 3464}
	//u := []float64{142.56, 149.69, 156.82, 164.00, 171.13, 178.28, 185.45, 192.60, 199.75, 203.33, 206.91, 210.49, 212.64, 214.07, 217.67, 221.25, 224.85, 228.42, 232.02, 235.61, 242.82, 250.02, 257.24, 264.46, 271.69, 278.93, 286.16, 293.43, 300.69, 307.99, 315.30, 322.62, 329.97, 337.32, 344.70, 352.08, 359.49, 366.92, 374.36, 381.84, 389.34, 396.86, 404.42, 411.97, 419.55, 427.15, 434.78, 442.42, 450.09, 457.78, 465.50, 473.25, 481.01, 488.81, 496.62, 504.45, 512.33, 520.23, 528.14, 536.07, 544.02, 551.99, 560.01, 576.12, 592.30, 608.59, 624.95, 641.40, 657.95, 674.58, 691.28, 708.08, 725.02, 741.98, 758.94, 776.10, 793.36, 810.62, 827.88, 845.33, 862.79, 880.35, 897.91, 915.57, 933.33, 951.09, 968.95, 986.90, 1004.76, 1022.82, 1040.88, 1058.94, 1077.10, 1095.26, 1113.52, 1131.77, 1150.13, 1168.49, 1186.95, 1205.41, 1223.87, 1242.43, 1260.99, 1279.65, 1298.30, 1316.96, 1335.72, 1354.48, 1373.24, 1392.7, 1439.8, 1487.2, 1534.9, 1582.6, 1630.6, 1678.7, 1726.8, 1775.3, 1823.8, 1872.4, 1921.3}
	//relative_specific_volume := []float64{1707.0, 1512.0, 1346.0, 1205.0, 1084.0, 979.0, 887.8, 808.0, 738.0, 706.1, 676.1, 647.9, 631.9, 621.2, 596.0, 572.3, 549.8, 528.6, 508.4, 489.4, 454.1, 422.2, 393.4, 367.2, 343.4, 321.5, 301.6, 283.3, 266.6, 251.1, 236.8, 223.6, 211.4, 200.1, 189.5, 179.7, 170.6, 162.1, 154.1, 146.7, 139.7, 133.1, 127.0, 121.2, 115.7, 110.6, 105.8, 101.2, 96.92, 92.84, 88.99, 85.34, 81.89, 78.61, 75.50, 72.56, 69.76, 67.07, 64.53, 62.13, 59.82, 57.63, 55.54, 51.64, 48.08, 44.84, 41.85, 39.12, 36.61, 34.31, 32.18, 30.22, 28.40, 26.73, 25.17, 23.72, 23.29, 21.14, 19.98, 18.896, 17.886, 16.946, 16.064, 15.241, 14.470, 13.747, 13.069, 12.435, 11.835, 11.275, 10.747, 10.247, 9.780, 9.337, 8.919, 8.526, 8.153, 7.801, 7.468, 7.152, 6.854, 6.569, 6.301, 6.046, 5.804, 5.574, 5.355, 5.147, 4.949, 4.761, 4.328, 3.994, 3.601, 3.295, 3.022, 2.776, 2.555, 2.356, 2.175, 2.012, 1.864}
	//s := []float64{1.29559, 1.34444, 1.39105, 1.43557, 1.47824, 1.51917, 1.55848, 1.59634, 1.63279, 1.65055, 1.66802, 1.68515, 1.69528, 1.70203, 1.71865, 1.73498, 1.75106, 1.76690, 1.78249, 1.79783, 1.82790, 1.85708, 1.88543, 1.91313, 1.94001, 1.96633, 1.99194, 2.01699, 2.04142, 2.06533, 2.08870, 2.11161, 2.13407, 2.15604, 2.17760, 2.19876, 2.21952, 2.23993, 2.25997, 2.27967, 2.29906, 2.31809, 2.33685, 2.35531, 2.37348, 2.39140, 2.40902, 2.42644, 2.44356, 2.46048, 2.47716, 2.49364, 2.50985, 2.52589, 2.54175, 2.55731, 2.57277, 2.58810, 2.60319, 2.61803, 2.63280, 2.64737, 2.66176, 2.69013, 2.71787, 2.74504, 2.77170, 2.79783, 2.82344, 2.84856, 2.87324, 2.89748, 2.92128, 2.94468, 2.96770, 2.99034, 3.01260, 3.03449, 3.05608, 3.07732, 3.09825, 3.11883, 3.13916, 3.15916, 3.17888, 3.19834, 3.21751, 3.23638, 3.25510, 3.27345, 3.29160, 3.30959, 3.32724, 3.34474, 3.36200, 3.37901, 3.39586, 3.41247, 3.42892, 3.44516, 3.46120, 3.47712, 3.49276, 3.50829, 3.52364, 3.53879, 3.55381, 3.56867, 3.58335, 3.5979, 3.6336, 3.6684, 3.7023, 3.7354, 3.7677, 3.7994, 3.8303, 3.8605, 3.8901, 3.9191, 3.9474}

	Enthalpi_h1 := interpolasi_linear(Temperature, Entalphi, T1)
	fmt.Println("Enthalpi H1 = ", Enthalpi_h1, "kJ/kg")

	//
	Enthalpi_h2 := interpolasi_linear(Temperature, Entalphi, T2)
	fmt.Println("Enthalpi H2 = ", Enthalpi_h2, "kJ/kg")

	Pr4 := interpolasi_linear(Temperature, relative_pressure, T4)
	Pr3 := Pr4 * (P3 / P4)
	Enthalpi_h3 := interpolasi_linear(relative_pressure, Entalphi, Pr3)
	T3 := interpolasi_linear(Entalphi, Temperature, Enthalpi_h3)
	fmt.Println("Enthalpi H3(Combuster) =", Enthalpi_h3, "kJ/kg")

	Entalphi_h4 := interpolasi_linear(Temperature, Entalphi, T4)
	fmt.Println("Enthalpi H4(Exhaust Turbin) =", Entalphi_h4, "kJ/kg")

	T_2 := teoritis_T(T1, P2, P4, K)
	Entalphi_h2_ := interpolasi_linear(Temperature, Entalphi, T_2)
	fmt.Println("Enthalpi H'2 = ", Entalphi_h2_, "kJ/kg")

	T_4 := teoritis_T(T3, P4, P3, K)
	Entalphi_h4_ := interpolasi_linear(Temperature, Entalphi, T_4)
	fmt.Println("Enthalpi H'4 = ", Entalphi_h4_, "kJ/kg")

	E_Kompresor := Efisiensi_kompresor(Entalphi_h2_, Enthalpi_h1, Enthalpi_h2)
	fmt.Println("Effesiensi compressor =", E_Kompresor*100, "%")

	E_Turbine := Efisiensi_turbine(Entalphi_h4_, Enthalpi_h3, Entalphi_h4)
	fmt.Println("Effesiensi Turbine =", E_Turbine*100, "%")

	kalor := kalor_masuk(Flow_Gas, LHV)
	fmt.Println("Suplay kalor ke ruang bakar =", kalor, "Kj/s")

	aliran_udara := Ma(kalor, Flow_Gas, Enthalpi_h3, Enthalpi_h2)
	fmt.Println("Laju aliran udara =", aliran_udara, "Kg/s")

	kerja_kompresor := Wca(aliran_udara, Enthalpi_h2, Enthalpi_h1, E_Kompresor)
	fmt.Println("Kerja kompresor =", kerja_kompresor, "Kj/s")

	kerja_turbine := Wta(aliran_udara, Flow_Gas, Enthalpi_h3, Entalphi_h4, E_Turbine)
	fmt.Println("Kerja Turbine =", kerja_turbine, "Kj/s")

	effesiensi_Thermal := Nth(kerja_turbine, kerja_kompresor, kalor)
	fmt.Println("Effesiensi Thermal =", effesiensi_Thermal*100, "%")

	energi_gas_buang_hrsg := energi_gas_buang(aliran_udara, Flow_Gas, Enthalpi_h3, Entalphi_h4, beban)
	fmt.Println("Energi Gas Buang untuk HRSG =", energi_gas_buang_hrsg, "Kj/s")

}

func energi_gas_buang(Ma float64, Mf float64, h3 float64, h4 float64, PT float64) (result_energi float64) {
	pt := PT * 1000
	result_energi = (Ma+Mf)*(h3-h4) - pt
	return
}

func Nth(Wta float64, Wca float64, Qin float64) (result_Nth float64) {
	result_Nth = (Wta - Wca) / Qin
	return
}
func Wta(Ma float64, Mf float64, h3 float64, h4 float64, Nta float64) (result_Wta float64) {
	result_Wta = (Ma + Mf) * (h3 - h4) * Nta
	return
}
func Wca(Ma float64, enthalpi2 float64, enthalpi1 float64, Nca float64) (result_Wca float64) {
	result_Wca = (Ma * (enthalpi2 - enthalpi1)) / Nca
	return
}

func Ma(Qin float64, Mf float64, enthalpi3 float64, enthalpi2 float64) (result_Ma float64) {
	result_Ma = (Qin - (Mf * enthalpi3)) / (enthalpi3 - enthalpi2)
	return
}

func kalor_masuk(flowGas float64, lhv float64) (result_kalor_masuk float64) {
	result_kalor_masuk = flowGas * lhv
	return
}

func teoritis_T(x float64, Pa float64, Pb float64, konstanta float64) (result_T float64) {
	k := (konstanta - 1) / konstanta
	result_T = x * math.Pow(Pa/Pb, k)
	return
}

func Efisiensi_kompresor(a float64, x float64, y float64) (R_Efisiensi_kompresor float64) {
	R_Efisiensi_kompresor = (a - x) / (y - x)
	return
}

func Efisiensi_turbine(a float64, x float64, y float64) (R_Efisiensi_turbine float64) {
	R_Efisiensi_turbine = (x - y) / (x - a)
	return
}

func interpolasi_linear(xArr []float64, yArr []float64, X float64) (result_Enthalpi float64) {
	// Flag to indicate if the x exists in the slice
	found := false

	// Loop through the slice and check if the x exists
	for _, v := range xArr {
		if v == X {
			found = true
			break
		}
	}

	if found {
		index := 0
		for i, v := range xArr {
			if v == X {
				index = i
				result_Enthalpi = yArr[index]
				break
			}
		}
	} else {
		var below_X, above_X float64

		//below_X = math.MinInt64
		//above_X = math.MaxInt64

		for _, value := range xArr {
			if value < X {
				below_X = value
			} else if value > X {
				above_X = value
				break
			}
		}

		var below_Y, above_Y float64

		i_below_Y := 0
		for c, j := range xArr {
			if j == below_X {
				i_below_Y = c
				below_Y = yArr[i_below_Y]
				break
			}
		}
		i_above_Y := 0
		for d, k := range xArr {
			if k == above_X {
				i_above_Y = d
				above_Y = yArr[i_above_Y]
				break
			}
		}

		//                  ( T1- Tbawah ) ( h atas – h bawah )
		// rumus : h1 = ---------------------------------------------- + h bawah
		// 			               (Tatas- Tbawah )

		result_Enthalpi = (X-below_X)*(above_Y-below_Y)/(above_X-below_X) + below_Y
	}
	return
}
