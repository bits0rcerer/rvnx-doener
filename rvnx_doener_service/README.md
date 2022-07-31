## Environment variables

| Key                        | Description                                                 | Default | Example(s)                           |
|----------------------------|-------------------------------------------------------------|---------|--------------------------------------|
| $DATABASE_URL              | Where is the data base?                                     | None    | postgresql://pg:pg@localhost:5433/pg |
| $PORT                      | Port the webserver will listen on                           | 8080    | 8080                                 |
| $DEBUG                     | Enables debug behavior                                      | FALSE   | FALSE / TRUE                         |
| $OSM_REQUEST_CACHING       | Cache OpenStreetMap requests                                | FALSE   | FALSE / TRUE                         |
| $OSM_SYNC                  | Enables an daily job to sync kebab shops from OpenStreetMap | TRUE    | FALSE / TRUE                         |
| $SESSION_SECRET            | Session secret                                              | None    | 123qweasd...dsaewq321                |
| $SESSION_ENCRYPTION_SECRET | Session encryption secret                                   | None    | 123qweasd...dsaewq321                |
| $TWITCH_CLIENT_ID          | Twitch client ID                                            | None    | 123qweasd...dsaewq321                |
| $TWITCH_CLIENT_SECRET      | Twitch client secret                                        | None    | 123qweasd...dsaewq321                |
