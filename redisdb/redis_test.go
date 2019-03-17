package redisdb

import (
	"github.com/gomodule/redigo/redis"
	"testing"
)

func TestConnection(t *testing.T){
	conn := GetPool().Get()
	defer conn.Close()
	resp,err := redis.String(conn.Do("PING"))
	if err != nil || resp != "PONG"{
		t.Error("Expected PONG got",resp)
	}
}

func TestGetValue(t *testing.T){
	conn := GetPool().Get()
	defer conn.Close()
	_,err := conn.Do("SET","test_value","Hello world")
	if err != nil{
		t.Error("An error occurred when setting value",err.Error())
		return
	}
	resp,err := redis.String(conn.Do("GET","test_value"))
	if err != nil || resp != "Hello world"{
		t.Error("Expected 'Hello world' got",resp)
	}
}