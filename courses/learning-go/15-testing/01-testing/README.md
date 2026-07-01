# Testing & Benchmarking

***Introducción***
- *Los tests deben*
    - Estar en un archivo que termina con “_test.go”
    - Estar en el mismo paquete el cual está siendo probado
    - Estar en una func con la firma “func TestXxx(*testing.T)”
- *Correr un test*
    - go test
- *Cómo tratar con un test que ha fallado*
    - usa t.Error para señalizar una falla
- *Idioma recomendado*
    - Expected (esperaba)
    - Got (obtuvo)