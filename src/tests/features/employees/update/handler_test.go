package tests

import (
	employees "employees-import/features/employees"
	employees_update "employees-import/features/employees/update"
	common "employees-import/tests/features/employees"
	"testing"

	"github.com/google/uuid"
)

func Test_Employees_Update_Handle_ValidData_Updated(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockUpdateEmployeeRepository{PayrollExists: false, IdExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_update.Handle(employee, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Update Employee Handler returned error %v`, err)
		return
	}

	employees_update.Match(result,
		func(updated employees_update.Updated) error {
			return nil
		},
		func(payrollExists employees_update.PayrollNumberAlreadyExists) error {
			t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
			return nil
		},
		func(validationErrors employees_update.ValidationErrors) error {
			t.Fatalf(`Unexpected result ValidationErrors %v`, validationErrors.Errors)
			return nil
		},
		func(notExists employees_update.EmployeeNotExists) error {
			t.Fatalf(`Unexpected result EmployeeNotExists`)
			return nil
		})
}

func Test_Employees_Update_Handle_InvalidValidData_ValidationErrors(t *testing.T) {
	employee := employees.Employee{}

	repository := MockUpdateEmployeeRepository{PayrollExists: false, IdExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_update.Handle(employee, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Update Employee Handler returned error %v`, err)
	}

	employees_update.Match(result,
		func(updated employees_update.Updated) error {
			t.Fatalf(`Unexpected result Updated`)
			return nil
		},
		func(payrollExists employees_update.PayrollNumberAlreadyExists) error {
			t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
			return nil
		},
		func(validationErrors employees_update.ValidationErrors) error {
			return nil
		},
		func(notExists employees_update.EmployeeNotExists) error {
			t.Fatalf(`Unexpected result EmployeeNotExists`)
			return nil
		})
}

func Test_Employees_Update_Handle_ExistingPayroll_PayrollNumberAlreadyExists(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockUpdateEmployeeRepository{PayrollExists: true, IdExists: true}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_update.Handle(employee, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Update Employee Handler returned error %v`, err)
	}

	employees_update.Match(result,
		func(updated employees_update.Updated) error {
			t.Fatalf(`Unexpected result Updated`)
			return nil
		},
		func(payrollExists employees_update.PayrollNumberAlreadyExists) error {
			return nil
		},
		func(validationErrors employees_update.ValidationErrors) error {
			t.Fatalf(`Unexpected result ValidationErrors %v`, validationErrors.Errors)
			return nil
		},
		func(notExists employees_update.EmployeeNotExists) error {
			t.Fatalf(`Unexpected result EmployeeNotExists`)
			return nil
		})
}

func Test_Employees_Update_Handle_NonExistingEmployee_EmployeeNotExists(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockUpdateEmployeeRepository{PayrollExists: false, IdExists: false}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_update.Handle(employee, &repository, &unitOfWork)

	if err != nil {
		t.Fatalf(`Update Employee Handler returned error %v`, err)
	}

	employees_update.Match(result,
		func(updated employees_update.Updated) error {
			t.Fatalf(`Unexpected result Updated`)
			return nil
		},
		func(payrollExists employees_update.PayrollNumberAlreadyExists) error {
			t.Fatalf(`Unexpected result PayrollNumberAlreadyExists`)
			return nil
		},
		func(validationErrors employees_update.ValidationErrors) error {
			t.Fatalf(`Unexpected result ValidationErrors %v`, validationErrors.Errors)
			return nil
		},
		func(notExists employees_update.EmployeeNotExists) error {
			return nil
		})
}

func Test_Employees_Update_Handle_IsPayrollNumberExistReturnsError_Error(t *testing.T) {
	employee := common.ValidExistingEmployee(uuid.New())

	repository := MockFailPayrollNumberExistsUpdateEmployeeRepository{}

	unitOfWork := common.MockSuccessUnitOfWork{}

	result, err := employees_update.Handle(employee, &repository, &unitOfWork)

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

	result, err := employees_update.Handle(employee, &repository, &unitOfWork)

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

	result, err := employees_update.Handle(employee, &repository, &unitOfWork)

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

	result, err := employees_update.Handle(employee, &repository, &unitOfWork)

	if err == nil {
		t.Fatalf(`Update Employee Handler didn't return error %v`, result)
	}

	if err.Error() != "FailUnitOfWork" {
		t.Fatalf(`Update Employee Handler returned different error %v`, err.Error())
	}
}
