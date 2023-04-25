migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("9s6fwccmgzormui")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "x4j8yhkb",
    "name": "chart",
    "type": "bool",
    "required": false,
    "unique": false,
    "options": {}
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("9s6fwccmgzormui")

  // remove
  collection.schema.removeField("x4j8yhkb")

  return dao.saveCollection(collection)
})
