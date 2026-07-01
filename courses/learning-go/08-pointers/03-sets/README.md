# Method Sets

El Method Sets determinan cuáles métodos adjuntar a un TIPO. Es exactamente lo que el
nombre dice: <strong style="color: red;">¿Cuál es la colección de métodos de un tipo dado? Es su colección de métodos.</strong>

## <span style="color: red;">IMPORTANTE: “el method set de un tipo determina las INTERFACES que el tipo implementa.....”</span>

El “importante” de arriba no fué mencionado en este video pero será discutido en la sección de
“concurrencia” en un video llamado “revisión de method sets”.

- Un RECEPTOR NO-POINTER
  - Funciona con valores que son POINTERS o NO-POINTERS
- Un RECEPTOR POINTER
  - Sólo funciona con valores que son POINTERS.

| Receptores | Valores |
| ---------- | ------- |
| (t T)      | T y *T  |
| (t *T)     | *T      |