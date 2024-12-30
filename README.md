# Gator

This project is a command-line application that allows users to manage user accounts, feeds, and follow feeds using a variety of commands. Below is a guide on how to use the available commands.

## Available Commands

### User Management
- **login `<name>`**: Log in as an existing user.
- **register `<name>`**: Create a new user account.
- **reset**: Reset the database by deleting all users.
- **users**: List all users. The current user will be marked.

### Feed Management
- **addfeed `<name>` `<url>`**: Add a new feed with a specified name and URL. Automatically follow the feed after creation.
- **feeds**: List all available feeds along with their details.

### Feed Actions
- **follow `<feed_url>`**: Follow a feed by its URL.
- **following**: List all feeds the current user is following.
- **unfollow `<feed_url>`**: Unfollow a feed by its URL.
- **browse `[limit]`**: Browse posts from the feeds the user is following. Optionally specify a limit for the number of posts to view (default: 2).

### Aggregation
- **agg `<time_between_reqs>`**: Continuously fetch feeds at a specified interval (e.g., "10s" for 10 seconds).

## Config File
- Project assume that a config file name `.gatorconfig.json` is available at the `Home` directory.
- Format of the config file
```json
{
    "DBURL": <PostgreSQL-db-connection-string>,
    "CurrentUserName": <Current-User>
}
```

## How to Run
1. Ensure you have the required dependencies installed.
2. Ensure PostgreSQL database installed.
3. Ensure config file with required parameters is available.
4. Build and run the application.
5. Use the above commands in the application to interact with users and feeds.

## Notes
- Some commands (e.g., `addfeed`, `follow`, `browse`) require the user to be logged in.
- Ensure proper usage of each command by following the syntax mentioned above.

