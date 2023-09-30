# Go-inventory
Welcome to PDM's take home coding project. The purpose of this project is to implement an application that allows users to manage inventory data at scale. Your task is to design and implement a data model along with read/write APIs for data retrieval and data management.

One of the challenges in the automotive aftermarket vehicle parts industry is managing an inventory of vehicle parts where each part is described by multiple data points. For example, below is a subset of the data points used to describe each vehicle part:

Part name

Image(s) of that part

SKU #

Description

Price

Attributes (ie. color, size, etc)

Fitment data (ie. All year/make/models this part can be used in)

Location (ie. Front left, front right, etc)

Shipment packaging (ie. Weight, size, hazardous, fragile, etc)

Additional metadata to describe each part
 
### Your deliverables include:

1. Design and implement a data model to represent this inventory

2. Implement APIs to support all CRUD operations

3. Save data to storage

4. Implement data versioning such that a user can retrieve any version

OPTIONAL: Implement search for an inventory of more than a million vehicle parts (this deliverable is not required)

## To run
Add your postgres connection string to .env file.

To create the corresponding database tables based on project models I used GORM's AutoMigrate function in main.go file:
db.AutoMigrate(&models.Part{}, &models.Attribute{}, &models.Fitment{}, &models.Image{}, &models.Metadata{}).

To launch: go run main.go

To test: go test
