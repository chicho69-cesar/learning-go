# Agrupando Datos

- Arreglo (Array)
    - Es una secuencia enumerada de elementos de un mismo tipo.
    - No cambia en tamaño
    - Usado por funcionalidades internas de Go; generalmente no es recomendado que lo uses en tu código.
    - [Arrays](https://golang.org/ref/spec#Array_types)

- Slice
    - Están construidos sobre el tipo arreglo.
    - Los valores que están contenidos en un slice son del mismo tipo.
    - Este si puede cambiar en tamaño.
    - Tiene una longitud y una capacidad.
    - [Slices](https://golang.org/ref/spec#Slice_types)

- map
    - Almacenamiento key / value
    - Es un grupo de elementos de un tipo, llamado element type, sin orden de agrupación, indexados por un conjunto único de keys de otro tipo, llamado key type.
    - [Maps](https://golang.org/ref/spec#Map_types)

- struct
    - Una estructura de datos
    - Un tipo compuesto
    - Nos permite poner juntos valores de diferentes tipos.
    - [Structs](https://golang.org/ref/spec#Struct_types)