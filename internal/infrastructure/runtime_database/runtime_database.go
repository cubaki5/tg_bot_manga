package runtime_database

type Database struct {
	db map[string]interface{} //потом здесь будет конкретный тип из моделс
}

func NewDatabase(db map[string]interface{}) *Database {
	return &Database{db: db}
}

func (d *Database) List() (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (d *Database) Add() error {
	//TODO implement me
	panic("implement me")
}

func (d *Database) Delete() error {
	//TODO implement me
	panic("implement me")
}
