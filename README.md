# Tender

When creating microservices, the need for a simple file server often arises to serve production assets from prior build steps. Tender was designed to meet this need with no dependencies and with an extremely permissive license ([ISC](https://choosealicense.com/licenses/isc/)).

The name of the project comes from purpose built boats that handle specific functions—some carry people from cruise ships to docks, some maintain buoys, some transport mail, etc.

This project utilizes the concept of tenders to provide different server functionality, the current tenders are:

- file

    Serves a single file at a specified route.

- directory

    Serves all files underneath a directory and provides a list of all files within the directory when the directory itself is requested.

- directory no listing

    Serves all files underneath a directory without providing a listing of all files within the directory.

## Notice

At this point in time, the API used within the Go code is absolutely prone to change. While the intent of this service is simply file serving, there may be additional tenders introduced. If you have suggestions for improving the API please feel free to open an issue or PR.
