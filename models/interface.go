package models


type Vms interface{
	List(region string)(map[string][]interface{}, error)
	Get(region string, id int)()
}
