# Chatbook (chat application clone)
Chatbook is a hobby project of mine. It aims to create a highly available chat system.

## Functional Requirements:

- [ ] **User Registration and Authentication** -allow users to register an account with a unique username and password.
  Provide authentication mechanisms to ensure secure access to the chat application.

- [ ] **Real-Time Chat Features** - enable users to create chat rooms or join existing rooms. Support real-time messaging with features like sending and receiving text messages and emojis.Allow users to view the list of online users in a chat room. Implement message history, allowing users to scroll back and view previous messages.

 - [ ] **User Management** - provide options for users to update their profile information. Allow users to add or remove friends and manage their contact list. Support blocking or reporting other users for inappropriate behavior.

- [ ] **Notifications** - implement notifications for new messages, friend requests, or other relevant events.
Allow users to customize notification preferences.

- [ ] **Search and Filtering** - enable users to search for specific chat rooms or messages.Implement filters to sort and display chat rooms based on criteria such as popularity or user activity.

- [ ] **Caching of Chat History** - implement a caching mechanism to store the chat history of the last hour. When a user requests chat history, check the cache first and retrieve the relevant messages if available. If the requested chat history is not present in the cache, fetch it from the database and update the cache for future requests. Set an expiration time for the cached chat history, ensuring it is automatically refreshed as per the configured duration (e.g., every minute).


## Non-Functional Requirements:

- Responsiveness - ensure the application is responsive across different devices and screen sizes.

- Scalability - estimate the expected number of users per hour and ensure the system can handle the load. For example, expect 1,000 users per hour initially, with room for growth. Design the system to scale horizontally by adding more servers or using load balancers to distribute the load effectively.

- Reliability and High Availability - ensure the chat application is highly available, with minimal downtime.
Implement fault-tolerant mechanisms to handle server failures or network disruptions. Use load balancers to distribute traffic and prevent bottlenecks.

- Performance - Define performance benchmarks such as the number of concurrent users the system can support, the average message delivery time, or the response time for various operations.
Optimize the system's performance by implementing efficient algorithms, caching frequently accessed data, and minimizing database queries.

- Security - implement secure authentication mechanisms to protect user accounts and prevent unauthorized access.
Enforce secure communication protocols to protect message content and user privacy.
Usability:

Design an intuitive user interface with features like chat bubbles, typing indicators, and online/offline indicators for a seamless user experience.


## System Design

