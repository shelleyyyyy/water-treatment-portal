migrate((db) => {
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

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("1dzo0fde7e85mac")

  // remove
  collection.schema.removeField("iqay0ks9")

  return dao.saveCollection(collection)
})
