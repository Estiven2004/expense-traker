# 🧾 Expense Tracker CLI (Go)

Una aplicación de línea de comandos para registrar, ver y administrar tus gastos de forma simple y rápida. Desarrollada en Go y pensada para uso personal o educativo.

---

## 🚀 Características

- Agregar gastos con descripción y monto
- Eliminar gastos por ID
- Ver lista completa de gastos
- Ver resumen de gastos totales
- Ver resumen de gastos por mes
- Persistencia de datos usando un archivo JSON

---

## 📦 Requisitos

- [Go](https://golang.org/doc/install) instalado (versión 1.23.4 o superior)

---

## 🛠️ Instalación

1. Clona este repositorio:

```bash
git clone https://github.com/tu-usuario/expense-tracker.git
cd expense-tracker
```

2. Inicializa el módulo si aún no lo hiciste:

```bash
go mod init expense-tracker
```

3. Ejecuta los comandos:

```bash
go run . <comando> [opciones]
```

---

## 📚 Uso

### Agregar un gasto

```bash
go run . add --description "Almuerzo" --amount 20
```

### Listar todos los gastos

```bash
go run . list
```

### Eliminar un gasto

```bash
go run . delete --id 2
```

### Ver resumen total de gastos

```bash
go run . summary
```

### Ver resumen por mes (ej: abril)

```bash
go run . summary --month 4
```

---

## 📁 Estructura del proyecto

```
expense-tracker/
├── main.go         # Lógica CLI
├── storage.go      # Funciones para cargar/guardar datos
├── expense.go      # Definición de estructura Expense
├── expenses.json   # Archivo donde se guardan los datos
├── go.mod
└── README.md
```

---

## ✅ Próximas mejoras

- Soporte para categorías
- Establecer presupuestos mensuales
- Exportar a CSV

---

## 🧑‍💻 Autor

- **Estiven Ospina** - [@Estiven2004](https://github.com/Estiven2004)

---

## 📄 Licencia

Este proyecto está bajo la licencia MIT. Libre para usar, modificar y compartir.
---