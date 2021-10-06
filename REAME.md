# Iniciando o projeto 

Utilize o comando go mod init project_name

# Bibliotecas utilizadas:

- gorm: ORM para o banco de dados - go get github.com/gin-gonic/gin
- gin: para criação do server, similar ao express - go get gorm.io/gorm
- mysql: driver do mysql - go get gorm.io/driver/mysql


gin.H 

H is a shortcut for map[string]interface{}
Em outras palavras, o trecho c.JSON(http. StatusOK, gin. H {"status": "successful login"})
Equivalente a c.JSON (http. StatusOK, map [string] interface {} {"status": "successful login"})
Isso simplifica a maneira como o JSON é gerado e, se você precisar aninhar json, aninhar gin.H. Tudo bem. Exemplo:

c.JSON(http.StatusOK, gin.H{
            "status":  gin.H{
                "code": http.StatusOK,
                "Status": "successful login"
            },
            "message": message
        })

