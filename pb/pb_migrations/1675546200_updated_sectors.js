migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("byqqrlylwbyp0sn")

  collection.viewRule = ""

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("byqqrlylwbyp0sn")

  collection.viewRule = null

  return dao.saveCollection(collection)
})
