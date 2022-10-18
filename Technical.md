- [1. Front matter](#1-front-matter)
- [2. Introduction](#2-introduction)
  - [a. Overview](#a-overview)
  - [b. Glossary  or Terminology](#b-glossary--or-terminology)
  - [c. Context or Background](#c-context-or-background)
  - [d. Product and Technical Requirements](#d-product-and-technical-requirements)
  - [e. Future Goals](#e-future-goals)
  - [f. Assumptions](#f-assumptions)
- [3. Solutions](#3-solutions)
  - [a. Current or Existing Solution / Design](#a-current-or-existing-solution--design)
  - [b. Suggested or Proposed Solution / Design](#b-suggested-or-proposed-solution--design)
  - [c. Test Plan](#c-test-plan)
  - [d. Release / Roll-out and Deployment Plan](#d-release--roll-out-and-deployment-plan)
  - [e. Alternate Solutions](#e-alternate-solutions)
- [4. Further Considerations](#4-further-considerations)
  - [a. Impact on other teams](#a-impact-on-other-teams)
  - [b. Third-party services and platforms considerations](#b-third-party-services-and-platforms-considerations)
  - [c. Cost analysis](#c-cost-analysis)
  - [d. Security considerations](#d-security-considerations)
  - [e. Privacy considerations](#e-privacy-considerations)
  - [f. Regional considerations](#f-regional-considerations)
  - [g. Accessibility considerations](#g-accessibility-considerations)
  - [h. Operational considerations](#h-operational-considerations)
  - [i. Risks](#i-risks)
  - [j. Support considerations](#j-support-considerations)
- [5. Work](#5-work)
  - [a. Work estimates and timelines](#a-work-estimates-and-timelines)
  - [b. Prioritization](#b-prioritization)
  - [c. Milestones](#c-milestones)
  - [d. Future work](#d-future-work)
- [6. Deliberation](#6-deliberation)
  - [a. Discussion](#a-discussion)
  - [b. Open Questions](#b-open-questions)
- [7. End Matter](#7-end-matter)
  - [a. Related Work](#a-related-work)
  - [b. References](#b-references)
  - [c. Acknowledgments](#c-acknowledgments)


# 1. Front matter
Smart Signage Project 1
<br>
Author: Thomas Planchard 
<br>
Team: Karine Vinette, Paul Nowak, Victor Leroy, Laura-Lee Hollande, Thomas Planchard
<br>
Created on : 
<br>
Last updated : 
<br>
Bug Tracker by Victor Leroy : https://docs.google.com/spreadsheets/d/12PCz3j1eYLg3Uv70rVtRlU4je5LbpEMlmCV1WTYXLxg/edit?usp=sharing

# 2. Introduction
## a. Overview
With the actual crisis we need to find a solution to decrease the consumption of electricity and one of the aspect that we can improve are LED's signs in the street. This is the reason why 
[SignAl](https://signall.com/) wants to invent a new product which would allow them to control the luminous display of the signs they propose remotly and to know the state of the LEDs. The objective behind this project is to enable the company to save electricity, reduce its ecological footprint and earn time for their technician. 

## b. Glossary  or Terminology

| Terms                        | Definition             |
| ---------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| LoRa|  LoRa (short for long range) is a spread spectrum modulation technique derived from chirp spread spectrum (CSS) technology. LoRa is a long range, low power wireless platform that became the de facto wireless platform of Internet of Things (IoT).|
| Gateway| It is a piece of networking hardware or software used in telecommunications networks that allows data to flow from one discrete network to another. |
| LEDs| It is a semiconductor light source that emits light when current flows through it. Consumes 10 times less electricity than an incandescent bulb and 6 to 8 times less than a halogen bulb.|
| STM32 Bluepill| It's an electronic board based on a STM32F103C8T6 microcontroller.|
| TTL module | Converter USB to TTL|



## c. Context or Background
At the moment, all the products sold by [SignAl](https://signall.com/) are not connected. Users can't switch on and off their signage with a remote. This poses a problem especially now with the increase of the electricity price. With this project the company wants their future users to be able to switch on or off their signs as they want. In addition, **SignAl** technicians cannot at the moment, identify if a sign is broken or not without visiting the site. This second issue wastes a lot of time for technicians and need to be solved. Our role is to create a solution to all the problems and allow the company to follow its ecological policy.  


## d. Product and Technical Requirements

Product requirements :
- Our product need to be a device that we can plug in a LED sign to transform the normal sign into a connected sign. Thanks to this the sign can be controlled remotely. In addition it will be possible to recognize if the sign is broken by controlling the intensity of the LEDs thanks to a sensor. 

Technical requirements :
- Following the company's need, **ALGOSUP** advised us by purchasing by itself a "LoRa-E5 Development Kit" and a "Bluepill STM32" to use this type of hardware. 
- We have at our disposal LEDs and different sensors to simulate a real panel. 
- Like in addition to be a real project it's also a school project. The school decided that the usage of TinyGo is mandatory. 

## e. Future Goals
At the moment, the customer only needs the device without a user interface. They just want to see our ideas and how we solve this problem technically. So the main goal is to bring a solution to the customer. If we have time, a user interface can be created to make our product more user-friendly, even if it is a beta version. If a user interface has to be made it will be in a second time.

## f. Assumptions
In order to create our solution product, we need like written above a LoRa-E5 Development Kit, an Arduino STM32 to flash the code on the board, some LEDs to simulate a sign and a captor of electric intensity. 
  
# 3. Solutions
## a. Current or Existing Solution / Design

[Trilux](https://www.trilux.com/fr/produits/gestion-declairage-livelink/livelink/) is a lighting company. They created a solution call **Livelink**. Livelink is an intelligent lighting network where all the LEDs are connected together. The lighting is biodynamic and controlled by sensors. On the official website it is written "Intelligent lighting installations, which thanks to their sensor control reduce energy consumption, self-report maintenance needs and are simple to use... the lighting market has been transformed since the arrival of LED technology, digital technology and major social trends, such as connectivity and big data.". <br>
Their product is the closest thing to our solution. The big difference is that it's not for advertising panels. 

## b. Suggested or Proposed Solution / Design 




First of all, this project needs hardware to work and this is how you need to connect the different device :
- some male and female pin cables  
- 1 breadboard
- a bluepill
- 1 TTL module
- LEDs
- 1 switch 
- Regular tools for hardware

//IMAGE MONTAGE PLUS EXPLICATION //

Our solution consists in making communicate the Lora-E5 board and the bluepill. The blue pill is the brain of this device. As a reminder, the programm need to be in Go and TinyGo for the bluepill and for the Lora-E5 we can communicate with it only by "AT COMMAND". To communicate between both boards we are going to use the protocole LoRa (refer to the glossary). <br>
**Why use LoRa and LoRaWAN Technologies?** <br>
There are several advantages of using LoRa and LoRaWAN technology.

**Long Range**<br>
LoRa enables wireless communication over far longer ranges compared to Wi-Fi or BLE.

**Low Power**<br>
Compared to WiFi, BLE or Satellite Communication, devices in a LoRa network consumes relatively less power. This allows them to run on renewable energy (eg. Solar power), and reduces battery replacement costs. Edge nodes can run on a single battery for a few years.

**Secure**<br>
LoRaWAN networks are protected by end-to-end AES128 encryption, mutual authentication, integrity protection, and confidentiality.

**Standardized**<br>
LoRa & LoRaWAN are widely accepted technologies and protocols, allowing you to capitalise on device interoperability and global availability of LoRaWAN networks for fast and convenient deployment of IoT applications anywhere.

**Low Cost**<br>
LoRa operates on unlicensed frequency spectrums, which reduces fees for network operations. In addition, a wide variety of pre-licensed LoRa development platforms reduces legislative costs.

**Flexible**<br>
LoRa & LoRaWAN combine the best of other technologies, and can be used in a variety of environments and networks. Like Wi-Fi, LoRaWAN operates in the unlicensed band and supports indoor applications; like Cellular, LoRa Technology is highly secure from end devices to the application server, and is suitable for outdoor applications.



After these explanations about why we use Lora protocole for this project, we gonna explain algo

Turn on the light
Turn off the light
Reduce the light intensity
Increase the light intensity
Know the led status and control it remotly
Turn off the light in accord with the law
In accord to Ecowatt, turn off or reduce the light
Send a notification when a led is down
Programable light wich light up on various hours


Business Logic
API changes
Pseudocode
Flowcharts
Error states
Failure scenarios
Conditions that lead to errors and failures
Limitations
Presentation Layer
User requirements
UX changes
UI changes
Wireframes with descriptions
Links to UI/UX designer’s work

UI states
Error handling
Other questions to answer
How will the solution scale?
What are the limitations of the solution?
How will it recover in the event of a failure?
How will it cope with future requirements?
## c. Test Plan
**(VICTOR)**

## d. Release / Roll-out and Deployment Plan

**SignAll** wants to release the product as soon as possible. They want to create a new line of product connected based on our product. They said " We want to be the first on this market". This product needs to be available for all the customers of **SignAll** but not only because they also want to sell this product to the competitor's customers. Indeed, if the product can be used on different signs of different brands, the market has no limit. In addition it's a really positive point for their brand image to be the precursor on this type of product.  

## e. Alternate Solutions 

We can see this problem from a different perspective and propose a different approach to solving it. One of the point that we can discuss and why use Lora instead of Wifi. 
# 4. Further Considerations
## a. Impact on other teams

How will this increase the work of other people?
## b. Third-party services and platforms considerations

Is it really worth it compared to building the service in-house?
What are some of the security and privacy concerns associated with the services/platforms?
How much will it cost?
How will it scale?
What possible future issues are anticipated? 
## c. Cost analysis

What is the cost to run the solution per day?
What does it cost to roll it out? 

## d. Security considerations

What are the potential threats?
How will they be mitigated?
How will the solution affect the security of other components, services, and systems?

## e. Privacy considerations

Does the solution follow local laws and legal policies on data privacy?
How does the solution protect users data privacy?
What are some of the tradeoffs between personalization and privacy in the solution? 

## f. Regional considerations

What is the impact of internationalization and localization on the solution?
What are the latency issues?
What are the legal concerns?
What is the state of service availability?
How will data transfer across regions be achieved and what are the concerns here? 

## g. Accessibility considerations

How accessible is the solution?
What tools will you use to evaluate its accessibility? 
## h. Operational considerations

Does this solution cause adverse aftereffects?
How will data be recovered in case of failure?
How will the solution recover in case of failure?
How will operational costs be kept low while delivering increased value to the users? 

## i. Risks

What risks are being undertaken with this solution?
Are there risks that once taken can’t be walked back?
What is the cost-benefit analysis of taking these risks? 
## j. Support considerations

How will the support team get across information to users about common issues they may face while interacting with the changes?
How will we ensure that the users are satisfied with the solution and can interact with it with minimal support?
Who is responsible for the maintenance of the solution?
How will knowledge transfer be accomplished if the project owner is unavailable? 

# 5. Work
## a. Work estimates and timelines

List of specific, measurable, and time-bound tasks
Resources needed to finish each task
Time estimates for how long each task needs to be completed
## b. Prioritization

Categorization of tasks by urgency and impact
## c. Milestones

Dated checkpoints when significant chunks of work will have been completed
Metrics to indicate the passing of the milestone
## d. Future work

List of tasks that will be completed in the future
# 6. Deliberation
## a. Discussion

Elements of the solution that members of the team do not agree on and need to be debated further to reach a consensus.
## b. Open Questions

Questions about things you do not know the answers to or are unsure that you pose to the team and stakeholders for their input. These may include aspects of the problem you don’t know how to resolve yet. 
# 7. End Matter
## a. Related Work

Any work external to the proposed solution that is similar to it in some way and is worked on by different teams. It’s important to know this to enable knowledge sharing between such teams when faced with related problems. 
## b. References

Links to documents and resources that you used when coming up with your design and wish to credit. 
## c. Acknowledgments
