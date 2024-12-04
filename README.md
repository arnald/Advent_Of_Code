
# Advent of Code 2024 CLI Application

A command-line application built in Go to solve the **2024 Advent of Code challenges**. This project provides a modular structure for solving daily coding puzzles, where each day is implemented as a separate command, and each part of the challenge can be run independently.

---

## **Table of Contents**
- [Purpose](#purpose)
- [Features](#features)
- [How It Works](#how-it-works)
- [Project Structure](#project-structure)
- [Usage](#usage)
  - [Running a Command](#running-a-command)
  - [Examples](#examples)
- [Contributing](#contributing)

---

## **Purpose**
The purpose of this project is to complete the **2024 Advent of Code challenges** in a structured and scalable way. Each day of the challenge presents two parts:
1. Part 1 calculates total distance between the left list and the right list, add up the distances between all of the pairs you found.
2. Part 2 calculates a total similarity score by adding up each number in the left list after multiplying it by the number of times that number appears in the right list.

This CLI application:
- Allows easy execution of challenges for specific days and parts.
- Provides a clean, modular codebase for working on multiple problems.
- Promotes the use of Go for algorithmic problem-solving.

---

## **Features**
- **Modular CLI Structure**:
  - Each day's challenge is implemented as a separate command (`day1`, `day2`, etc.).
  - Each part of a challenge can be specified using a flag (`--part=1` or `--part=2`).
- **Scalable Architecture**:
  - Easily add new days and parts as the challenge progresses.
- **Clean Output**:
  - Provides clear results for each challenge part.
- **Reusable Logic**:
  - Challenge logic is implemented in separate packages (`partOne`, `partTwo`) to enable testing and reuse.

---

## **How It Works**

1. **Command-Line Structure**:
   - The CLI application is powered by the **Cobra** library.
   - Each day is implemented as a **subcommand** (e.g., `day1`, `day2`).
   - The `--part` flag is used to specify which part of the challenge to run.

2. **Execution Flow**:
   - The program starts in `main.go` and calls the `cmd.Execute()` function.
   - The root command (`adventofcode`) initializes the CLI, and subcommands like `day1` and `day2` register themselves during runtime.
   - When a user specifies a day and part (e.g., `day1 --part=1`), the corresponding logic from `challenges/dayOne/partOne` or `challenges/dayOne/partTwo` is executed.

---

## **Project Structure**

```plaintext
adventofcode/
├── challenges/
│   ├── dayOne/
│   │   ├── partOne/
│   │   │   └── distance.go         # Logic for Day 1 Part 1
│   │   ├── partTwo/
│   │   │   └── similarity_score.go # Logic for Day 1 Part 2
│   │   └── distances_list.txt      # Input data for Day 1
│   ├── dayTwo/
│   │   ├── partOne/
│   │   │   └── part_one.go         # Logic for Day 2 Part 1
│   │   ├── partTwo/
│   │   │   └── part_two.go         # Logic for Day 2 Part 2
│   └── dayThree/                   # Placeholder for Day 3
├── cmd/
│   ├── day1.go                     # CLI command for Day 1
│   ├── day2.go                     # CLI command for Day 2
│   ├── root.go                     # Root command
├── main.go                         # Entry point for the CLI
├── go.mod                          # Go module file
└── README.md                       # Documentation
```

### **Explanation of Key Components**
- **`cmd/`**:
  - Contains CLI commands (`day1.go`, `day2.go`, etc.).
  - Each command registers itself to the root command in its `init()` function.
- **`challenges/`**:
  - Houses the logic for each challenge, organized by day and part.
  - Separate files for Part 1 and Part 2 to keep the code modular and testable.

---

## **Usage**

### **1. Running a Command**

The general syntax for running a command is:

```bash
go run main.go <day> --part=<part>
```

- `<day>`: The day of the challenge (e.g., `day1`, `day2`).
- `<part>`: The part of the challenge (either `1` or `2`).

---

### **2. Examples**

#### **Day 1 Part 1**:
To calculate the total distance (Part 1 of Day 1):
```bash
go run main.go day1 --part=1
```
**Output:**
```plaintext
Day 1 Part 1 - Total Distance: 12345
```

#### **Day 1 Part 2**:
To calculate the similarity score (Part 2 of Day 1):
```bash
go run main.go day1 --part=2
```
**Output:**
```plaintext
Day 1 Part 2 - Similarity Score: 67890
```

#### **Day 2 Part 1**:
```bash
go run main.go day2 --part=1
```
**Output:**
```plaintext
Day 2 Part 1 - Result: 54321
```

#### **Day 2 Part 2**:
```bash
go run main.go day2 --part=2
```
**Output:**
```plaintext
Day 2 Part 2 - Result: 98765
```

---

## **Adding New Challenges**

1. Create a new directory under `challenges/` for the day (e.g., `challenges/dayThree`).
2. Implement the logic for Part 1 and Part 2 in `partOne` and `partTwo` subdirectories.
3. Add a new command file in `cmd/` (e.g., `cmd/day3.go`) and register it to the root command in its `init()` function.
4. Run the command using:
   ```bash
   go run main.go day3 --part=1
   ```

---

## **Contributing**

Contributions are welcome! If you'd like to contribute:
1. Fork the repository.
2. Create a new branch (`feature/new-challenge`).
3. Submit a pull request with your changes.

---

Let me know if there's anything to tweak or expand further!
