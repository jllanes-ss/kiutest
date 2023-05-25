# Hello service

Servicio saludador de alta concurrencia.

* Path: `/hello`
* Port: `8080`

RQ 

```json
{
    "name": "<string>"
}
```

## Escribir un microservicio con las siguientes funcionalidades

1. Si se recibe un request con método POST, devolver:

```json
{
    "message":"Hello, <name>!",
    "exists": false
}
```

1.1 En caso de que el nombre ya haya sido utilizado, devuelva:

```json
{
    "message":"Hello, <name>! Welcome back!",
    "exists": true
}
```

2. Si se recibe un request con método GET, devolver la lista de nombres ya utilizados.

```json
{
    "names": ["Diego", "Simona", "Ada"]
}
```

2.1 En caso de estar vacía la lista:

```json
{
    "names": null
}
```



3. Cualquier otro método HTTP debe devovler:

```
405 Method not allowed
```

4. Si el `Content-Type` no es `application/json` devolver:

```
422 Unprocessable entity
```
