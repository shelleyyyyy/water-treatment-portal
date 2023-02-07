migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("9s6fwccmgzormui")

  collection.updateRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("9s6fwccmgzormui")

  collection.updateRule = null

  return dao.saveCollection(collection)
})
