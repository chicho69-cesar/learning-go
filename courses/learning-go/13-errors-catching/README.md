# Entendiendo

* https://golang.org/doc/faq#exceptions
    - Go no apoya el idioma try / catch / finally

### ¿Porqué Go no tiene excepciones?
Creemos que acoplar excepciones a una estructura de control, como en la expresión try-catch-finally, da como resultado un código complicado. También tiende a incentivar a los programadores a etiquetar demasiados errores comunes, como por ejemplo no abrir un archivo, como excepcionales.
Go toma un enfoque diferente. Para el manejo simple de errores, los retornos multi-valor de las funciones en Go facilitan el reporte de un error sin sobrecargar el valor de retorno. Un tipo de error canónico, junto con otras características de Go, hace que el manejo de errores sea agradable, pero bastante diferente del de otros lenguajes.
Go también tiene un algunas de funciones integradas para señalar y recuperarse de condiciones realmente excepcionales. El mecanismo de recuperación se ejecuta solamente como parte del estado de una función que es generado después de un error, el cual es suficiente para manejar una catástrofe pero no requiere estructuras de control adicionales y, cuando se usa bien, puede dar como resultado un código limpio de manejo de errores.

* https://en.wikipedia.org/wiki/Exception_handling#Criticism
    - Notice Hoare’s work also influenced goroutines and channels
* https://blog.golang.org/error-handling-and-go