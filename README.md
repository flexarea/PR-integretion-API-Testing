# **Integrated Automation System for GitHub PR Actions, Trello Cards, and Slack**
-------------------------

### **Overview**
**This integration is an internal tool to enhance CI/CD workflows by automating the retrieval of information from GitHub PR actions to update Trello cards and send notifications to a target Slack channel**

### 1. **Technologies stack and API Documentation**


![Static Badge](https://img.shields.io/badge/Trello-API-blue?logo=Trello&link=https%3A%2F%2Fdeveloper.atlassian.com%2Fcloud%2Ftrello%2Frest%2Fapi-group-actions%2F%23api-group-actions)
![Static Badge](https://img.shields.io/badge/Slack-API-orange?logo=Slack&link=https%3A%2F%2Fapi.slack.com%2Fweb)
![Static Badge](https://img.shields.io/badge/Github-REST-white?logo=Github&link=https%3A%2F%2Fdocs.github.com%2Fen%2Frest%2Factions%2Fworkflow-jobs%3FapiVersion%3D2022-11-28)
![Static Badge](https://img.shields.io/badge/Cloud-platform-red?logo=Google%20Cloud&cacheSeconds=https%3A%2F%2Fcloud.google.com%2F)
[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/) 

- github.com/joho/godotenv v1.5.1 -> to handle .env variables
- Database
  
### 2. **Requirements**
   - **Functional Requirements**
     - Bot Token configuration for slack authentication
     - API token and API key configuration for Trello authentication
     - Error Handling
   - **Non-Functional Requirements**
     - Performance and Scalability
     - Security
     - Usability
     - Maintainability
     - Reliability

### 3. **System Architecture**
```
API-Server-Integration/
├── cmd/
│   └── server/
│       └── main.go            # Entry point of the application
├── internal/
│   ├── trello/
│   │   ├── trello.go          # Functions for interacting with Trello API
│   │   ├── webhook.go         # Handles GitHub webhook events
│   ├── slack/
│   │   └── slack.go           # Functions for interacting with Slack API
│   ├── config/
│   │   └── config.go          # Configuration and environment variable handling
│   ├── http/
│   │   └── server.go          # HTTP server setup and route handling
├── pkg/
│   ├── models/
│   │   └── models.go          # Data models and types used in the application
├── scripts/
│   └── deploy.sh              # Deployment scripts and infrastructure setup
├── .env                       # Environment variables file ( API keys, main endpoints, Bot and api token)
├── .gitignore                 # Git ignore file (ignore .env)
├── go.mod                     # Go module file
├── go.sum                     # Go module dependencies file
└── README.md                  # Project documentation
```
   - **Deployment Architecture**
     - Google Cloud -> Server & cloud service
     - (Deployment Pipeline)

### 4. **API Specifications**
   - **Endpoints**
- Trello main endpoint: https://api.trello.com/1/
- Slack main endpoint: https://slack.com/api/

     - HTTP Methods (GET, POST, PUT, DELETE, etc.)
   - **Request Parameters**
     - Path Parameters Trello: {API_TOKEN, API_KEY}
     - The Slack API requires an authorization header for is (Bearer) authentication with the BOT token
     - Will use JSON payloads for POST & PUT methods instead of query parameters 
   - **Response Format**
     -Data format: application/json
     - Success Response
     - Error Response
   - **Examples**
  ```go
  //making a request
  req, err := http.NewRequest(method, URL, body io.Reader)
  //handle error here

  //response
  res, err := http.DefautClient(req)
  //handle error here
  databyte, err := io.ReadAll(res.Body)
  /handle error here
  ```

### 5. **Data Model**
   - **Database Schema**
     - Tables/Collections
     - Relationships
   - **Data Flow**
     - Data Input
     - Data Processing
     - Data Output

### 6. **Security Considerations**
   - **Authentication Mechanisms**
   - **Authorization Levels**
   - **Data Encryption**
   - **Vulnerability Management**

### 7. **Error Handling and Logging**
   - **Error Types and Codes**
   - **Error Responses**
   - **Logging Strategy**
     - Log Levels
     - Log Format
     - Log Storage and Rotation

### 8. **Testing Strategy**
   - **Unit Testing**
   - **Integration Testing**
   - **End-to-End Testing**
   - **Performance Testing**
   - **Security Testing**

### 9. **Performance and Scalability**
   - **Performance Metrics**
   - **Load Balancing**
   - **Caching Strategy**
   - **Database Optimization**
   - **Scalability Plan**

### 10. **Deployment and Maintenance**
   - **Deployment Process**
   - **Environment Configuration**
     - Development
     - Staging
     - Production
   - **Monitoring and Alerts**
   - **Backup and Recovery**
   - **Maintenance Plan**

### 11. **Documentation and Training**
   - **User Documentation**
   - **Developer Documentation**
   - **API Documentation**
   - **Training Materials**

### 12. **Project Management**
   - **Timeline and Milestones**
   - **Resource Allocation**
   - **Risk Management**
   - **Communication Plan**

### 13. **Appendices**
   - **Glossary**
   - **References**
   - **Additional Notes**

---

This outline covers all essential aspects of a backend software project specification, ensuring a comprehensive and clear documentation for developers, stakeholders, and other involved parties.

