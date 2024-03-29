# LCG Random Number Generator in Go 🎲

## About

This repository contains an implementation of the Linear Congruential Generator (LCG) for generating pseudorandom numbers in Go. LCG is a simple and efficient algorithm used in various applications due to its ease of implementation and reasonably good statistical properties.

## Features

- **Customizable Parameters:** Users can specify their values for the modulus, multiplier, increment, and seed.
- **Default Configuration:** A default LCG setup is provided for quick usage.
- **Range-Based Random Numbers:** Generate random integers within a specified range.

## Installation

### Prerequisites

Before running this project, ensure you have Go installed on your system. If not, follow the installation steps below:

#### Installing Go

1. **Download Go:**
   - Visit the [official Go downloads page](https://golang.org/dl/).
   - Select the appropriate package for your operating system and architecture.

2. **Install Go:**
   - Follow the installation instructions specific to your platform.

3. **Verify Installation:**
   - Open your terminal or command prompt.
   - Run `go version`.
   - You should see the installed version of Go displayed.

### Setting Up the Project

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/yourusername/your-repo-name.git
   cd your-repo-name
   ```

2. **Navigate to the Project Directory:**
   ```bash
   cd path/to/your-project
   ```

## Usage

To use the random number generator:

1. **Import the Package:**
   - Include the RNG package in your Go file.
     ```go
     import "path/to/your-project/rng"
     ```

2. **Create an RNG Instance:**
   - You can create a new instance of the RNG with or without custom parameters.
     ```go
     generator := rng.NewRandomNumberGenerator() // With default settings
     // OR
     generator := rng.NewRandomNumberGenerator(modulus, multiplier, increment, seed) // With custom settings
     ```

3. **Generate Random Numbers:**
   - Use the `GetRandomInt` method to generate random integers within a specified range.
     ```go
     randomNumber := generator.GetRandomInt(min, max)
     ```

4. **Add your own Random Number Generator:**
    - Add a different function that uses the paramters and generates random number in a different way and replace with LCG in Next function, it should seemlessly use this new method
    ```go
        func (rng *randomNumberGenerator) Next() int64 {
            rng.seed = LCG(rng)
            return rng.seed
        }
    ```
### Output
![RNG Output](output.png)

## Contributing

Contributions to improve this project are welcome! Feel free to fork the repository, make changes, and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the [MIT License](LICENSE).