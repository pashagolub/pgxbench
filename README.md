# **"Writing your first Go app for PostgreSQL" Training**

Table of Contents

- [Introduction](#introduction)
- [Task: Hello World!](#task-hello-world)
- [Task: Implement Database-Specific Functions](#task-implement-database-specific-functions)
- [Task: Implement Different Insert Methods](#task-implement-different-insert-methods)
- [Task: Implement Different Fetch Methods](#task-implement-different-fetch-methods)
- [Task: Refactor Code to Use Interfaces](#task-refactor-code-to-use-interfaces)
- [Task: Add Unit Tests Using `pgxmock`](#task-add-unit-tests-using-pgxmock)
- [Final Words](#final-words)

## **Introduction**  

Welcome to the **Writing your first Go app for PostgreSQL** training! This repository is designed to help you step-by-step in learning how to build a Go application for PostgreSQL. Each task corresponds to a specific commit in the repository, guiding you through the development process while allowing you to learn and explore independently.

## **How It Works**

1. **Task and Commit Mapping**  
   Each task in the training corresponds to a specific commit in this repository. By checking out the respective commit, you will get a pre-configured starting point for the task.

2. **Read and Understand the Task**  
   - Read the task description provided in this document.
   - Each task includes hints, templates (if necessary), and subtasks to help you focus on specific concepts.

3. **Implement the Task**  
   - Switch to the corresponding commit for the task using:

     ```bash
     git checkout <commit_hash>
     ```

   - Implement the required code locally based on the task description.

4. **Save Your Progress**  
   - Once you're done with the task, stash your changes or push them to a branch. For example:

     ```bash
     git stash   # if you don't want to save progress to a branch
     # OR
     git checkout -b my-task-branch-name && git add . && git commit -m "Implement <Task Name>"
     
     git checkout <next_task_commit_hash>
     ```

5. **Compare Your Solution**  
   - After completing the task, compare your implementation with the one provided in the next commit.
   - Analyze the differences, rethink your approach, and improve your understanding.

6. **Branching for Your Work**  
   - If you’d like to continue with your implementation, switch back to your branch.
   - Otherwise, follow the repository's commit history for guidance.

7. **Repeat for All Tasks**  
   - Progress through the tasks by repeating the above steps.

## **Before You Start**

### **Prerequisites**  

Ensure you have the following tools installed on your system:  

- **Go:** Version 1.23 is recommended ([Download](https://go.dev/dl/)).
- **Git:** ([Download](https://git-scm.com/downloads)).
- **PostgreSQL:** Either installed locally or via Docker ([Official Guide](https://www.postgresql.org/download/)).

### **Clone the Repository**  

Clone the `pgxbench` repository to your local machine, if you want to follow along with the tasks:  

```bash
git clone https://github.com/pashagolub/pgxbench.git
cd pgxbench
```

Or create a new folder to work on the tasks, while referring to the code in the web interface.

With this workflow in mind, you're ready to dive into the tasks. Follow each task step-by-step, and don't hesitate to experiment and explore as you go. Let’s get started!

---

# **Task: Hello World!**

## **Objective**  

Create a simple Go application that connects to a PostgreSQL database using a hardcoded connection string and retrieves the value of `SELECT version()` from the database.

## **Description**  

In this task, you'll build a basic Go application to establish a connection to a PostgreSQL database and execute a simple query to retrieve the PostgreSQL server version. This will introduce you to connecting and querying the database with the `pgx` library.

## **Subtasks**

1. **Initialize the Go Module**  
   - Run the following command to initialize the Go module (if you start from scratch with an empty folder):

     ```bash
     go mod init github.com/yourusername/pgxbench
     ```

2. **Add Dependencies**  
   - Use the following command to add the `pgx` library:

     ```bash
     go get github.com/jackc/pgx/v5
     ```

3. **Set Up the Connection**  
   - Create a `main.go` file.
   - Establish a connection to PostgreSQL using `pgx.Connect`. Replace placeholders with your database details:

     ```go
     conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/database_name")
     ```

4. **Execute a Query**  
   - Write a query to retrieve the database version using the `SELECT version()` SQL statement.

5. **Return and Print the Result**  
   - Fetch the result from the query and print it to the console.

## **Hints**

- Use `conn.QueryRow` to execute the query and fetch a single result. Example:

  ```go
  var version string
  err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
  ```

- Handle errors properly for both the query execution and result scanning.
- Print the version string to verify the output.

## **Example Output**  

After running the program, you should see output similar to this:  

```plaintext
PostgreSQL version: PostgreSQL 15.2 (Debian 15.2-1) on x86_64-pc-linux-gnu
```

## **Next Steps**  

Once you complete this task, you'll understand the basics of connecting to a PostgreSQL database, executing a query, and retrieving a result. This will be the foundation for more complex tasks ahead. To check your implementation, proceed to the commit [#1aadeb](https://github.com/pashagolub/pgxbench/commit/1aadeb8f748ade333799ea1eb803b43db7b338bc) in the repository.

---

# **Task: Implement Database-Specific Functions**

## **Objective**

Refactor the existing Go application by separating database-related code into a dedicated `database.go` file and enhance the application to accept a custom PostgreSQL connection string via a command-line parameter.

## **Description**

In this task, you'll improve the structure of your Go application by isolating database-specific functions into a separate file. Additionally, you'll modify the application to accept a custom PostgreSQL connection string through a command-line flag, enhancing its flexibility and usability.

---

## **Subtasks**

1. **Create `database.go` File**
   - Move all database-related functions from `main.go` to a new file named `database.go`.
   - Ensure that the `database.go` file belongs to the same package as `main.go`.

2. **Implement the Following Functions in `database.go`:**
   - **`ConnectToDB`**
     - Establish a connection to the PostgreSQL database using the provided connection string.
   - **`InitDB`**
     - Perform any necessary database initialization, such as creating tables or seeding data (leave as a placeholder if not implemented yet).
   - **`RunBenchmarks`**
     - Placeholder function where benchmarking logic will later be implemented. For now, it can be left empty.
   - **`CloseDB`**
     - Handle cleanup tasks, such as dropping tables or closing the database connection.

3. **Add Command-Line Parameter for Connection String**
   - In `main.go`, import the `flag` package to handle command-line arguments.
   - Define a command-line flag `-c` (or `--conn`) that allows users to specify a custom PostgreSQL connection string.
   - Update the database connection logic to utilize the connection string provided via the command-line flag.

4. **Update `main.go` to Use `database.go` Functions**
   - Modify the `main` function in `main.go` to:
     1. Parse the connection string from the command-line flag.
     2. Call `ConnectToDB` to establish a database connection.
     3. Call `InitDB` to perform database setup.
     4. Call `RunBenchmarks` to execute the main logic.
     5. Call `CloseDB` to handle cleanup tasks.

## **Hints**

- **Organizing Code:**
  - Ensure that both `main.go` and `database.go` are part of the `main` package to allow seamless function calls between them.

- **Handling Command-Line Flags:**
  - Use the `flag.String` function to define the `-c` flag for the connection string:

    ```go
    connStr := flag.String("c", "default_connection_string", "PostgreSQL connection string")
    ```

  - Remember to call `flag.Parse()` in the `main` function to parse the command-line arguments.

- **Default Connection String:**
  - Provide a sensible default connection string that can be used if the user does not specify one via the command-line flag.

---

## **Example Usage**

After implementing the above changes, you should be able to run your application with a custom connection string as follows:

```bash
go run . -c "postgresql://username:password@localhost:5432/database_name"
```

If no connection string is provided, the application should use the default connection string specified in the code.

## **Next Steps**

Upon completing this task, your Go application will have a cleaner structure with database-specific code separated into its own file. Additionally, the application will be more flexible, allowing users to specify different PostgreSQL databases via command-line parameters. This setup prepares the foundation for implementing the actual benchmarking logic in subsequent tasks.

To check your implementation, proceed to the commit [#7085a3](https://github.com/pashagolub/pgxbench/commit/7085a337c60ebeeac1ef9987e4bc47fae5402f0b) in the repository.

---

# **Task: Implement Different Insert Methods**

## **Objective**
Enhance the Go application by implementing various data insertion methods into the PostgreSQL database, including row-by-row INSERT, batched INSERT, and COPY. Additionally, introduce command-line parameters to control the total number of rows to insert and the batch size.

---

## **Description**
In this task, you'll expand the application's functionality by adding three distinct methods for inserting data into the PostgreSQL database:

1. **Row-by-Row INSERT (`InsertSimple`)**: Inserts data one row at a time.
2. **Batched INSERT (`InsertBatch`)**: Inserts data in batches to improve performance.
3. **COPY (`InsertCopy`)**: Utilizes PostgreSQL's COPY command for bulk data insertion.

You'll also add command-line parameters to specify the total number of rows to insert and the batch size, allowing for flexible benchmarking.

---

## **Subtasks**

1. **Define the `DbUser` Struct**
   - Create a struct named `DbUser` to represent the user data to be inserted.
   - Define a variable `TestUser` with constant values to be used for insertion.

2. **Add Command-Line Parameters**
   - Introduce two new command-line flags:
     - `-n` (or `--num-rows`): Specifies the total number of rows to insert.
     - `-b` (or `--batch-size`): Specifies the number of rows per batch for batched inserts.

3. **Implement Insertion Methods in `database.go`**
   - **`InsertSimple(ctx context.Context, conn *pgx.Conn) error`**
     - Inserts `numRows` instances of `TestUser` into the database one at a time.
   - **`InsertBatch(ctx context.Context, conn *pgx.Conn) error`**
     - Inserts `numRows` instances of `TestUser` into the database in batches of `batchSize`.
   - **`InsertCopy(ctx context.Context, conn *pgx.Conn) error`**
     - Uses the COPY protocol to insert a data into the database efficiently.

4. **Update `main.go` to Handle New Parameters and Methods**
   - Parse the new command-line flags for `numRows` and `batchSize`.
   - Call the appropriate insertion method based on user input or predefined logic.
  
5. **Benchmark the Insertion Methods**
   - Implement the `RunBenchmarks` function to benchmark the three insertion methods.
   - Measure the execution time for each method and print the results.

## **Hints**

- **Struct Definition:**
  - Define the `DbUser` struct with fields corresponding to the database columns.
  - Example:

    ```go
    type DbUser struct {
        Id   int
        Name string
        Age  int
        Meta string
    }
    ```

- **Command-Line Flags:**
  - Use the `flag` package to define and parse the new command-line parameters.
  - Example:

    ```go
    numRows := flag.Int("n", 1000, "Total number of rows to insert")
    batchSize := flag.Int("b", 100, "Number of rows per batch")
    ```

- **Insertion Methods:**
  - For `InsertSimple`, execute individual INSERT statements in a loop.
  - For `InsertBatch`, utilize `pgx.Batch` to group multiple INSERT statements and send them in a single round-trip to the database.
  - For `InsertCopy`, use the `CopyFrom` method provided by `pgx` to perform bulk inserts efficiently.

## **Example Usage**

After implementing the changes, you can run the application with the new parameters as follows:

```bash
go run . -c "postgresql://username:password@localhost:5432/database_name" -n 5000 -b 500
```

This command sets the connection string, inserts a total of 5000 rows, and uses a batch size of 500 for batched inserts.

## **Next Steps**

Upon completing this task, your Go application will support multiple data insertion methods, allowing you to benchmark and compare their performance. This functionality is crucial for understanding the trade-offs between different insertion strategies in PostgreSQL.

To check your implementation, proceed to the commit [#ab521f](https://github.com/pashagolub/pgxbench/commit/ab521fc3a3e64bf308c3a196c408de4887ac2d26) in the repository.

---

# **Task: Implement Different Fetch Methods**

## **Objective**
Enhance the Go application by implementing various methods to fetch data from the PostgreSQL database, including using regular `SELECT` statements with `Scan` and `pgx.CollectRows`. This will demonstrate the functionality provided by the `pgx` library for data retrieval.

## **Description**
In this task, you'll expand the application's functionality by adding two distinct methods for fetching data from the PostgreSQL database:

1. **Regular `SELECT` with `Scan` (`FetchSelectScan`)**: Retrieves data using a standard `SELECT` query and processes each row individually with the `Rows.Scan()` method.
2. **`SELECT` with `pgx.CollectRows` (`FetchSelectCollect`)**: Retrieves data using a standard `SELECT` query but collects all rows into a slice using `pgx.CollectRows` for more concise handling.

## **Subtasks**

1. **Implement Fetch Methods in `benchmark.go`**
   - **`FetchSelectScan(ctx context.Context, conn *pgx.Conn) error`**
     - Executes a `SELECT` query to retrieve data from the `test` table.
     - Iterates over the rows and scans the results into variables using `Rows.Scan()`.
   - **`FetchSelectCollect(ctx context.Context, conn *pgx.Conn) error`**
     - Executes a `SELECT` query to retrieve data from the `test` table.
     - Collects all rows into a slice of `DbUser` structs using `pgx.CollectRows`.

2. **Update `main.go` to Call Fetch Methods**
   - Modify the `main` function to call the appropriate fetch method based on user input or predefined logic.

## **Hints**

- **Fetching Data with `Rows.Scan`:**
  - Use the `Query` method to execute the `SELECT` statement.
  - Iterate over the rows with `Next()` and use `Scan()` to read column values into variables.
- **Fetching Data with `pgx.CollectRows`:**
  - Use the `pgx.CollectRows` function to collect all rows into a slice of `DbUser` structs.
- Ensure that the test table contains data to retrieve. For that make sure insert benchmark was run before.

## **Example Usage**

After implementing the changes, you can run the application to benchmark all the implemented methods. 

```terminal
$ go run . -c "host=localhost dbname=Test user=pasha" -n 100000 -b 1000
2025/01/27 18:22:40 Connected to database Test on PostgreSQL 16.1, compiled by Visual C++ build 1937, 64-bit
2025/01/27 18:22:40 Starting Insert row by row
2025/01/27 18:23:02 Finished Insert row by row in 21824ms
2025/01/27 18:23:02 Starting Insert in batch
2025/01/27 18:23:03 Finished Insert in batch in 971ms
2025/01/27 18:23:03 Starting Insert using copy
2025/01/27 18:23:03 Finished Insert using copy in 172ms
2025/01/27 18:23:03 Starting Select, then Scan()
2025/01/27 18:23:04 Finished Select, then Scan() in 76ms
2025/01/27 18:23:04 Starting Select, then CollectRows()
2025/01/27 18:23:04 Finished Select, then CollectRows() in 83ms
```

## **Next Steps**

Upon completing this task, your Go application will support multiple data retrieval methods, allowing you to benchmark and compare their performance. This functionality is crucial for understanding the trade-offs between different data fetching strategies in PostgreSQL.

To check your implementation, proceed to the commit [#d11f9a](https://github.com/pashagolub/pgxbench/commit/d11f9a4a7368903037c9f7025a4838947e6d7328) in the repository.

---

# **Task: Refactor Code to Use Interfaces**

## **Objective**
Refactor the Go application to improve flexibility and maintainability by introducing interfaces to abstract database operations. This will decouple the code from specific implementations, allowing for easier testing and future extensions.

## **Description**
In this task, you'll refactor the existing codebase to introduce a `Database` interface that abstracts database operations. This will enable the application to work with different database connection types (e.g., `*pgx.Conn`, `pgx.Tx`, or `*pgpool.Pool`) without modifying the core logic. Additionally, you'll implement type assertions in the `CloseDB` function to handle different database connection types appropriately.

## **Subtasks**

1. **Define the `Database` Interface**
   - Create an interface named `Database` that includes methods for executing queries and commands. For example:

     ```go
     type Database interface {
         Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
         Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
         QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
         // Add other methods as needed
     }
     ```

2. **Update Functions to Use the `Database` Interface**
   - Modify existing functions to accept the `Database` interface instead of concrete types. For instance:

     ```go
     func InsertSimple(ctx context.Context, conn Database) error {
         // Function implementation
     }
     ```

3. **Refactor the `CloseDB` Function**
   - Update the `CloseDB` function to handle different database connection types using type assertions or type switches. For example:

     ```go
     func CloseDB(ctx context.Context, conn Database) error {
         switch c := conn.(type) {
         case *pgx.Conn:
             return c.Close(ctx)
         case *pgpool.Pool:
             conn.Close()
         }
        // ...
     }
     ```

## **Hints**

- **Defining Interfaces:**
  - Focus on the methods that your application requires and define them in the `Database` interface.
  - Ensure that the method signatures in the interface match those of the `pgx` library or any other database libraries you are using.

- **Implementing Interfaces:**
  - In Go, an interface is implemented implicitly. Ensure that your connection types have methods that match the `Database` interface's method signatures.

- **Type Assertions and Switches:**
  - Use type assertions or type switches to determine the underlying type of the `Database` interface when specific behavior is needed for different connection types.

## **Example Usage**

After refactoring, your functions should work seamlessly with any type that implements the `Database` interface. For example:

```go
conn, _ := pgpool.New(context.Background(), connStr)
InsertSimple(context.Background(), conn)
tx, _ := conn.Begin(context.Background())
InsertBatch(context.Background(), tx)
```

This approach allows for greater flexibility, as you can now pass different database connection types to your functions without changing their implementations.

## **Next Steps**

Upon completing this task, your Go application will be more flexible and maintainable, with decoupled code that is easier to test and extend. This refactoring sets the stage for implementing mock databases for testing purposes and integrating additional database connection types in the future.

To check your implementation, proceed to the commit [#1b27ad](https://github.com/pashagolub/pgxbench/commit/1b27ad89890dd0b02e6051c946e94ab0e6207026) in the repository.

---

# **Task: Add Unit Tests Using `pgxmock`**

## **Objective**
Enhance the Go application by incorporating unit tests for database operations using the `pgxmock` library. This will ensure that your database interactions are functioning as intended without requiring a live PostgreSQL instance.

## **Description**
In this task, you'll implement unit tests for your database-related functions using `pgxmock`. This mock library simulates `pgx` behavior, allowing you to test various scenarios and code paths efficiently.

## **Subtasks**

1. **Set Up `pgxmock`**
   - Add the `pgxmock` library to your project dependencies.
   - Import `pgxmock` in your test files.

2. **Write Unit Tests for Database Functions**
   - **`InsertSimple` Function:**
     - Test successful insertion of a single row.
     - Test handling of insertion errors.
   - **`InsertBatch` Function:**
     - Test successful batch insertion.
     - Test handling of batch insertion errors.
   - **`InsertCopy` Function:**
     - Test successful bulk insertion using the `COPY` protocol.
     - Test handling of `COPY` insertion errors.
   - **`FetchSelectScan` Function:**
     - Test successful data retrieval using `SELECT` and `Scan`.
     - Test handling of retrieval errors.
   - **`FetchSelectCollect` Function:**
     - Test successful data retrieval using `pgx.CollectRows`.
     - Test handling of retrieval errors.

3. **Validate Mock Expectations**
   - Ensure that all expected database interactions are met in your tests.
   - Use `pgxmock`'s `ExpectationsWereMet` method to verify that all expectations were fulfilled.

## **Hints**

- **Using `pgxmock`:**
  - Initialize a new mock connection using `pgxmock.NewConn()`.
  - Set up expectations for database operations (e.g., `ExpectExec`, `ExpectQuery`).
  - Use the mock connection in place of a real `pgx.Conn` in your tests.

- **Example Test Structure:**
  - Begin a new test by creating a mock connection.
  - Define the expected database interactions.
  - Call the function under test with the mock connection.
  - Assert the results and check for errors.
  - Verify that all expectations were met.

## **Next Steps**

Upon completing this task, your Go application will have a robust suite of unit tests for its database operations, ensuring reliability and facilitating future development and refactoring.

To check your implementation, proceed to the commit [#33361b](https://github.com/pashagolub/pgxbench/commit/33361bdbde7f30baad4a076bc56ee3259ca890f4) in the repository.

---

# Final Words

Congratulations on completing the **Writing your first Go app for PostgreSQL** training! You've learned how to build a Go application for PostgreSQL, covering essential concepts such as database connections, data insertion, retrieval methods, code refactoring, and unit testing. I hope this training has provided you with valuable insights and hands-on experience in working with PostgreSQL databases using Go.

Do not hesitate to explore further, experiment with different scenarios, and continue enhancing your Go programming skills. Feel free to refer back to this repository for reference or practice, and share your feedback or questions with [me (Pavlo Golub)](https://pashagolub.github.io/).