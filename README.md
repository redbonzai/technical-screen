### NOTE: 
This code can be accessed via the following REPL: [ACCESS IT HERE](https://replit.com/@redbonzai/algorithms-structures)
# Golang Project

```markdown
# Golang Sort Package

This is a Golang project that implements a package to categorize packages as `STANDARD`, `SPECIAL`, or `REJECTED` based on their volume and mass.

## Directory Structure

```terminal
golang/
├── go.mod              # Go module file
├── main.go             # Entry point of the application
├── sortpkg/
│   ├── sort.go         # Contains the core sorting logic
│   ├── sort_test.go    # Unit tests for the sorting logic
```

## Installation

1. Navigate to the `golang` directory (assuming project is in the ~/Code directory):
   ```bash
   cd ~/Code/technical-screen/golang
   ```

2. Initialize Go modules if you haven't already:
   ```bash
   go mod tidy
   ```

This will install all necessary dependencies and ensure your Go module is ready to use.

## Building the Application

To build the application, run the following command from the `golang` directory:

```bash
go build -o sortpkg
```

This will create an executable named `sortpkg` in the current directory.

## Running the Application

After building the application, run it with:

```bash
./sortpkg
```

The output will categorize packages based on their dimensions and mass.

## Running Tests

To run the unit tests for the `sortpkg` package, run the following command:

```bash
go test ./sortpkg
```

This will execute the tests in `sort_test.go` and ensure the sorting logic is functioning correctly.

---

# TypeScript Project

```markdown
# TypeScript Sort Package

This is a TypeScript project that implements a function to categorize packages as `STANDARD`, `SPECIAL`, or `REJECTED` based on their volume and mass.

## Directory Structure

```terminal
typescript/
├── src/
│   ├── sort.ts         # Contains the core sorting logic
│   ├── sort.spec.ts    # Unit tests for the sorting logic
├── dist/               # Directory where the built distributable will be placed
├── tsconfig.json       # TypeScript configuration file
├── package.json        # Node.js dependencies and scripts
```

## Installation

1. Navigate to the `typescript` directory:
   ```bash
   cd ~/Code/technical-screen/typescript
   ```

2. Install the project dependencies:
   ```bash
   npm install
   ```

This will install all necessary Node.js dependencies for the project.

## Building the Application

To compile the TypeScript code into JavaScript, use the following command:

```bash
npm run build
```

The distributable files will be placed in the `dist` directory.

## Running the Application

Once the project is built, you can run the application by executing the built file in the `dist` directory. Example:

```bash
node dist/sort.js
```

Make sure that `node` is installed on your machine.

## Running Tests

To run the unit tests for the project, execute the following command:

```bash
npm run test
```

This will run the test suite using Jest and ensure that the sorting logic works as expected.

---
### SAMPLE
To sample the sort function, just run the src/index.ts file by first building `npm run build`, and then executing `node dist/index.js`

With these `README.md` files in place, both your Golang and TypeScript projects will be well-documented for installation, building, testing, and running the applications.

Let me know if you need any changes or additions!
