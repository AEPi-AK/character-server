#character-server

## Data
Character (JSON):
{"id": unique_hash, "pro_id": unique_hash2, "number": int, "points": int, "strength": int, "wisdom": int, "dexterity": int,  "name": name}

The hashes are generated from string representations of the data we receive from the card readers. number is an autoincrementing player number, and should never be updated. 

A pro_id differs from an _id only in that pro_id's are from the cards we distribute. As long as you have either the id, prod_id, or player num, you can query for the rest of the player data.

## Endpoints
GET /characters/{identifier}
Gets the character for a given identifier. The identifier can be the id, pro_id, or player num.
Returns a Character in JSON

GET /character-leaderboards
Gets the characters in the DB sorted in decreasing order of points. Returns a JSON array of JSON Characters

POST /characters/create
Creates a new character. We expect JSON of the form {"id": identifier, "race": some_race_string, "strength": int, "wisdom": int, "dexterity": int}. Returns the created character.

POST /characters/update
Updates an existing character. We expect JSON of the form {"id": identifier, "pro_id": data_from_card_reader, "points": int, "name": name}. Returns the updated character

The values given points are not added to the current value, they completely replace them. If omitted, the previous data is kept.
The same applies to pro_id. pro_id should only be set when registering someone's pro card with their account. After that, never use that field in this request. 

