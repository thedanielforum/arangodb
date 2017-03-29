package arangodb

import (
	"encoding/json"
	"github.com/apex/log"
	"fmt"
	"github.com/thedanielforum/arangodb/types"
	"github.com/thedanielforum/arangodb/errc"
)

type EdgeProp struct {
	From      string    `json:"_from"`
	To        string    `json:"_to"`
}

func (c *Connection) Create(collectionName string, doc interface{}) error {

	checkType := c.validateCollection(collectionName, doc)

	if checkType != TypeDoc && checkType != TypeEdge {
		println(checkType,"Fatal Error")
	}
	//Collection Confirm Exist Now , Proceed to perform save document/edge
	urlStack := fmt.Sprintf("/_db/%s/_api/document/%s",c.db,collectionName)
	encoded,err := json.Marshal(doc)
	if err != nil {
		return err
	}
	_, err = c.post(urlStack,encoded)
	if err != nil {
		if checkType == TypeEdge {
			log.WithError(err).Info(errc.ErrorCodeInvalidEdgeAttribute.Msg())
			return err
		}
		log.WithError(err).Info(err.Error())
		return err
	}
	return nil
}

//CHECK CONNECTION EXIST AND TYPE
func (c *Connection) validateCollection(collectionName string, doc interface{}) int {
	urlStack := fmt.Sprintf("/_db/%s/_api/collection/%s",c.db,collectionName)

	checkEdge := new(EdgeProp)
	encodedDoc, _ := json.Marshal(doc)

	json.Unmarshal(encodedDoc, checkEdge)

	colProp, err := c.get(urlStack)
	if err != nil {
		//Create Collection if Not Exist
		if checkEdge.To != "" && checkEdge.From != "" {
			c.CreateColEdge(collectionName)
			return TypeEdge
		}
		c.CreateColDoc(collectionName)
		return TypeDoc
	}
	sorter := new(types.CollectionInfo)
	json.Unmarshal(colProp, sorter)
	return sorter.Type
}

func (c *Connection) cacheValidation(collectionName string, doc interface{}) error{
	//true means that collection exist
	if c.colCache[c.db][collectionName] {
		return nil
	}
	return errc.ErrorCodeCollectionNotExist
}

func (c *Connection) GetAllColProp() {
	urlStack := fmt.Sprintf("/_db/%s/_api/collection",c.db)
	allProp, err := c.get(urlStack)
	if err != nil {
		//db does not exist
		log.WithError(err).Info(errc.ErrorCodeNoDatabaseSelected.Msg())
	}
	sorter := new(types.ColInfo)
	json.Unmarshal(allProp, sorter)

	for _, fish := range sorter.Result{
		if fish.IsSystem == false {
			//println(fish.Name)
			cacheAdd(c.colCache, c.db, fish.Name)
		}
	}
	return
}

func (c *Connection) stubFunc() {
	println(c.colCache[c.db]["DemoTest0"])
	//fmt.Printf("%v\n", c.colCache)
	//fmt.Printf("%v\n", c.colCache["yello"]["DemoTest01"])
}