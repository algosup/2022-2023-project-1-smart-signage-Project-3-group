# Smart Signage 
<sub> Last update : October 24, 2022  
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
    - [Error reporting](#error-reporting-1)
  - [Out of scope](#out-of-scope)
    - [Future commands](#future-commands)
      - [cons](#cons)

## Introduction
[SignAll](https://signall.com/), in Vierzon needs a new connected product able to know the state and control the LED's remotly. The objective with this project is to enable the company to save on electricity, reduce its ecological footprint and limited unnecessary travel.

## Glossary
| Word                         	| Definition                                                                                                                                                                                                                                                                                                                                                                      	|
|------------------------------	|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| Gateway                      	| It is a piece of networking hardware or software used in telecommunications networks that allows data to flow from one discrete network to another                                                                                                                                                                                                                              	|
| LEDs                         	| It is a semiconductor light source that emits light when current flows through it.  Consumes 10 times less electricity than an incandescent bulb and 6 to 8 times less than a halogen bulb                                                                                                                                                                                      	|
| Signage                      	| Commercial or public display signs                                                                                                                                                                                                                                                                                                                                              	|
| Carte météo de l'électricité 	| A map that will inform in real time about the level of electricity available in the country, thanks to a map of France established by the [Ecowatt](https://www.monecowatt.fr/) device.  Green for a normal situation, orange for a tense electrical situation and red for a very tense situation synonymous with inevitable blackouts if nothing is done to reduce consumption |
| Lux | It is the unit of illuminance in the International System of Units |
| LoRa  | LoRa (short for long range) is a spread spectrum modulation technique derived from chirp spread spectrum (CSS) technology. LoRa is a long range, low power wireless platform that became the de facto wireless platform of Internet of Things (IoT). |

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
- Send a notification when a led grouping is down <!-- (out of scope) -->
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

- Login page
- Home page
- Commands page

### Login
To more security, to access our project the user needs to connect itself with :
- user (user section allows for up to 20 characters to be typed.)
- password (The password section allows for up to 12 characters to be typed. To disguise them and prevent hacking, as the user types in the password box, asterisks (*) will appear instead of the characters that they type.)

### Home page
To avoid any mistakes we will display the home page according to the login entered.
**If the user is connected as a member of the maintenance team:** 
After login the member of the maintenance team have an overview of all the brands and signages of his sector, he can see the the state of each signage and get an historic with information about the life of the signage (his electricity consumption, if there was an overheating... )

**If the user is connected as employees of the place where the signage is:** 
After login the employee have an overview of all the signages of the brand where he is. He can see which signage is off or if there is a problem with a signage. He can select one panel by his number to get more information and control it. He can not have access to the other brand signage.

### Commands
To control the LEDs and to have access to various information the user will have different commands. After logging in and select a signage, a summary table of all the commands available for the user and with their description will be visible in the terminal:
<!-- voire pour ajouter screen du terminal -->
| Commands 	| Description                        	| Access restricted to           	|
|----------	|------------------------------------	|--------------------------------	|
| check    	| Check the state of the LEDs and gives some information       	| Maintenance team and the employees of the brand              	|
| off      	| Turn off the LEDs                  	| The brand where the signage is 	|
| on       	| Turn on the LEDs                   	| The brand where the signage is 	|
| up       	| Increase the intensity of the LEDs 	| The brand where the signage is 	|
| down     	| Decrease the intensity of the LEDs 	| The brand where the signage is 	|

<!-- mode auto ? -->

#### check
<!-- selected panel or not? -->
Gives information on the status of the different LED grouping of a selected panel. The following information can be found: 
- if the panel is currently on or off
- the global status of the panel
  - panels off because a LED grouping is down (precise which LED grouping is down by his number : "LED grouping X is down" )
  - nothing to report
- the date of the last maintenance and the currently date (use the **american date format** : mm-dd-yyyy)
- if there has been any suspicious activity and the source
  - unintentional shutdown of the panel
  - higher than average consumption
  - overheating
- give a prediction of the state of the LEDs over a month
  - can be impacted by the weather

<!-- screen de la page terminal check -->

#### off
Turns off the LEDs.

#### on
Turns on the LEDs, by default it lights up at 50% of their capacity.

#### up
Increases the intensity of the entire panel. By using this command several times in a row, the panel intensity increases by 10%. It can go up to the maximum capacity of the LEDs, i.e. 100%. 

#### down
Decreases the intensity of the entire panel. By using this command several times in a row, the panel intensity decreases by 10%. It can go up to the minimum capacity of the LEDs, i.e. 0%. 

### Error reporting
If an error occurs, a text will appear in the terminal. The error message may be different depending on the error encountered.

| Error                           	| Description                                                                  	|
|---------------------------------	|------------------------------------------------------------------------------	|
| command not found               	| Can not execute the command                                                  	|
| impossible to turn on the light 	| If the user try to turn on the light between 1:00AM and 6:00AM (French time) 	|
| access denied                   	| Login problem, the password or/and the user job title is not correct         	|


## Out of scope
For a better user experience, create a web interface will be easier. It will be used to manage the LEDs and see their states.
We will use all the features mentioned above and new one to develop our web interface.

### Future commands

#### cons
Using this command, the user can choose between two periods: **"month"** to see consumption for the current month or **"year"** to see consumption for the past year and compare.
<!-- #### check -->

<!-- Questions
- même interface pour maintenance -> login
- qui sont les clients et les différences en conséquences 
- vérifier les entrées du login pour être le plus clair : intitulé du job, entreprise, user title... ?
        -> changer les "owner" en conséquence
- différencier les différents panneaux ?-->

<!-- 0.72W max led -->
