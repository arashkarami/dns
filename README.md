# Housing-Anywhere
## code challenge. 

This application returns calculated location for the drones for locating data store point.

## Notes

* Go modules used when building the app
* App has 1 post endpoint to calculate and return location and 1 healthcheck to ensure if it is up and ready
* Default port is 5001
* Application handles most of the errors by checking posted values and logs accordingly.
* `SectorID` for observed sector has been kept in environment variables and does not change at the runtime.   
It can be editable through: `.env` 

## How-to-run-app
`docker build -t dns .`  
`docker run -p 5001:5001 -d dns:latest`

Location endpoint serves under swagger document http://localhost:5001/docs
Healtcheck endpoint serves under http://localhost:5001/api/v1/health

FUTURE