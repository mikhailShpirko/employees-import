package tests

import (
	"employees-import/features/employees"
	import_handler "employees-import/features/employees/import"
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

	result, err := import_handler.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, *newEmployee}, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Import Employee Handler returned error %v`, err)
	}

	switch importResult := result.(type) {
	case import_handler.SuccessfullyImported:
		if len(importResult.Result) != 2 {
			t.Fatalf(`Expected 2 successful imports, but returned %v`, importResult.Result)
		}

		if !slices.Contains(importResult.Result, import_handler.EmployeeImportResult{PayrollNumber: existingEmployee.PayrollNumber, Status: import_handler.EMPLOYEE_IMPORT_STATUS_UPDATED, Id: existingEmployee.Id}) {
			t.Fatalf(`%v expected to be updated. Full result %v`, existingEmployee.PayrollNumber, importResult.Result)
		}

		hasCreated := false
		for _, res := range importResult.Result {
			if res.PayrollNumber == newEmployee.PayrollNumber && res.Status == import_handler.EMPLOYEE_IMPORT_STATUS_CREATED {
				hasCreated = true
				break
			}
		}

		if !hasCreated {
			t.Fatalf(`%v expected to be created. Full result %v`, existingEmployee.PayrollNumber, importResult.Result)
		}
	case import_handler.ValidationErrors:
		t.Fatalf(`Unexpected result ValidationErrors %v`, importResult.Errors)
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Update_Import_InvalidData_ValidationErrors(t *testing.T) {

	employee1 := employees.EmployeeData{PayrollNumber: "1"}
	employee2 := employees.EmployeeData{PayrollNumber: "2"}

	payrollIdMap := map[string]uuid.UUID{}

	repository := MockImportEmployeeRepository{PayrollIdMap: payrollIdMap}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := import_handler.Handle([]employees.EmployeeData{employee1, employee2}, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Import Employee Handler returned error %v`, err)
	}

	switch importResult := result.(type) {
	case import_handler.SuccessfullyImported:
		t.Fatalf(`Unexpected result SuccessfullyImported %v`, importResult.Result)

		return
	case import_handler.ValidationErrors:
		if len(importResult.Errors) != 2 {
			t.Fatalf(`Expected 2 successful imports, but returned %v`, importResult.Errors)
		}
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Update_Import_PartiallyInvalidData_ValidationErrors(t *testing.T) {

	employee1 := employees.EmployeeData{PayrollNumber: "1"}
	employee2 := employees.EmployeeData{PayrollNumber: "2"}
	validEmployee := common.ValidNewEmployee()

	payrollIdMap := map[string]uuid.UUID{}

	repository := MockImportEmployeeRepository{PayrollIdMap: payrollIdMap}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := import_handler.Handle([]employees.EmployeeData{employee1, employee2, *validEmployee}, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Import Employee Handler returned error %v`, err)
	}

	switch importResult := result.(type) {
	case import_handler.SuccessfullyImported:
		t.Fatalf(`Unexpected result SuccessfullyImported %v`, importResult.Result)
		return
	case import_handler.ValidationErrors:
		if len(importResult.Errors) != 2 {
			t.Fatalf(`Expected 2 successful imports, but returned %v`, importResult.Errors)
		}
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Update_Import_GetPayrollNumberToIdMapReturnsError_Error(t *testing.T) {

	existingEmployee := common.ValidExistingEmployee(uuid.New())
	newEmployee := common.ValidNewEmployee()

	newEmployee.PayrollNumber = "NEW EMPLOYEE"

	repository := MockFailGetPayrollNumberToIdMapImportEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := import_handler.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, *newEmployee}, &repository, &unitOfWork)

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

	result, err := import_handler.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, *newEmployee}, &repository, &unitOfWork)

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

	result, err := import_handler.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, *newEmployee}, &repository, &unitOfWork)

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

	result, err := import_handler.Handle([]employees.EmployeeData{existingEmployee.EmployeeData, *newEmployee}, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Import Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUnitOfWork" {
		t.Fatalf(`Import Employee Handler returned different error %v`, err.Error())
	}
}
