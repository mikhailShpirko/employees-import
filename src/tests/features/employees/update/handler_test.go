package tests

import (
	employees "employees-import/features/employees"
	update_handler "employees-import/features/employees/update"
	common "employees-import/tests/features/employees"
	"testing"

	"github.com/google/uuid"
)

func Test_Employees_Update_Handle_ValidData_Updated(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockUpdateEmployeeRepository{PayrollExists: false, IdExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := update_handler.Handle(employee, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Update Employee Handler returned error %v`, err)
		return
	}

	switch updateResult := result.(type) {
	case update_handler.Updated:
		return
	case update_handler.PayrollNumberAlreadyExists:
		t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
	case update_handler.ValidationErrors:
		t.Fatalf(`Unexpected result ValidationErrors %v`, updateResult.Errors)
	case update_handler.EmployeeNotExists:
		t.Fatalf(`Unexpected result EmployeeNotExists`)
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Update_Handle_InvalidValidData_ValidationErrors(t *testing.T) {
	employee := employees.Employee{}

	repository := MockUpdateEmployeeRepository{PayrollExists: false, IdExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := update_handler.Handle(employee, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Update Employee Handler returned error %v`, err)
	}

	switch result.(type) {
	case update_handler.Updated:
		t.Fatalf(`Unexpected result Updated`)
	case update_handler.PayrollNumberAlreadyExists:
		t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
	case update_handler.ValidationErrors:
		return
	case update_handler.EmployeeNotExists:
		t.Fatalf(`Unexpected result EmployeeNotExists`)
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Update_Handle_ExistingPayroll_PayrollNumberAlreadyExists(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockUpdateEmployeeRepository{PayrollExists: true, IdExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := update_handler.Handle(employee, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Update Employee Handler returned error %v`, err)
	}

	switch updateResult := result.(type) {
	case update_handler.Updated:
		t.Fatalf(`Unexpected result Updated`)
	case update_handler.PayrollNumberAlreadyExists:
		return
	case update_handler.ValidationErrors:
		t.Fatalf(`Unexpected result ValidationErrors %v`, updateResult.Errors)
	case update_handler.EmployeeNotExists:
		t.Fatalf(`Unexpected result EmployeeNotExists`)
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Update_Handle_NonExistingEmployee_EmployeeNotExists(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockUpdateEmployeeRepository{PayrollExists: false, IdExists: false}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := update_handler.Handle(employee, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Update Employee Handler returned error %v`, err)
	}

	switch updateResult := result.(type) {
	case update_handler.Updated:
		t.Fatalf(`Unexpected result Updated`)
	case update_handler.PayrollNumberAlreadyExists:
		t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
	case update_handler.ValidationErrors:
		t.Fatalf(`Unexpected result ValidationErrors %v`, updateResult.Errors)
	case update_handler.EmployeeNotExists:
		return
	default:
		t.Fatalf("Unsupported result")
	}
}

func Test_Employees_Update_Handle_IsPayrollNumberExistReturnsError_Error(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockFailPayrollNumberExistsUpdateEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := update_handler.Handle(employee, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Update Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailPayrollNumberExists" {
		t.Fatalf(`Update Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Update_Handle_IsIdExistReturnsError_Error(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockFailIdExistsUpdateEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := update_handler.Handle(employee, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Update Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailIdExists" {
		t.Fatalf(`Update Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Update_Handle_UpdateReturnsError_Error(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockFailUpdateUpdateEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := update_handler.Handle(employee, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Update Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUpdate" {
		t.Fatalf(`Update Employee Handler returned different error %v`, err.Error())
	}
}

func Test_Employees_Update_Handle_UnitOfWorkReturnsError_Error(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockUpdateEmployeeRepository{PayrollExists: false, IdExists: true}

	unitOfWork := common.MockFailUnitOfWork{}

	result, err := update_handler.Handle(employee, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Update Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUnitOfWork" {
		t.Fatalf(`Update Employee Handler returned different error %v`, err.Error())
	}
}
