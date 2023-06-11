package torm

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/suifengpiao14/torm/parser/ddlparser"
)

func getBuilder() *Builder {
	return NewBuilder("example", getDDL(), *getDBConfig())
}

func readFile(filename string) (contnent string) {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	contnent = string(b)
	return contnent
}

func getDBConfig() (dbcfg *ddlparser.DatabaseConfig) {
	cfgFile := "./example/dbconfig.json"
	content := readFile(cfgFile)
	dbcfg = &ddlparser.DatabaseConfig{}
	err := json.Unmarshal([]byte(content), dbcfg)
	if err != nil {
		panic(err)
	}
	return dbcfg
}

func getDDL() (ddl string) {
	ddlsql := "./example/ddl.sql"
	ddl = readFile(ddlsql)
	return ddl
}

func getTormMetaTpl() (tormMetaTpl string) {
	filename := "./example/torm.meta.tpl"
	content := readFile(filename)
	return content
}
func getTorm() (tormMetaTpl string) {
	filename := "./example/torm.tpl"
	content := readFile(filename)
	return content
}

func TestGenerateModel(t *testing.T) {
	builder := getBuilder()
	buf, err := builder.GenerateModel()
	require.NoError(t, err)
	fmt.Println(buf.String())
}
func TestGenerateTorm(t *testing.T) {
	builder := getBuilder()
	tormMap := TormMetaMap{
		"server": getTormMetaTpl(),
	}
	buf, err := builder.GenerateTormFromMeta(tormMap)
	require.NoError(t, err)
	fmt.Println(buf.String())
}
func TestGenerateSQLEntity(t *testing.T) {
	builder := getBuilder()
	buf, err := builder.GenerateSQLEntity(getTorm())
	require.NoError(t, err)
	fmt.Println(buf.String())
}
