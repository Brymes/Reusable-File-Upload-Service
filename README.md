# Reusable File Upload Service

This project is a File Upload backend designed to be plug and play(just pull the
docker image and deploy). I built it as a step to ease workload for developers
building MircroServices.

Kindly leave a Star and share so Others who may Find it useful.

Feel Free to Open Issues and Leave PRs.

My Contact:

Twitter:

LinkedIn:

## Table of Contents

- PreRequisites
- Setup
    - Setup for S3
    - Setup for Cloudinary
- Usage
- Testing
- TODOs
- NOTE

## PreRequisites

- Cloudinary Account (Optional)
    - CLOUDINARY_URL
    - CLOUDINARY_API_SECRET
- AWS S3 Account(Optional)
    - AWS_ACCESS_KEY_ID
    - AWS_SECRET_ACCESS_KEY
    - AWS_REGION
- MongoDB URL(Required)
- Docker Installed on Server/Local Machine

## Setup

- Setup Environment Variables(see env.example)
    - ### Setup for Cloudinary
        - Get AWS Access Keys
        - https://aws.github.io/aws-sdk-go-v2/docs/getting-started#get-your-aws-access-keys
        - Add Keys to .env file

### Setup for S3

- Get AWS Access Keys
    - https://aws.github.io/aws-sdk-go-v2/docs/getting-started#get-your-aws-access-keys
    - Add Keys to .env file

## Usage

- `docker pull brymes/reusable-file-upload-service:stable`
- `docker run --platform=linux/amd64 --env-file .env reusable-file-upload-service`

## Testing

- WIP

## TODOs

- Unit tests
- Add New Services
    - Azure Blob Storage
    - Google Cloud Storage
- Authentication Workflow (A POC for how to use in a Microservice Environent)
- Empty upload folder after upload complete
- Consider using goroutines to perform uploads in background
- Error Management with Sentry or similar/ Distributing logs
- Eager Transformations for Cloudinary

## NOTE

Tags always get parsed as strings