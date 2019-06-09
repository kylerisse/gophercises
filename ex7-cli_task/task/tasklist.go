package task

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"strings"

	bolt "go.etcd.io/bbolt"
)

// Tasks holds a pointer to the bolt db
type Tasks struct {
	db *bolt.DB
}

// OpenTaskList is called prior to any DB action and will
// create a new DB as ~/.task/task.db if it doesn't exist
func OpenTaskList() *Tasks {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	_, err = os.Stat(usr.HomeDir + "/.task/")
	if os.IsNotExist(err) {
		err := os.MkdirAll(usr.HomeDir+"/.task/", 0700)
		if err != nil {
			panic(err)
		}
	}
	db, err := bolt.Open(usr.HomeDir+"/.task/task.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Tasks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return &Tasks{
		db: db,
	}
}

// Add a new task to the list
func (tl *Tasks) Add(name string) {
	tl.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		id, _ := b.NextSequence()
		err := b.Put([]byte(strconv.Itoa(int(id))), []byte(name))
		return err
	})
}

// List all tasks
func (tl *Tasks) List() string {
	var ret strings.Builder
	ret.WriteString("ID\tTask\n")
	tl.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			ret.WriteString(string(k) + "\t" + string(v) + "\n")
		}
		return nil
	})
	return ret.String()
}

// Done marks a task complete, removing it from the DB
func (tl *Tasks) Done(id string) error {
	tl.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Tasks"))
		err := b.Delete([]byte(id))
		return err
	})
	return nil
}
