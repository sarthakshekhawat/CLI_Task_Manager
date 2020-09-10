package db

import (
	"encoding/binary"
	"log"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

// Tasks struct, for key-value pair
type Tasks struct {
	Key   int
	Value string
}

var db *bolt.DB
var taskBucket = []byte("tasks")
var completedTaskBucket = []byte("completedTasks")
var deleteCompletedTask = []byte("delete")

// Init is used to intialize the buckets and database if does not exist
func Init(Path string) error {
	var err error
	db, err = bolt.Open(Path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(completedTaskBucket)
		return err
	})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err = tx.CreateBucketIfNotExists(deleteCompletedTask)
		return err
	})
	return err
}

// AddTask is used to add the tasks
func AddTask(task string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id := int(id64)
		return b.Put(itob(id), []byte(task))
	})
	return err
}

// ListAllTasks is used to display all the incomplete tasks in the list
func ListAllTasks() ([]Tasks, error) {
	var tasks []Tasks
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Tasks{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	return tasks, err
}

// ListAllCompletedTasks shows the list of all the completed tasks
func ListAllCompletedTasks() ([]Tasks, error) {
	var keyString []int
	var tasks []Tasks
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(completedTaskBucket)
		bd := tx.Bucket(deleteCompletedTask)
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			timeStampOld, _ := strconv.Atoi(string(bd.Get(k)))
			timeStampNow := int(time.Now().Unix())
			// OneDay := 86400
			OneDay := 10
			if (timeStampNow - timeStampOld) > OneDay {
				keyString = append(keyString, btoi(k))
			} else {
				tasks = append(tasks, Tasks{
					Key:   btoi(k),
					Value: string(v),
				})
			}
		}
		return nil
	})
	for _, k := range keyString {
		_ = DeleteCompletedTask(k)
	}
	return tasks, err
}

// DeleteTask deletes the task from the list
func DeleteTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
	return err
}

// DeleteCompletedTask deletes the tasks from the completed task list
// after 24hrs (after adding the task in Completed tasks list)
func DeleteCompletedTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(completedTaskBucket)
		return b.Delete(itob(key))
	})
	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(deleteCompletedTask)
		return b.Delete(itob(key))
	})
	return err
}

// CompleteTask deletes the task from the list and add the task in the completed task list
func CompleteTask(key int) error {
	var value string
	// Deleting the key value pair from the "taskBucket"
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		value = string(b.Get(itob(key)))
		return b.Delete(itob(key))
	})
	if err != nil {
		return err
	}
	// Adding the deleted value from "taskBucket" to "completedTaskBucket"
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(completedTaskBucket)
		return b.Put(itob(key), []byte(value))
	})
	if err != nil {
		return err
	}

	timeStamp := strconv.Itoa(int(time.Now().Unix()))
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(deleteCompletedTask)
		return b.Put(itob(key), []byte(timeStamp))
	})
	return err
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(v []byte) int {
	return int(binary.BigEndian.Uint64(v))
}
