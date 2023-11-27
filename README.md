# cse-4050-project

cse-4050-project (MedConnect Dashboard) is a web application designed to facilitate seamless communication and appointment scheduling between patients and doctors. Developed as part of a web development class, this platform aims to enhance the overall healthcare experience by providing a centralized space for patients and doctors to connect.

## Features

- **Appointment Scheduling:** Patients can easily view and choose available time slots provided by doctors for appointments.
- **Availability Management:** Doctors can manage their availability, making it convenient for patients to find suitable appointment times.
- **Post-Appointment Summaries:** Doctors can efficiently write and share summaries after appointments, allowing patients to understand and reference the details of their medical visits.

## Installation

Before you begin, ensure you have met the following requirements:

- [Go](https://golang.org/) 1.20 or higher installed.

1. Clone the repository:
   ```bash
   git clone https://github.com/lumix103/cse-4050-project.git
   ```
2. Navigate to the project:
   ```bash
   cd cse-4050-project
   ```
3. Install packages:
   ```bash
   # To download and install packages
   go get ./...
   ```

## Environment Variables

To run this project, you will need to add the following environment variables to your `.env` file. If you don't have an `.env` file, create one in the root folder.

`MONGODB_URI`
`JWT_SECRET_KEY`

## Deployment

To deploy this project run

```bash
  go run ./cmd/dashboard/
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
