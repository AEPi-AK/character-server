# booth-2016-character-server

## Data
Character (JSON):
{"_id": unique_hash, "pro_id": unique_hash2, "num": int, gold": int, "experience": int}

The hashes are generated from string representations of the data we receive from the card readers. num is an autoincrementing player number, and should never be updated. 

A pro_id differs from an _id only in that pro_id's are from the cards we distribute. As long as you have either the id, prod_id, or player num, you can query for the rest of the player data.

## Endpoints
GET /characters/{identifier}
Gets the character for a given identifier. The identifier can e the id, pro_id, or player num.

POST /characters
Creates a new character. We expect JSON of the form {"data": string_from_card_reader}

POST /characters/update
Updates an existing character. We expect JSON of the form {"_id": unique_hash, "pro_id": data_from_card_reader, "gold": int, "experience": int}

The values given for gold and experience are not added to the current value, they completely replace them. If omitted, the previous data is kept.
The same applies to pro_id. pro_id should only be set when registering someone's pro card with their account. After that, never use that field in this request.

## Todo
  * More queries?
    * Top experience
    * Top gold
  
