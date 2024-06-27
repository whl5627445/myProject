# Project Name

Thanks to [JetBrains](https://www.jetbrains.com/) for sponsorship

## Introduction
This is a web API project based on the OpenModelica project's backend omc.

## Project Goals and Vision
The goal of this project is to provide users with a backend API based on omc, separating the graphical display of Modelica models, model operations, simulation execution, and Modelica code parsing. In short, it aims to enhance and simplify functionalities based on omc, making it easier for other developers to use.

## Features
1. **Model Compilation**
2. **Model Simulation**
3. **Model Source Code Parsing**
4. **Model Instantiation**
5. **More features are under development**

## Dependencies
- Official installation of omc, minimum version 1.21
- Golang minimum version 1.21
- MySQL database for storing user information, model information, etc.
- This project does not include a graphical display component; users need to develop their own graphical rendering code

## Installation Guide
1. **Install omc**: [Official Download Link](https://openmodelica.org/download/download-linux/)
2. **Install Golang**: [Official Download Link](https://go.dev/doc/install)
3. Clone this project: `git clone <repository-url>`
4. Configure the database address in the `config` file
5. Compile the project: `go build main.go`

## Deployment Guide
1. If you have completed all the configurations and installations according to the installation guide, you just need to run the compiled binary file.
2. It is currently recommended to run on a Linux server.

## Example
### Configuration File
Ensure your configuration file correctly points to the MySQL database and configures other necessary options, such as the path to omc.

Starting the Project

In the project’s root directory, run:

    go build main.go
    ./main

Now, your API server should be up and running, and you can access the respective endpoints to use various features.

FAQ 
#### 1.	How to resolve omc path errors?
#### Ensure you have correctly specified the path to omc in the configuration file and that the omc executable has the appropriate permissions.
#### 2.	How to handle database connection errors?
#### Check your database configuration, make sure the database server is running, and the credentials in the configuration file are correct.

Contribution

Contributions to this project are welcome! Please submit Pull Requests or Issues.

If you have any other questions or need further assistance, feel free to contact me.