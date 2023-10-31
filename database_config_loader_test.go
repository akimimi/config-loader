package config_loader

import "testing"

func TestDatasourceConfigGroup_LoadByBytes(t *testing.T) {
	config := DatasourceConfigGroup{}
	yamlstr := `
    test:
        dbtype: mongo
        server: dev.db.cc
        port: 27017
        username: "akimimi"
        password: "123456"
        query: connectTimeoutMS=10000
        database: dev
        prefix: test
    `

	config.LoadByBytes([]byte(yamlstr))

	c, ok := config["test"]
	if !ok || c.Server != "dev.db.cc" {
		t.Error("datasource config YAML unmarshal failed!")
	}
	if c.Prefix != "test" {
		t.Errorf("datasource prefix expected to be %s, actual %s", "test", c.Prefix)
	}
	expected := "mongodb://akimimi:123456@dev.db.cc:27017/?connectTimeoutMS=10000"
	actual := config.ConnectString("test")
	if actual != expected {
		t.Errorf("connect string expected %s, actual %s", expected, actual)
	}
	if config.DatabaseName("test") != "dev" {
		t.Errorf("connect database expected %s, actual %s", "dev", c.Database)
	}
}

func TestDatasourceConfigGroup_LoadByBytesForReplicaSet(t *testing.T) {
	config := DatasourceConfigGroup{}
	yamlstr := `
    test:
        dbtype: mongo
        server: dev.db.cc
        port: 27017
        username: "akimimi"
        password: "123456"
        query: connectTimeoutMS=10000&authSource=dev&authMechanism=SCRAM-SHA-1
        replica_set: replica 
        replica_servers: db0.dev.cc:27017,db1.dev.cc:27017
        replica_read_preference: secondary
        database: dev
    `

	config.LoadByBytes([]byte(yamlstr))

	c, ok := config["test"]
	if !ok || c.ReplicaSet != "replica" {
		t.Error("datasource config YAML unmarshal failed!")
	}
	expected := "mongodb://akimimi:123456@db0.dev.cc:27017,db1.dev.cc:27017/?replicaSet=replica&readPreference=secondary&connectTimeoutMS=10000&authSource=dev&authMechanism=SCRAM-SHA-1"
	actual := config.ConnectString("test")
	if actual != expected {
		t.Errorf("connect string expected %s, actual %s", expected, actual)
	}
	if config.DatabaseName("test") != "dev" {
		t.Errorf("connect database expected %s, actual %s", "dev", c.Database)
	}
}

func BenchmarkDatasourceConfigGroup_LoadByBytes(b *testing.B) {
	yamlstr := `
    test:
        dbtype: mongo
        server: dev.db.cc
        port: 27017
        username: "akimimi"
        password: "123456"
        query: connectTimeoutMS=10000
        database: dev
    `

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config := DatasourceConfigGroup{}
		config.LoadByBytes([]byte(yamlstr))
	}
}

func TestDatasourceConfig_String(t *testing.T) {
	config := DatasourceConfig{
		Server:         "dev.db.cc",
		Port:           27017,
		Username:       "akimimi",
		Password:       "123456",
		Query:          "connectTimeoutMS=10000",
		Database:       "dev",
		Dbtype:         DataTypeMongo,
		ReplicaSet:     "",
		ReplicaServers: "",
		ReadPreference: "",
		Prefix:         "",
	}
	expected := "mongodb://akimimi:123456@dev.db.cc:27017/?connectTimeoutMS=10000"
	actual := config.String()
	if actual != expected {
		t.Errorf("connect string expected %s, actual %s", expected, actual)
	}

	config.Dbtype = DataTypeMysql
	expected = "akimimi:123456@dev.db.cc:27017/dev?connectTimeoutMS=10000"
	actual = config.String()
	if actual != expected {
		t.Errorf("connect string expected %s, actual %s", expected, actual)
	}

	config.Dbtype = DataTypeMongo
	config.ReplicaSet = "replica"
	config.ReplicaServers = "db0.dev.cc:27017,db1.dev.cc:27017,db2.dev.cc:27017"
	config.ReadPreference = "secondaryPreferred"
	expected = "mongodb://akimimi:123456@db0.dev.cc:27017,db1.dev.cc:27017,db2.dev.cc:27017/?replicaSet=replica&readPreference=secondaryPreferred&connectTimeoutMS=10000"
	actual = config.String()
	if actual != expected {
		t.Errorf("connect string expected %s, actual %s", expected, actual)
	}
}
