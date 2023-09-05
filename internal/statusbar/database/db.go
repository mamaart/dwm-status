package database

import "github.com/mamaart/statusbar/internal/models"

type DB struct {
	tasks map[int]models.Task
	addCh chan models.Task
	delCh chan int
}

func New() *DB {
	return &DB{
		tasks: make(map[int]models.Task),
		addCh: make(chan models.Task),
		delCh: make(chan int),
	}
}

func (db *DB) Run(callback chan<- []models.Task) {
	for {
		select {
		case task := <-db.addCh:
			db.tasks[task.Id] = task
		case id := <-db.delCh:
			delete(db.tasks, id)
		}
		callback <- db.List()
	}
}

func (db *DB) List() (out []models.Task) {
	for _, e := range db.tasks {
		out = append(out, e)
	}
	return out
}

func (db *DB) Delete(id int) {
	db.delCh <- id
}

func (db *DB) Add(task models.Task) {
	db.addCh <- task
}