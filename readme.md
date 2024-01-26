# Go Demo

Pequeña demo de una REST API con Go.


# Nueva API REST

> [!NOTE]
> Instalar fiber
> * Fiber es un web framework para construir REST API basado en express js

```bash
go get -u github.com/gofiber/fiber/v2
```

> Instalar air
> * Air nos ayudar a reiniciar el servidor después de detectar cambios en nuestro
codigo

```bash
go get -u github.com/cosmtrek/air
```

* Ejecutar la api

```bash
cd ./api/
```

* Ejecutar air

```bash
air
```

```bash
http://localhost:8000
```

## Enlaces de interes
* [fiber](https://docs.gofiber.io/)
* [Air repositorio oficial](https://github.com/cosmtrek/air)
* [Como usar air with go project](https://medium.com/@relia/guide-for-implementing-live-reload-using-golang-air-3b0535b92784)