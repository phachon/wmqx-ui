package models

import "github.com/snail007/go-activerecord/mysql"

const Table_UserNode_Name = "user_node"

type UserNode struct {

}

var UserNodeModel = UserNode{}

func (userNode *UserNode) Insert(userNodeValue map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_UserNode_Name, userNodeValue))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

func (userNode *UserNode) InsertBatch(insertValues []map[string]interface{}) (id int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().InsertBatch(Table_UserNode_Name, insertValues))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

func (userNode *UserNode) GetUserNodeByNodeId(nodeId string) (userNodes []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_UserNode_Name).Where(map[string]interface{}{
		"node_id": nodeId,
	}))
	if err != nil {
		return
	}
	userNodes = rs.Rows()
	return
}

func (userNode *UserNode) GetUserNodeByUserId(userId string) (userNodes []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_UserNode_Name).Where(map[string]interface{}{
		"user_id": userId,
	}))
	if err != nil {
		return
	}
	userNodes = rs.Rows()
	return
}

func (userNode *UserNode) DeleteUserNodeByUserNodeId(userNodeId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_UserNode_Name, map[string]interface{}{
		"user_node_id": userNodeId,
	}))
	return
}

func (userNode *UserNode) DeleteUserNodeByNodeId(nodeId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_UserNode_Name, map[string]interface{}{
		"node_id": nodeId,
	}))
	return
}

func (userNode *UserNode) DeleteUserNodeByUserId(userId string) (err error) {
	db := G.DB()
	_, err = db.Exec(db.AR().Delete(Table_UserNode_Name, map[string]interface{}{
		"user_id": userId,
	}))
	return
}

func (userNode *UserNode) GetUserNodeByUserNodeIds(userNodeIds []string) (userNodes []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_UserNode_Name).Where(map[string]interface{}{
		"user_node_id": userNodeIds,
	}))
	if err != nil {
		return
	}
	userNodes = rs.Rows()
	return
}

func (userNode *UserNode) GetUserNodeByUserNodeId(userNodeId string) (userNodes []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_UserNode_Name).Where(map[string]interface{}{
		"user_node_id": userNodeId,
	}))
	if err != nil {
		return
	}
	userNodes = rs.Rows()
	return
}