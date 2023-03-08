# Self Payroll System

This is a web service that allows employees to withdraw their salaries independently every month. The system includes the following features:

1. Position Management: CRUD operations (Create, Read, Update, Delete) to manage position data.
2. Employee Management: CRUD operations (Create, Read, Update, Delete) to manage employee data.
3. Admin Balance Top-up: Admin can top up the company balance.
4. Salary Withdrawals: Employees can withdraw their salaries by providing their Employee ID and Secret ID. The salary amount is based on the position held by each employee.
5. Transaction History: Transaction history of top-ups and reductions of the company's balance.

## Getting Started

To run the application, follow the instructions below:

1. Clone this repository and navigate to the project directory.
2. Create a `.env` file by running `cp .env.example .env`. Then fill in the `.env` file according to your local environment.
3. Run `go mod tidy && go mod vendor`.
4. Run `go run *.go`.

The list of endpoints is available in the [documenter](https://documenter.getpostman.com/view/4080490/2s83Ychhk4).

## Technologies Used

- Go programming language

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
