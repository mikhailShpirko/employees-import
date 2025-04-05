package tests

import (
	"employees-import/features/employees"
	employees_import "employees-import/features/employees/import"
	common "employees-import/tests/features/employees"
	"slices"
	"testing"

	"github.com/google/uuid"
)

func Test_Employees_Update_Import_ValidData_SuccessfullyImported(t *testing.T) {

	existingEmployee := common.ValidExistingEmployee(uuid.New())
	newEmployee := common.ValidNewEmployee()

	newEmployee.PayrollNumber = "NEW EMPLOYEE"

	payrollIdMap := map[string]uuid.UUID{existingEmployee.PayrollNumber: existingEmployee.Id}

	repository := MockImportEmployeeRepository{PayrollIdMap: payrollIdMap}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_import.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, newEmployee}, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Import Employee Handler returned error %v`, err)
	}

	employees_import.Match(result,
		func(imported employees_import.SuccessfullyImported) error {
			if len(imported.Result) != 2 {
				t.Fatalf(`Expected 2 successful imports, but returned %v`, imported.Result)
			}

			if !slices.Contains(imported.Result, employees_import.EmployeeImportResult{PayrollNumber: existingEmployee.PayrollNumber, Status: employees_import.EMPLOYEE_IMPORT_STATUS_UPDATED, Id: existingEmployee.Id}) {
				t.Fatalf(`%v expected to be updated. Full result %v`, existingEmployee.PayrollNumber, imported.Result)
			}

			hasCreated := false
			for _, res := range imported.Result {
				if res.PayrollNumber == newEmployee.PayrollNumber && res.Status == employees_import.EMPLOYEE_IMPORT_STATUS_CREATED {
					hasCreated = true
					break
				}
			}

			if !hasCreated {
				t.Fatalf(`%v expected to be created. Full result %v`, existingEmployee.PayrollNumber, imported.Result)
			}
			return nil
		},
		func(validationErrors employees_import.ValidationErrors) error {
			t.Fatalf(`Unexpected result ValidationErrors %v`, validationErrors.Errors)
			return nil
		})
}

func Test_Employees_Update_Import_InvalidData_ValidationErrors(t *testing.T) {

	employee1 := employees.EmployeeData{PayrollNumber: "1"}
	employee2 := employees.EmployeeData{PayrollNumber: "2"}

	payrollIdMap := map[string]uuid.UUID{}

	repository := MockImportEmployeeRepository{PayrollIdMap: payrollIdMap}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_import.Handle([]employees.EmployeeData{employee1, employee2}, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Import Employee Handler returned error %v`, err)
	}

	employees_import.Match(result,
		func(imported employees_import.SuccessfullyImported) error {
			t.Fatalf(`Unexpected result SuccessfullyImported %v`, imported.Result)
			return nil
		},
		func(validationErrors employees_import.ValidationErrors) error {
			if len(validationErrors.Errors) != 2 {
				t.Fatalf(`Expected 2 validation errors, but returned %v`, validationErrors.Errors)
			}
			return nil
		})
}

func Test_Employees_Update_Import_PartiallyInvalidData_ValidationErrors(t *testing.T) {

	employee1 := employees.EmployeeData{PayrollNumber: "1"}
	employee2 := employees.EmployeeData{PayrollNumber: "2"}
	validEmployee := common.ValidNewEmployee()

	payrollIdMap := map[string]uuid.UUID{}

	repository := MockImportEmployeeRepository{PayrollIdMap: payrollIdMap}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_import.Handle([]employees.EmployeeData{employee1, employee2, validEmployee}, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Import Employee Handler returned error %v`, err)
	}

	employees_import.Match(result,
		func(imported employees_import.SuccessfullyImported) error {
			t.Fatalf(`Unexpected result SuccessfullyImported %v`, imported.Result)
			return nil
		},
		func(validationErrors employees_import.ValidationErrors) error {
			if len(validationErrors.Errors) != 2 {
				t.Fatalf(`Expected 2 validation errors, but returned %v`, validationErrors.Errors)
			}
			return nil
		})
}

func Test_Employees_Update_Import_GetPayrollNumberToIdMapReturnsError_Error(t *testing.T) {

	existingEmployee := common.ValidExistingEmployee(uuid.New())
	newEmployee := common.ValidNewEmployee()

	newEmployee.PayrollNumber = "NEW EMPLOYEE"

	repository := MockFailGetPayrollNumberToIdMapImportEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_import.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, newEmployee}, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Import Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailGetPayrollNumberToIdMap" {
		t.Fatalf(`Import Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Update_Import_CreateReturnsError_Error(t *testing.T) {

	existingEmployee := common.ValidExistingEmployee(uuid.New())
	newEmployee := common.ValidNewEmployee()

	newEmployee.PayrollNumber = "NEW EMPLOYEE"

	payrollIdMap := map[string]uuid.UUID{existingEmployee.PayrollNumber: existingEmployee.Id}

	repository := MockFailCreateImportEmployeeRepository{PayrollIdMap: payrollIdMap}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_import.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, newEmployee}, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Import Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailCreate" {
		t.Fatalf(`Import Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Update_Import_UpdateReturnsError_Error(t *testing.T) {

	existingEmployee := common.ValidExistingEmployee(uuid.New())
	newEmployee := common.ValidNewEmployee()

	newEmployee.PayrollNumber = "NEW EMPLOYEE"

	payrollIdMap := map[string]uuid.UUID{existingEmployee.PayrollNumber: existingEmployee.Id}

	repository := MockFailUpdateImportEmployeeRepository{PayrollIdMap: payrollIdMap}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_import.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, newEmployee}, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Import Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUpdate" {
		t.Fatalf(`Import Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Update_Import_UnitOfWorkReturnsError_Error(t *testing.T) {

	existingEmployee := common.ValidExistingEmployee(uuid.New())
	newEmployee := common.ValidNewEmployee()

	newEmployee.PayrollNumber = "NEW EMPLOYEE"

	payrollIdMap := map[string]uuid.UUID{existingEmployee.PayrollNumber: existingEmployee.Id}

	repository := MockImportEmployeeRepository{PayrollIdMap: payrollIdMap}

	unitOfWork := common.MockFailUnitOfWork{}

	result, err := employees_import.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, newEmployee}, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Import Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUnitOfWork" {
		t.Fatalf(`Import Employee Handler returned different error %v`, err.Error())
	}
}
