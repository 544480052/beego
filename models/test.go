package models

var(
	TestList map[string]*Test
)

func init()  {
	TestList = make(map[string]*Test)
	u := Test{"test_id","test_name","test_password",Trofile{"man",18,"address","email"}}
	TestList["test_1"] = &u
	}

	
type Test struct {
	Id string
	Name string
	Password string
	Trofile Trofile
}

type Trofile struct {
	Gender string
	Age int
	Address string
	Email string
}

func GetAllTests() map[string]*Test  {
	return TestList
}


