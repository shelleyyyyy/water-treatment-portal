migrate((db) => {
  const collection = new Collection({
    "id": "1dzo0fde7e85mac",
    "created": "2023-02-04 21:15:53.413Z",
    "updated": "2023-02-04 21:15:53.413Z",
    "name": "devices",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "yvqdo0bj",
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
        "id": "xkrutzmz",
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
        "id": "tthrsndv",
        "name": "controls",
        "type": "json",
        "required": false,
        "unique": false,
        "options": {}
      },
      {
        "system": false,
        "id": "lj4xhrin",
        "name": "sector",
        "type": "relation",
        "required": false,
        "unique": false,
        "options": {
          "collectionId": "byqqrlylwbyp0sn",
          "cascadeDelete": false,
          "maxSelect": 1,
          "displayFields": []
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
  const collection = dao.findCollectionByNameOrId("1dzo0fde7e85mac");

  return dao.deleteCollection(collection);
})
