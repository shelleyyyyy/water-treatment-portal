migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("1dzo0fde7e85mac")

  // remove
  collection.schema.removeField("iqay0ks9")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "kwuk482m",
    "name": "subs",
    "type": "relation",
    "required": false,
    "unique": false,
    "options": {
      "collectionId": "9s6fwccmgzormui",
      "cascadeDelete": false,
      "maxSelect": 1,
      "displayFields": []
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("1dzo0fde7e85mac")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "iqay0ks9",
    "name": "subs",
    "type": "json",
    "required": false,
    "unique": false,
    "options": {}
  }))

  // remove
  collection.schema.removeField("kwuk482m")

  return dao.saveCollection(collection)
})
