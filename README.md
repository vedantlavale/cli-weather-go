# Weather CLI

This is a simple command-line application written in Go that fetches and displays weather information for a given location using the WeatherAPI.

## Features
- Fetches current weather information for a specified location.
- Displays hourly weather forecasts, including temperature, chance of rain, and conditions.
- Highlights hours with a high chance of rain in red.

## Prerequisites
- Go installed on your system (version 1.16 or later).
- A WeatherAPI account to obtain an API key.
- A `.env` file containing your API key.

## Installation
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd cli
   ```

2. Create a `.env` file in the project directory and add your WeatherAPI key:
   ```env
   API_KEY=your_weatherapi_key_here
   ```

3. Build the application:
   ```bash
   go build -o weather-cli
   ```

## Usage
Run the application with the following command:
```bash
./weather-cli [location]
```
- Replace `[location]` with the name of the city or location you want to query (e.g., `Pune`).
- If no location is provided, the application will prompt you to enter one.

### Example
```bash
./weather-cli Pune
```
Output:
```
Pune,India: 30C,Sunny
12:00 - 30C, 10%, Sunny
13:00 - 32C, 20%, Partly Cloudy
14:00 - 33C, 50%, Rain
```
- Hours with a high chance of rain (e.g., 50%) will be highlighted in red.

## Dependencies
This project uses the following Go packages:
- [fatih/color](https://github.com/fatih/color): For colored terminal output.
- [joho/godotenv](https://github.com/joho/godotenv): For loading environment variables from a `.env` file.

Install dependencies using:
```bash
go mod tidy
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments
- [WeatherAPI](https://www.weatherapi.com/) for providing the weather data.
