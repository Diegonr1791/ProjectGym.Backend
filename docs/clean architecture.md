├── /cmd/ # Punto de entrada principal (main.go)
│ └── main.go
│
├── /internal/ # Código privado de tu app (dominio + lógica)
│ ├── /domain/ # Reglas de negocio puras (entidades + interfaces)
│ │ ├── model/ # Entidades del dominio (e.g., Usuario, Rutina)
│ │ └── repository/ # Interfaces de repositorios (abstracciones)
│ │
│ ├── /usecase/ # Casos de uso (aplicación, lógica orquestadora)
│ │ └── usuario/ # Casos de uso por feature
│ │ ├── service.go # Implementación de casos de uso
│ │ └── interface.go
│ │
│ └── /config/ # Configuración de entorno, conexión DB, etc.
│
├── /infrastructure/ # Implementaciones concretas (DB, APIs externas)
│ └── /persistence/ # Adaptadores para DB (ej. PostgreSQL)
│ └── usuario_pg.go
│
├── /interfaces/ # Entradas y salidas (HTTP, gRPC, CLI, etc.)
│ └── /http/ # Adaptador HTTP (Gin controllers)
│ └── usuario_handler.go
│
├── /pkg/ # Librerías comunes reutilizables
│
├── go.mod
└── go.sum
