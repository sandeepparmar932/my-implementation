#About Project 

This is the implementation of basic wallet system.It was created in the machine coding round of flipkart. The basic functionality is implemented here.


Asked question -

You are supposed to make a digital wallet system that allows people to transfer money among their wallets. The wallet system uses its own currency called FkRupee (F₹). The smallest amount that the users can transfer is F₹ 1.

Wallet account can be of 2 types -

Personal
Username
Balance amount
Business 
Company name
GST
Balance amount
Wallet allows following operations -  

CreateWallet <Personal/business> <user,gst> <amount> - creates a new wallet of type specifiedwith a balance of F₹ <amount> in the name of <accountHolder1>.
AddMoney <user> <amount>- add the money to the wallet.
TransferMoney <user1> <user2> <amount> - transfer money from user1 account to user2 account.
Business wallet can’t transfer money to the Personal account.
Business wallet can transfer money to business account
Personal wallet can transfer money to personal account or business account.
Balance <user>- should display the balance amount for the user.
Overview - should display the current balance of all the accounts.
Conditions imposed on the wallet -

Personal account need to maintain a minimum balance of F₹ 50
Business accounts can maintain the zero balance.
Offer - When User A transfers money to user B and both the account holders have the same balance after the transaction then both the customers get F₹ 10 as a reward.

Bonus Question (Only if time permits) :-

Apply a fixed  F₹ 5 transaction fee for all transfers from Personal account to Personal account. Display the total transaction fee earned by Flipkart.

The examples below are just to understand the functionalities and may not necessarily be used in the same format as input to your driver program.

Sample Inputs
(Bold is user input)

CreateWallet Personal Harry 100

CreateWallet Personal Ron 95

CreateWallet Personal Hermione 104

CreateWallet Business Albus GST1234 200

CreateWallet Business Draco GST5678 500

Overview

Harry        100

Ron 95

Hermione 104

Albus 200

Draco 500

TransferMoney Albus Draco 30

TransferMoney Hermione Harry 2

TransferMoney Albus Ron 5

Failed

Overview

Harry        112

Ron 95

Hermione 112

Albus 170

Draco 530

TransferMoney Harry Albus 100

Failed

Balance Harry

Harry 112

Balance Albus

Albus 170



Guidelines:

Use language of your choice.
Feel free to store all interim/output data in-memory. Use of any external persistent datastore is not required.
Restrict internet usage to looking up syntax or API references
Save your code/project by your name and email it to the email address provided by the interviewer. Your program will be executed on another machine. So, explicitly specify any dependencies in your email.
Do not create rest apis, use simple driver class to trigger functionalities.
Executing test cases or simple Main function should be sufficient
Do not use any database store, use in-memory data structure.


Expectations:

Have working and demonstrable code
Feature requirements should be strictly followed. Work on the required functionalities first and then work on bonus requirements.
Use of proper abstraction, modularity and separation of concerns is required.
Code should easily accommodate new requirements with minimal changes.
Proper exception handling is required.
Writing test cases is encouraged.
Code should be modular, readable and unit-testable.
