# controle-de-ponto-back
Back-end of web app "Controle de Ponto", developed with GoLang + MongoDB. It serves static files developed with Quasar.
To start to use this application is necessary to create a database on MongoDB, default named "dateRegister". In this database, create a collection called "Users", with the following document:

    {
      "_id" : ObjectId("XXXXXXXXXXXXXXXXXX"),
      "login" : "admin",
      "password" : "61646d696ee3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
      "session" : "",
      "level" : 1
    }
    
It will create an user with the credentials:
    
    Login: "admin"
    Senha (password): "admin"
    
OBS: To use another DB name is necessary to set the environment variable called "DB_NAME" with the choiced DB name.
