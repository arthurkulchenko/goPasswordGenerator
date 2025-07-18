# Password Generator
This is a small project and it generates passwords

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation
1. Clone the repository:
```bash
 git clone https://github.com/arthurkulchenko/goPasswordGenerator
```

2.a. To build and run locally:
```bash
  go build -o bin/pass_gen
  ./bin/pass_gen
 ```

2.b. To run from source:
```bash
  go run main.go
 ```

2.c. To build and run in container:
```bash
  GOOS=linux GOARCH=amd64 go build -o bin/pass_gen
  podman build -t pass_gen:latest -f docker/Dockerfile .
  podman run -it --name pass_gen pass_gen
```

## Usage
To run the project, use the following command:
```bash
  ./bin/pass_gen
```

## Contributing
1. Fork the repository.
2. Create a new branch: `git checkout -b feature-name`.
3. Make your changes.
4. Push your branch: `git push origin feature-name`.
5. Create a pull request.

## License
This project is licensed under the [MIT License](LICENSE).

![Build Status](https://github.com/arthurkulchenko/goPasswordGenerator/actions/workflows/test.yml/badge.svg)
