# NearLuckyRobot
This is a project for NEAR Certified Developer Program demo day.
## Indroduction
this is a telegram red packet bot for near. The bot is based on https://github.com/tucnak/telebot and written in go language, and the contract is written in rust while i use the near-js-api to call the contract.
## Use
you need mysql, and config the database, details can be found in config directory.
## Functions
all the funtions and commands are following:
### /bind
the bind command can help you bind your telegram account and near account, before you use the bot, you must bind your near account
### /deposit
if you want to send the red packet, you must transfer to the robot, and give the robot transactions hash
### /lucky
you can use this command to send the lucky drop, the amount of near is random divided, and robot will give you the private key, people who own the key can get the near in the red packet
### /claim
after input this command, input your private key then you will get a random near drop
### /withdraw
if you want to get your money back, use this command
### /show
this command will show how many near you own
