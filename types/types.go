package types

type GBUser struct {
	GOBID    string   `json:"gobid"`
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Email    string   `json:"email"`
	Role     []GBRole `json:"role"`
}

type GBRole struct {
	ID       string `json:"id"`
	RoleName string `json:"role_name"`
}

type GBCommongResponse struct {
	Token   string `json:"tokren"`
	Message string `json:"message"`
}

type JwtClaims struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
	Iat  int64  `json:"iat"`
	Exp  int64  `json:"exp"`
}

// Valid implements jwt.Claims.
func (*JwtClaims) Valid() error {
	return nil
}

type GBWorkspace struct {
	GOBID       string `json:"id"`
	Name        string `json:"workspace_name"`
	Description string `json:"description"`
	User        string `json:"user"`
}

type GBTestSuite struct {
	GOBID             string             `json:"id"`
	Name              string             `json:"testsuite_name"`
	Description       string             `json:"description"`
	User              string             `json:"user"`
	Workspace         string             `json:"workspace"`
	Application       string             `json:"application"`
	LoaderCss         string             `json:"locaderCss"`
	TestCaseModuleGrp []GBTestCaseModule `json:"testCaseModuleGrp"`
}

type GBTestCaseModule struct {
	GOBID       string       `json:"id"`
	Name        string       `json:"testcasemodule_name"`
	Description string       `json:"description"`
	TestCaseGrp []GBTestCase `json:"testCaseGrp"`
}

type GBTestCase struct {
	GOBID           string             `json:"id"`
	TestSuite       string             `json:"test_suite"`
	TestCaseName    string             `json:"testcase_name"`
	TestCaseType    string             `json:"test_case_type"`
	TestCaseDetails TestCaseDetailsVal `json:"test_case_details"`
}

type TestCaseDetailsVal struct {
	SelectorGroup SelectorGroupVal `json:"selector_group"`
	Action        ActionVal        `json:"action"`
}

type SelectorGroupVal struct {
	CssSelector   SelectorVal `json:"css_selector"`
	IdSelector    SelectorVal `json:"id_selector"`
	XpathSelector SelectorVal `json:"xpath_selector"`
}

type SelectorVal struct {
	IsSelected string `json:"is_selected"`
	Selector   string `json:"selector"`
}

type ActionVal struct {
	EnterValue string         `json:"eneter_value"`
	Read       string         `json:"read"`
	Assertion  []AssertionVal `json:"assertion"`
}

type AssertionVal struct {
	AssertionType   string `json:"assertion_type"`
	FieldToValidate string `json:"field_to_validate"`
	TrueValue       string `json:"true_value"`
	FindBu          string `json:"find_by"`
}
