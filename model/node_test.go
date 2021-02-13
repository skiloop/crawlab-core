package model

import (
	"github.com/crawlab-team/crawlab-db/mongo"
	"github.com/stretchr/testify/require"
	"testing"
)

func setupNodeTest() (err error) {
	return mongo.InitMongo()
}

func cleanupNodeTest() {
	_ = mongo.GetMongoCol(NodeColName).Delete(nil)
	_ = mongo.GetMongoCol(ArtifactColName).Delete(nil)
}

func TestNode_Add(t *testing.T) {
	err := setupNodeTest()
	require.Nil(t, err)

	n := Node{}

	err = n.Add()
	require.Nil(t, err)
	require.NotNil(t, n.Id)

	a, err := n.GetArtifact()
	require.Nil(t, err)
	require.Equal(t, n.Id, a.Id)
	require.NotNil(t, a.CreateTs)
	require.NotNil(t, a.UpdateTs)

	col := mongo.GetMongoCol(NodeColName)
	col.GetContext()

	cleanupNodeTest()
}

func TestNode_Save(t *testing.T) {
	err := setupNodeTest()
	require.Nil(t, err)

	n := Node{}

	err = n.Add()
	require.Nil(t, err)

	name := "test_node"
	n.Name = name
	err = n.Save()
	require.Nil(t, err)

	err = mongo.GetMongoCol(NodeColName).FindId(n.Id).One(&n)
	require.Nil(t, err)
	require.Equal(t, name, n.Name)

	cleanupNodeTest()
}

func TestNode_Delete(t *testing.T) {
	err := setupNodeTest()
	require.Nil(t, err)

	n := Node{
		Name: "test_node",
	}

	err = n.Add()
	require.Nil(t, err)

	err = n.Delete()
	require.Nil(t, err)

	var a Artifact
	col := mongo.GetMongoCol(ArtifactColName)
	err = col.FindId(n.Id).One(&a)
	require.Nil(t, err)
	require.NotNil(t, a.Obj)
	require.True(t, a.Del)

	cleanupNodeTest()
}
