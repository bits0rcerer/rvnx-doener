## Environment variables

| Key                  | Description                                                 | Default | Example(s)                           |
|----------------------|-------------------------------------------------------------|---------|--------------------------------------|
| $DATABASE_URL        | Where is the data base?                                     | None    | postgresql://pg:pg@localhost:5433/pg |
| $PORT                | Port the webserver will listen on                           | 8080    | 8080                                 |
| $DEBUG               | Enables debug behavior                                      | FALSE   | FALSE / TRUE                         |
| $OSM_REQUEST_CACHING | Cache OpenStreetMap requests                                | FALSE   | FALSE / TRUE                         |
| $OSM_SYNC            | Enables an daily job to sync kebab shops from OpenStreetMap | TRUE    | FALSE / TRUE                         |
