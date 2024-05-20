# GoSecurePath

GoSecurePath is a project that helps you create a secure authorization server using the OAuth2 protocol with the Go programming language. This project aims to demonstrate how to set up an OAuth2 authorization server and a client application that securely grants access to third-party applications.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Introduction
OAuth2 (Open Authorization) is a protocol that allows users to grant third-party applications access to their resources without sharing their credentials. This project provides a step-by-step guide to creating an OAuth2 authorization server and client using Go.

## Features
- Secure authorization server implementation
- Memory-based token storage
- Support for authorization code flow
- Client application for testing authorization

## Getting Started
Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- Go programming language installed (version 1.16 or later)
- Git installed
- A GitHub account

### Installation
1. **Clone the repository:**


    git clone https://github.com/TolgaTD/GoSecurePath.git
    cd GoSecurePath


2. **Initialize the Go module and download dependencies:**


    go mod tidy


## Usage
1. **Run the authorization server:**


    go run main.go


2. **Run the client application:**

    Open another terminal window and navigate to the project directory, then run:


    go run client.go


3. **Send an authorization request:**

    Open your web browser and navigate to the following URL:


    http://localhost:9096/authorize?response_type=code&client_id=client_id&redirect_uri=http://localhost:9094/callback&scope=all&state=xyz


4. **Exchange the authorization code for an access token:**

    After receiving the authorization code, use the following `curl` command to get the access token:


    curl -X POST -d "grant_type=authorization_code&code=AUTH_CODE&redirect_uri=http://localhost:9094/callback&client_id=client_id&client_secret=client_secret" http://localhost:9096/token


5. **Access a protected resource:**

    Use the obtained access token to access a protected resource:


    curl -H "Authorization: Bearer ACCESS_TOKEN" http://localhost:9096/protected-resource


## Project Structure
- `main.go`: The main file that runs the authorization server.
- `client.go`: The client application that handles the authorization code flow.
- `go.mod`: Go module file that defines the project dependencies.

## Contributing
Contributions are welcome!

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
