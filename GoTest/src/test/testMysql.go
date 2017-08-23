package test

import (
        "fmt"
        "database/sql"
        "time"
)

type TestMysql struct {
        db *sql.DB
}

func NewTestMysql(driver, link string) (*TestMysql){
        if db, err := sql.Open(driver, link); err != nil{
                panic(err.Error())
        }else {
                return &TestMysql{
                        db:db,
                }
        }

        return nil
}

func (this *TestMysql)CreateUser(){
        stmtOut, err := this.db.Prepare("insert user values (?,?,?,?,?,?,?,?,?,?,?,?)")
        if err != nil{
                panic(err.Error())
        }

        defer stmtOut.Close()

        now := time.Now().Unix()
        for i := 0; i < 30000; i++{
                if _, err = stmtOut.Exec(i,
                        fmt.Sprintf("nickname%d", i),
                        fmt.Sprintf("passwd%d", i),
                        fmt.Sprintf("%d@163.com", i),
                        fmt.Sprintf("status%d", i),
                        fmt.Sprintf("%d", i%2),
                        fmt.Sprintf("qq%d", i),
                        fmt.Sprintf("sign%d", i),
                        fmt.Sprintf("%d", now),
                        fmt.Sprintf("hobby%d", i),
                        fmt.Sprintf("location%d", i),
                        fmt.Sprintf("desc%d", i),
                );err != nil{
                        panic(err.Error())
                }
        }
}

func (this *TestMysql) CreateGroup(){
        stmtOut, err := this.db.Prepare("insert groups values (?,?,?,?,?,?)")
        if err != nil{
                panic(err.Error())
        }

        defer stmtOut.Close()

        now := time.Now().Unix()
        for i := 1; i <= 300; i++{
                if _, err = stmtOut.Exec(i,
                        fmt.Sprintf("%d", now),
                        fmt.Sprintf("%d", now),
                        fmt.Sprintf("group%d", i),
                        fmt.Sprintf("status%d", i),
                        fmt.Sprintf(""),
                );err != nil{
                        panic(err.Error())
                }
        }
}

func (this *TestMysql) CreateUserGroup(){
        stmtOut, err := this.db.Prepare("insert user_group values (?,?,?,?,?,?)")
        if err != nil{
                panic(err.Error())
        }

        defer stmtOut.Close()

        now := time.Now().Unix()
        for i := 0; i < 30000; i++{
                if _, err = stmtOut.Exec(fmt.Sprintf("%d", i),
                        fmt.Sprintf("%d", i % 300 + 1),
                        fmt.Sprintf("%d", i%3),
                        fmt.Sprintf("%d", now),
                        fmt.Sprintf("%d", now),
                        fmt.Sprintf(""),
                );err != nil{
                        panic(err.Error())
                }
        }
}

func (this *TestMysql) CreateGroupMessage(){
        stmtOut, err := this.db.Prepare("insert group_message values (?,?,?,?,?,?,?)")
        if err != nil{
                panic(err.Error())
        }

        defer stmtOut.Close()

        now := time.Now().Unix()
        for i := 0; i < 10000; i++{
                if _, err = stmtOut.Exec(fmt.Sprintf("%d", i),
                        fmt.Sprintf("%d", now),
                        fmt.Sprintf("%d", now),
                        fmt.Sprintf("%d", i%300),
                        fmt.Sprintf("%d", i*3),
                        fmt.Sprintf("subject%d", i),
                        fmt.Sprintf("content%d", i),
                );err != nil{
                        panic(err.Error())
                }
        }
}