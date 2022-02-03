package utils

var Constants = struct {
	EmailDomain string
}{
	EmailDomain: "xyz.com",
}

var Paths = struct {
	DbPath string
}{
	DbPath: "/Users/tejasds/go/src/github.com/dstejas19/HRManagementSoftware-SoftwareEngineering/Backend/Database/HR_DB",
}

var Queries = struct {
	RegisterHR string
}{
	RegisterHR: "INSERT INTO USERS(FIRSTNAME, LASTNAME, BUSINESSUNIT, MANAGERID, GRADE,	LOCATION, COUNTRY, TITLE, TYPE, PERSONALEMAIL, EMAIL, CREATED_TS, UPDATED_TS) VALUES(",
}

var EmployeeType = struct {
	GradHire string
}{
	GradHire: "000",
}
