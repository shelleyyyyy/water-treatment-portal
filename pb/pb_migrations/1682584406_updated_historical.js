migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("3gagiuxqhykozkl")

  // remove
  collection.schema.removeField("kfp6n5zp")

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("3gagiuxqhykozkl")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "kfp6n5zp",
    "name": "value",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
})
