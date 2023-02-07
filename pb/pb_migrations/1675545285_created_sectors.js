migrate((db) => {
  const collection = new Collection({
    "id": "byqqrlylwbyp0sn",
    "created": "2023-02-04 21:14:45.831Z",
    "updated": "2023-02-04 21:14:45.831Z",
    "name": "sectors",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "mbw2mzbw",
        "name": "title",
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
        "id": "afobrefb",
        "name": "description",
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
        "id": "5ebhgd5j",
        "name": "img",
        "type": "file",
        "required": false,
        "unique": false,
        "options": {
          "maxSelect": 1,
          "maxSize": 5242880,
          "mimeTypes": [],
          "thumbs": []
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
  const collection = dao.findCollectionByNameOrId("byqqrlylwbyp0sn");

  return dao.deleteCollection(collection);
})
