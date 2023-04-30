migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("3gagiuxqhykozkl")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "njkrrz1k",
    "name": "msg",
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
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("3gagiuxqhykozkl")

  // remove
  collection.schema.removeField("njkrrz1k")

  return dao.saveCollection(collection)
})
