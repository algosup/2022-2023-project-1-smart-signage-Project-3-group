<font size="150"> Appsolu </font> 
<sub> Last update : October 24, 2022  
Author : Laura-Lee Hollande
Team : Victor Leroy, Karine Vinette, Thomas Planchard, Paul Nowak</sub>

<font size="5"> Table of Contents </font>

- [Introduction](#introduction)
- [Glossary](#glossary)
- [Scenario/Uses cases](#scenariouses-cases)
- [Risks and assumptions](#risks-and-assumptions)
  - [About the laws](#about-the-laws)
- [Non goals](#non-goals)
- [Requirements](#requirements)
- [Terminal](#terminal)
  - [Commands](#commands)
    - [check](#check)
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

<!-- Patrick Pop is the owner of Burger King restaurant in Bourges and have a signage on his restaurant. He wants to reduce his electricity consumption to save money. He needs to control and see how much electricity is consummed.-->
<img width="1133" alt="personae-oscar" src="https://user-images.githubusercontent.com/71769490/197851394-2eae68e7-f02d-42cc-997e-2709795f705c.png">


**Scenario 1** : Patrick is the owner of a Burger King restaurant. He is concerned about the environment and wants to save energy this winter. He would like to be able to control the consumption of his signs and add options to save money. 

**Scenario 2** : Paul is a member of the maintenance team. He would like to know the status of the LEDs and be able to control them remotely to know when to move and intervene. This would avoid unnecessary travel.

# Risks and assumptions
## About the laws
In France, according to the law the LEDs must be off during the night, between 1:00AM and 6:00AM (French time).

<!-- ## Need to verify what the law tells about this type of project. -->

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
<!-- voire pour ajouter screen du terminal -->
| Commands 	| Description                        	| Access restricted to           	|
|----------	|------------------------------------	|--------------------------------	|
| check    	| Check the state of the LEDs and gives some information       	| Maintenance team and the employees of the brand              	|
| off      	| Turn off the LEDs                  	| The brand where the signage is 	|
| on       	| Turn on the LEDs                   	| The brand where the signage is 	|
| up       	| Increase the intensity of the LEDs 	| The brand where the signage is 	|
| down     	| Decrease the intensity of the LEDs 	| The brand where the signage is 	|

### check
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

### off command
Turns off the LEDs.

### on command
Turns on the LEDs, by default it lights up at 50% of their capacity.

### up command
Increases the intensity of the entire panel. By using this command several times in a row, the panel intensity increases by 10%. It can go up to the maximum capacity of the LEDs, i.e. 100%. 

### down command
Decreases the intensity of the entire panel. By using this command several times in a row, the panel intensity decreases by 10%. It can go up to the minimum capacity of the LEDs, i.e. 0%. 

## Error reporting
If an error occurs, a text will appear in the terminal. The error message may be different depending on the error encountered.

| Error                           	| Description                                                                  	|
|---------------------------------	|------------------------------------------------------------------------------	|
| command not found               	| Can not execute the command                                                  	|
| impossible to turn on the light 	| If the user try to turn on the light between 1:00AM and 6:00AM (French time) 	|
| access denied                   	| Login problem, the password or/and the user job title is not correct         	|


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

<!-- 0.72W max led -->
