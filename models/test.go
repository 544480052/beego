package models

var(
	TestList map[string]*Test
)

func init()  {
	TestList = make(map[string]*Test)
	u := Test{"test_id","test_name","test_password",Profile{"man",18,"address","email"}}
	TestList["test_1"] = &u
	}

	
type Test struct {
	Id string
	Name string
	Password string
	Profile Profile
}

type Profile struct {
	Gender string
	Age int
	Address string
	Email string
}

func getAllTests() map[string]*Test  {
	return TestList
}


