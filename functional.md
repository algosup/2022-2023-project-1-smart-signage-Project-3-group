# Smart Signage 
<sub> Last update : October 21, 2022  
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
  - [What the project is not](#what-the-project-is-not)
  - [Requirements](#requirements)
  - [Error reporting](#error-reporting)
  - [Terminal](#terminal)
    - [Login](#login)
    - [Home page](#home-page)
  - [Out of scope](#out-of-scope)

<a  name="overview"/></a>

## Introduction
[SignAll](https://signall.com/), in Vierzon needs a new connected product able to know the state and control the LED's remotly. The objective with this project is to enable the company to save on electricity, reduce its ecological footprint and limited unnecessary travel.

<a name="voc"/></a>

## Glossary
| Word                         	| Definition                                                                                                                                                                                                                                                                                                                                                                      	|
|------------------------------	|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| Gateway                      	| It is a piece of networking hardware or software used in telecommunications networks that allows data to flow from one discrete network to another                                                                                                                                                                                                                              	|
| LEDs                         	| It is a semiconductor light source that emits light when current flows through it.  Consumes 10 times less electricity than an incandescent bulb and 6 to 8 times less than a halogen bulb                                                                                                                                                                                      	|
| Signage                      	| Commercial or public display signs                                                                                                                                                                                                                                                                                                                                              	|
| Carte météo de l'électricité 	| A map that will inform in real time about the level of electricity available in the country, thanks to a map of France established by the [Ecowatt](https://www.monecowatt.fr/) device.  Green for a normal situation, orange for a tense electrical situation and red for a very tense situation synonymous with inevitable blackouts if nothing is done to reduce consumption |

<a name="scenario"/></a>

## Scenario/Uses cases

<!-- Patrick Pop is the owner of Burger King restaurant in Bourges and have a signage on his restaurant. He wants to reduce his electricity consumption to save money. He needs to control and see how much electricity is consummed.-->

**Scenario 1** : Patrick is the owner of a Burger King restaurant. He is concerned about the environment and wants to save energy this winter. He would like to be able to control the consumption of his signs and add options to save money. 

**Scenario 2** : Paul is a member of the maintenance team. He would like to know the status of the LEDs and be able to control them remotely to know when to move and intervene. This would avoid unnecessary travel.


<a name="risk"/></a>

## Risks and assumptions
### About the laws
In France, according to the law the LEDs must be off during the night, between 1:00AM and 6:00AM (French time).

<!-- ## Need to verify what the law tells about this type of project. -->

### Energy crisis <!--(out of scope ?) (requirements) -->
According to the energy crisis, a new tool will maybe appears to keep up to date about the level of electricity available in France in real time. It is called "carte météo de l'éléctricité".
It will be important to program the LEDs according to the estimates of this card.

<!-- ## Check privacy -->


<a name="is-not"/></a>

## What the project is not
This new product do not send information by wifi for security reasons.

<a name="needs"/></a>

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

<a name="error"/></a>

## Error reporting
If an error occurs, a text will appear in the terminal. The error message may be different depending on the error encountered.

| Error                           	| Description                                                                  	|
|---------------------------------	|------------------------------------------------------------------------------	|
| command not found               	| Can not execute the command                                                  	|
| impossible to turn on the light 	| If the user try to turn on the light between 1:00AM and 6:00AM (French time) 	|
| access denied                   	| Login problem, the password or/and the user job title is not correct         	|

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

<!-- suivie de la consommation -> maintenance + restaurant -->

<a name="out-of-scope"/></a>

## Out of scope
For a better user experience, create a web interface will be easier. It will be used to manage the LEDs and see their states.
We will use all the features mentioned above to develop our web interface.

<!-- Questions
- même interface pour maintenance -> login
- qui sont les clients et les différences en conséquences 
- vérifier les entrées du login pour être le plus clair : intitulé du job, entreprise, user title... ?
        -> changer les "owner" en conséquence-->
