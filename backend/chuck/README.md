Frasero Chuck - Back End
========================

Back End of the Frasero Chuck Project using GO AppEngine

Instructions
------------

1. Download and install the [Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads)
2. Execute the goapp script to start the server:
```bash
goapp serve dispatch.yaml chuck/chuck.yaml
```
3. Access [http://localhost:8080](http://localhost:8080) to access the app

- **http://localhost:8080/frases/add**: Add a new frase.
- **http://localhost:8080/frases**: List a random quote from the list of frases.
