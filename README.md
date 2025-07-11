# Cash Log - CLI Based Application

A simple terminal-based finance tracker written in Go, using PostgreSQL for persistent storage. Track your income, expenses, and view detailed reports — all from your command line.

---

## Features

- Add, edit, delete transactions (income/expense)
- Manage categories and transaction types dynamically
- View transaction history with pagination, sorting, and filtering
- View current balance summary
- Generate monthly, weekly, yearly financial reports
- Interactive CLI prompts for ease of use
- Store data securely in PostgreSQL
- Configurable via `.env` file

---

## Requirements

- Go 1.20+  
- PostgreSQL  
- `make` (optional, if you use a Makefile)

---

## Setup

1. Clone the repo:
    ```bash
    git clone https://github.com/yourusername/cash-log-cli.git
    cd cash-log-cli
    ```
2. Create a PostgreSQL database:
    ```bash
    CREATE DATABASE cashlog;
    ```
3. Create a `.env` file in the project root with your database credentials:
    ```bash
    DB_HOST=hostname
    DB_PORT=port
    DB_USER=yourusername
    DB_PASSWORD=yourpassword
    DB_NAME=cashlog
    ```

4. Install dependencies:

    ```bash
    go mod tidy
    ```
5. Build the application:
    ```bash
    go build -o cash-log
    ```
6. Run the app:
    ```bash
    ./cash-log
    ```
### Using Makefile

- Build the app:
    ```bash
    make build
    ```
- Run the app:
    ```bash
    make run
    ```
- Clean dependencies:
    ```bash
    make tidy
    ```
- Remove binary:
    ```bash
    make clean
    ```

## Usage
Run the app with one of these commands:

```bash
./cash-log add           # Add a new transaction
./cash-log view          # View transaction history
./cash-log edit          # Edit a transaction by ID
./cash-log delete        # Delete a transaction by ID
./cash-log balance       # Show income, expenses, and balance summary
./cash-log add-category  # Add new category
./cash-log add-type      # Add new transaction type
./cash-log report        # Generate detailed financial reports
```

Each command will prompt you interactively for required inputs.

## Contributing
Contributions and suggestions are welcome! Feel free to open issues or submit pull requests.

## License
MIT License © Bishal Shrestha

## Contact
Created by Bishal Shrestha
Feel free to reach out!