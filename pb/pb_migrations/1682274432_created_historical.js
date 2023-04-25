migrate((db) => {
  const collection = new Collection({
    "id": "3gagiuxqhykozkl",
    "created": "2023-04-23 18:27:12.146Z",
    "updated": "2023-04-23 18:27:12.146Z",
    "name": "historical",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "qj95ozvj",
        "name": "topic",
        "type": "text",
        "required": false,
        "unique": false,
        "options": {
          "min": null,
          "max": null,
          "pattern": ""
        }
      },
      {
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
      }
    ],
    "listRule": null,
    "viewRule": null,
    "createRule": null,
    "updateRule": null,
    "deleteRule": null,
    "options": {}
  });

  return Dao(db).saveCollection(collection);
}, (db) => {
  const dao = new Dao(db);
  const collection = dao.findCollectionByNameOrId("3gagiuxqhykozkl");

  return dao.deleteCollection(collection);
})
