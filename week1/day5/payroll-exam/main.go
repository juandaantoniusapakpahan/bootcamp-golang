package main

import (
	"fmt"
	"payroll-exam/employee"
	"payroll-exam/payroll"
	"payroll-exam/salarymatrix"
)

func main() {

	employ := employee.NewEmployee()
	salaryMatrix := salarymatrix.NewSalaryMatrix()
	payRool := payroll.NewPayroll(salaryMatrix, employ)

	salaryMatrix.Add(1, 5000000, 100000, 150000, 1000000)
	salaryMatrix.Add(2, 9000000, 200000, 300000, 2000000)
	salaryMatrix.Add(3, 15000000, 400000, 600000, 3000000)

	//  L = Laki-laki; P = Perempuan

	// TASK 1 ADD EMPLOYEE
	// employ.Add("Rizal Automation", "L", 1, true)
	// employ.Add("Bozokal Pipeline", "L", 1, false)
	// employ.Add("Cocok Engine", "P", 1, false)
	// employ.Add("Duku Asem", "P", 1, false)
	// employ.Add("Epososo Good game well play", "P", 2, true)
	// employ.Add("Food Ku Mantap", "L", 2, true)

	// TASK 3 ADD PAYROLL
	// fmt.Println("==== Show Payroll By User Id ====")
	// payRool.Add("employ-0", 19, 1)
	// payRool.Add("employ-1", 19, 1)
	// payRool.Add("employ-2", 20, 0)
	// payRool.Add("employ-3", 20, 0)
	// payRool.Add("employ-4", 20, 0)
	// payRool.Add("employ-5", 20, 0)

	// TASK 4 SHOW PAYROLL BY USER ID
	// payRool.ShowPayrollById("employ-0")
	// payRool.ShowPayrollById("employ-1")
	// payRool.ShowPayrollById("employ-2")
	// payRool.ShowPayrollById("employ-3")
	// payRool.ShowPayrollById("employ-4")
	// payRool.ShowPayrollById("employ-5")

	// fmt.Println()

	// TASK 2 SHOW ALL EMPLOYEE
	// employ.ShowAll()
	// fmt.Println()

	// TASK 5 SHOW SALARY MATRIX
	// salaryMatrix.GetAll()

	/* ===== RESULT PAYROll ===== */
	// employ-0 : {{employ-0 Rizal Automation L 1 true} [{payroll-0 5e+06 100000 3.85e+06 July-2023 8.75e+06}]}
	// employ-1 : {{employ-1 Bozokal Pipeline L 1 false} [{payroll-1 5e+06 100000 2.85e+06 July-2023 7.75e+06}]}
	// employ-2 : {{employ-2 Cocok Engine P 1 false} [{payroll-2 5e+06 0 3e+06 July-2023 8e+06}]}
	// employ-3 : {{employ-3 Duku Asem P 1 false} [{payroll-3 5e+06 0 3e+06 July-2023 8e+06}]}
	// employ-4 : {{employ-4 Epososo Good game well play P 2 true} [{payroll-4 9e+06 0 6e+06 July-2023 1.5e+07}]}
	// employ-5 : {{employ-5 Food Ku Mantap L 2 true} [{payroll-5 9e+06 0 8e+06 July-2023 1.7e+07}]}

	/* ===== RESULT EMPLOYEE ==== */
	// 	EmployeId: employ-0, Name: Rizal Automation, Gender: L, Grade: 1, IsMarried: true
	// EmployeId: employ-1, Name: Bozokal Pipeline, Gender: L, Grade: 1, IsMarried: false
	// EmployeId: employ-2, Name: Cocok Engine, Gender: P, Grade: 1, IsMarried: false
	// EmployeId: employ-3, Name: Duku Asem, Gender: P, Grade: 1, IsMarried: false
	// EmployeId: employ-4, Name: Epososo Good game well play, Gender: P, Grade: 2, IsMarried: true
	// EmployeId: employ-5, Name: Food Ku Mantap, Gender: L, Grade: 2, IsMarried: true

	/* ===== SalaryMatrix ===== */
	// {matrix-0 1 5e+06 100000 150000 1e+06}
	// {matrix-1 2 9e+06 200000 300000 2e+06}
	// {matrix-2 3 1.5e+07 400000 600000 3e+06}

	var choose int

admin:
	for {
		fmt.Println("1. ADD EMPLOYEE")
		fmt.Println("2. Tampilkan semua karyawan")
		fmt.Println("3. Add payroll By user Id")
		fmt.Println("4. SHOW PAYROLL BY USER ID")
		fmt.Println("5. SalaryMatrix")
		fmt.Println("6. Exit")

		fmt.Scan(&choose)

		switch choose {
		case 1:
			// employ.Add("Rizal Automation", "L", 1, true)
			var (
				name      string
				gender    string
				grade     int
				isMarried bool
			)
			fmt.Println("Masukan: nama, gender(L/P), grade(1/2/3), ismarried(true/false)")
			fmt.Scanln(&name, &gender, &grade, &isMarried)
			employ.Add(name, gender, grade, isMarried)
			continue

		case 2:
			fmt.Println("==== Show All Employee ====")
			employ.ShowAll()
			fmt.Println()
			continue
		case 3:
			// payRool.Add("employ-0", 19, 1)
			var (
				emploId string
				hadir   float64
				absen   float64
			)
			fmt.Println("Masukan: employeid, jumlah hadir, jumlah tidak hadir")
			fmt.Scan(&emploId, &hadir, &absen)
			payRool.Add(emploId, hadir, absen)
			continue
		case 4:
			var emploId string
			fmt.Println("Masukan user Id:")
			fmt.Scanln(&emploId)
			fmt.Println("==== Show Payroll By ID ====")
			payRool.ShowPayrollById(emploId)
			fmt.Println()
			fmt.Println()

			continue
		case 5:
			salaryMatrix.GetAll()
			fmt.Println()
			continue
		case 6:
			break admin
		default:
			continue
		}

	}

}
