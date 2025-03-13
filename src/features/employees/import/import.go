package employees_import

import (
	common "employees-import/features/common"
	employees "employees-import/features/employees"

	"time"

	"github.com/google/uuid"
)

func Handle(employeesData []employees.EmployeeData,
	repository IImportEmployeeRepository,
	unitOfWork common.IUnitOfWork) (IImportEmployeesResult, error) {

	var invalidEmployeeData []InvalidEmployeeData

	now := time.Now()
	for _, employeeData := range employeesData {
		validationErrors := employees.Validate(&employeeData, now)

		if validationErrors != nil {
			invalidEmployeeData = append(invalidEmployeeData,
				InvalidEmployeeData{PayrollNumber: employeeData.PayrollNumber, ValidationErrors: validationErrors})
		}
	}

	if len(invalidEmployeeData) > 0 {
		return ValidationErrors{Errors: invalidEmployeeData}, nil
	}

	existingEmployeesPayrollIdMap, getPayrollNumberToIdMapError := repository.GetPayrollNumberToIdMap()

	if getPayrollNumberToIdMapError != nil {
		return nil, getPayrollNumberToIdMapError
	}

	var importResult []EmployeeImportResult

	for _, employeeData := range employeesData {
		employee := employees.Employee{EmployeeData: employeeData}

		id, exists := existingEmployeesPayrollIdMap[employeeData.PayrollNumber]

		if exists {
			employee.Id = id
			updateError := repository.Update(&employee)

			if updateError != nil {
				return nil, updateError
			}

			importResult = append(importResult, EmployeeImportResult{PayrollNumber: employeeData.PayrollNumber, Status: EMPLOYEE_IMPORT_STATUS_UPDATED, Id: id})

		} else {
			employee.Id = uuid.New()
			createError := repository.Create(&employee)

			if createError != nil {
				return nil, createError
			}

			importResult = append(importResult, EmployeeImportResult{PayrollNumber: employeeData.PayrollNumber, Status: EMPLOYEE_IMPORT_STATUS_CREATED, Id: employee.Id})
		}
	}

	saveChangesError := unitOfWork.SaveChanges()

	if saveChangesError != nil {
		return nil, saveChangesError
	}

	return SuccessfullyImported{Result: importResult}, nil
}
