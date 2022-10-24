# Smart Signage 
<sub> Last update : October 23, 2022  
Author : Laura-Lee Hollande
Team : Victor Leroy, Karine Vinette, Thomas Planchard, Paul Nowak</sub>

## Table of Contents
- [Smart Signage](#smart-signage)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Glossary](#glossary)
  - [Scenario/Uses cases](#scenariouses-cases)
  - [Risks and assumptions](#risks-and-assumptions)
    - [About the laws](#about-the-laws)
    - [Energy crisis](#energy-crisis)
  - [Non goals](#non-goals)
  - [Requirements](#requirements)
  - [Error reporting](#error-reporting)
  - [Terminal](#terminal)
    - [Login](#login)
    - [Home page](#home-page)
    - [Commands](#commands)
      - [check](#check)
      - [off](#off)
      - [on](#on)
      - [up](#up)
      - [down](#down)
      - [prediction (out of scope)](#prediction-out-of-scope)
      - [cons](#cons)
    - [Error reporting](#error-reporting-1)
  - [Out of scope](#out-of-scope)

## Introduction
[SignAll](https://signall.com/), in Vierzon needs a new connected product able to know the state and control the LED's remotly. The objective with this project is to enable the company to save on electricity, reduce its ecological footprint and limited unnecessary travel.

## Glossary
| Word                         	| Definition                                                                                                                                                                                                                                                                                                                                                                      	|
|------------------------------	|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| Gateway                      	| It is a piece of networking hardware or software used in telecommunications networks that allows data to flow from one discrete network to another                                                                                                                                                                                                                              	|
| LEDs                         	| It is a semiconductor light source that emits light when current flows through it.  Consumes 10 times less electricity than an incandescent bulb and 6 to 8 times less than a halogen bulb                                                                                                                                                                                      	|
| Signage                      	| Commercial or public display signs                                                                                                                                                                                                                                                                                                                                              	|
| Carte météo de l'électricité 	| A map that will inform in real time about the level of electricity available in the country, thanks to a map of France established by the [Ecowatt](https://www.monecowatt.fr/) device.  Green for a normal situation, orange for a tense electrical situation and red for a very tense situation synonymous with inevitable blackouts if nothing is done to reduce consumption |

## Scenario/Uses cases

<!-- Patrick Pop is the owner of Burger King restaurant in Bourges and have a signage on his restaurant. He wants to reduce his electricity consumption to save money. He needs to control and see how much electricity is consummed.-->

**Scenario 1** : Patrick is the owner of a Burger King restaurant. He is concerned about the environment and wants to save energy this winter. He would like to be able to control the consumption of his signs and add options to save money. 

**Scenario 2** : Paul is a member of the maintenance team. He would like to know the status of the LEDs and be able to control them remotely to know when to move and intervene. This would avoid unnecessary travel.

## Risks and assumptions
### About the laws
In France, according to the law the LEDs must be off during the night, between 1:00AM and 6:00AM (French time).

<!-- ## Need to verify what the law tells about this type of project. -->

### Energy crisis <!--(out of scope ?) (requirements) -->
According to the energy crisis, a new tool will maybe appears to keep up to date about the level of electricity available in France in real time. It is called "carte météo de l'éléctricité".
It will be important to program the LEDs according to the estimates of this card.

## Non goals
This version will **not** support the following features:
- Send information by wifi for security reasons.
- A product just for the new SignAll panels. It must be compatible with all SignAll panels, both new and old, but also with those of the competition.
- It must not make the LEDs flash (visible to the naked eye).


## Requirements
The most important points for our customers are its ecological impact, its electricity consumption and a remote controlits electricity consumption. To answer his needs we can develop some features.

- Turn on the light
- Turn off the light
- Reduce the light intensity
- Increase the light intensity
- Know the led status and control it remotly
- Turn off the light in accord with the law
- In accord to [Ecowatt](https://www.monecowatt.fr/), turn off or reduce the light <!-- (out of scope) -->
- Send a notification when a led is down <!-- (out of scope) -->
- Programable light wich light up on various hours <!-- (out of scope) -->


## Error reporting
If an error occurs, a text will appear in the terminal. The error message may be different depending on the error encountered.

| Error                           	| Description                                                                  	|
|---------------------------------	|------------------------------------------------------------------------------	|
| command not found               	| Can not execute the command                                                  	|
| impossible to turn on the light 	| If the user try to turn on the light between 1:00AM and 6:00AM (French time) 	|
| access denied                   	| Login problem, the password or/and the user job title is not correct         	|

## Terminal
Displayed when the terminal is open, the home page serves three purposes:

- login page
- home page

### Login
To more security, to access our project the user needs to connect itself with :
- user (user section allows for up to 20 characters to be typed.)
- password (The password section allows for up to 12 characters to be typed. To disguise them and prevent hacking, as the user types in the password box, asterisks (*) will appear instead of the characters that they type.)

### Home page
To avoid any mistakes we will display the home page according to the login entered.
- If the user is connected as a member of the maintenance team, he will see the state of all the signages and their LEDs.
- If the user is connected as employees of the place where the signage is, he will be able to control the LEDs of his panel ang get some information. He can not have access to the other brand signage.

### Commands
To control the LEDs and to have access to various information the user will have different commands. After logging in, a summary table of all the commands available for the user and with their description will be visible in the terminal:
<!-- voire pour ajouter screen du terminal -->
| Commands 	| Description                        	| Access restricted to           	|
|----------	|------------------------------------	|--------------------------------	|
| check    	| Check the state of the LEDs and gives some information       	| Maintenance team and the employees of the brand              	|
| off      	| Turn off the LEDs                  	| The brand where the signage is 	|
| on       	| Turn on the LEDs                   	| The brand where the signage is 	|
| up       	| Increase the intensity of the LEDs 	| The brand where the signage is 	|
| down     	| Decrease the intensity of the LEDs 	| The brand where the signage is 	|
| prediction  <!--out of scope-->   	| Gives a prediction of the state of the led over the coming months depending on various factors. 	| The brand where the signage is 	|
| cons    	| Displays a curve of the panel's electricity consumption over the week or month. 	| The brand where the signage is 	|

<!-- mode auto ? -->

#### check
<!-- selected panel or not? -->
Gives information on the status of the different LED grouping of a selected panel. The following information can be found: 
- the global status of the panel
  - panels off because an LED grouping is down
  - nothing to report
- the date of the last maintenance (use the **american date format** : mm-dd-yyyy)
- if there has been any suspicious activity and the source
  - unintentional shutdown of the panel
  - higher than average consumption

<!-- screen de la page terminal check -->

#### off
Turns off the LEDs.

#### on
Turns on the LEDs, by default it lights up at 50% of their capacity.

#### up
Increases the intensity of the entire panel. By using this command several times in a row, the panel intensity increases by 10%. It can go up to the maximum capacity of the LEDs, i.e. 100%. 

#### down
Decreases the intensity of the entire panel. By using this command several times in a row, the panel intensity decreases by 10%. It can go up to the minimum capacity of the LEDs, i.e. 0%. 

#### prediction (out of scope)


#### cons
Using this command, the user can choose between two periods: **"month"** to see consumption for the current month or **"year"** to see consumption for the past year and compare.
<!-- #### check -->

### Error reporting
If an error occurs, a text will appear in the terminal. The error message may be different depending on the error encountered.

| Error                           	| Description                                                                  	|
|---------------------------------	|------------------------------------------------------------------------------	|
| command not found               	| Can not execute the command                                                  	|
| impossible to turn on the light 	| If the user try to turn on the light between 1:00AM and 6:00AM (French time) 	|
| access denied                   	| Login problem, the password or/and the user job title is not correct         	|


## Out of scope
For a better user experience, create a web interface will be easier. It will be used to manage the LEDs and see their states.
We will use all the features mentioned above to develop our web interface.

<!-- Questions
- même interface pour maintenance -> login
- qui sont les clients et les différences en conséquences 
- vérifier les entrées du login pour être le plus clair : intitulé du job, entreprise, user title... ?
        -> changer les "owner" en conséquence
- différencier les différents panneaux ?-->
