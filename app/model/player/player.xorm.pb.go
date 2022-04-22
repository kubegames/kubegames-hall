// code generated by github.com/kubegames/protoc-gen-xorm. DO NOT EDIT.

package player

import "xorm.io/xorm"

//insert one
func InsertPlayer(session *xorm.Session, model Player) (int64, error) {
	return session.InsertOne(&model)
}

//insert array
func InsertPlayers(session *xorm.Session, models []*Player) (int64, error) {
	return session.Insert(&models)
}

//delete
func DeletePlayer(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	for _, where := range wheres {
		session = where(session)
	}
	return session.Delete(&Player{})
}

//update
func UpdatePlayer(session *xorm.Session, updates map[string]interface{}, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	session = session.Table(&Player{})
	for _, where := range wheres {
		session = where(session)
	}
	return session.Update(updates)
}

//find
func FindPlayer(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result Player, ok bool, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	ok, err = session.Get(&result)
	if err != nil {
		return result, false, err
	}
	return result, ok, nil
}

//finds
func FindPlayers(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result []*Player, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	if err = session.Find(&result); err != nil {
		return nil, err
	}
	return result, nil
}

//page
func PagePlayer(session *xorm.Session, page, size int, wheres ...func(*xorm.Session) *xorm.Session) (total int64, result []*Player, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	total, err = session.Limit(size, (page-1)*size).FindAndCount(&result)
	if err != nil {
		return total, nil, err
	}
	return total, result, nil
}

//insert one
func InsertStatus(session *xorm.Session, model Status) (int64, error) {
	return session.InsertOne(&model)
}

//insert array
func InsertStatuss(session *xorm.Session, models []*Status) (int64, error) {
	return session.Insert(&models)
}

//delete
func DeleteStatus(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	for _, where := range wheres {
		session = where(session)
	}
	return session.Delete(&Status{})
}

//update
func UpdateStatus(session *xorm.Session, updates map[string]interface{}, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	session = session.Table(&Status{})
	for _, where := range wheres {
		session = where(session)
	}
	return session.Update(updates)
}

//find
func FindStatus(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result Status, ok bool, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	ok, err = session.Get(&result)
	if err != nil {
		return result, false, err
	}
	return result, ok, nil
}

//finds
func FindStatuss(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result []*Status, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	if err = session.Find(&result); err != nil {
		return nil, err
	}
	return result, nil
}

//page
func PageStatus(session *xorm.Session, page, size int, wheres ...func(*xorm.Session) *xorm.Session) (total int64, result []*Status, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	total, err = session.Limit(size, (page-1)*size).FindAndCount(&result)
	if err != nil {
		return total, nil, err
	}
	return total, result, nil
}

//insert one
func InsertPlayerRecord(session *xorm.Session, model PlayerRecord) (int64, error) {
	return session.InsertOne(&model)
}

//insert array
func InsertPlayerRecords(session *xorm.Session, models []*PlayerRecord) (int64, error) {
	return session.Insert(&models)
}

//delete
func DeletePlayerRecord(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	for _, where := range wheres {
		session = where(session)
	}
	return session.Delete(&PlayerRecord{})
}

//update
func UpdatePlayerRecord(session *xorm.Session, updates map[string]interface{}, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	session = session.Table(&PlayerRecord{})
	for _, where := range wheres {
		session = where(session)
	}
	return session.Update(updates)
}

//find
func FindPlayerRecord(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result PlayerRecord, ok bool, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	ok, err = session.Get(&result)
	if err != nil {
		return result, false, err
	}
	return result, ok, nil
}

//finds
func FindPlayerRecords(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result []*PlayerRecord, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	if err = session.Find(&result); err != nil {
		return nil, err
	}
	return result, nil
}

//page
func PagePlayerRecord(session *xorm.Session, page, size int, wheres ...func(*xorm.Session) *xorm.Session) (total int64, result []*PlayerRecord, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	total, err = session.Limit(size, (page-1)*size).FindAndCount(&result)
	if err != nil {
		return total, nil, err
	}
	return total, result, nil
}

//insert one
func InsertPlayerBill(session *xorm.Session, model PlayerBill) (int64, error) {
	return session.InsertOne(&model)
}

//insert array
func InsertPlayerBills(session *xorm.Session, models []*PlayerBill) (int64, error) {
	return session.Insert(&models)
}

//delete
func DeletePlayerBill(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	for _, where := range wheres {
		session = where(session)
	}
	return session.Delete(&PlayerBill{})
}

//update
func UpdatePlayerBill(session *xorm.Session, updates map[string]interface{}, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	session = session.Table(&PlayerBill{})
	for _, where := range wheres {
		session = where(session)
	}
	return session.Update(updates)
}

//find
func FindPlayerBill(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result PlayerBill, ok bool, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	ok, err = session.Get(&result)
	if err != nil {
		return result, false, err
	}
	return result, ok, nil
}

//finds
func FindPlayerBills(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result []*PlayerBill, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	if err = session.Find(&result); err != nil {
		return nil, err
	}
	return result, nil
}

//page
func PagePlayerBill(session *xorm.Session, page, size int, wheres ...func(*xorm.Session) *xorm.Session) (total int64, result []*PlayerBill, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	total, err = session.Limit(size, (page-1)*size).FindAndCount(&result)
	if err != nil {
		return total, nil, err
	}
	return total, result, nil
}