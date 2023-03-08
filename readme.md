# Self Payroll System

## Description

Self Payroll System is a web service built with Go that allows employees to withdraw their salaries independently every month. This project was created as the final task for the Project-Based Virtual Internship at the Core Initiative organized by Rakamin Academy as a Backend Developer.

## Features

The following features are included in the system:

1. Position Management: CRUD operations (Create, Read, Update, Delete) to manage position data.
2. Employee Management: CRUD operations (Create, Read, Update, Delete) to manage employee data.
3. Admin Balance Top-up: Admin can top up the company balance.
4. Salary Withdrawals: Employees can withdraw their salaries by providing their Employee ID and Secret ID. The salary amount is based on the position held by each employee.
5. Transaction History: Transaction history of top-ups and reductions of the company's balance.

## Tools

The following tools were used to build this project:

- Echo Framework
- PostgreSQL as database
- GORM as ORM
- ozzo-validation as input validation

## Getting Started

To run the application, follow the instructions below:

1. Clone this repository and navigate to the project directory.

   ```bash
   git clone https://github.com/szczynk/self-payroll.git
   ```

1. Create a `.env` file by running `cp .env.example .env`. Then fill in the `.env` file according to your local environment.

   ```bash
     cp .env.example .env
   ```

1. Run `go mod tidy && go mod vendor`.

   ```bash
     go mod tidy && go mod vendor
   ```

1. Run `go run *.go`.

   ```bash
     go run *.go
   ```

The list of endpoints is available in the [documenter](https://documenter.getpostman.com/view/4080490/2s83Ychhk4).

## Technologies Used

- Go programming language

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
