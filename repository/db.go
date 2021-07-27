package repository

type Db struct {
}

func (d *Db) Connection() string {
	return "med_test"
}
