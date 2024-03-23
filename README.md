# Discord-Mass-DM

Discord-Mass-DM is a command-line tool designed to send direct messages (DMs) to all members of a Discord server. It offers the ability to specify multiple messages, which then get cycled through for each recipient.

## Installation

To install the tool, run the below commands in the location you want to install it:

```
git clone https://github.com/AlexJarrah/Discord-Mass-DM.git
cd Discord-Mass-DM
```

## Configuration

1. Open the data/config.json file in a text editor.
2. Replace the "discord_token" value with your Discord account token. This token can be obtained from either a user or bot account.
3. Populate the "message_pool" array with the messages you want to send. Each message should be a separate string within the array.

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
