# num-guess-go

A simple interactive terminal-based Number Guessing Game written in **Go (Golang)**. The program generates a random number between `0` and `10`, and prompts the user to guess it, providing helpful hints based on how close the guess is.

---

## Features

- **Random Number Generation**: Uses Go's modern `math/rand/v2` package to generate a target number between `0` and `10`.
- **Interactive Terminal Input**: Accepts user guesses via `fmt.Scan`.
- **Smart Feedback/Hints**:
  - **"Near, a little higher"**: When the guess is within 2 units below the target number.
  - **"Higher"**: When the guess is more than 2 units below the target number.
  - **"Near, a little lower"**: When the guess is within 2 units above the target number.
  - **"Lower"**: When the guess is more than 2 units above the target number.
  - **"!!You Won!!"**: When the guess matches the target number.
- **Error Handling**: Gracefully handles input reading errors and prompts again.

---

## Prerequisites

- **Go**: Version `1.22` or higher (required for `math/rand/v2`).
---

## 🎮 How to Run

1. **Clone or create the code file**:
   Save the code in a file named `main.go`.

2. **Run the application directly**:
   ```bash
   go run main.go
   ```

3. **Or build and run the executable**:
   ```bash
   go build main.go
   ./main
   ```

---

## 💡 Example Usage

```text
Ingresa un valor: 5
Mas alto
Ingresa un valor: 8
Cerca,Un poco mas alto
Ingresa un valor: 9
!!Ganaste!!
```

---

