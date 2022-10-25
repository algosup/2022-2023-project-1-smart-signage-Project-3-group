<font size="150"> Appsolu </font> 
<sub> Last update : October 24, 2022  
Author : Laura-Lee Hollande
Team : Victor Leroy, Karine Vinette, Thomas Planchard, Paul Nowak</sub>

<font size="5"> Table of Contents </font>

- [Introduction](#introduction)
- [Glossary](#glossary)
- [Scenario/Uses cases](#scenariouses-cases)
  - [Anya Manya - Owner of a Burger King](#anya-manya---owner-of-a-burger-king)
  - [Oscar Pali - Maintenance technician at SignAll](#oscar-pali---maintenance-technician-at-signall)
  - [Kévin Ohama - Employee at Burger King](#kévin-ohama---employee-at-burger-king)
  - [Katia Soneta - Works at SignAll](#katia-soneta---works-at-signall)
  - [Clarisse Polia - Citizen](#clarisse-polia---citizen)
- [Risks and assumptions](#risks-and-assumptions)
  - [About the laws](#about-the-laws)
- [Non goals](#non-goals)
- [Requirements](#requirements)
- [Terminal](#terminal)
  - [Commands](#commands)
    - [check command](#check-command)
    - [off command](#off-command)
    - [on command](#on-command)
    - [up command](#up-command)
    - [down command](#down-command)
  - [Error reporting](#error-reporting)
- [Out of scope](#out-of-scope)
  - [Energy crisis](#energy-crisis)
  - [Information about the electricity consumption](#information-about-the-electricity-consumption)
  - [Progamable lights and auto mode](#progamable-lights-and-auto-mode)

# Introduction
[SignAll](https://signall.com/), in Vierzon needs a new connected product able to know the state and control the LED's remotly. The objective with this project is to enable the company to save on electricity, reduce its ecological footprint and limited unnecessary travel.

# Glossary
| Word                         	| Definition                                                                                                                                                                                                                                                                                                                                                                      	|
|------------------------------	|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------	|
| Gateway                      	| It is a piece of networking hardware or software used in telecommunications networks that allows data to flow from one discrete network to another                                                                                                                                                                                                                              	|
| LEDs                         	| It is a semiconductor light source that emits light when current flows through it.  Consumes 10 times less electricity than an incandescent bulb and 6 to 8 times less than a halogen bulb                                                                                                                                                                                      	|
| Signage                      	| Commercial or public display signs                                                                                                                                                                                                                                                                                                                                              	|
| Carte météo de l'électricité 	| A map that will inform in real time about the level of electricity available in the country, thanks to a map of France established by the [Ecowatt](https://www.monecowatt.fr/) device.  Green for a normal situation, orange for a tense electrical situation and red for a very tense situation synonymous with inevitable blackouts if nothing is done to reduce consumption |
| Lux | It is the unit of illuminance in the International System of Units |
| LoRa  | LoRa (short for long range) is a spread spectrum modulation technique derived from chirp spread spectrum (CSS) technology. LoRa is a long range, low power wireless platform that became the de facto wireless platform of Internet of Things (IoT). |

# Scenario/Uses cases

## Anya Manya - Owner of a Burger King
<img width="1133" alt="personae-Anya" src="https://user-images.githubusercontent.com/71769490/197859835-fdf93644-0433-4d6c-9a73-7596f2171acd.png">

**Use case :** Anya Manya is the owner of Burger King restaurant. She is very involved in her restaurant and likes to keep an eye on everything. This winter she would like to reduce her electricity consumption to save money. She thinks the signages in her restaurant is a good idea to save money by dimming the lights or turning them off at off-peak times. Knowing the state of the signages and their consumption are key elements in the management of her business in order to act accordingly and not lose money. In addition, she could reduce the light pollution of its establishment.

## Oscar Pali - Maintenance technician at SignAll
<img width="1133" alt="personae-oscar" src="https://user-images.githubusercontent.com/71769490/197859864-5d5dcbd4-cc65-4a0a-9473-6064bedcde21.png">

**Use case :** Oscar Pali has just become a father and makes his child his priority. He would like to spend less time on the roads and only move when he is sure that there is a need for maintenance. He could work from home and have an overview of the signages in his area and therefore anticipate maintenance operations.

## Kévin Ohama - Employee at Burger King
<img width="1133" alt="personae-kevin" src="https://user-images.githubusercontent.com/71769490/197868262-67a9b1a8-d84c-4529-9a6f-0c9442f4ca61.png">

**Use case :** Kévin Ohama is a student who works to pay his study. he had a lot of responsabilies in his work like the signages management which is manage manually. He does not have lot of knowledge in technologies so he will need a little formation if the way to manage all the signages will change.

## Katia Soneta - Works at SignAll
<img width="1133" alt="personae-katia" src="https://user-images.githubusercontent.com/71769490/197868287-be392e1a-9ecc-433d-8157-c0cc12ed2421.png">

**Use case :** Katia Soneta works at SignAll and she is very implicated in the development of the brand. Her first goal is to earn new customer for SignAll, for that she needs to offer something different and usefull for the potentials customers.

## Clarisse Polia - Citizen
<img width="1133" alt="personae-clarisse" src="https://user-images.githubusercontent.com/71769490/197873007-0800e75c-03e1-4dc6-9fb4-95133dfa32c5.png">

**Use case :** Clarisse Polia has just arrived in Lamotte Beuvron, she has no mark in the city and is a little bit lost. She walks a lot around the city to discover new place. For her, signage is very important because it is the only thing which serves as a reference point that is why she does not like when a signage does not work correctly ("Burger King" who becomes "Brger Kin" for example).

# Risks and assumptions
## About the laws
In France, according to the law the LEDs must be off during the night, between 1:00AM and 6:00AM (French time).

# Non goals
This version will **not** support the following features:
- Send information by wifi for security reasons.
- A product just for the new SignAll panels. It must be compatible with all SignAll panels, both new and old, but also with those of the competition.
- It must not make the LEDs flash (visible to the naked eye).

# Requirements
The most important points for our customers are its ecological impact, its electricity consumption and a remote controlits electricity consumption. To answer his needs we can develop some features.

- Turn on the light
- Turn off the light
- Reduce the light intensity
- Increase the light intensity
- Know the led status and control it remotly
- Turn off the light in accord with the law

# Terminal
Displayed when the terminal is open:
- Commands page

## Commands
To control the LEDs and to have access to various information the user will have different commands. After logging in and select a signage, a summary table of all the commands available for the user and with their description will be visible in the terminal:


### check command
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
Accessible by employees and the maintenance team.

### off command
Turns off the LEDs.
Accessible by employees and the maintenance team.

### on command
Turns on the LEDs, by default it lights up at 50% of their capacity.
Accessible by employees and the maintenance team.

### up command
Increases the intensity of the entire panel. By using this command several times in a row, the panel intensity increases by 10%. It can go up to the maximum capacity of the LEDs, i.e. 100%. 
Accessible by employees and the maintenance team.

### down command
Decreases the intensity of the entire panel. By using this command several times in a row, the panel intensity decreases by 10%. It can go up to the minimum capacity of the LEDs, i.e. 0%. 
Accessible by employees and the maintenance team.

Here a summary of all orders for this version of the project:
| Commands 	| Description                        	| Access restricted to           	|
|----------	|------------------------------------	|--------------------------------	|
| check    	| Check the state of the LEDs and gives some information       	| Employees and the maintenance team             	|
| off      	| Turn off the LEDs                  	| Employees and the maintenance team 	|
| on       	| Turn on the LEDs                   	| Employees and the maintenance team 	|
| up       	| Increase the intensity of the LEDs 	| Employees and the maintenance team 	|
| down     	| Decrease the intensity of the LEDs 	| Employees and the maintenance team 	|

## Error reporting
If an error occurs, a text will appear in the terminal. The error message may be different depending on the error encountered.

| Error                           	| Description                                                                  	|
|---------------------------------	|------------------------------------------------------------------------------	|
| command not found               	| Can not execute the command                                                  	|
| impossible to turn on the light 	| If the user try to turn on the light between 1:00AM and 6:00AM (French time) 	|
| on-going maintenance 	| maintenance is in progress, the panel cannot be used. 	|

# Out of scope
Other features can be implemented in the module for a better control of the module, energy and to reduce its ecological impact.
Here some ideas for new features :
- In accord to [Ecowatt](https://www.monecowatt.fr/), turn off or reduce the light
- Send a notification when a led grouping is down
- Programable light wich light up on various hours
- See the electricity consumption of signage
- Auto mode

## Energy crisis
According to the energy crisis, a new tool will maybe appears to keep up to date about the level of electricity available in France in real time. It is called "carte météo de l'éléctricité".
It will be important to program the LEDs according to the estimates of this card.

## Information about the electricity consumption
With a command (cons for exemple), the user can choose between two periods: **"month"** to see consumption for the current month or **"year"** to see consumption for the past year and compare.

## Progamable lights and auto mode 
Programable light could be an idea for the company where the signage is, the employees could choose a time period when the signage need to be turn on and avoid the case that an employee miss to turn off the signge during the night. 
To see further we can imagine an auto mode, the signage will be able to lights from external information like the weather, the the use of the site <!--(fréquentation du lieu)--> or the luminosity. This could be possible thanks to different sensors.

