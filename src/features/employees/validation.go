package employees

import (
	common "employees-import/features/common"
	"net/mail"
	"time"
)

type EMPLOYEE_VALIDATION_ERROR int

const (
	EMPLOYEE_VALIDATION_ERROR_EMPTY_PAYROLL_NUMBER EMPLOYEE_VALIDATION_ERROR = iota
	EMPLOYEE_VALIDATION_ERROR_EMPTY_FORENAMES
	EMPLOYEE_VALIDATION_ERROR_EMPTY_SURNAME
	EMPLOYEE_VALIDATION_ERROR_EMPTY_DATE_OF_BIRTH
	EMPLOYEE_VALIDATION_ERROR_EMPTY_TELEPHONE_NUMBER
	EMPLOYEE_VALIDATION_ERROR_INVALID_TELEPHONE_NUMBER
	EMPLOYEE_VALIDATION_ERROR_EMPTY_MOBILE_NUMBER
	EMPLOYEE_VALIDATION_ERROR_INVALID_MOBILE_NUMBER
	EMPLOYEE_VALIDATION_ERROR_EMPTY_EMAIL
	EMPLOYEE_VALIDATION_ERROR_INVALID_EMAIL
	EMPLOYEE_VALIDATION_ERROR_EMPTY_START_DATE
	EMPLOYEE_VALIDATION_ERROR_FUTURE_DATE_OF_BIRTH
	EMPLOYEE_VALIDATION_ERROR_DATE_OF_BIRTH_AFTER_START_DATE
)

func (validationError EMPLOYEE_VALIDATION_ERROR) String() string {
	return [...]string{"EMPTY_PAYROLL_NUMBER",
		"EMPTY_FORENAMES",
		"EMPTY_SURNAME",
		"EMPTY_DATE_OF_BIRTH",
		"EMPTY_TELEPHONE_NUMBER",
		"INVALID_TELEPHONE_NUMBER",
		"EMPTY_MOBILE_NUMBER",
		"INVALID_MOBILE_NUMBER",
		"EMPTY_EMAIL",
		"INVALID_EMAIL",
		"EMPTY_START_DATE",
		"FUTURE_DATE_OF_BIRTH",
		"DATE_OF_BIRTH_AFTER_START_DATE"}[validationError]
}

func Validate(employeeData *EmployeeData, today time.Time) []EMPLOYEE_VALIDATION_ERROR {
	validationErrors := []EMPLOYEE_VALIDATION_ERROR{}

	if common.IsStringEmptyOrWhiteSpace(employeeData.PayrollNumber) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_EMPTY_PAYROLL_NUMBER)
	}

	if common.IsStringEmptyOrWhiteSpace(employeeData.Forenames) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_EMPTY_FORENAMES)
	}

	if common.IsStringEmptyOrWhiteSpace(employeeData.Surname) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_EMPTY_SURNAME)
	}

	isDateOfBirthValid := true
	if employeeData.DateOfBirth.IsZero() {
		isDateOfBirthValid = false
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_EMPTY_DATE_OF_BIRTH)
	} else if !common.IsTimeBefore(employeeData.DateOfBirth, today) {
		isDateOfBirthValid = false
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_FUTURE_DATE_OF_BIRTH)
	}

	if common.IsStringEmptyOrWhiteSpace(employeeData.TelephoneNumber) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_EMPTY_TELEPHONE_NUMBER)
	} else if !isValidPhoneNumber(employeeData.TelephoneNumber) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_INVALID_TELEPHONE_NUMBER)
	}

	if common.IsStringEmptyOrWhiteSpace(employeeData.MobileNumber) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_EMPTY_MOBILE_NUMBER)
	} else if !isValidPhoneNumber(employeeData.MobileNumber) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_INVALID_MOBILE_NUMBER)
	}

	if common.IsStringEmptyOrWhiteSpace(employeeData.Email) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_EMPTY_EMAIL)
	} else if !isValidEmail(employeeData.Email) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_INVALID_EMAIL)
	}

	if employeeData.StartDate.IsZero() {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_EMPTY_START_DATE)
	} else if isDateOfBirthValid && !common.IsTimeBefore(employeeData.DateOfBirth, employeeData.StartDate) {
		validationErrors = append(validationErrors, EMPLOYEE_VALIDATION_ERROR_DATE_OF_BIRTH_AFTER_START_DATE)
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func isValidPhoneNumber(str string) bool {
	//for now assuming that we check the full number (12 digit) strinctly formatter as numeric
	//no spaces dashes brackets allowed
	//for now assume that formatting will be a frontend task
	return len(str) == 12 && common.IsStringNumeric(str)
}

func isValidEmail(str string) bool {
	_, err := mail.ParseAddress(str)
	return err == nil
}
