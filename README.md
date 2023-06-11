# Weather API 🌤️

The Weather API is a RESTful API that provides weather information for different cities. It uses the OpenWeatherMap API to fetch real-time weather data based on the provided city name.

## Prerequisites

Before running the Weather API, make sure you have the following prerequisites installed:

- Go programming language (version 1.16+)
- Docker (optional, if you want to run the API inside a Docker container)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/weather-api.git
   
2. Change into the project directory:

   ```bash
   cd weather-api
   
3. Rename the .env.example file to .env and provide your OpenWeatherMap API key inside the .env file:

   ```bash
   mv .env.example .env

4. Build and run the API:

   
  - If you have Go installed:
   ```bash
   go run main.go
   ```
 
  - If you have Docker installed:
   ```bash
   docker build -t weather-api .
   docker run -p 8000:8000 -d weather-api
   ```
   
## Usage  🚀

Once the API is up and running, you can make HTTP GET requests to retrieve weather information for a specific city.

## Endpoint
  ```bash
GET /weather/{city}
  ```

## Example

Request:
   ```bash
   GET /weather/London
   ```

Response:
   ```bash
   {
  "temperature": 19.5,
  "description": "Cloudy"
}
  ```
  
## Contributing 🤝
Contributions are welcome! If you find any issues or want to add new features to the Weather API, feel free to open an issue or submit a pull request.

Please make sure to follow the existing code style and write appropriate tests for your changes.

## License 📝
  This project is licensed under the MIT License.
  
## Acknowledgements 🙏
   The Weather API is built using the following libraries:
 -  Gorilla Mux: https://github.com/gorilla/mux
 -  Resty: https://github.com/go-resty/resty
 -  godotenv: https://github.com/joho/godotenv


