# MaaS  [![CircleCI](https://circleci.com/gh/jumaallan/MaaS.svg?style=shield)](https://circleci.com/gh/jumaallan/MaaS) [![Maintainability](https://api.codeclimate.com/v1/badges/798db88c95d4fef56e5f/maintainability)](https://codeclimate.com/github/jumaallan/MaaS/maintainability)  [![Test Coverage](https://api.codeclimate.com/v1/badges/798db88c95d4fef56e5f/test_coverage)](https://codeclimate.com/github/jumaallan/MaaS/test_coverage)

MaaS provides convenient access to the [Safaricom MPESA Daraja API](https://developer.safaricom.co.ke/apis-explorer) for Developers, Tech Ethusiasts, Companies and Corporates. 

MaaS makes this possible by making all the API's available on Daraja accessible as a Service. The platform also offers a Dashboard to monitor your activity on Daraja, an Android SDK and inbuilt Webhooks configurations to hook it to your own system with ease! 

MaaS is built using [Golang]() , to make sure it is as reliable as it can be due to its great memory utilization and speed!  

It’s easy. We promise. :rocket:

> MaaS does not replace Daraja in any way, but offer a managed access to Daraja for FREE!

> The idea is still in Brainstorming & Development Mode, and this repo acts as the project host :rocket:

> MaaS need to run on a server with SSL enabled, in `PRODUCTION ENV`!

## Getting Started

Setting up MaaS in your own Infrastructure is relatively easy! You only need to install [Golang]() and [Docker]() if you want to run the setup in Developer Mode.

If you don't intend to modify the setup, you only need to setup [Docker]() in your deployment server. 

## MaaS Modules - For Developers

For the developers who feel that they need to work on additional features on the platform, this is a breakdown of how the service is setup. The following modules make up the MaaS:

* mod_authentication
* mod_safaricom
* mod_reports
* mod_callbacks
* mod_health
* mod_timeouts
* mod_api

It is recommended that you run MaaS in SSL Mode, but there is a `DEV ENV` setup that allows you to run the setup in Localhost with SSL disabled for both the server and database level.

## System Design and Setup

MaaS is built using the following tools:

* Golang - Go is a low level programming language with great memory utilization and speed.
* Postgres DB - Is a common and fast sql database that acts as MaaS datastore
* Redis - Acts as a cache
* Docker - Docker takes care of MaaS Deployment with ease without knowing a thing!


## Future Additions

We are thinking about adding the following features/tools in the next versions of MaaS

* Support Document Database's e.g MongoDB
* GraphQL

