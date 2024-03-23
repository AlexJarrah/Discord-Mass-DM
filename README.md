# Discord-Mass-DM

Discord-Mass-DM is a command-line tool designed to send direct messages (DMs) to all members of a Discord server. This tool offers the ability to create a message pool and specify server roles to look for or stay away from when messaging server members. You can also add a wildcard (\*) into the included or excluded role list to include or exclude all members by default. With a message pool, messages will be cycled through after each use for each recipient.

## Installation

To install the tool, run the below commands in the location you want to install it:

```
git clone https://github.com/AlexJarrah/Discord-Mass-DM.git
cd Discord-Mass-DM
```

## Configuration

1. Open the data/config.json file in a text editor.
2. Replace the "discord_token" value with your Discord account token. This token can be obtained from either a user or bot account. You can find your Discord account token by following the steps below:
    - Log into your Discord account using a web browser
    - Press F12 to open Chrome DevTools & navigate to the network tab
    - Filter log by typing `url:discord.com/api` into the filter field near the top of the page
    - Click on any request (you might need to refresh the page)
    - Scroll down to the request headers section and find a header named "Authorization", this is your account token. Keep this token safe, do not share it with anyone.
3. Populate the "message_pool" array with the messages you want to send. Each message should be a separate string within the array.
4. Modify the "roles" field with all roles you want to include or exclude to filter server members.

## Usage

1. Run the main script using the command:
```
go run cmd/Discord-Mass-DM/main.go
```
2. Follow the on-screen instructions to select the server and confirm the operation.

## Disclaimer

Automating a user account is against the Discord Terms of Service. By using this tool, you acknowledge and understand that the developer is not responsible for any consequences or actions taken against your account due to violation of the Discord Terms of Service.

This tool is not intended for spamming, harassment, or any malicious activities, and such behavior is not endorsed by the developer. Use this tool responsibly at your own risk.

## License

This project is licensed under the MIT License.
