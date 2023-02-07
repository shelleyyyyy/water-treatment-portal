migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("byqqrlylwbyp0sn")

  collection.listRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("byqqrlylwbyp0sn")

  collection.listRule = null

  return dao.saveCollection(collection)
})
