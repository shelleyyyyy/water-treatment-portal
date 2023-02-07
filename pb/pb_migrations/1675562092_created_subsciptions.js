migrate((db) => {
  const collection = new Collection({
    "id": "9s6fwccmgzormui",
    "created": "2023-02-05 01:54:52.265Z",
    "updated": "2023-02-05 01:54:52.265Z",
    "name": "subsciptions",
    "type": "base",
    "system": false,
    "schema": [
      {
        "system": false,
        "id": "zvfucv5l",
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
        "id": "op0tw234",
        "name": "value",
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
        "id": "hbllltbr",
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
  const collection = dao.findCollectionByNameOrId("9s6fwccmgzormui");

  return dao.deleteCollection(collection);
})
