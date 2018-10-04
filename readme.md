<h1 align="center">
    Define
</h1>

<h4 align="center">
    cli word lookup - written in go
</h4>



## To do

- Parse response JSON.
- If query fails, find closest-to and retry.


## Spell Correction API

 * Payload

 ```
 curl --get --include 'https://montanaflynn-spellcheck.p.mashape.com/check/?text=epistomology' \
     -H 'X-Mashape-Key: $SPELL_CHECK_KEY \
     -H 'Accept: application/json'
 ```

 * Response

 ```json
 HTTP/1.1 200 OK
 Connection: keep-alive
 Content-Length: 272
 Content-Type: application/json; charset=utf-8
 Date: Thu, 04 Oct 2018 22:15:10 GMT
 Server: Mashape/5.0.6
 X-Powered-By: koa
 
 {
   "original": "epistomology",
   "suggestion": "epistemology",
   "corrections": {
     "epistomology": [
       "epistemology",
       "epistemological",
       "entomologist",
       "epistolatory",
       "entomology",
       "epidemiologist",
       "epidemiology"
     ]
   }
 }
 ```

## Oxford Dictionary API

```python
LANGUAGE = 'en'
URL = 'https://od-api.oxforddictionaries.com:443/api/v1/entries/{language}/{word}'

APP_ID = os.environ.get('DICTIONARY_APP_ID')
APP_KEY = os.environ.get('DICTIONARY_APP_KEY')

HEADERS = {
    'app_id': APP_ID,
    'app_key': APP_KEY,
}
```
