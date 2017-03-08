package main_test

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/shinji62/concourse-demo-meetup-5"
)

var dsnTest string

var _ = BeforeSuite(func() {
	addr := os.Getenv("MYSQL_PORT_3306_TCP_ADDR")
	port := os.Getenv("MYSQL_PORT_3306_TCP_PORT")
	password := os.Getenv("MYSQL_ENV_MYSQL_ROOT_PASSWORD")
	fmt.Println(os.Getenv("MYSQL_PORT_3306_TCP_ADDR"))
	if addr == "" {
		dsnTest = "root:root@tcp(localhost:3306)/meetup"
	} else {
		dsnTest = fmt.Sprintf("root:%s@tcp(%s:%s)/meetup", password, addr, port)
	}

	Migrate("mysql://" + dsnTest)

})

var _ = Describe("ConcourseDemoMeetup5", func() {

	var db *sql.DB
	var err error

	BeforeEach(func() {
		db, err = sql.Open("mysql", dsnTest)
		Expect(err).ShouldNot(HaveOccurred())

	})
	Context("With Valid DB", func() {
		It("Should be able to insert Data", func() {
			stmtIns, err := db.Prepare("INSERT INTO meetup (message) VALUES ( ? )")
			Expect(err).ShouldNot(HaveOccurred())
			defer stmtIns.Close()
			_, err = stmtIns.Exec("Hello World")
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should be able to select Data", func() {
			var message string
			err = db.QueryRow("select message from meetup where id = ?", 1).Scan(&message)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(message).To(Equal("Hello World"))
		})
	})

	AfterEach(func() {
		db.Close()
	})

})
