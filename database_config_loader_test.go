package config_loader

import "testing"

func TestDatasourceConfigGroup_LoadByBytes(t *testing.T) {
	config := DatasourceConfigGroup{}
	yamlstr := `
    test:
        dbtype: mongo
        server: dev.mimixiche.cc
        port: 27017
        username: "akimimi"
        password: "123456"
        query: connectTimeoutMS=10000
        database: dev
        prefix: test
    `

	config.LoadByBytes([]byte(yamlstr))

	c, ok := config["test"]
	if !ok || c.Server != "dev.mimixiche.cc" {
		t.Error("datasource config YAML unmarshal failed!")
	}
	if c.Prefix != "test" {
		t.Errorf("datasource prefix expected to be %s, actual %s", "test", c.Prefix)
	}
	expected := "mongodb://akimimi:123456@dev.mimixiche.cc:27017/?connectTimeoutMS=10000"
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
        server: dev.mimixiche.cc
        port: 27017
        username: "akimimi"
        password: "123456"
        query: connectTimeoutMS=10000&authSource=dev&authMechanism=SCRAM-SHA-1
        replica_set: mimiset-1801
        replica_servers: db-member-0.mimixiche.cc:35017,db-member-1.mimixiche.cc:35017
        replica_read_preference: secondary
        database: dev
    `

	config.LoadByBytes([]byte(yamlstr))

	c, ok := config["test"]
	if !ok || c.ReplicaSet != "mimiset-1801" {
		t.Error("datasource config YAML unmarshal failed!")
	}
	expected := "mongodb://akimimi:123456@db-member-0.mimixiche.cc:35017,db-member-1.mimixiche.cc:35017/?replicaSet=mimiset-1801&readPreference=secondary&connectTimeoutMS=10000&authSource=dev&authMechanism=SCRAM-SHA-1"
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
        server: dev.mimixiche.cc
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
