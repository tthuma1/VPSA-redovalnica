# VPSA-redovalnica

VPSA-redovalnica is a small Go application for storing student grades and calculating final results.

**Install dependencies**
- Run the following to download and tidy module dependencies:

```bash
go mod tidy
```

**Build & Run**
- Run directly from source:

```bash
go run ./cmd
```

- Build a binary and run it:

```bash
go build -o main.out ./cmd
./main.out
```

**CLI flags**
- `--stOcen` : Minimum number of grades required to consider a result positive (default: `6`).
- `--minOcena` : Minimum allowed grade value (default: `1`).
- `--maxOcena` : Maximum allowed grade value (default: `10`).

Example with custom flags:

```bash
go run ./cmd -- --stOcen 5 --minOcena 1 --maxOcena 10
```

Package redovalnica provides a simple gradebook system for managing students,
storing their grades, enforcing grade boundaries, and computing final results
based on configurable grading rules.

Example usage:

```go
r := redovalnica.UstvariRedovalnico(3, 1, 10) // need 3 grades, valid range 1..10
r.DodajStudenta("1001", redovalnica.Student{Ime: "Ana", Priimek: "Horvat"})
r.DodajOceno("1001", 9)
r.DodajOceno("1001", 8)
r.DodajOceno("1001", 10)
fmt.Println("=== Izpis redovalnice ===")
r.IzpisVsehOcen()
fmt.Println("\n=== Končni uspeh ===")
r.IzpisiKoncniUspeh()
```