package tests

import (
	"context"
	employees "employees-import/features/employees"
	persistence "employees-import/persistence"
	"path/filepath"
	"slices"

	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type EmployeeRepositoryTestSuite struct {
	suite.Suite
	connectionString        string
	container               *postgres.PostgresContainer
	employees               []employees.Employee
	updatedEmployeeIndex    int
	deletedEmployeeIndex    int
	unmodifiedEmployeeIndex int
}

func createPostgresContainer(ctx context.Context) (*postgres.PostgresContainer, string, error) {
	pgContainer, err := postgres.Run(ctx,
		"postgres:17.4-alpine3.21",
		postgres.WithInitScripts(filepath.Join("..", "..", "persistence", "migrations", "1.Init.sql")),
		postgres.WithDatabase("test-employees-import"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, "", err
	}
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, "", err
	}

	return pgContainer, connStr, nil
}

func (suite *EmployeeRepositoryTestSuite) SetupSuite() {

	t := suite.T()

	container, connectionString, err := createPostgresContainer(t.Context())

	if err != nil {
		t.Fatalf(`Failed to create test container %v`, err)
	}

	suite.connectionString = connectionString
	suite.container = container

	for i := 1; i <= 10; i++ {
		suite.employees = append(suite.employees, createEmployee(uuid.New(), strconv.Itoa(i)))
	}

	suite.updatedEmployeeIndex = 2
	suite.deletedEmployeeIndex = 7
	suite.unmodifiedEmployeeIndex = 0
}

func (suite *EmployeeRepositoryTestSuite) TearDownSuite() {
	t := suite.T()
	if err := suite.container.Terminate(t.Context()); err != nil {
		t.Fatalf("Error terminating container: %s", err)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_01_Create() {

	t := suite.T()
	repo, uow, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Write Repository is nil`)
	}

	if uow == nil {
		t.Fatalf(`Unit of Work is nil`)
	}

	for _, emp := range suite.employees {
		repoErr := repo.Create(&emp)

		if repoErr != nil {
			t.Fatalf(`Failed to Create %v : %v`, emp, repoErr)
		}
	}

	uowErr := uow.SaveChanges()

	if uowErr != nil {
		t.Fatalf(`Unit of Work failed to save changes %v`, uowErr)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_02_Update() {

	t := suite.T()
	repo, uow, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Write Repository is nil`)
	}

	if uow == nil {
		t.Fatalf(`Unit of Work is nil`)
	}

	newDateOfBirth, _ := time.Parse(time.DateOnly, "1991-11-12")
	newStartDate, _ := time.Parse(time.DateOnly, "2023-05-16")

	suite.employees[suite.updatedEmployeeIndex] = *employees.CreateEmployee(suite.employees[suite.updatedEmployeeIndex].Id,
		"99",
		"Changed Forename",
		"Changed Surname",
		newDateOfBirth,
		"0129384756",
		"6758493021",
		"Changed address line 1",
		"Changed address line 2",
		"0497",
		"changed.email@test.go",
		newStartDate)

	repoErr := repo.Update(&suite.employees[suite.updatedEmployeeIndex])

	if repoErr != nil {
		t.Fatalf(`Failed to Update %v : %v`, suite.employees[suite.updatedEmployeeIndex], repoErr)
	}

	uowErr := uow.SaveChanges()

	if uowErr != nil {
		t.Fatalf(`Unit of Work failed to save changes %v`, uowErr)
	}
}

// Testing that GetById works and Update executed previously updated all the fields as expected
// as we get employee that was updated on previous step and compare the fields
func (suite *EmployeeRepositoryTestSuite) Test_03_GetById_ExistingEmployee() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	expected := suite.employees[suite.updatedEmployeeIndex]

	exists, employee, repoErr := repo.GetById(expected.Id)

	if repoErr != nil {
		t.Fatalf(`Failed to GetById %v`, repoErr)
	}

	if !exists {
		t.Fatalf(`Employee %v should be existing`, expected.Id)
	}

	if *employee != expected {
		t.Fatalf(`Wrong employee returned. Expected %v, Actual %v`, expected, employee)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_04_GetById_NonExistingEmployee() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	exists, employee, repoErr := repo.GetById(uuid.New())

	if repoErr != nil {
		t.Fatalf(`Failed to GetById %v`, repoErr)
	}

	if exists {
		t.Fatalf(`Employee %v should not exist`, employee)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_05_Delete() {

	t := suite.T()
	repo, uow, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	if uow == nil {
		t.Fatalf(`Unit of Work is nil`)
	}

	repoErr := repo.Delete(suite.employees[suite.deletedEmployeeIndex].Id)

	if repoErr != nil {
		t.Fatalf(`Failed to Delete %v`, repoErr)
	}

	uowErr := uow.SaveChanges()

	if uowErr != nil {
		t.Fatalf(`Unit of Work failed to save changes %v`, uowErr)
	}
}

// a;so testing that Id of deleted employee in previous test not exists
func (suite *EmployeeRepositoryTestSuite) Test_06_IsIdExist_NonExistingEmployee() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	deletedEmployeeId := suite.employees[suite.deletedEmployeeIndex].Id

	exists, repoErr := repo.IsIdExist(deletedEmployeeId)

	if repoErr != nil {
		t.Fatalf(`Failed to IsIdExist %v`, repoErr)
	}

	if exists {
		t.Fatalf(`Employee %v should not exist`, deletedEmployeeId)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_07_IsIdExist_ExistingEmployee() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	existingEmployeeId := suite.employees[suite.unmodifiedEmployeeIndex].Id

	exists, repoErr := repo.IsIdExist(existingEmployeeId)

	if repoErr != nil {
		t.Fatalf(`Failed to IsIdExist %v`, repoErr)
	}

	if !exists {
		t.Fatalf(`Employee %v should exist`, existingEmployeeId)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_08_IsPayrollNumberExist_ExistingPayrollNumber() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	existingPayrollNumber := suite.employees[suite.unmodifiedEmployeeIndex].PayrollNumber

	exists, repoErr := repo.IsPayrollNumberExist(existingPayrollNumber)

	if repoErr != nil {
		t.Fatalf(`Failed to IsPayrollNumberExist %v`, repoErr)
	}

	if !exists {
		t.Fatalf(`Employee %v should exist`, existingPayrollNumber)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_09_IsPayrollNumberExist_NonExistingPayrollNumber() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	nonExistingPayrollNumber := "987654321"
	exists, repoErr := repo.IsPayrollNumberExist(nonExistingPayrollNumber)

	if repoErr != nil {
		t.Fatalf(`Failed to IsPayrollNumberExist %v`, repoErr)
	}

	if exists {
		t.Fatalf(`Employee %v should not exist`, nonExistingPayrollNumber)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_10_IsPayrollNumberExistExclusive_ExistingPayrollNumberAndNonExistsingId() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	existingPayrollNumber := suite.employees[suite.unmodifiedEmployeeIndex].PayrollNumber

	exists, repoErr := repo.IsPayrollNumberExistExclusive(existingPayrollNumber, uuid.New())

	if repoErr != nil {
		t.Fatalf(`Failed to IsPayrollNumberExistExclusive %v`, repoErr)
	}

	if !exists {
		t.Fatalf(`Employee %v should exist`, existingPayrollNumber)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_11_IsPayrollNumberExistExclusive_ExistingPayrollNumberAndExistsingId() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	existingEmployeeId := suite.employees[suite.unmodifiedEmployeeIndex].Id
	existingPayrollNumber := suite.employees[suite.unmodifiedEmployeeIndex].PayrollNumber

	exists, repoErr := repo.IsPayrollNumberExistExclusive(existingPayrollNumber, existingEmployeeId)

	if repoErr != nil {
		t.Fatalf(`Failed to IsPayrollNumberExistExclusive %v`, repoErr)
	}

	if exists {
		t.Fatalf(`Employee %v should not exist`, existingPayrollNumber)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_12_IsPayrollNumberExistExclusive_NonExistingPayrollNumberAndExistsingId() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	existingEmployeeId := suite.employees[suite.unmodifiedEmployeeIndex].Id
	nonPayrollNumber := "107452732344"

	exists, repoErr := repo.IsPayrollNumberExistExclusive(nonPayrollNumber, existingEmployeeId)

	if repoErr != nil {
		t.Fatalf(`Failed IsPayrollNumberExistExclusive %v`, repoErr)
	}

	if exists {
		t.Fatalf(`Employee %v should not exist`, nonPayrollNumber)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_13_IsPayrollNumberExistExclusive_NonExistingPayrollNumberAndNonExistsingId() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	nonPayrollNumber := "107452732344"

	exists, repoErr := repo.IsPayrollNumberExistExclusive(nonPayrollNumber, uuid.New())

	if repoErr != nil {
		t.Fatalf(`Failed to IsPayrollNumberExistExclusive %v`, repoErr)
	}

	if exists {
		t.Fatalf(`Employee %v should not exist`, nonPayrollNumber)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_14_GetPayrollNumberToIdMap() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	expected := make(map[string]uuid.UUID)

	for i, emp := range suite.employees {
		if i == suite.deletedEmployeeIndex {
			continue
		}

		expected[emp.PayrollNumber] = emp.Id
	}

	actual, repoErr := repo.GetPayrollNumberToIdMap()

	if repoErr != nil {
		t.Fatalf(`Failed to GetPayrollNumberToIdMap %v`, repoErr)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf(`Mapping does not match. Expected %v. Actual %v`, expected, actual)
	}
}

func (suite *EmployeeRepositoryTestSuite) Test_15_GetAll() {

	t := suite.T()
	repo, _, err := persistence.CreateEmployeeRepository(suite.connectionString, t.Context())

	if err != nil {
		t.Fatalf(`Failed to create Employee Repository %v`, err)
	}

	if repo == nil {
		t.Fatalf(`Employee Repository is nil`)
	}

	actual, repoErr := repo.GetAll()

	if repoErr != nil {
		t.Fatalf(`Failed to Get Payroll Number to Id Map %v`, repoErr)
	}

	//one employee was deleted as a part of the test
	if len(actual) != (len(suite.employees) - 1) {
		t.Fatalf(`Employees does not match. Expected %v. Actual %v`, len(suite.employees), len(actual))
	}

	//order may not be consistent as create is transactional and most records will ahve same timestamp
	//so check overall length and that all expected employees present
	for i, emp := range suite.employees {
		if i == suite.deletedEmployeeIndex {
			continue
		}

		if slices.Index(actual, emp) == -1 {
			t.Fatalf(`Employee %v was not returned`, emp)
		}
	}
}

func Test_EmployeeRepository(t *testing.T) {
	suite.Run(t, new(EmployeeRepositoryTestSuite))
}

func createEmployee(id uuid.UUID, payrollNumber string) employees.Employee {
	dateOfBirth, _ := time.Parse(time.DateOnly, "1990-12-11")
	startDate, _ := time.Parse(time.DateOnly, "2025-10-24")

	return *employees.CreateEmployee(id,
		payrollNumber,
		"Forename",
		"Surname",
		dateOfBirth,
		"998991234567",
		"998991234589",
		"address line 1",
		"address line 2",
		"109345",
		"test@test.go",
		startDate)
}
