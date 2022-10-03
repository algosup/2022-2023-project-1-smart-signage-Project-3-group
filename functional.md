# Functionnal specifications

- name and role

## Table of Contents

-  [Introduction](#overview)


-  [Scenario](#scenario)


-  [Requirements](#needs)


-  [Risks and assumptions](#risk)


-  [Error reporting](#error)


-  [Terminal](#terminal)


-  [Out of scope](#out-of-scope)


<a  name="overview"/></a>

## Introduction

[SignAl](https://signall.com/) wants to invent a new product which would allow them to control the luminous display of the signs they propose remotly and to know the state of the LEDs. The objective behind this project is to enable the company to save on electricity and reduce its ecological footprint.

<a name="scenario"/></a>

## Scenario

<u> Scenario 1 </u> : Patrick is the owner of a Burger King restaurant. He is concerned about the environment and wants to save energy this winter. He would like to be able to control the consumption of his signs and add options to save money. 

<u> Scenario 2 </u> : Paul is a member of the maintenance team. He would like to know the status of the LEDs and be able to control them remotely to know when to move and intervene. This would avoid unnecessary travel.

<a name="needs"/></a>

## Requirements

The two most important points for our customers are its ecological impact and its electricity consumption. To answer his needs we can develop some feature. He ask for remote control too.
To reduce his electricity consumption :

- Turn on the light
- Turn off the light
- In accord to the electricity price, turn off or reduce the light
- Reduce the light intensity
- Know the led status and control it remotly
- turn off the light in accord with the law
- Send a notif when a led is down (out of scope)
- Programable light wich light up on various hours (out of scope)

<a name="risk"/></a>

## Risks and assumptions

### About the laws
According to the law the LEDs must be off during the night, between 1:00AM and 6:00AM.

<!-- ## Need to verify what the law tells about this type of project. -->

### Energy crisis <!--(out of scope ?) (requirements) -->
According to the energy crisis, a new tool will maybe appears to keep up to date about the level of electricity available in France in real time. It is called "carte météo de l'éléctricité".
It will be important to program the LEDs according to the estimates of this card.


## Check privacy

<a name="error"/></a>

## Error reporting

If an error occurs, a text will appear in the terminal. The error message may be different depending on the error encountered.

| Error             	| Description                 	|
|-------------------	|-----------------------------	|
| command not found 	| Can not execute the command 	|

<a name="terminal"/></a>

## Terminal

Displayed when the terminal is open, the home page serves three purposes:

- login page
- home page

### Login
To more security, to access our project the user needs to connect itself with :
- user job title (job title section allows for up to 20 characters to be typed.)
- password (The password section allows for up to 12 characters to be typed. To disguise them and prevent hacking, as the user types in the password box, asterisks (*) will appear instead of the characters that they type.)

### Home page
To avoid any mistakes we will display the home page according to the login entered.
- If the user is connected as a member of the maintenance team, he will see the state of all the signages and thier LEDs.
- If the user is connected as owner of the place where the signage is, he will be able to control the LEDs of his panel. He can not have access to the other brand signage.
Then they will find all the commands they can use and their description :

| Commands 	| Description                        	| Access restricted to           	|
|----------	|------------------------------------	|--------------------------------	|
| check    	| Check the state of the LEDs        	| Maintenance team               	|
| off      	| Turn off the LEDs                  	| The brand where the signage is 	|
| on       	| Turn on the LEDs                   	| The brand where the signage is 	|
| up       	| Increase the intensity of the LEDs 	| The brand where the signage is 	|
| down     	| Decrease the intensity of the LEDs 	| The brand where the signage is 	|

<a name="out-of-scope"/></a>

## Out of scope

To a better user experience, create a web interface will be easier. It will be use to manage the LEDs and see their states.
We will use all the features mentioned above to develop our web interface.

<!-- Questions
- même interface pour maintenance -> login
- qui sont les clients et les différences en conséquences 
- vérifier les entrées du login pour être le plus clair : intitulé du job, entreprise, user title... ?
        -> changer les "owner" en conséquence-->