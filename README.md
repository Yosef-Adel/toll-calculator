# Toll Calculator

A distributed system for calculating toll fees based on vehicle travel data. The system processes real-time location data from vehicles, calculates distances traveled, and generates invoices for toll fees.

## System Overview

The Toll Calculator is a microservices-based application that processes vehicle location data to calculate toll fees. It simulates a real-world scenario where vehicles equipped with On-Board Units (OBU) transmit their location data for toll calculation purposes.

## Architecture

The project is structured as a mono-repo containing the following microservices:

### Core Services

1. **OBU (On-Board Unit Simulator)**

   - Generates random sample location data
   - Simulates real vehicles sending GPS coordinates
   - Helps in testing the system without actual vehicle data

2. **Data Receiver**

   - Receives real-time location data (latitude, longitude) from vehicles
   - Publishes received data to Kafka for further processing
   - Acts as the entry point for all vehicle location data

3. **Distance Calculator**

   - Consumes location data from Kafka
   - Calculates the distance traveled by each vehicle
   - Stores calculated distances for aggregation

4. **Aggregator**

   - Processes calculated distances
   - Generates invoices containing toll fees and travel information
   - Maintains vehicle travel history

5. **Gateway**
   - Provides REST API endpoints for the system
   - Handles requests for invoice generation
   - Offers cost estimation for planned travel routes

## Prerequisites

- Docker
- Kafka
- Go (latest version recommended)

## Setup and Installation

1. Start Kafka locally using Docker:

```bash
    docker-compose up -d
```

2. Run the services in the following order:
   - Aggregator
   - Data Receiver
   - Distance Calculator
   - OBU (Simulator)

## Usage

### Basic Operations

1. Start all services as described in the setup section
2. The OBU service will automatically start generating sample vehicle data
3. Access the system through the Gateway API

### API Endpoints

The Gateway service provides the following endpoints:

1. **Get Invoice**

   - Endpoint: `/invoice/{obuID}`
   - Method: GET
   - Returns toll fee calculations and travel details for a specific vehicle

2. **Estimate Travel Cost**
   - Endpoint: `/estimate`
   - Method: POST
   - Allows users to estimate toll costs for planned routes

## Data Flow

1. OBU → Sends location data (lat, long)
2. Data Receiver → Receives and publishes to Kafka
3. Distance Calculator → Processes location data and calculates distances
4. Aggregator → Generates invoices and maintains records
5. Gateway → Provides API access to the system

## Testing

To test the system:

1. Ensure all services are running
2. Use the Gateway API with a valid OBU ID
3. Monitor the data flow through Kafka topics

## Development

### Project Structure

```
toll_calculator/
├── obu/
├── data_receiver/
├── distance_calculator/
├── aggregator/
└── gateway/
```
