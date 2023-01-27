

### Description
This microservice is responsible for consuming a third-party API to get data for game tournaments and matches, including Dota 2, LoL, CS:GO, and Valorant. It provides a simple interface to access this data, allowing other applications to easily use it.

### Installation
Clone the repository: **git clone https://github.com/duel-me-next-level/api-consumer**
Enter the project directory: **cd api-consumer**
Install the dependencies: **go get -d**
Build the project: **go build**
Run the project: **./api-consumer**

### Usage
This microservice exposes a series of endpoints to access tournament and match data. Here are some examples of how to use them:

Get all tournaments
```
GET /tournaments
```
Get a specific match
```
```
GET /match/{id}
```
Get all matches for a tournament
```
GET /tournament/{id}/matches
```
Get all available games
```
GET /games
###   Configuration
Before using the microservice, you will need to configure the API key to access the third-party API. This can be done by editing the **config.yml** file and inserting your API key in the **api_key** property.

### Contributions
This is an closed project, contributions are always welcome! If you have any suggestions or find any bugs, please open an issue or submit a pull request.

### Credits
This project was developed by [blarth](https://github.com/blarth) with contributions from [lodustitan](https://github.com/lodustitan).

### License
This project is licensed under the [License name](link to license).

### Note
Make sure you have permission to access the third-party API before using this microservice.
The third-party API may have daily usage limits, so make sure to use this microservice efficiently and avoid making unnecessary calls.
It's recommended to cache the returned data to avoid making frequent calls to the third-party API.
Please let me know if you have any questions or need further help.
