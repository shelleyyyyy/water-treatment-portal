migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("9s6fwccmgzormui")

  // remove
  collection.schema.removeField("tuciay7d")

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("9s6fwccmgzormui")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "tuciay7d",
    "name": "device",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "1dzo0fde7e85mac",
      "cascadeDelete": false,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
})
