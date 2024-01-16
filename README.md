# ParkingLot
## Problem Statement

- I own a parking lot that can hold up to `n` cars at any given point in time. Each slot is given a number starting at 1 increasing with increasing distance from the entry point in steps of one. I want to create an automated ticketing system that allows my customers to use my parking lot without human intervention.
- When a car enters my parking lot, I want to have a ticket issued to the driver. The ticket issuing process includes us documenting the registration number (number plate) and allocating an available parking slot to the car before actually handing over the ticket to the driver. We assume that our customers are nice enough to always park in the slots allocated
to them. The customer should be allocated a parking slot which is nearest to the entry. At the exit the customer returns the ticket with the time the car was parked in the lot, which then marks the slot they were using as being available.
- Total parking charge should be calculated as per the parking time. Charge applicable is $10 for the first 2 hours and $10 for every additional hour.
- We interact with the system via a simple set of commands which produce a specific output. Please take a look at the example below, which includes all the commands you need to support - they're self-explanatory.

## Command
- `create [size]` -> Creates parking lot of size n
- `park [car-number]` -> Parks a car (time start counting when customer park)
- `leave [car-number] ` -> Removes (unpark) a car and charges based on the number of hours the car is parked
- `status [car-number]` -> Prints status of the parking lot (slot, parking time, charges based on the number of hours the car is parked)
- `exit` -> Exit program
